package main

import (
	"github.com/fatih/color"
	"log"
	"toukyou/analysis"
	"toukyou/data/db"
	"toukyou/data/fetcher"
	"toukyou/data/orchestrator"
	"toukyou/data/pipeline"
)

func main() {
	color.New(color.FgRed, color.Bold).Println("\n___________           __                         \n\\__    ___/___  __ __|  | _____.__. ____  __ __  \n  |    | /  _ \\|  |  \\  |/ <   |  |/  _ \\|  |  \\ \n  |    |(  <_> )  |  /    < \\___  (  <_> )  |  / \n  |____| \\____/|____/|__|_ \\/ ____|\\____/|____/  \n                          \\/\\/                   \n")

	// Initialize and open database
	if _, err := db.OpenDatabase(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Create orchestrator and pipeline
	o := orchestrator.NewOrchestrator()
	p := pipeline.NewPipeline()

	sI := pipeline.Stage{StageFunc: func() {
		fetcher.GetStockData()
	}}

	s1 := pipeline.Stage{StageFunc: func() {
		err := analysis.RunAnalysis("dcf")
		if err != nil {
			return
		}
	}}

	p.AddStage(sI)
	p.AddStage(s1)

	// Add pipeline to orchestrator
	o.AddPipeline(*p)

	// Deploy orchestrator
	o.Deploy()
}
