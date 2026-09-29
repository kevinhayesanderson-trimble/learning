why ?
	- at what point the replication is not working to kick in the partitioning
	- cannot scale the writes with replication
	- for write heavy workload, it should be partitioned
	- or for locality of related data, related data in the single partition(local sequential access)
	- scale reads with replication [[Replication]] and writes with Partitioning

Challenges with partitioning:
- how do you route the request ?
- repartitioning should not change application logic
- local or global index 

Manual vs Automatic Partitioning:

Index in Partitioning:
- do we store the index on each node/partition ?
- global index - which is also be partitioned
- updating the index on write transactionally

Hot Partitioning:

Distribute based on hash:
- hash(key) % n
	- n - no of node
- mod by no of node is bad:
	- add or remove node, all the previous hashed items are inaccessible 

Consistent hashing:
- ids in a ring
- 2^32 bit identifier
- logical and physical partition
- 

Papers:
- https://arxiv.org/pdf/1406.2294
- https://arxiv.org/pdf/1505.00062

