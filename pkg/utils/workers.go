package utils

type Worker[T interface{}] struct {
	Job chan func() T
	Res chan T
}

func (w Worker[T]) CreateWorkers(id int, job <-chan func() T, res chan<- T) {
	for j := range job {

		doJob := j
		go func() {
			result := doJob()
			res <- result
		}()
	}
}

func (w Worker[T]) RunWorkers(doFunc func() T, numJobs int) {
	for i := 0; i < numJobs; i++ {
		w.Job <- doFunc
	}
}


