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

	if len(cfg.Peers) != 3 {
		t.Errorf("Expected PeerAddresses length to be 3, got %d", len(cfg.Peers))
	}

	if cfg.ListenerAddress != "localhost:2461" {
		t.Errorf("Expected ListenerAddress to be localhost:2460, got %s", cfg.ListenerAddress)
	}

	// Test the PeerAddresses field
	if cfg.Peers[0].PeerAddress != "localhost:2461" {
		t.Errorf("Expected PeerAddress to be localhost:2461, got %s", cfg.Peers[0].PeerAddress)
	}

	if cfg.Peers[1].PeerAddress != "localhost:2462" {
		t.Errorf("Expected PeerAddress to be localhost:2462, got %s", cfg.Peers[1].PeerAddress)
	}

	if cfg.Peers[2].PeerAddress != "localhost:2463" {
		t.Errorf("Expected PeerAddress to be localhost:2463, got %s", cfg.Peers[2].PeerAddress)
	}
}
