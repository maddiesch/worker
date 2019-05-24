package worker

import (
	"sync/atomic"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

type testJob struct {
	count *int32
}

func (j *testJob) Run() {
	atomic.AddInt32(j.count, 1)
}

func TestNewPool(t *testing.T) {
	pool := NewPool(4, 64)

	var rawCount int32
	count := &rawCount

	for i := 0; i < 16; i++ {
		pool.Dispatch(&testJob{count: count})
	}

	t.Log(pool)

	pool.Wait()

	assert.Equal(t, true, rawCount == 16)
}
