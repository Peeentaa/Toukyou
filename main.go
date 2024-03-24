package main

import (
	"fmt"
	"log"
	"toukyou/data/db"
	"toukyou/data/orchestrator"
	"toukyou/data/pipeline"
)

func main() {

	_, err := db.OpenDatabase()
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	o := orchestrator.NewOrchestrator()
	p := pipeline.NewPipeline()
	s := pipeline.Stage{Stage: func() {
		fmt.Printf("hello")
	}}

	p.AddStage(s)
	o.AddPipeline(*p)
	o.Deploy()
}
