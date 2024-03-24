package orchestrator

import (
	"toukyou/data/pipeline"

	log "github.com/sirupsen/logrus"
)

type Orchestrator struct {
	Pipelines []pipeline.Pipeline
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		Pipelines: []pipeline.Pipeline{},
	}
}

func (o *Orchestrator) AddPipeline(p pipeline.Pipeline) {
	o.Pipelines = append(o.Pipelines, p)
}

func (o *Orchestrator) Deploy() {
	log.Info("Deploying pipelines")
	for index, p := range o.Pipelines {
		log.Infof("Deploying pipeline: %d", index+1)
		p.Run()
	}
}
