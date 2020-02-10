package engine

//Request : Task struct
type Request struct {
	URL       string
	ParseFunc func([]byte) ParseResult
}

//ParseResult : result
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//NilParse : Nil Func
func NilParse([]byte) ParseResult {
	return ParseResult{}
}
