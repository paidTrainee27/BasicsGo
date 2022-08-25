package goroutine

import (
	"fmt"
	"sync"
	"time"
)

// add to another package

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	semC chan struct{}
}

func NewSemaphore(maxConcurrency int) Semaphore {
	return &semaphore{
		semC: make(chan struct{}, maxConcurrency),
	}
}
func (s *semaphore) Acquire() {
	s.semC <- struct{}{}
}
func (s *semaphore) Release() {
	<-s.semC
}

func SemaphoreTask() {
	sem := NewSemaphore(3)
	wg := &sync.WaitGroup{}
	totProcess := 10
	for i := 1; i <= totProcess; i++ {
		sem.Acquire()
		wg.Add(1)
		go func(wgr *sync.WaitGroup, v int) {
			defer func() {
				sem.Release()
				wgr.Done()
			}()
			longRunningProcess(v)
		}(wg, i)
	}
	wg.Wait()
}
func longRunningProcess(taskID int) {
	fmt.Println(
		time.Now().Format("15:04:05"),
		"Running task with ID",
		taskID)
	time.Sleep(1 * time.Second)
}
