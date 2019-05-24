package worker

import (
	"sync"
)

type worker struct {
	ID   int
	Jobs chan *jobRunner
}

// Pool executes the
type Pool struct {
	jobs    chan *jobRunner
	depth   int
	waiter  sync.WaitGroup
	Workers []*worker
}

// Job is performed on a worker
type Job interface {
	Run()
}

type jobRunner struct {
	job  Job
	done func()
}

// Run begins executing jobs
func (w *worker) Run() {
	for r := range w.Jobs {
		r.job.Run()
		r.done()
	}
}

// Dispatch pushes a job onto an available pool
func (p *Pool) Dispatch(j Job) {
	p.waiter.Add(1)
	p.jobs <- &jobRunner{
		job:  j,
		done: p.waiter.Done,
	}
}

// Wait blocks and waits for all jobs to be completed
func (p *Pool) Wait() {
	p.waiter.Wait()
}

// NewPool returns a new pool with a count
//
// The pool allows a depth between 64 & 256
func NewPool(count, depth int) *Pool {
	dep := min(max(depth, 64), 256)

	pool := &Pool{
		depth:   dep,
		jobs:    make(chan *jobRunner, depth),
		Workers: make([]*worker, 0, count),
	}
	for i := 1; i <= count; i++ {
		wkr := &worker{
			ID:   i,
			Jobs: pool.jobs,
		}

		go wkr.Run()

		pool.Workers = append(pool.Workers, wkr)
	}
	return pool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
