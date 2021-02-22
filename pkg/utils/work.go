package utils

import (
	"sync"
	"time"
)

func Work(workNum, workerNum int, do func()) time.Duration {
	start := time.Now()

	distributor, done := make(chan bool), &sync.WaitGroup{}

	for i := 0; i < workerNum; i++ {
		done.Add(1)
		go worker(do, distributor, done)
	}

	go func() {
		for i := 0; i < workNum; i++ {
			distributor <- true
		}
		close(distributor)
	}()

	done.Wait()

	return time.Since(start)
}

func worker(do func(), distributor chan bool, done *sync.WaitGroup) {
	defer done.Done()

loop:
	for {
		_, ok := <-distributor
		if !ok {
			break loop
		}
		do()
	}
}
