package scheduler

import "imooc/craw/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	//send request down to woker chan
	go func() {
		s.workerChan <- request
	}()

}
