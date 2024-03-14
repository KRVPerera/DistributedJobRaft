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

class LogAnalyzer:
    def __init__(self, filename):
        self.filename = filename
        self.tests = []
        self.listeningPorts = []
        self.testStarted = False
        self.firstTime = None
        self.currentTime = datetime.now()
        self.appendEntries = defaultdict(lambda: defaultdict(lambda: defaultdict(int)))

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
                print(f"Eval line: {evalLine.group(0)}")
                time = f"{evalLine.group(1)}:{evalLine.group(2)}:{evalLine.group(3)}.{evalLine.group(4)}"
                time_obj = datetime.strptime(time, "%H:%M:%S.%f")
                if not self.firstTime:
                    self.firstTime = time_obj
                    self.currentTime = time_obj
                    print(f"First time: ", end="")
                    self.printTime(time_obj)
                # compare if the time difference is greater than 1 second
                if (time_obj - self.currentTime).seconds > 0:
                    self.printTime(time_obj)
                    self.currentTime = time_obj

    def printTime(self, time):
        print(f"{time.hour}:{time.minute}:{time.second}.{time.microsecond}")

    def printReport(self):
        print(f"Total tests: {len(self.tests)}")
        print(f"Listening Ports: {self.listeningPorts}")

def readLineByLine(filename):
    with open(filename, 'r') as file:
        for line in file:
            yield line

def parseFile(filename):
    for line in readLineByLine(filename):
        print(line)

def main():
    # Get the filename from command line arguments
    filename = "./logs/sample.log"
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