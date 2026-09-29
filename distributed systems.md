goals:
- scalable(we have hit the limits of single node hardware)
- reliable(fail-over mitigation)
- maintainable

OZ framing:
2 tools:
- replications(copying, to prevent data loss, replication lag)
	- synchronous or asynchronous ?
	- asynchronous - best effort
	- more replicas decreases reliability ,on synchronous replication, what is happens to reliability on async replication
	- semi-synchronous replica
	- primary is down:
		- automatic fail-over to secondary ? or a person should be pinged ?
		- how do you decide even the primary is down ? probe ? heartbeat ?
		- cost of fail-over ?
	- cold cache ?
	- on fail-over, on some failure, be aware of write on both primary and secondary db
	- full replication give fault tolerance
- partitioning(sharding or splitting)
	- need ?
	- why ?
	- how ?
	- partition on write throughput ?
	- partition on read throughput ?

Big Costs to consider in tradeoffs:
- consistency
- complexity/maintainability

Consistency:

Complex Systems:


Auto-incrementing integer id:
- bad
- uuid
- primary id/internal id auto-incrementing and secondary key as uuid - db specific optimisation

cache updating and consolidation in distributed systems, problem ?

Papers:
- dynamo paper(trade-off read consistency) 
	- https://cdn.amazon.science/ac/1d/eb50c4064c538c8ac440ce6a1d91/dynamo-amazons-highly-available-key-value-store.pdf
- https://how.complexsystems.fail/
- https://github.blog/news-insights/github-availability-this-week/
- post-mortems:
	- https://github.com/danluu/post-mortems


Goals:
- Scalability
- reliability
- maintainability

Scalability:
- cpu vs ram vs storage vs i/o
- throughput vs latency
- read vs write


	

