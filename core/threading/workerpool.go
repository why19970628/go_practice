package threading

import "sync"

type Pool struct {
	work chan func()
	wg   sync.WaitGroup
}

func NewWorker(workerNum int) *Pool {
	p := Pool{
		work: make(chan func()),
	}
	p.wg.Add(workerNum)
	//group := NewRoutineGroup()
	for i := 0; i < workerNum; i++ {
		go func() {
			for w := range p.work {
				//group.RunSafe(w)
				w()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Add(w func()) {
	p.work <- w
}

func (p *Pool) ShutDown() {
	close(p.work)
	p.wg.Wait()
}
