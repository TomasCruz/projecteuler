package projecteuler

func makeRecursiveOneResult(f workerFunctionType) OneResultSolver {
	r := recursiveOneResult{
		solutionFound: false,
		workerFunc:    f,
	}

	return &r
}

func (ro *recursiveOneResult) OneResultSolve(jobs <-chan interface{}) (interface{}, error) {
	for j := range jobs {
		if ro.solutionFound {
			break
		}

		r, err := ro.workerFunc(j)
		if err != nil {
			return r, err
		}

		if r.Ok() {
			ro.solutionFound = true
			var result interface{} = r
			return result, nil
		}
	}

	return nil, nil
}
