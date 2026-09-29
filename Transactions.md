https://zookeeper.apache.org/
https://etcd.io/
https://cloud.google.com/spanner
https://raft.github.io/


Distributed Transactions:
- Two Phase Transactions:
	- 1st phase:
		- ask from coordinator
		- confirm/abort from participants
	- 2nd phase:
		- commit from coordinator
		- commit ack from participants
	- between the 1st phase and 2nd commit , there has to be locks on the data to prevent any other transactions
	- holding long time locks - not ideal
	- single point of failures in coordinator - synchronous replication of coordinator - coordinator failing
	- no leader or two leader scenario's
	- 

