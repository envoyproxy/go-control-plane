package types

type DataInputAvailability int

// Matcher describes a type that can produce a match result from a set of matching data.
type Matcher interface {
	Match(MatchingData) (Result, error)
}

// OnMatch is a node in the match tree, either describing an action (leaf node) or
// the start of a subtree (internal node).
type OnMatch struct {
	Matcher Matcher
	Action  Action
}

// Action describes an opaque action that is the final result of a match. Implementations would likely
// need to cast this to a more appropriate type.
type Action interface{}

// MatchingData describes an opaque set of input data.
type MatchingData interface{}

// Result describes the result of evaluating the match tree.
type Result struct {
	// MatchResult is the final result, if NeedMoreData is false. This can be nil if the match tree completed
	// but no action was resolved.
	MatchResult  *OnMatch
	// NeedMoreData specified whether the match tree failed to resolve due to input data not being available yet.
	// This can imply that as more data is made available, a match might be found.
	NeedMoreData bool
}

const (
	// NotAvailable indicates that the data input is not available.
	NotAvailable             DataInputAvailability = iota
	// MoreDataMightBeAvailable indicates that there might be more data available.
	MoreDataMightBeAvailable DataInputAvailability = iota
	// AllDataAvailable indicates that all data is present, no more data will be added.
	AllDataAvailable         DataInputAvailability = iota
)

// DataInputResult describes the result of evaluating a DataInput.
type DataInputResult struct {
	// Availability describes the kind of data availability the associated data has.
	Availability DataInputAvailability
	// Data is the resulting data. This might be nil if the data is not available or if the
	// backing data is available but the requested value does not.
	Data         *string
}

// DataInput describes a type that can extract an input value from the MatchingData.
type DataInput interface {
	Input(MatchingData) (DataInputResult, error)
}
