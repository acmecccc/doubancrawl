package parse

import (
	"doubancrawl/engine"
	"doubancrawl/model"
	"regexp"
	"strconv"
)

var authorRe = regexp.MustCompile(`<span class="pl">作者:</span>&nbsp;[\d\D]*?<a.*?>[\d\D*?].*?]([^<]+)</a>[\d\D]*?<br>`)
var publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span>([^<]+)<br/>`)
var bookpagesRe = regexp.MustCompile(`<span class="pl">页数:</span> (.*)<br/>`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>(.*)元<br/>`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> (.*) </strong>`)
var infoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>(.*)</p>`)

// ParseBookDetail : Parse book details
func ParseBookDetail(contents []byte) engine.ParseResult {
	//fmt.Printf("%s", contents)
	bookdetail := model.Bookdetail{}
	bookdetail.Author = ExactString(contents, authorRe)
	bookdetail.Publisher = ExactString(contents, publisherRe)
	pages, err := strconv.Atoi(ExactString(contents, bookpagesRe))
	if err == nil {
		bookdetail.Bookpages = pages
	}
	bookdetail.Price = ExactString(contents, priceRe)
	bookdetail.Score = ExactString(contents, scoreRe)
	bookdetail.Info = ExactString(contents, infoRe)
	result := engine.ParseResult{
		Items: []interface{}{bookdetail},
	}
	return result
}

//ExactString : to exact all the regular expression strings for the BookDetail struct
func ExactString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return " "

}
