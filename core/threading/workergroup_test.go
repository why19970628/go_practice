package threading

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/lang"
)

func TestWorkerGroup(t *testing.T) {
	m := make(map[string]lang.PlaceholderType)
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())
	group := NewWorkerGroup(func() {
		lock.Lock()
		m[fmt.Sprint(RoutineId())] = lang.Placeholder
		lock.Unlock()
		wg.Done()
	}, runtime.NumCPU())
	go group.Start()
	wg.Wait()
	assert.Equal(t, runtime.NumCPU(), len(m))
}

func TestWorkerGroups(t *testing.T) {
	m := make(map[string]lang.PlaceholderType)
	var (
		task = rand.Intn(10)
		jobs []func()
		lock sync.Mutex
	)
	for i := 0; i < task; i++ {
		jobs = append(jobs, func() {
			lock.Lock()
			m[fmt.Sprint(RoutineId())] = lang.Placeholder
			lock.Unlock()
		})
	}
	group := NewWorkersGroups(jobs)
	group.Start()
	//assert.Equal(t, len(jobs), len(m))
}
