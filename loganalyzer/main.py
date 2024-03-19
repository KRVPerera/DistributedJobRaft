from collections import defaultdict
import sys
import os
import re
from datetime import datetime

mathStartLine = re.compile(r"^=== RUN\s+(\w+)")
mathEndLine = re.compile(r"^--- PASS")
matchListeningPort = re.compile(r".*\[(\d+)\] listening at \[.*\]:(\d+)")
# 13:21:11.892239 EvalDump::[0]::Status::[Candidate]::Name::[0]
matchEvalLine = re.compile(r"(\d+):(\d+):(\d+)\.(\d+)\sEvalDump::\[(\d+)\]::Status::\[(\w+)\]::Name::\[(\d+)\]")
matchAppendEntries = re.compile(r".*AppendEntries: {Term:\d+\sLeaderId:\d+\sPrevLogIndex:\d+\sPrevLogTerm:\d+\sEntries:\[(.*)\]\sLeaderCommit:\d+}.*")
matchRequestVotes = re.compile(r".*RequestVote: {Term:\d+ CandidateId:(\d+) LastLogIndex:-?\d+ LastLogTerm:-?\d+}.*")


class LogAnalyzer:
    def __init__(self, filename):
        self.filename = filename
        self.tests = []
        self.listeningPorts = []
        self.leaderSet = ()
        self.currentLeader = None
        self.testStarted = False
        self.firstTime = None
        self.currentTime = datetime.now()
        self.appendEntries = defaultdict(lambda: defaultdict(lambda: defaultdict(int)))
        self.leaderCount = defaultdict(int)
        self.appendEntriesCounts = 0
        self.requestVoteCounts = 0

    def readLineByLine(self):
        with open(self.filename, 'r') as file:
            for line in file:
                yield line

    def parseFile(self):
        for line in self.readLineByLine():
            startLineMatch = mathStartLine.match(line)
            if startLineMatch:
                self.tests.append(startLineMatch.group(1))
                self.testStarted = True
            if not self.testStarted:
                continue
            if mathEndLine.match(line):
                self.testStarted = False
            ports = matchListeningPort.match(line)
            if ports:
                nodeAndPort = (int(ports.group(1)), int(ports.group(2)))
                self.listeningPorts.append(nodeAndPort)

            evalLine = matchEvalLine.match(line)
            if evalLine:
                # print(f"Eval line: {line}")
                appendEntry = matchAppendEntries.match(line)
                if appendEntry:
                    self.appendEntriesCounts += 1
                requestVoteLine = matchRequestVotes.match(line)
                if requestVoteLine:
                    self.requestVoteCounts += 1
                node = int(evalLine.group(5))
                status = evalLine.group(6)
                if (status == 'Leader') and self.currentLeader != node:
                    self.leaderCount[node] += 1
                    self.currentLeader = node
                time = f"{evalLine.group(1)}:{evalLine.group(2)}:{evalLine.group(3)}.{evalLine.group(4)}"
                time_obj = datetime.strptime(time, "%H:%M:%S.%f")
                if not self.firstTime:
                    self.firstTime = time_obj
                    self.currentTime = time_obj
                    # print(f"First time: ", end="")
                    # self.printTime(time_obj)
                # compare if the time difference is greater than 1 second
                if (time_obj - self.currentTime).seconds > 0:
                    # self.printTime(time_obj)
                    self.currentTime = time_obj

    def printTime(self, time):
        print(f"{time.hour}:{time.minute}:{time.second}.{time.microsecond}")

    def printReport(self):
        print("=" * 10 + " Report " + "=" * 10)
        print("File Name: " + self.filename)
        print(f"Total tests: {self.tests}")
        print(f"Node Count : {len(self.listeningPorts)}")
        print(f"Listening Ports: {self.listeningPorts}")
        for k, v in self.leaderCount.items():
            print(f"Node : {k} - leader for {v} times")
        print(f"AppendEntries Count : {self.appendEntriesCounts}")
        print(f"RequestVote Count : {self.requestVoteCounts}")


def readLineByLine(filename):
    with open(filename, 'r') as file:
        for line in file:
            yield line


def parseFile(filename):
    for line in readLineByLine(filename):
        print(line)


def main():
    # Get the filename from command line arguments
    filename = "../logs/TestEval_PayLoadLarge.log"
    if len(sys.argv) >= 2:
        filename = sys.argv[1]

    # check if the file exists
    try:
        with open(filename, 'r') as file:
            pass
    except FileNotFoundError:
        print(f"File {filename} not found")
        sys.exit(1)

    # Call the python_main function with the filename
    la = LogAnalyzer(filename)
    la.parseFile()
    la.printReport()


if __name__ == '__main__':
    main()
