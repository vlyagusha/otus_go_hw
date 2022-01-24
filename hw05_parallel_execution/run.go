package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		m = len(tasks)
	}
	errorsCount := int32(m)

	tasksCh := make(chan Task)
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksCh {
				err := task()
				if err != nil {
					atomic.AddInt32(&errorsCount, -1)
				}
			}
		}()
	}

	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) <= 0 {
			break
		}
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()

	if atomic.LoadInt32(&errorsCount) <= 0 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
