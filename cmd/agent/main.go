package main

import (
	"log"

	"github.com/nikpet14321/Arithmetic-expression-calculation-service2/internal/application"
)

func main() {
	agent := application.NewAgent()
	log.Println("Starting Agent...")
	agent.Run()
}
