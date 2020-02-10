package engine

import (
	"fmt"
	"log"

	"doubancrawl/fetcher"
)

//Concurrentengine : 并发引擎结构体
type Concurrentengine struct {
	Scheduler   Scheduler
	Workcounter int
	ItemChan chan interface{}
}

//Scheduler : 调度器接口
type Scheduler interface {
	Submit(Request)
	//ConfigureWorkChan(chan Request)
	Run()
	WorkReady(chan Request)
	Workchannel() (chan Request)
}

//Run ：并发引擎函数
func (e *Concurrentengine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.Workcounter; i++ {
		CreateWrok(e.Scheduler.Workchannel(),e.Scheduler,out)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(){e.ItemChan <-item}()
		}

	}
}

//CreateWrok :
func CreateWrok(in chan Request,s Scheduler, out chan ParseResult) {

	go func() {
		for {
			s.WorkReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}() //这样的形式开启go协程必须加（）以运行
}

//worker :对Request进行实质性处理的函数
func Worker(r Request) (ParseResult, error) {
	fmt.Printf("Fetching URL: %s \n", r.URL)
	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		log.Printf("Fetch Error: %s", err)
		return ParseResult{}, nil
	}
	return r.ParseFunc(body), nil
}
