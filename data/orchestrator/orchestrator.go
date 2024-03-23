package orchestrator

type orcherstrator struct {
	pipelines []pipeline
}

func newOrchestrator() *orcherstrator {
	return &orcherstrator{
		pipelines: []pipeline{},
	}
}

func (o *orcherstrator) deploy() {
	for _, p := range o.pipelines {
		p
	}
}
