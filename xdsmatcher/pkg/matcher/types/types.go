package types

type DataInputAvailability int

type Matcher interface {
	Match(MatchingData) (Result, error)
}

type OnMatch struct {
	Matcher Matcher
	Action  Action
}
type Action interface{}
type MatchingData interface{}
type Result struct {
	MatchResult  *OnMatch
	NeedMoreData bool
}

const (
	NotAvailable             DataInputAvailability = iota
	MoreDataMightBeAvailable DataInputAvailability = iota
	AllDataAvailable         DataInputAvailability = iota
)

type DataInputResult struct {
	Availability DataInputAvailability
	Data         *string
}

type DataInput interface {
	Input(MatchingData) (DataInputResult, error)
}
