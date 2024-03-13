package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// Create a sample config
	config := ClusterConfig{
		MyID:             1,
		ElectionTimeout:  1000,
		HeartbeatTimeout: 500,
		PeerIDs:          []int{2, 3, 4},
	}

	// Test the ID field
	if config.MyID != 1 {
		t.Errorf("Expected ID to be 1, got %d", config.MyID)
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

func TestLoadConfigFromXML(t *testing.T) {
	// Call the function with the path to your XML file
	cfg, err := LoadConfigFromXML("test_config.xml")
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	t.Logf("Loaded config: %+v", cfg)

	// Assert that the returned ClusterConfig has the expected values
	if cfg.MyID != 1 {
		t.Errorf("Expected myID to be 1, got %d", cfg.MyID)
	}

	if cfg.ElectionTimeout != 5000000000 {
		t.Errorf("Expected ElectionTimeout to be 5000000000, got %d", cfg.ElectionTimeout)
	}

	if cfg.HeartbeatTimeout != 1000000000 {
		t.Errorf("Expected HeartbeatTimeout to be 1000000000, got %d", cfg.HeartbeatTimeout)
	}

	if cfg.CommitTimeout != 2000000000 {
		t.Errorf("Expected CommitTimeout to be 2000000000, got %d", cfg.CommitTimeout)
	}

	if len(cfg.PeerIDs) != 3 {
		t.Errorf("Expected PeerIDs length to be 3, got %d", len(cfg.PeerIDs))
	}

	expectedPeerIDs := []int{1, 2, 3}
	for i, id := range cfg.PeerIDs {
		if id != expectedPeerIDs[i] {
			t.Errorf("Expected PeerID at index %d to be %d, got %d", i, expectedPeerIDs[i], id)
		}
	}

	if len(cfg.PeerAddresses) != 3 {
		t.Errorf("Expected PeerAddresses length to be 3, got %d", len(cfg.PeerAddresses))
	}

	expectedPeerAddresses := []string{"localhost:2461", "localhost:2462", "localhost:2463"}
	for i, addr := range cfg.PeerAddresses {
		if addr != expectedPeerAddresses[i] {
			t.Errorf("Expected PeerAddress at index %d to be %s, got %s", i, expectedPeerAddresses[i], addr)
		}
	}

	if cfg.ListenerAddress != "localhost:2461" {
		t.Errorf("Expected ListenerAddress to be localhost:2460, got %s", cfg.ListenerAddress)
	}
}
