package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// Create a sample config
	config := Config{
		ID:               1,
		Peers:            []string{"peer1", "peer2", "peer3"},
		ElectionTimeout:  1000,
		HeartbeatTimeout: 500,
	}

	// Test the ID field
	if config.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", config.ID)
	}

	// Test the Peers field
	expectedPeers := []string{"peer1", "peer2", "peer3"}
	if len(config.Peers) != len(expectedPeers) {
		t.Errorf("Expected Peers to have length %d, got %d", len(expectedPeers), len(config.Peers))
	}
	for i, peer := range config.Peers {
		if peer != expectedPeers[i] {
			t.Errorf("Expected Peers[%d] to be %s, got %s", i, expectedPeers[i], peer)
		}
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
