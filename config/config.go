package config

import (
	"encoding/xml"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

type ClusterConfig struct {
	XMLName          xml.Name      `xml:"ClusterConfig"`
	MyID             int           `xml:"MyID"`
	ElectionTimeout  time.Duration `xml:"ElectionTimeout"`
	HeartbeatTimeout time.Duration `xml:"HeartbeatTimeout"`
	CommitTimeout    time.Duration `xml:"CommitTimeout"`
	PeerIDs          []int         `xml:"PeerIDs>PeerID"`
	PeerAddresses    []string      `xml:"PeerAddresses>PeerAddress"`
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

	// sort the peer IDs
	sort.Ints(cfg.PeerIDs)

	// sort addresses
	sort.Strings(cfg.PeerAddresses)

	// Log the loaded config
	log.Printf("Loaded config: %+v", cfg)
	return &cfg, nil
}
