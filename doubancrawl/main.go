package main

import (
	"doubancrawl/engine"
	"doubancrawl/parse"
	"doubancrawl/persist"
	"doubancrawl/scheduler"
)

func main() {
	e := engine.Concurrentengine{
		Scheduler: &scheduler.QueueScheduler{},
		Workcounter: 100,
		ItemChan:persist.ItemSave(),
	}
	e.Run(engine.Request{
		URL:       "https://book.douban.com/tag/",
		ParseFunc: parse.ParseTag,
	})
}

/* Single Engine Testing passed
func main() {
	e := engine.SingleEngine{}
	e.Run(engine.Request{
		URL:       "https://book.douban.com/tag/",
		ParseFunc: parse.ParseTag,
	})
}
*/

