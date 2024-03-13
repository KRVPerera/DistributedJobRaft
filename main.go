package main

import (
	"encoding/json"
	"fmt"
	"github.com/KRVPerera/DistributedJobRaft/config"
	"github.com/KRVPerera/DistributedJobRaft/raft"
	"log"
	"net/http"
)

type CommandRequest struct {
	ID      int         `json:"id"`
	Command interface{} `json:"command"`
}

var ns *raft.Server

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var cmdReq CommandRequest
	err := json.NewDecoder(r.Body).Decode(&cmdReq)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	isLeader := ns.Submit(cmdReq.Command)
	if !isLeader {
		http.Error(w, "Not a leader", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Create a sample config
	cfg, err := config.LoadConfigFromXML("config/config.xml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Printf("Loaded config: %+v\n", cfg)

	readyChan := make(chan interface{})
	storageForServer := raft.NewSQLiteStorage()
	commitChannel := make(chan raft.CommitEntry)
	ns = raft.NewServer(cfg.MyID, config.ExtractPeerIDs(cfg), storageForServer, readyChan, commitChannel)
	ns.Serve(cfg.ListenerAddress)

	cfg2, err2 := config.LoadConfigFromXML("config/config2.xml")
	if err2 != nil {
		log.Fatalf("Failed to load config: %v", err2)
	}
	log.Printf("Loaded config2: %+v\n", cfg2)
	storageForServer2 := raft.NewSQLiteStorage()
	commitChannel2 := make(chan raft.CommitEntry)
	singleServer2 := raft.NewServer(cfg2.MyID, config.ExtractPeerIDs(cfg2), storageForServer2, readyChan, commitChannel2)
	singleServer2.Serve(cfg2.ListenerAddress)

	cfg3, err3 := config.LoadConfigFromXML("config/config3.xml")
	if err3 != nil {
		log.Fatalf("Failed to load config3: %v", err3)
	}
	log.Printf("Loaded config3: %+v\n", cfg3)
	storageForServer3 := raft.NewSQLiteStorage()
	commitChannel3 := make(chan raft.CommitEntry)
	singleServer3 := raft.NewServer(cfg3.MyID, config.ExtractPeerIDs(cfg3), storageForServer3, readyChan, commitChannel3)
	singleServer3.Serve(cfg3.ListenerAddress)

	// go through peer addresses and connect to them
	for _, peer := range cfg.Peers {
		log.Printf("Connecting to peer id : %d, peer address : %s\n", peer.PeerID, peer.PeerAddress)
		// break address to host and port
		err := ns.ConnectToPeerStringAddress(peer.PeerID, peer.PeerAddress)
		if err != nil {
			log.Fatalf("Failed to connect to peer : %d, peer address : %s, error : %v\n", peer.PeerID, peer.PeerAddress, err)
		}
	}

	for _, peer := range cfg2.Peers {
		log.Printf("Connecting to peer id : %d, peer address : %s\n", peer.PeerID, peer.PeerAddress)
		// break address to host and port
		err := singleServer2.ConnectToPeerStringAddress(peer.PeerID, peer.PeerAddress)
		if err != nil {
			log.Fatalf("Failed to connect to peer : %d, peer address : %s, error : %v\n", peer.PeerID, peer.PeerAddress, err)
		}
	}

	for _, peer := range cfg3.Peers {
		log.Printf("Connecting to peer id : %d, peer address : %s\n", peer.PeerID, peer.PeerAddress)
		// break address to host and port
		err := singleServer3.ConnectToPeerStringAddress(peer.PeerID, peer.PeerAddress)
		if err != nil {
			log.Fatalf("Failed to connect to peer : %d, peer address : %s, error : %v\n", peer.PeerID, peer.PeerAddress, err)
		}
	}
	close(readyChan)
	//singleServer.ConnectToPeer(1, cfg.PeerAddresses[1])

	//
	//clusterSize := 3
	//connected := make([]bool, clusterSize)
	//alive := make([]bool, clusterSize)
	//commitChans := make([]chan raft.CommitEntry, clusterSize)
	//ready := make(chan interface{})
	//storage := make([]*raft.MapStorage, clusterSize)
	//
	//// Create all Servers in this cluster, assign ids and peer ids.
	//for i := 0; i < clusterSize; i++ {
	//	peerIds := make([]int, 0)
	//	for p := 0; p < clusterSize; p++ {
	//		if p != i {
	//			peerIds = append(peerIds, p)
	//		}
	//	}
	//
	//	storage[i] = raft.NewMapStorage()
	//	commitChans[i] = make(chan raft.CommitEntry)
	//	ns[i] = raft.NewServer(i, peerIds, storage[i], ready, commitChans[i])
	//	ns[i].Serve(":0")
	//	alive[i] = true
	//}
	//
	//// Connect all peers to each other.
	//for i := 0; i < clusterSize; i++ {
	//	for j := 0; j < clusterSize; j++ {
	//		if i != j {
	//			ns[i].ConnectToPeer(j, ns[j].GetListenAddr())
	//			log.Println("Connected", i, "to", ns[j].GetListenAddr())
	//		}
	//	}
	//	connected[i] = true
	//}
	//close(ready)

	//singleServer.Shutdown()
	//singleServer2.Shutdown()
	//singleServer3.Shutdown()
	http.HandleFunc("/submit", SubmitHandler)
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}
