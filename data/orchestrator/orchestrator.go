package orchestrator

import (
	"toukyou/data/pipeline"
)

type Orchestrator struct {
	pipelines []pipeline.Pipeline
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		pipelines: []pipeline.Pipeline{},
	}
}

func (o *Orchestrator) AddPipeline(p pipeline.Pipeline) {
	o.pipelines = append(o.pipelines, p)
}

func (o *Orchestrator) Deploy() {
	for _, p := range o.pipelines {
		p.Run()
	}
}
