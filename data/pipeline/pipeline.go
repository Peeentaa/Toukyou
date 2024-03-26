package pipeline

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type Stage struct {
	StageFunc func()
}

type Pipeline struct {
	Stages []Stage
}

func NewPipeline() *Pipeline {
	log.Info("Initializing new pipeline")
	return &Pipeline{
		Stages: []Stage{},
	}
}

func (p *Pipeline) AddStage(s Stage) {
	p.Stages = append(p.Stages, s)
}

func (p *Pipeline) Run() error {
	log.Info("Starting pipeline execution")
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	runStages := func() {
		for index, s := range p.Stages {
			log.Infof("Executing stage: %d", index+1)
			s.StageFunc() // Execute the stage function
		}
		return
	}

	runStages()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				runStages()
			}
		}
	}()

	wg.Wait()

	return nil
}
