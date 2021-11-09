package matcher

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	pbmatcher "github.com/envoyproxy/go-control-plane/envoy/config/common/matcher/v3"
	pbtypematcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	"github.com/lyft/xdsmatcher/pkg/matcher/registry"
	"github.com/lyft/xdsmatcher/pkg/matcher/types"
)

type MatcherTree struct {
	onNoMatch *types.OnMatch
	matchRoot types.Matcher
}

func (m MatcherTree) Match(data types.MatchingData) (types.Result, error) {
	result, err := m.matchRoot.Match(data)
	if err != nil {
		return types.Result{}, err
	}

	if result.MatchResult != nil {
		if result.MatchResult.Matcher != nil {
			return result.MatchResult.Matcher.Match(data)
		}

		return result, nil
	}

	if result.NeedMoreData {
		return result, nil
	}

	return types.Result{
		MatchResult:  m.onNoMatch,
		NeedMoreData: false,
	}, nil
}

func Create(matcher *pbmatcher.Matcher) (*MatcherTree, error) {
	on_no_match, err := createOnMatch(matcher.OnNoMatch)
	if err != nil {
		return nil, err
	}

	switch m := matcher.MatcherType.(type) {
	case *pbmatcher.Matcher_MatcherList_:
		matcher, err := createMatcherList(m.MatcherList)
		if err != nil {
			return nil, err
		}
		return &MatcherTree{
			onNoMatch: on_no_match,
			matchRoot: matcher,
		}, nil
	case *pbmatcher.Matcher_MatcherTree_:
		matcher, err := createMatchTree(m.MatcherTree)
		if err != nil {
			return nil, err
		}

		return &MatcherTree{
			onNoMatch: on_no_match,
			matchRoot: matcher,
		}, nil
	default:
		return nil, errors.New("not implemented: matcher tree")

	}
}

func createOnMatch(onMatch *pbmatcher.Matcher_OnMatch) (*types.OnMatch, error) {
	// TODO move this to call sites, this is part of validation
	if onMatch == nil {
		return nil, nil
	}
	switch om := onMatch.OnMatch.(type) {
	case *pbmatcher.Matcher_OnMatch_Matcher:
		matcher, err := Create(om.Matcher)
		if err != nil {
			return nil, err
		}

		return &types.OnMatch{
			Matcher: matcher,
		}, nil
	case *pbmatcher.Matcher_OnMatch_Action:
		return &types.OnMatch{Action: nil}, nil
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

type InputMatcher interface {
	match(value *string) (matcherResult, error)
}

type simpleMatcher interface {
	match(value *string) bool
}

type singlePredicate struct {
	input   types.DataInput
	matcher simpleMatcher
}

func (s singlePredicate) match(data types.MatchingData) (matcherResult, error) {
	matchFunc := func(input *string) *types.OnMatch {
		log.Printf("attempting to match %v", input)
		if s.matcher.match(input) {
			return &types.OnMatch{}
		}

		return nil
	}

	result, err := handlePossiblyMissingInput(s.input, data, matchFunc)
	if err != nil {
		return matcherResult{}, err
	}

	return matcherResult{
		match:        result.MatchResult != nil,
		needMoreData: result.NeedMoreData,
	}, nil

}

type conjunctionPredicate struct {
	predicates []predicateFunc
}

func (a conjunctionPredicate) match(data types.MatchingData) (matcherResult, error) {
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
}

func (a disjunctionPredicate) match(data types.MatchingData) (matcherResult, error) {
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

func createSinglePredicate(predicate *pbmatcher.Matcher_MatcherList_Predicate_SinglePredicate) (*singlePredicate, error) {
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
			input:   input,
			matcher: matcher,
		}, nil
	default:
		return nil, errors.New("not implemented: custom match")
	}
}

func createValueMatcher(valueMatcher *pbtypematcher.StringMatcher) (simpleMatcher, error) {
	switch m := valueMatcher.MatchPattern.(type) {
	case *pbtypematcher.StringMatcher_SafeRegex:
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
}

func (m matcherList) Match(data types.MatchingData) (types.Result, error) {
	for _, m := range m.matchEntries {
		log.Printf("predicate type %T", m.predicate)
		result, err := m.predicate.match(data)
		if err != nil {
			return types.Result{}, err
		}
		log.Printf("list matcher result: %v", result)
		if result.needMoreData {
			return types.Result{
				MatchResult:  nil,
				NeedMoreData: true,
			}, nil
		}
		if result.match {
			return types.Result{
				MatchResult:  m.onMatch,
				NeedMoreData: false,
			}, nil
		}
	}

	return types.Result{
		MatchResult:  nil,
		NeedMoreData: false,
	}, nil
}

func createPredicate(predicate *pbmatcher.Matcher_MatcherList_Predicate) (predicateFunc, error) {
	switch predicate := predicate.MatchType.(type) {
	case *pbmatcher.Matcher_MatcherList_Predicate_AndMatcher:
		predicates := []predicateFunc{}
		for _, p := range predicate.AndMatcher.Predicate {
			pp, err := createPredicate(p)
			if err != nil {
				return nil, err
			}
			predicates = append(predicates, pp)
		}

		return conjunctionPredicate{
			predicates: predicates,
		}, nil
	case *pbmatcher.Matcher_MatcherList_Predicate_OrMatcher:
		predicates := []predicateFunc{}
		for _, p := range predicate.OrMatcher.Predicate {
			pp, err := createPredicate(p)
			if err != nil {
				return nil, err
			}
			predicates = append(predicates, pp)
		}

		return disjunctionPredicate{
			predicates: predicates,
		}, nil
	case *pbmatcher.Matcher_MatcherList_Predicate_SinglePredicate_:
		return createSinglePredicate(predicate.SinglePredicate)
	default:
		return nil, fmt.Errorf("unexpected predicate type %T", predicate)
	}
}

func createMatcherList(list *pbmatcher.Matcher_MatcherList) (*matcherList, error) {
	ml := &matcherList{}
	for _, m := range list.Matchers {
		onMatch, err := createOnMatch(m.OnMatch)
		if err != nil {
			return nil, err
		}

		entry := matchEntry{onMatch: onMatch}
		entry.predicate, err = createPredicate(m.Predicate)
		if err != nil {
			return nil, err
		}

		ml.matchEntries = append(ml.matchEntries, entry)
	}

	return ml, nil
}

type exactMatchTree struct {
	tree  map[string]*types.OnMatch
	input types.DataInput
}

func (m exactMatchTree) Match(data types.MatchingData) (types.Result, error) {
	matchFunc := func(data *string) *types.OnMatch {
		if data == nil {
			return nil
		}
		if m, ok := m.tree[*data]; ok {
			return m
		}

		return nil
	}

	return handlePossiblyMissingInput(m.input, data, matchFunc)
}

func handlePossiblyMissingInput(dataInput types.DataInput, data types.MatchingData, matchFunc func(*string) *types.OnMatch) (types.Result, error) {
	input, err := dataInput.Input(data)
	if err != nil {
		return types.Result{}, err
	}

	log.Printf("handling input %v", *input.Data)

	switch input.Availability {
	case types.AllDataAvailable:
		return types.Result{
			MatchResult:  matchFunc(input.Data),
			NeedMoreData: false,
		}, nil
	case types.MoreDataMightBeAvailable:
		result := matchFunc(input.Data)
		return types.Result{
			MatchResult:  result,
			NeedMoreData: result == nil,
		}, nil
	case types.NotAvailable:
		return types.Result{MatchResult: nil, NeedMoreData: true}, nil
	default:
		return types.Result{}, errors.New("unexpected return value")
	}

}

func createMatchTree(matcherTree *pbmatcher.Matcher_MatcherTree) (*exactMatchTree, error) {
	i, err := registry.ResolveInputExtension(matcherTree.Input)
	if err != nil {
		return nil, err
	}

	switch m := matcherTree.TreeType.(type) {
	case *pbmatcher.Matcher_MatcherTree_ExactMatchMap:
		tree := &exactMatchTree{
			tree:  map[string]*types.OnMatch{},
			input: i,
		}
		for k, v := range m.ExactMatchMap.Map {
			onMatch, err := createOnMatch(v)
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
