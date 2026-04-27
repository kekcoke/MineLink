package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/kekcoke/minelink/supervisor/internal/state"
)

type Orchestrator struct {
	client mqtt.Client
	state  *state.SupervisorState
	id     string
}

// TacticalAssignment represents the schema defined in asyncapi.yaml
type TacticalAssignment struct {
	CommandID    string `json:"commandId"`
	SupervisorID string `json:"supervisorId"`
	Action       string `json:"action"` // SCALE_UP, SCALE_DOWN, ROTATE_SHIFTS
	Parameters   struct {
		WorkerCount int `json:"workerCount"`
	} `json:"parameters"`
}

func NewOrchestrator(brokerURL, clientID string, s *state.SupervisorState) (*Orchestrator, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientID)
	
	// Basic TLS config could be added here for mTLS

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Orchestrator{
		client: client,
		state:  s,
		id:     strings.TrimPrefix(clientID, "supervisor-"),
	}, nil
}

func (o *Orchestrator) Start(ctx context.Context) error {
	topic := fmt.Sprintf("c2/tactical/supervisor/%s", o.id)
	
	token := o.client.Subscribe(topic, 1, o.handleTacticalAssignment)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}

	log.Printf("Supervisor [%s] orchestrator listening on topic: %s\n", o.id, topic)
	return nil
}

func (o *Orchestrator) handleTacticalAssignment(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received Tactical Assignment: %s", msg.Payload())

	var assignment TacticalAssignment
	if err := json.Unmarshal(msg.Payload(), &assignment); err != nil {
		log.Printf("Error unmarshalling assignment: %v\n", err)
		return
	}

	// Route the action to actualize the intent
	ctx := context.Background()
	switch assignment.Action {
	case "SCALE_UP":
		log.Printf("Executing SCALE_UP: Incrementing workers by %d", assignment.Parameters.WorkerCount)
		// In a real system, this would spawn C++ Operator processes. 
		// For now, we simulate actualization by updating Redis state.
		// TODO: Implement actual process spawning via gRPC or local OS commands
	case "ROTATE_SHIFTS":
		log.Printf("Executing ROTATE_SHIFTS for Shift B")
		err := o.state.SetShiftStatus(ctx, "shift-b", true)
		if err != nil {
			log.Printf("Failed to rotate shift: %v\n", err)
		}
	default:
		log.Printf("Unknown action: %s\n", assignment.Action)
	}
}
