package projecteuler

import (
	"runtime"
)

func makeWorkerPoolOneResult(f workerFunctionType) OneResultSolver {
	wp := workerPoolOneResult{
		solutionFound: false,
		workerFunc:    f,
	}

	return &wp
}

func (wp *workerPoolOneResult) worker(
	jobs <-chan interface{},
	done chan<- bool,
	result chan<- Oker,
	errCh chan<- error) {

	for j := range jobs {
		if wp.solutionFound {
			break
		}

		r, err := wp.workerFunc(j)
		if err != nil {
			errCh <- err
			return
		}

		if r.Ok() {
			wp.m.Lock()
			if !wp.solutionFound {
				wp.solutionFound = true
				result <- r
			}
			wp.m.Unlock()

			break
		}
	}

	done <- true
}

func (wp *workerPoolOneResult) OneResultSolve(jobs <-chan interface{}) (interface{}, error) {
	numWorkers := runtime.NumCPU()

	done := make(chan bool, numWorkers)
	result := make(chan Oker)
	errCh := make(chan error)
	resFound := make(chan bool)

	var res Oker
	go func() {
		res = <-result
		resFound <- true
	}()

	for w := 1; w <= numWorkers; w++ {
		go wp.worker(jobs, done, result, errCh)
	}

	var err error
	for i := 1; i <= numWorkers; i++ {
		select {
		case <-done:
		case err = <-errCh:
			break
		}
	}

	<-resFound
	return res, err
}
