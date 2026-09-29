
benefits:
- increasing the throughput
- improving responsiveness
	- task fast-switching and interleaving system is implemented by using a combination of hardware interrupts and operating system traps
- performance scaling
	- measure of how well a program speeds up in proportion to the increase in the number of resources available to the program

Scalability Limit:
- Amdahl’s law
- Gustafson’s law

## *Concurrency with processes*
- provides isolation
- on UNIX systems, ```fork()``` system call to create a copy of an execution
- ```Copy on write (COW)``` is an optimization introduced to the ```fork()``` system call
- https://pkg.go.dev/golang.org/x/sys
- 
  
## *Concurrency with threads*
- threads are another execution context (kind of a microprocess) within a process
- each single execution a _thread_ (or _thread of execution_)
- On Linux, we can use the ```clone()``` system call with the CLONE_THREAD option



