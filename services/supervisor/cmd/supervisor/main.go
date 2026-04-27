package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/kekcoke/minelink/supervisor/internal/mqtt"
	"github.com/kekcoke/minelink/supervisor/internal/state"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 1. Initialize Redis Connection (Edge State)
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "localhost:6379" // Default for local dev
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Edge Redis: %v", err)
	}
	log.Println("Connected to Edge Redis state store.")

	supervisorState := state.NewSupervisorState(rdb)

	// 2. Initialize MQTT Orchestrator (C2 Link)
	brokerURL := os.Getenv("MQTT_BROKER_URL")
	if brokerURL == "" {
		brokerURL = "tcp://localhost:1883" // Default for local dev
	}
	
	supervisorID := os.Getenv("SUPERVISOR_ID")
	if supervisorID == "" {
		supervisorID = "pit-a-alpha"
	}
	clientID := "supervisor-" + supervisorID

	orchestrator, err := mqtt.NewOrchestrator(brokerURL, clientID, supervisorState)
	if err != nil {
		log.Fatalf("Failed to initialize MQTT Orchestrator: %v", err)
	}

	if err := orchestrator.Start(ctx); err != nil {
		log.Fatalf("Failed to start Orchestrator: %v", err)
	}

	// 3. Graceful Shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	log.Printf("Supervisor Agent [%s] is running. Press Ctrl+C to exit.", supervisorID)
	<-sigChan

	log.Println("Shutting down Supervisor Agent...")
	// Cleanup logic here (close Redis, disconnect MQTT)
	rdb.Close()
}
