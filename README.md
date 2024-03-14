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



Participating nodes must:
- Exchange information (messages): RPC,
- Log their behavior understandably: messages, events, actions, etc.

### Naming
- Each node is given uniques `integer` number upon creation of the cluster.
- Each node is also introduced to its peers in the cluster upon setup.

### Communication
- Each node is aware of the other nodes in the cluster and can communicate with them using RPC.
- Upon cluster creation stage each node opens and keep a RPC channel between itself and the other nodes in the cluster.

RPC calls are exactly the ones described in the Raft paper:
- RequestVote
- AppendEntries
- RequestVoteResponse

### Fault tolerance

In this approach, fault tolerance is achieved by Leader election, Log replication and term-based approach. 

- Leader Election - Raft elects a new leader if the current leader fails, ensuring continuous operation through majority vote.

- Log replication - Raft replicates logs for fault tolerance. Leader waits for majority to copy entries before applying them, ensuring data consistency even with node failures.

- Term-based approach - Raft uses "terms", where each term start with leader election. This helps spot leader failures and trigger new elections if needed.

### Synchronization

Raft achieves synchronization by bringing all nodes in the cluster into a consistent state.

- Leader is responsible for cordinating log entries and maintaining consistency accross the cluster.

- After the leader is elected, it replicate it's log entries to the other nodes in the cluster. Before commiting a log entry, it ensures that it is safely replicated to majority of nodes. This makes sure that all nodes have the same log entries.

- Commit log entries - // TODO

- Also, raft uses heartbeats and Append Entries RPCs to maintain synchronization and detect failures.

- Term-based approch ensures that nodes are synchronized within the same term.

Detailed descriptions of relevant principles covered in the course (architecture, processes, communication, naming, synchronization, consistency and replication, fault tolerance); irrelevant principles can be left out.

## Built with:
Detailed description of the system functionality and how to run the implementation.

- If you are familiar with a particular container technology, feel free to use it (Docker is not mandatory)
- Any programming language can be used, such as: Python, Java, JavaScript, ..
- Any communication protocol / Internet protocol suite can be used: HTTP(S), MQTT, AMQP, CoAP, ..

## Getting Started:
You can run the system on your local machine using docker. Look at `setup.sh` and `run.sh` for more details.



## Results of the tests:
Detailed description of the system evaluation
Evaluate your implementation using selected criteria, for example:
- Number of messages / lost messages, latencies, ...
- Request processing with different payloads, ..
- System throughput, ..


Design two evaluation scenarios that you compare with each other, for example:
- Small number / large number of messages
- Small payload / big payload

Collect numerical data of test cases:
- Collecting logs of container operations
- Conduct simple analysis for documentation purposes (e.g. plots or graphs)

## Acknowledgments:
- [Raft](https://raft.github.io/)
- [Eli Bendersky's Raft Guide](https://eli.thegreenplace.net/2020/implementing-raft-part-0-introduction/)

