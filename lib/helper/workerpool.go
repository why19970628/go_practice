package helper

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
	for i := 0; i < workerNum; i++ {
		go func() {
			for w := range p.work {
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

func RunMultiFunc(funcs []func()) {
	num := len(funcs)
	if num <= 0 {
		return
	}
	var wg sync.WaitGroup
	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(index int) {
			funcs[index]()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
