Consensus Algorithms
-

* Paxos.
* Chandra–Toueg.
* Raft.

Paxos operates as sequence of proposals,
which may or may not be accepted by majority of the processes in the system.
If proposer receives agreement from majority of the acceptors,
it can commit the proposal by sending commit message with value.
