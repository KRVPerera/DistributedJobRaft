package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// Create a sample config
	config := Config{
		ID:               1,
		ElectionTimeout:  1000,
		HeartbeatTimeout: 500,
	}

	// Test the ID field
	if config.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", config.ID)
	}

	// Test the ElectionTimeout field
	if config.ElectionTimeout != 1000 {
		t.Errorf("Expected ElectionTimeout to be 1000, got %d", config.ElectionTimeout)
	}

	// Test the HeartbeatTimeout field
	if config.HeartbeatTimeout != 500 {
		t.Errorf("Expected HeartbeatTimeout to be 500, got %d", config.HeartbeatTimeout)
	}
}
