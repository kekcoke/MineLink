package state

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SupervisorState struct {
	redis *redis.Client
}

func NewSupervisorState(r *redis.Client) *SupervisorState {
	return &SupervisorState{redis: r}
}

// SetShiftStatus updates the local Redis state for a shift rotation
func (s *SupervisorState) SetShiftStatus(ctx context.Context, shiftID string, active bool) error {
	key := fmt.Sprintf("shift:%s:status", shiftID)
	status := "inactive"
	if active {
		status = "active"
	}
	return s.redis.Set(ctx, key, status, 24*time.Hour).Err()
}

// GetActiveWorkerCount retrieves the number of active operator agents from local state
func (s *SupervisorState) GetActiveWorkerCount(ctx context.Context) (int, error) {
	val, err := s.redis.Get(ctx, "stats:active_workers").Int()
	if err == redis.Nil {
		return 0, nil
	}
	return val, err
}
