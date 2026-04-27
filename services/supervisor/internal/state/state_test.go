package state

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

// This requires a local Redis instance running or a mock.
// For demonstration of TDD, we assume a local redis is available on default port.
func TestSupervisorState_SetShiftStatus(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	// Quick ping to check if redis is up, skip test if not
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		t.Skip("Redis not available, skipping state test.")
	}

	stateMgr := NewSupervisorState(rdb)
	ctx := context.Background()

	// Act
	err := stateMgr.SetShiftStatus(ctx, "test-shift-1", true)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assert
	val, err := rdb.Get(ctx, "shift:test-shift-1:status").Result()
	if err != nil {
		t.Fatalf("Expected to find key, got error %v", err)
	}
	if val != "active" {
		t.Errorf("Expected status to be 'active', got '%s'", val)
	}

	// Cleanup
	rdb.Del(ctx, "shift:test-shift-1:status")
}
