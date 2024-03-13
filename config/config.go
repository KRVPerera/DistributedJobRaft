package config

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"time"
)

type ClusterConfig struct {
	XMLName          xml.Name      `xml:"ClusterConfig"`
	MyID             int           `xml:"MyID"`
	ElectionTimeout  time.Duration `xml:"ElectionTimeout"`
	HeartbeatTimeout time.Duration `xml:"HeartbeatTimeout"`
	CommitTimeout    time.Duration `xml:"CommitTimeout"`
	Peers            []Peer        `xml:"Peers>Peer"`
	ListenerAddress  string        `xml:"ListenerAddress"`
}

type Peer struct {
	XMLName     xml.Name `xml:"Peer"`
	PeerID      int      `xml:"PeerID"`
	PeerAddress string   `xml:"PeerAddress"`
}

// LoadConfigFromXML loads the configuration settings from an XML file.
// It takes a file path as a parameter and returns a pointer to a ClusterConfig struct and an error.
func LoadConfigFromXML(filePath string) (*ClusterConfig, error) {
	// Log the file being loaded
	log.Printf("Loading config from %s", filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return nil, err
	}
	defer file.Close()

	// Read the file content
	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		return nil, err
	}

	// Create an instance of ClusterConfig
	var cfg ClusterConfig

	// Parse the XML content into the ClusterConfig instance
	err = xml.Unmarshal(data, &cfg)
	if err != nil {
		log.Printf("Failed to unmarshal XML: %v", err)
		return nil, err
	}

	// Log the loaded config
	log.Printf("Loaded config: %+v", cfg)
	return &cfg, nil
}
