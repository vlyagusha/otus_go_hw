package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksCh := make(chan Task)
	wg := sync.WaitGroup{}

	go func() {
		for _, task := range tasks {
			tasksCh <- task
		}
		close(tasksCh)
	}()

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasksCh {
				task()
			}
		}()
	}

	wg.Wait()

	return nil
}
