package main

import (
	"log"

	"github.com/nikpet14321/Arithmetic-expression-calculation-service2/internal/application"
)

func main() {
	app := application.NewOrchestrator()
	log.Println("Starting Orchestrator on port", app.Config.Addr)
	if err := app.RunServer(); err != nil {
		log.Fatal(err)
	}
}
