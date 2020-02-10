package fetcher

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var ratelimit = time.Tick(2000 * time.Millisecond)

//Fetch : get the full content of a web page
func Fetch(weburl string) ([]byte, error) {
	<-ratelimit
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:57096")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}
	req, err := http.NewRequest("GET", weburl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//req.Close = true
	req.Header.Set("User-Agent", "Mozilla/5.0 ")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyreader := bufio.NewReader(resp.Body)
	e := MyDetermineEncoding(bodyreader)
	utf8Reader := transform.NewReader(bodyreader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)

}

//MyDetermineEncoding : determine the encoding type of web pages and transform it to `utf8`
func MyDetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetching error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
