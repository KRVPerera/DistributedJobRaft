package raft

import (
	"sync"
	"time"
	"log"
	"fmt"
)

type RaftState uint8

const (
	Follower RaftState = iota
	Candidate
	Leader
)

type LogEntry struct {
	Command interface{}
	Term    int
}

type PersistentState struct {
	currentTerm        	int
	votedFor           	int
	log                	[]LogEntry
}

type VolatileState struct {
	commitIndex        	int
	lastApplied        	int
}

type VolatileStateOnLeader struct {
	nextIndex          	[]int
	matchIndex         	[]int
}

type RaftConsensusModule struct {
	mutex             		sync.Mutex
	state             		RaftState
	ID                 		int
	peers              		[]int
	electionTimeout    		int
	heartbeatTimeout   		int
	persistentState 		PersistentState
	volatileState   		VolatileState
	VolatileStateOnLeader 	VolatileStateOnLeader
	electionResetEvent 		time.Time
}

func InitRaftConsensusModule(id int, peers []int, electionTimeout int, heartbeatTimeout int, ready <-chan interface{}) *RaftConsensus {
	raftConsensus := &RaftConsensusModule{
		state:            Follower,
		ID:               id,
		peers:            peers,
		electionTimeout:  electionTimeout,
		heartbeatTimeout: heartbeatTimeout,
		currentTerm:      0,
		votedFor:         -1,
	}

	go func() {
		<-ready	
		raftConsensus.mutex.Lock()
		raftConsensus.electionResetEvent = time.Now().Add(time.Duration(raftConsensus.electionTimeout) * time.Millisecond)
		raftConsensus.mutex.Unlock()
		raftConsensus.runStateMachine()
	}()
	return raftConsensus
}

func (raftConsensus *RaftConsensusModule) Debug(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s][%d] ", time.Now().String(), cm.id) + format
	log.Printf(format, args...)
}

func (raftConsensus *RaftConsensusModule) isLeader() bool {
	return raftConsensus.state == Leader
}

func (raftConsensus *RaftConsensusModule, termStarted: int) termChanged () bool {
	return raftConsensus.currentTerm != termStarted
}

func (raftConsensus *RaftConsensusModule) runStateMachine() {
  timeoutDuration := time.Duration(cm.electionTimeout + rand.Intn(cm.electionTimeout)) * time.Millisecond
  raftConsensus.mutex.Lock()
  termStarted := cm.currentTerm
  raftConsensus.mutex.Unlock()

  raftConsensus.Debug("Statemachine timer started (%v), term=%d", timeoutDuration, termStarted)

  ticker := time.NewTicker(10 * time.Millisecond)
  defer ticker.Stop()

  for {
    <-ticker.C

    raftConsensus.mutex.Lock()
    if raftConsensus.isLeader() {
      raftConsensus.mutex.Unlock()
      return
    }

    if raftConsensus.termChanged() {
      raftConsensus.Debug("Term changed termStarted :  %d, currentTerm : %d", termStarted, raftConsensus.currentTerm)
      raftConsensus.mutex.Unlock()
      return
    }


    if elapsed := time.Since(raftConsensus.electionResetEvent); elapsed >= timeoutDuration {
      raftConsensus.startElection()
      raftConsensus.mutex.Unlock()
      return
    }
    raftConsensus.mutex.Unlock()
  }
}
