#!/usr/bin/env bash

go test -v -c -o github.com/KRVPerera/DistributedJobRaft/raft
go tool test2json -t -test.v -test.paniconexit0 -test.run ^\QTestNNodesWithLeaderElection\E$