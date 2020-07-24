package projecteuler

import "sync"

type (
	// OneResultSolver has OneResultSolve(jobs <-chan interface{}) (interface{}, error)
	OneResultSolver interface {
		OneResultSolve(jobs <-chan interface{}) (interface{}, error)
	}

	workerFunctionType func(input interface{}) (Oker, error)

	workerPoolOneResult struct {
		solutionFound bool
		workerFunc    workerFunctionType
		m             sync.Mutex
	}

	recursiveOneResult struct {
		solutionFound bool
		workerFunc    workerFunctionType
	}
)

const (
	// RecursiveKind is self-explanatory, you stupid Lint
	RecursiveKind = iota
	// ConcurrentKind is self-explanatory, you stupid Lint
	ConcurrentKind
)

// OneResultSolverFactory produces OneResultSolver appropriate per kind
func OneResultSolverFactory(kind int, f workerFunctionType) OneResultSolver {
	switch kind {
	case ConcurrentKind:
		return makeWorkerPoolOneResult(f)
	case RecursiveKind:
		return makeRecursiveOneResult(f)
	default:
		return nil
	}
}
