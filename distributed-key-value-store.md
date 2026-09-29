 Distributed Key-Value Store

design space:
  requirements:
    - single node system - local
    - cmd line interface
    - get/set api

independent variable:

Planning and Design:

Considerations TODO:
- what is the actual use case ?
- what are the throughput requirement ?
	- what will be the bottleneck ? e.g. disk, RAM, NIC etc
- what is the scale and volume of data ? volume of access of data ?
- what are our latency requirement ? primary requirement ?
- what is our consistency model ? probably eventual consistent
- what are the specific semantic we are supporting ?
	- get/set ?
	- delete ?
	- update ?
	- others, e.g. transactional updates ?
- what about a secondary index ?

Step 1 scope:
- just GET/SET
	- don't worry about DELETE; and SET will overwrite an existing key
- placeholder wire protocol
	- keys and values cannot have spaces
	- messages are operation(GET or SET) <space\> key <space\> value
	- first 3bytes are just "GET" or "SET"
	- next 10 bytes are key
	- rest of message is the value
- use UDP over localhost to not worry about connections
- aim to support two clients on one machine speaking to the same serve, which simply blocks while servicing each client
- do it all in memory: worry about persistence later

Step 1 steps/milestone:
- server which echos over UDP
- server which sets on SET and gets on GET
- test with two clients

