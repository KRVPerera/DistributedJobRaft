#!/usr/bin/env bash

go test -v ./raft -run TestEval_MessageCountLow > ./logs/TestEval_MessageCountLow.log
go test -v ./raft -run TestEval_MessageCountHigh > ./logs/TestEval_MessageCountHigh.log
go test -v ./raft -run TestEval_PayLoadSmall > ./logs/TestEval_PayLoadSmall.log
go test -v ./raft -run TestEval_PayLoadLarge > ./logs/TestEval_PayLoadLarge.log