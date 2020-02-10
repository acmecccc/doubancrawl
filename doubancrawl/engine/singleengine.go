package engine

import (
	"fmt"
	"log"

	"doubancrawl/fetcher"
)

//SingleEngine :单任务引擎结构体
type SingleEngine struct {
}

//Run : 单任务引擎函数
func (s SingleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, e := range seeds {
		requests = append(requests, e)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching: %s", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetch error: %s", err)
		}
		parseresult := r.ParseFunc(body)
		requests = append(requests, parseresult.Requests...) //将元素打散逐个使用
		for _, item := range parseresult.Items {
			fmt.Printf("Got Item: %s \n", item)
		}
	}

}
