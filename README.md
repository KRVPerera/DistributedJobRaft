# Raft Based Distributed Job Queue

## About the project

This project implements a scalable distributed system using the Raft consensus algorithm. The system can grow to any number of nodes and ensures consistency through leader election and log replication. Nodes are containerized for easy deployment and communicate via RPC. This document outlines the design, implementation, and evaluation methods for the Raft cluster.

## Implemented components

We have a raft based cluster which can be scaled to N number of nodes.
Cluster is RPC connected and setup statically. Which meansEach node is aware of the other nodes in the cluster. This can be configured by `config/config.xml`

- [x] Leader Election
- [x] Log Replication
- [x] Raft persistent states on SQlite DB
- [x] Docker containerization

Our system can have any number of nodes and the leader election is done using Raft consensus algorithm. The leader is responsible for the log replication.
Each node goes through the following states:
- Follower
- Candidate
- Leader

*This image is from the raft paper*
![Node state changes](docs%2Fimages%2Fraft_node_state_changes.png)


### Naming
- Each node is given uniques `integer` number upon creation of the cluster.
- Each node is also introduced to its peers in the cluster upon setup.

### Communication
- Each node is aware of the other nodes in the cluster and can communicate with them using RPC.
- Upon cluster creation stage each node opens and keep a RPC channel between itself and the other nodes in the cluster.

RPC calls are exactly the ones described in the Raft paper:
- RequestVote
- AppendEntries

### Fault tolerance

In this approach, fault tolerance is achieved by Leader election, Log replication and term-based approach. 

- Leader Election - Raft elects a new leader if the current leader fails, ensuring continuous operation through majority vote.

- Log replication - Raft replicates logs for fault tolerance. Leader waits for majority to copy entries before applying them, ensuring data consistency even with node failures.

- Term-based approach - Raft uses "terms", where each term start with leader election. This helps spot leader failures and trigger new elections if needed.

- Raft persistent states are saved in SQLite database so that when a node restarts from failure it does not have to start fresh.

### Synchronization

Raft achieves synchronization by bringing all nodes in the cluster into a consistent state.

- Leader is responsible for cordinating log entries and maintaining consistency accross the cluster.

- After the leader is elected, it replicate it's log entries to the other nodes in the cluster. Before commiting a log entry, it ensures that it is safely replicated to majority of nodes. This makes sure that all nodes have the same log entries.

- Raft has a commit mechanism to make sure that all nodes agree on which log entries are considered committed. Once a log entry has been replicated to a majority of nodes, the leader can commit the entry, and all nodes in the cluster will also commit it. This ensures that all nodes agree on the state of the replicated log.

- Also, raft uses heartbeats and Append Entries RPCs to maintain synchronization and detect failures.

- Term-based approch ensures that nodes are synchronized within the same term.

### Consistency

Raft nodes mostly rely on the leader to tell the nodes to commit the entries. Therefore we can say the cluster achieves eventual consistency.

### Architecture and Processes

Architecture is completely based on Raft paper. We dont have log compaction and cluster is static. Almost all the functionalities are written as go routienes and shared data is protected by mutex locks.


## Built with:
   - Go language
   - SQlite
   - Docker
   - TCP

## Getting Started:
Currently our containerized solution does not work properly. Therefore we have evaluated by creating servers in separate threads. They have full functionaluty except that they run on the same machine.

To test docker containerized version you can run the system on your local machine using docker.
`setup.sh` creates docker images and `run.sh` initializes the cluster.

Configuration files for each node is in `config` folder. Code and scale to few number of nodes (tested upto 13) but only three docker images are created using config files.

## Results of the tests:

### Test 1 : number of messages

- scenario 1 : 100 message to the leader with 100 ms delays and 500 ms sleep after

```
go test -v ./raft -run TestEval_MessageCountLow
```

- scenario 2 : 2000 message to the leader with 100 ms delays and 500 ms sleep after

```
go test -v ./raft -run TestEval_MessageCountHigh
```
|Message count | Time Taken(s) | Through put(messages/s)|
|--------------|---------------|------------------------|
|100           | 13.19         | 7.6                    |
|2000          |  266.83       | 7.5                    |


### Test 2 (Failed): payload size

- scenario 1 : 1000 integer messages to the leader with 100 ms delays and 500 ms sleep after
```
go test -v ./raft -run TestEval_PayLoadSmall
```

- scenario 2 : 1000, 200 character long messages to the leader with 100 ms delays and 500 ms sleep after
```
go test -v ./raft -run TestEval_PayLoadLarge
```

|Message size  | Count | Time Taken(s) | Through put(messages/s)|
|--------------|--------------|--------|------------------------|
|4 bytes       |1000          | 123.44         | 8.10                    |
|200 bytes     |1000 (Crashed around 475)  |  600       | 0.79 |


### Test 3 (Failed cannot handle): payload size

- scenario 1 : 1000 integer messages to the leader with 100 ms delays and 500 ms sleep after
```
go test -v ./raft -run TestEval_PayLoadSmall
```

- scenario 2 : 1000, 640 character long messages to the leader with 100 ms delays and 500 ms sleep after
```
go test -v ./raft -run TestEval_PayLoadLarge
```

|Message size  | Count | Time Taken(s) | Through put(messages/s)|
|--------------|--------------|--------|------------------------|
|4 bytes       |1000          | 123.44         | 8.10                    |
|640 bytes     |1000 (Crashed around 475)  |  600       | 0.79 |

Collect numerical data of test cases:
- Collecting logs of container operations
- Conduct simple analysis for documentation purposes (e.g. plots or graphs)

## Future Enhancements

By considering some listed enhancements, the Raft-based distributed job queue can evolve into a more robust, scalable, and feature-rich system capable of meeting the needs of a wide range of use cases.

- Dynamic Clustering: 
  By making nodes to join or leave the group whenever needed, making the system more adaptable to changes in node availability.

- Fault Recovery Mechanisms:
  Build in automatic recovery features to deal with node failures, ensuring that the system can quickly recover and continue functioning without interruption.

- Performance Optimization:
  Identify and fix any slowdowns in the system, such as reducing unnecessary communication overhead or making data replication more efficient,
  to make the system faster and more responsive.

- Enhanced Monitoring and Logging:
   Incorporate external tools for monitor  how the system is doing and for tracking down and fixing any problems that come up.

- Security Enhancements:
   Can improve security by adding features like user authentication and encryption to protect sensitive data from unauthorized access.

- Job Prioritization and Scheduling:
   Give users more control over how jobs are handled, allowing them to set priorities and schedule tasks based on their needs.

- Support for Event-Driven Architecture:
   Enable the system to respond to events or messages, allowing for more flexible and efficient job handling.

- Documentation and User Interface Improvements:
   by improving documentation and creating a more user-friendly interfaces
	

## Acknowledgments:
- [Raft](https://raft.github.io/)
- [Eli Bendersky's Raft Guide](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/)

