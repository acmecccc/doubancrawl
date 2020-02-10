package parse

import (
	"doubancrawl/engine"
	"regexp"
)

//ParseBookList : Parse the book list deatails
func ParseBookList(contents []byte) engine.ParseResult {
	const BooklistRe = `<a href="([^"]+)" title="([^"]+)"`
	re := regexp.MustCompile(BooklistRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			URL:       string(m[1]),
			ParseFunc: ParseBookDetail,
		})
	}
	return result
}
