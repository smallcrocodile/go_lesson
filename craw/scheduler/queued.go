package scheduler

import (
	"imooc/craw/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	wokerChan   chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WokerReady(w chan engine.Request) {
	s.wokerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.wokerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWoker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWoker = workerQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.wokerChan:
				workerQ = append(workerQ, w)
			case activeWoker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}

	}()

}
