gettimeofday gives a wall-clock snapshot
getrusage gives cumulative CPU counters for this process

 tv -> timeval
 tv_sec -> seconds
 tv_usec -> microseconds
 tv_nsec -> nanoseconds

user: CPU time spent executing programs instructions outside the kernel
sys: CPU time the kernel spent doing the work for the process(i.e syscalls)

Multiplication:
real: 0.405s user: 0.405s sys: 0.000s
CPU time:0.404770 , Wall time:0.404781
    nearly equal, the program actively computing

Allocation:
real: 0.037s user: 0.001s sys: 0.036s
CPU time:0.037371 , Wall time:0.037385
  kernel(sys) did most of the work allocating
  malloc is a C-library allocator, and many calls can be handled in user space by reusing memory it already manages.
  But to grow/manage the process’s memory, the allocator can need help from the kernel;
  touching newly backed memory can also involve kernel page handling.
  That work can contribute to sys.

Sleeping:
real: 3.000s user: 0.000s sys: 0.000s
CPU time:0.000034 , Wall time:3.000124
  sleep makes a system call; the kernel blocks the process until the timer expires,
  so the scheduler can run the other runnable processes.
  if none are runnable, the cpu can be idle
  little cpu time is used to enter the kernel, set up the sleep and resume afterwards

Key difference in single threaded vs multi threaded:
    Process CPU time adds up across threads(running concurrently), while the wall time measures time passing once.
    4 sec of CPU time in 1 sec of wall time

PID - process identity
CPU number - logical cpu observed at the time we asked the OS, the process could migrate to another CPU

get_pid() -  returns the process ID, stays same for the process
sched_getcpu() -  returns the logical cpu number where the calling thread is running at the instant of the call
    scheduler may move it before the next call

Concurrent run: ./cputime & ./cputime
output:
```sh
[1] 15156 #shell's job number
[pid 15156, cpu 13] 100000000 fmuls
[pid 15157, cpu 14] 100000000 fmuls
[pid 15157, cpu 14] real: 0.411s user: 0.411s sys: 0.000s
[pid 15157, cpu 14] CPU time:0.410565 , Wall time:0.410596
[pid 15157, cpu 14] 100000 mallocs of size 1000
[pid 15156, cpu 12] real: 0.412s user: 0.412s sys: 0.000s # Linux moving 15156 between cpu 13 to 12
[pid 15156, cpu 12] CPU time:0.412156 , Wall time:0.412298
[pid 15156, cpu 12] 100000 mallocs of size 1000
[pid 15157, cpu 14] real: 0.041s user: 0.013s sys: 0.028s
[pid 15157, cpu 14] CPU time:0.040971 , Wall time:0.040971
[pid 15157, cpu 14] sleeping for 3
[pid 15156, cpu 12] real: 0.042s user: 0.005s sys: 0.036s
[pid 15156, cpu 12] CPU time:0.041746 , Wall time:0.041746
[pid 15156, cpu 12] sleeping for 3
[pid 15157, cpu 14] real: 3.000s user: 0.000s sys: 0.000s
[pid 15157, cpu 14] CPU time:0.000033 , Wall time:3.000400
[pid 15156, cpu 12] real: 3.000s user: 0.000s sys: 0.000s
[pid 15156, cpu 12] CPU time:0.000043 , Wall time:3.000239
```

the process keeps the user-space state, between the thread CPU migration
linux preserves the process address space and saves its CPU state(registers and instruction pointer), so that it can continue on another CPU


Time:
: Epoch Timestamp
	- CLOCK_REALTIME gives the seconds since 1970-01-01 UTC
	- what time ?
	- can be adjusted by the system clock
: Elapsed since program start
	- CLOCK_MONOTOMIC gives a steadily increasing clock with an arbitrary origin
	- how long has this program been running
