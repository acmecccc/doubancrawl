package scheduler

import "doubancrawl/engine"


//SimpleScheduler :引入通道
type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) Run() {
	s.WorkChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	return
}

func (s *SimpleScheduler) Workchannel() chan engine.Request {
	return s.WorkChan
}

//Submit :--开启协程，将Request实例传入WorkChan通道
func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}