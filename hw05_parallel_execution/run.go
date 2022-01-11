package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksCh := make(chan Task, n)
	wg := sync.WaitGroup{}

	go func() {
		for _, task := range tasks {
			tasksCh <- task
		}
		close(tasksCh)
	}()

	for task := range tasksCh {
		wg.Add(1)
		task := task
		go func() {
			defer wg.Done()
			task()
		}()
	}

	wg.Wait()

	return nil
}
