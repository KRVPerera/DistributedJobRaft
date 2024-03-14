/**
 * Created by Rukshan Perera (rukshan.perera@student.oulu.fi)
 */

package raft

import (
	"testing"
)

func TestNNodesWithLeaderElection(t *testing.T) {
	h := NewSQliteDBHarness(t, 3)
	defer h.Shutdown()

	sleepMs(5000)
	//h.CheckSingleLeader()
}

func TestEval_MessageCountLow(t *testing.T) {
	N := 3
	h := NewSQliteDBHarness(t, N)
	defer h.Shutdown()

	origLeaderId, _ := h.CheckSingleLeader()
	for i := 0; i < 100; i++ {
		h.SubmitToServer(origLeaderId, i)
		sleepMs(100)
	}
	sleepMs(500)
}

func TestEval_MessageCountHigh(t *testing.T) {
	N := 3
	h := NewSQliteDBHarness(t, N)
	defer h.Shutdown()

	origLeaderId, _ := h.CheckSingleLeader()
	for i := 0; i < 2000; i++ {
		h.SubmitToServer(origLeaderId, i)
		sleepMs(100)
	}
	sleepMs(500)
}

func TestEval_PayLoadSmall(t *testing.T) {
	N := 3
	h := NewSQliteDBHarness(t, N)
	defer h.Shutdown()

	origLeaderId, _ := h.CheckSingleLeader()
	for i := 0; i < 500; i++ {
		h.SubmitToServer(origLeaderId, i)
		sleepMs(100)
	}
	sleepMs(500)
}

func TestEval_PayLoadLarge(t *testing.T) {
	N := 3
	h := NewSQliteDBHarness(t, N)
	defer h.Shutdown()

	origLeaderId, _ := h.CheckSingleLeader()
	for i := 0; i < 500; i++ {
		h.SubmitToServer(origLeaderId, "wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
		sleepMs(100)
	}
	sleepMs(500)
}

func TestSQliteElectionLeaderAndAnotherDisconnect(t *testing.T) {
	h := NewHarness(t, 3)
	defer h.Shutdown()

	origLeaderId, _ := h.CheckSingleLeader()

	h.DisconnectPeer(origLeaderId)
	otherId := (origLeaderId + 1) % 3
	h.DisconnectPeer(otherId)

	// No quorum.
	sleepMs(450)
	h.CheckNoLeader()

	// Reconnect one other server; now we'll have quorum.
	h.ReconnectPeer(otherId)
	h.CheckSingleLeader()
}

func TestSimpleDisconnectAllThenRestore(t *testing.T) {
	h := NewHarness(t, 3)
	defer h.Shutdown()

	sleepMs(100)
	//	Disconnect all servers from the start. There will be no leader.
	for i := 0; i < 3; i++ {
		h.DisconnectPeer(i)
	}
	sleepMs(450)
	h.CheckNoLeader()

	// Reconnect all servers. A leader will be found.
	for i := 0; i < 3; i++ {
		h.ReconnectPeer(i)
	}
	h.CheckSingleLeader()
}
