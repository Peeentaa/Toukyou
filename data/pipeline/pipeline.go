package pipeline

type stage struct {
	stage func()
}

type pipeline struct {
	stages []stage
}

func newPipeline() *pipeline {
	return &pipeline{
		stages: []stage{},
	}
}

func (p *pipeline) addStage(s stage) {
	p.stages = append(p.stages, s)
}

func (p *pipeline) run() {

	for _, s := range p.stages {
		s.stage()
	}
}
