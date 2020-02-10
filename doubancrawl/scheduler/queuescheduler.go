package scheduler

import (
	"fmt"

	"doubancrawl/engine"
)

//QueueScheduler :
type QueueScheduler struct {
	RequestChan chan engine.Request
	Workerchanchan  chan chan engine.Request
}

func (q *QueueScheduler) Workchannel() chan engine.Request {
	return make(chan engine.Request)
}

//Submit :将request提交进入通道
func (q *QueueScheduler) Submit(r engine.Request) {
	q.RequestChan <- r
}

//WorkReady : 把·request通道放到另一个通道中，构成一个通道的通道
func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.Workerchanchan <- w

}

//Run :
func (q *QueueScheduler) Run() {
	q.Workerchanchan = make(chan chan engine.Request)
	q.RequestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeRequest = requestQ[0]
				activeWork = workQ[0]
			}
			select {
			case r := <-q.RequestChan:
				fmt.Printf("Request monitoring: %v",r)
				requestQ = append(requestQ, r)
			case w := <-q.Workerchanchan:
				//fmt.Printf("Request monitoring: %v",w)
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}

		}

	}()

}
