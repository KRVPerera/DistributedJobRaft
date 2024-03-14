# Industry track

## Raft Based Distributed Job Queue

## About the project
## Implemented components:

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
- each node is given uniques `integer` number upon creation of the cluster
- Each node is also introduced to its peers in the cluster upon setup

### Communication
- Each node is aware of the other nodes in the cluster and can communicate with them using RPC
- Upon cluster creation stage each node opens and keep a RPC channel between itself and the other nodes in the cluster

RPC calls are exactly the ones described in the Raft paper:
- RequestVote
- AppendEntries
- RequestVoteResponse

Detailed descriptions of relevant principles covered in the course (architecture, processes, communication, naming, synchronization, consistency and replication, fault tolerance); irrelevant principles can be left out.

## Built with:
Detailed description of the system functionality and how to run the implementation 

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

