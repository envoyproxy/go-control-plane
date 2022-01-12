package matcher

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	pbmatcher "github.com/cncf/xds/go/xds/type/matcher/v3"
	pblegacymatcher "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	"github.com/envoyproxy/go-control-plane/xdsmatcher/pkg/matcher/registry"
	"github.com/envoyproxy/go-control-plane/xdsmatcher/pkg/matcher/types"
	"google.golang.org/protobuf/proto"
)

// MatcherTree is the constructed match tree, allowing for matching input data against the match tree.
type MatcherTree struct {
	onNoMatch *types.OnMatch
	matchRoot types.Matcher

	logger *tracingLogger
}

// Match attempts to match the input data against the match tree.
// An error is returned if something went wrong during match evaluation.
func (m MatcherTree) Match(data types.MatchingData) (types.Result, error) {
	result, err := m.matchRoot.Match(data)
	if err != nil {
		return types.Result{}, err
	}
	if result.MatchResult != nil && result.NeedMoreData {
		return types.Result{}, errors.New("invalid result, invariant violation")
	}

	// We arrived at a result.
	if result.MatchResult != nil {
		// The result is a recursive matcher, so recurse.
		if result.MatchResult.Matcher != nil {
			m.logger.log("recursively matching")
			m.logger.push()
			defer m.logger.pop()
			return result.MatchResult.Matcher.Match(data)
		}

		m.logger.log("arrived at result %+v, %+v", result, result.MatchResult)

		// Otherwise pass through the result.
		return result, nil
	}

	// We failed to arrive at a result due to the data not being available,
	// report this back up.
	if result.NeedMoreData {
		m.logger.log("matcher tree requires more data")
		return result, nil
	}

	m.logger.log("matcher tree failed to match, falling back to OnNoMatch (%v)", m.onNoMatch)
	// We have all the data, so we know that we failed to match. Return OnNoMatch.
	return types.Result{
		MatchResult:  m.onNoMatch,
		NeedMoreData: false,
	}, nil
}

// Create creates a new match tree from the provided match tree configuration.
func Create(matcher *pbmatcher.Matcher) (*MatcherTree, error) {
	return create(matcher, &tracingLogger{})
}

// CreateLegacy creates a new match tree from the provided match tree configuration, specified by the legacy
// format.
func CreateLegacy(matcher *pblegacymatcher.Matcher) (*MatcherTree, error) {
	// Wire cast the matcher to the new matcher format.
	b, err := proto.Marshal(matcher)
	if err != nil {
		return nil, err
	}
	m := &pbmatcher.Matcher{}
	err = proto.Unmarshal(b, m)
	if err != nil {
		return nil, err
	}

	return Create(m)
}

func create(matcher *pbmatcher.Matcher, logger *tracingLogger) (*MatcherTree, error) {
	onNoMatch, err := createOnMatch(matcher.OnNoMatch, logger)
	if err != nil {
		return nil, err
	}

	switch m := matcher.MatcherType.(type) {
	case *pbmatcher.Matcher_MatcherList_:
		matcher, err := createMatcherList(m.MatcherList, logger)
		if err != nil {
			return nil, err
		}
		return &MatcherTree{
			onNoMatch: onNoMatch,
			matchRoot: matcher,
			logger:    logger,
		}, nil
	case *pbmatcher.Matcher_MatcherTree_:
		matcher, err := createMatchTree(m.MatcherTree, logger)
		if err != nil {
			return nil, err
		}

		return &MatcherTree{
			onNoMatch: onNoMatch,
			matchRoot: matcher,
			logger:    logger,
		}, nil
	default:
		return nil, errors.New("not implemented: matcher tree")
	}
}

func createOnMatch(onMatch *pbmatcher.Matcher_OnMatch, logger *tracingLogger) (*types.OnMatch, error) {
	// TODO move this to call sites, this is part of validation
	if onMatch == nil {
		return nil, nil
	}
	switch om := onMatch.OnMatch.(type) {
	case *pbmatcher.Matcher_OnMatch_Matcher:
		matcher, err := create(om.Matcher, logger)
		if err != nil {
			return nil, err
		}

		return &types.OnMatch{
			Matcher: matcher,
		}, nil
	case *pbmatcher.Matcher_OnMatch_Action:
		action, err := registry.ResolveActionExtension(om.Action)
		if err != nil {
			return nil, err
		}
		return &types.OnMatch{Action: action}, nil
	default:
		return nil, errors.New("unexpected on match type")
	}
}

type matcherResult struct {
	match        bool
	needMoreData bool
}

type predicateFunc interface {
	match(types.MatchingData) (matcherResult, error)
}

// InputMatcher is the interface for a matcher implementation, specifying how
// an input string should be matched.
type InputMatcher interface {
	match(value *string) (matcherResult, error)
}

type simpleMatcher interface {
	match(value *string) bool
}

type singlePredicate struct {
	input   types.DataInput
	matcher simpleMatcher
	logger  *tracingLogger
}

func (s singlePredicate) match(data types.MatchingData) (matcherResult, error) {
	matchFunc := func(input *string) *types.OnMatch {
		s.logger.log("attempting to match '%v'", safePrint(input))
		if s.matcher.match(input) {
			return &types.OnMatch{}
		}

		return nil
	}

	result, err := handlePossiblyMissingInput(s.input, data, matchFunc, s.logger)
	if err != nil {
		return matcherResult{}, err
	}

	matchResult := matcherResult{
		match:        result.MatchResult != nil,
		needMoreData: result.NeedMoreData,
	}

	s.logger.log("single predicate evaluation result: %+v", matchResult)

	return matchResult, nil
}

type conjunctionPredicate struct {
	predicates []predicateFunc
	logger     *tracingLogger
}

func (a conjunctionPredicate) match(data types.MatchingData) (matcherResult, error) {
	a.logger.log("performing and matching")
	a.logger.push()
	defer a.logger.pop()
	for _, p := range a.predicates {
		result, err := p.match(data)
		if err != nil {
			return matcherResult{}, err
		}

		if result.needMoreData {
			return result, nil
		}

		if !result.match {
			return result, nil
		}
	}

	return matcherResult{
		match:        true,
		needMoreData: false,
	}, nil
}

type disjunctionPredicate struct {
	predicates []predicateFunc
	logger     *tracingLogger
}

func (a disjunctionPredicate) match(data types.MatchingData) (matcherResult, error) {
	a.logger.log("performing or matching")
	a.logger.push()
	defer a.logger.pop()

	resultCouldChange := false
	for _, p := range a.predicates {
		result, err := p.match(data)
		if err != nil {
			return matcherResult{}, err
		}

		if result.needMoreData {
			resultCouldChange = true
			continue
		}

		if result.match {
			return result, nil
		}
	}

	if resultCouldChange {
		return matcherResult{
			match:        false,
			needMoreData: true,
		}, nil
	}

	return matcherResult{
		match:        false,
		needMoreData: false,
	}, nil
}

type matchEntry struct {
	onMatch   *types.OnMatch
	predicate predicateFunc
}

type regexMatcher struct {
	regex *regexp.Regexp
}

func (r regexMatcher) match(s *string) bool {
	if s == nil {
		return false
	}

	return r.regex.Match([]byte(*s))
}

func createSinglePredicate(predicate *pbmatcher.Matcher_MatcherList_Predicate_SinglePredicate, logger *tracingLogger) (*singlePredicate, error) {
	input, err := registry.ResolveInputExtension(predicate.Input)
	if err != nil {
		return nil, err
	}

	switch m := predicate.Matcher.(type) {
	case *pbmatcher.Matcher_MatcherList_Predicate_SinglePredicate_ValueMatch:
		matcher, err := createValueMatcher(m.ValueMatch)
		if err != nil {
			return nil, err
		}

		return &singlePredicate{
			logger:  logger,
			input:   input,
			matcher: matcher,
		}, nil
	default:
		return nil, errors.New("not implemented: custom match")
	}
}

func createValueMatcher(valueMatcher *pbmatcher.StringMatcher) (simpleMatcher, error) {
	switch m := valueMatcher.MatchPattern.(type) {
	case *pbmatcher.StringMatcher_SafeRegex:
		r, err := regexp.Compile(m.SafeRegex.Regex)
		if err != nil {
			return nil, err
		}
		return regexMatcher{regex: r}, nil
	default:
		return nil, errors.New("not implemented: value match")
	}
}

type matcherList struct {
	matchEntries []matchEntry
	logger       *tracingLogger
}

func (m matcherList) Match(data types.MatchingData) (types.Result, error) {
	m.logger.log("attempting to match against match list")
	m.logger.push()
	defer m.logger.pop()

	for _, me := range m.matchEntries {
		m.logger.log("predicate type %T", me.predicate)
		result, err := me.predicate.match(data)
		if err != nil {
			return types.Result{}, err
		}

		if result.match {
			m.logger.log("found match entry matching input: %+v")
			return types.Result{
				MatchResult:  me.onMatch,
				NeedMoreData: false,
			}, nil
		}
		if result.needMoreData {
			m.logger.log("found match entry that needs more data, bailing")
			return types.Result{
				MatchResult:  nil,
				NeedMoreData: true,
			}, nil
		}
	}

	m.logger.log("no matching match entry was found")

	return types.Result{
		MatchResult:  nil,
		NeedMoreData: false,
	}, nil
}

func createPredicate(predicate *pbmatcher.Matcher_MatcherList_Predicate, logger *tracingLogger) (predicateFunc, error) {
	switch predicate := predicate.MatchType.(type) {
	case *pbmatcher.Matcher_MatcherList_Predicate_AndMatcher:
		predicates := []predicateFunc{}
		for _, p := range predicate.AndMatcher.Predicate {
			pp, err := createPredicate(p, logger)
			if err != nil {
				return nil, err
			}
			predicates = append(predicates, pp)
		}

		return conjunctionPredicate{
			predicates: predicates,
			logger:     logger,
		}, nil
	case *pbmatcher.Matcher_MatcherList_Predicate_OrMatcher:
		predicates := []predicateFunc{}
		for _, p := range predicate.OrMatcher.Predicate {
			pp, err := createPredicate(p, logger)
			if err != nil {
				return nil, err
			}
			predicates = append(predicates, pp)
		}

		return disjunctionPredicate{
			predicates: predicates,
			logger:     logger,
		}, nil
	case *pbmatcher.Matcher_MatcherList_Predicate_SinglePredicate_:
		return createSinglePredicate(predicate.SinglePredicate, logger)
	default:
		return nil, fmt.Errorf("unexpected predicate type %T", predicate)
	}
}

func createMatcherList(list *pbmatcher.Matcher_MatcherList, logger *tracingLogger) (*matcherList, error) {
	ml := &matcherList{logger: logger}
	for _, m := range list.Matchers {
		onMatch, err := createOnMatch(m.OnMatch, logger)
		if err != nil {
			return nil, err
		}
		if onMatch == nil {
			return nil, errors.New("matcher list requires on match")
		}
		predicate, err := createPredicate(m.Predicate, logger)
		if err != nil {
			return nil, err
		}

		entry := matchEntry{onMatch: onMatch, predicate: predicate}

		ml.matchEntries = append(ml.matchEntries, entry)
	}

	return ml, nil
}

type exactMatchTree struct {
	tree   map[string]*types.OnMatch
	input  types.DataInput
	logger *tracingLogger
}

func (m exactMatchTree) Match(data types.MatchingData) (types.Result, error) {
	m.logger.log("attempting to match against exact tree")
	m.logger.push()
	defer m.logger.pop()

	matchFunc := func(data *string) *types.OnMatch {
		if data == nil {
			m.logger.log("attempted to exact match nil, no match")
			return nil
		}
		if r, ok := m.tree[*data]; ok {
			m.logger.log("input '%s' found in exact match map", *data)
			return r
		}

		m.logger.log("input '%s' not found in exact match map", *data)
		return nil
	}

	return handlePossiblyMissingInput(m.input, data, matchFunc, m.logger)
}

func safePrint(s *string) string {
	if s == nil {
		return "<nil>"
	}

	return *s
}

func handlePossiblyMissingInput(dataInput types.DataInput, data types.MatchingData, matchFunc func(*string) *types.OnMatch, logger *tracingLogger) (types.Result, error) {
	input, err := dataInput.Input(data)
	if err != nil {
		return types.Result{}, err
	}

	logger.log("handling input '%v'", safePrint(input.Data))

	switch input.Availability {
	case types.AllDataAvailable:
		logger.log("data is available, performing match")
		return types.Result{
			MatchResult:  matchFunc(input.Data),
			NeedMoreData: false,
		}, nil
	case types.MoreDataMightBeAvailable:
		logger.log("more data might be available, performing eager match")
		result := matchFunc(input.Data)
		return types.Result{
			MatchResult:  result,
			NeedMoreData: result == nil,
		}, nil
	case types.NotAvailable:
		logger.log("data not available, bailing")
		return types.Result{MatchResult: nil, NeedMoreData: true}, nil
	default:
		return types.Result{}, errors.New("unexpected return value")
	}

}

func createMatchTree(matcherTree *pbmatcher.Matcher_MatcherTree, logger *tracingLogger) (*exactMatchTree, error) {
	i, err := registry.ResolveInputExtension(matcherTree.Input)
	if err != nil {
		return nil, err
	}

	switch m := matcherTree.TreeType.(type) {
	case *pbmatcher.Matcher_MatcherTree_ExactMatchMap:
		tree := &exactMatchTree{
			logger: logger,
			tree:   map[string]*types.OnMatch{},
			input:  i,
		}
		for k, v := range m.ExactMatchMap.Map {
			onMatch, err := createOnMatch(v, logger)
			if err != nil {
				return nil, err
			}
			tree.tree[k] = onMatch
		}
		return tree, nil
	default:
		return nil, errors.New("not implemented: tree type")
	}
}

// Internal logger that helps debugging by providing indention during matching to provide some visual guidance to the current nesting level.
// TODO(snowp): Make this optional, not ideal if this ends up being used outside of testing scenarios.
type tracingLogger struct {
	indent int
}

func (t *tracingLogger) push() {
	t.indent++
}

func (t *tracingLogger) pop() {
	t.indent--
}

func (t *tracingLogger) log(formatString string, args ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s%s\n", strings.Repeat("  ", t.indent), formatString), args...)
}
