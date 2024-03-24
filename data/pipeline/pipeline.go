package pipeline

import (
	"sync"
	"time"
)

type Stage struct {
	Stage func()
}

type Pipeline struct {
	Stages []Stage
}

func NewPipeline() *Pipeline {
	return &Pipeline{
		Stages: []Stage{},
	}
}

func (p *Pipeline) AddStage(s Stage) {
	p.Stages = append(p.Stages, s)
}

func (p *Pipeline) Run() {
	ticker := time.NewTicker(5 * time.Second)

	runStages := func() {
		for _, s := range p.Stages {
			s.Stage()
		}
	}

	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				runStages()
			case <-done:
				return
			}
		}
	}()

	wg.Wait()
}
