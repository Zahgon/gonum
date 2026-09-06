package optimize

type localOptimizer struct{}

func (l localOptimizer) run(method localMethod, gradThresh float64, operation chan<- Task, result <-chan Task, tasks []Task) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (localOptimizer) initialOperation(task Task, n needser) Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (l localOptimizer) initialLocation(operation chan<- Task, result <-chan Task, task Task, needs needser) Task {
	_ = "STUB: not implemented"
	return *new(Task)
}

func (l localOptimizer) checkStartingLocation(task Task, gradThresh float64) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (localOptimizer) checkGradientConvergence(gradient []float64, gradThresh float64) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func (localOptimizer) finish(operation chan<- Task, result <-chan Task) {
	_ = "STUB: not implemented"
	return
}

func (l localOptimizer) finishMethodDone(operation chan<- Task, result <-chan Task, task Task) {
	_ = "STUB: not implemented"
	return
}
