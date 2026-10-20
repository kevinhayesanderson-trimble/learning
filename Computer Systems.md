fetch decode execute:

- pipelining
 	- around 20 stages
  		- initial latency: 15-20 CPU cycles
- Branch Prediction
 	- Speculative execution
- Instruction fetching
 	- batched
 	- pre-fetching
- Micro Ops
 	- Machine code instructions are decoded into micro-ops
 	- one assembly instruction(machine code) can be converted into multiple micro-ops
 	- micro-ops is internal to the cpu, it's the language for cpu
 	- Instruction freezing:
 	- efficient design
- Out of order execution
- Numerous execution units organized into ports
- Register renaming
- SIMD(single instruction multiple data):
 	- single instruction like vector add acting on multiple data(two vectors), two streams of data and two SIMD registers - pair wise operations, thus enabling parallel execution
 	- wider registers, more parallel operations
 	-

| General Purpose registers | AUX - 512     |
| ------------------------- | ------------- |
| 64 bits                   | 512bits = 2^9 |
| 1 int ops per cycle       | 16 32bit int  |

## CPU Cache

- going to the ram is 100x longer than the operation itself
- Cache lines:
 	- notional unit of ram, cpu will pull in the line of data(64 byte) from ram
- CPU speed has improved over time, but ram latency has not been increasing as par
- one cache miss will pull a cache line(64 bytes)
 	- 64 byte cache line
 	- 4 byte integer
 	- 16 ints per cache line
 	- 15 hits per 1 miss
 	- cost of looking up 16 int is 100ns for the miss and 1ns for cache hits
 	-
- typically, L1 cache latency 4 cpu cycles -> 1ns

Array of structs vs Structs of array problem:

- instead of having one array with structs, we can have one struct with array of each data field
- make sense for aggregation operations, not for each user analysis
- column store motivation

Pointer Chasing Problem:

- even worse array of different types(array of pointer to different types)
- each of the pointer lookup is a cache miss

Levels of CPU cache:

- 1 cpu cycle  ~ 0.25ns for 4Ghz
- one cache miss will fill all the levels, benefit

Code Optimizations:

- faster sum
 -
