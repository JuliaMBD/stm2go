package stm2go

import (
	"sync"
	"time"
)

type Tasktype int

const (
	Done Tasktype = iota
	Continue
)

type TestEnv struct {
	task      []func()
	wg        *sync.WaitGroup
	timescale time.Duration
}

func NewTestEnv() *TestEnv {
	return &TestEnv{
		task:      []func(){},
		wg:        &sync.WaitGroup{},
		timescale: 1,
	}
}

func (t *TestEnv) Add(tasktype Tasktype, x func()) {
	switch tasktype {
	case Done:
		t.wg.Add(1)
		t.task = append(t.task, func() {
			defer t.wg.Done()
			x()
		})
	case Continue:
		t.task = append(t.task, x)
	}
}

func (t *TestEnv) Set(i time.Duration) {
	t.timescale = i
}

func (t *TestEnv) Go() {
	for _, x := range t.task {
		go x()
	}
	t.wg.Wait()
}

func (t *TestEnv) Sleep(h time.Duration) {
	time.Sleep(h * t.timescale)
}

func (t *TestEnv) Tick(i time.Duration) <-chan time.Time {
	return time.Tick(i * t.timescale)
}

func (t *TestEnv) After(i time.Duration) <-chan time.Time {
	return time.After(i * t.timescale)
}
