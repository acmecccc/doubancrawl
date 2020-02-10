package parse

import (
	"regexp"

	"doubancrawl/engine"
)

const regexpStr = `<td><a href="([^"]+)">([^<]+)</a>` //<td><a href="/tag/科普">科普</a>
//Parsecontent : used to filter out the content we need by regular expressions
func ParseTag(content []byte) engine.ParseResult {
	re := regexp.MustCompile(regexpStr)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, m[2])
		result.Requests = append(result.Requests, engine.Request{
			URL:       "https://book.douban.com" + string(m[1]),
			ParseFunc: ParseBookList,
		})
	}
	return result
}
