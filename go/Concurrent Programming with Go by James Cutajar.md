
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

User-Level threads:
	- threads running in the user-space
	- memory space is part of our application

Kernel-level threads:
 - OS managed threads and context(registers, stack and state)
 - OS memory space

User-level threads executing within a single kernel-level thread

User-level threads:
-  needs it's own thread management system, mimics OS
- context switching is fast
- Disadvantages:
	- blocking I/O reads
	- one process at a time in multicore system

Green Threads:
- JAVA - user-level thread name
- later virtual threads introduced in JAVA- kernel-level


## Hybrid Technique:
- kernel-level threads containing user-level threads
- M:N hybrid threading
- M:N threading model:
	- M user-level threads (goroutines) mapped to N kernel-level threads
	- normal user-level threads -> N:1 threading model, N user-level threads with 1 kernel level threads
	- go's runtime determines how many kernel-level threads to use based on the no of logical processors -> set in GOMAXPROCS(no of cpus)
	- go's runtime assigns a Local run queue(LRQ) to each kernel-level thread , this will contain the subset of the goroutines
	- Global run queue(GRQ) has go-routines which doesn't have the LRQ assigned
	-  if any goroutine blocks a kernel-level thread on  I/O, go's runtime creates a new kernel-level thread and moves the queue of go-routines to the new thread
	- Work Stealing:
		- also used to balance the LRQs

Locking to a kernel-level thread:
- ```
  runtime.LockOSThread()
  runtime.UnlockOSThread()
  ```

- specialized control over kernel-level threads eg: interfacing with c library


Scheduling goroutines:


Concurrency vs Parallelism:
- Concurrency is about _planning_ how to do many tasks at the same time. 
- Parallelism is about _performing_ many tasks at the same time.
- A concurrent tasks may execute in parallel
-  parallelism is a subset of concurrency
