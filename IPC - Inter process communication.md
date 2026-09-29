two process cannot share memory


Links:
- https://capnproto.org/

---
topic: Linux Inter-Process Communication (IPC)
tags: [linux, ipc, system-programming, architecture, devops, database]

---

ipc_overview:
  definition: Kernel-managed or shared memory mechanisms allowing local processes to exchange data or events.

uds_vs_localhost:
  unix_domain_sockets:
    addressing: Filesystem path (e.g., `/tmp/.s.PGSQL.5432`)
    overhead: Zero network stack; direct kernel RAM ring buffer copies
    performance: 30-50% lower latency, reduced CPU cycle consumption
    security: OS file permissions (`chmod`/`chown`)
    scope: Strictly local host
  localhost_tcp:
    addressing: IP + Port (`127.0.0.1:5432`)
    overhead: Full TCP/IP stack (framing, ACKs, checksums, loopback processing)
    performance: Higher latency and network CPU overhead
    security: IP binding, firewall rules, MD5/password auth

ipc_paradigms:
  stream_and_message:
    anonymous_pipes:
      type: Unidirectional stream
      scope: Parent-child processes (e.g., `ps | grep`)
      buffer: Kernel ring buffer (~64 KB default)
    named_pipes_fifos:
      type: Unidirectional stream via filesystem path (`mkfifo`)
      behavior: Blocks on `openat()` until writer attaches; returns 0 bytes (EOF) when all writers disconnect
    unix_domain_sockets:
      type: Full-duplex stream (`SOCK_STREAM`) or datagram (`SOCK_DGRAM`)
      behavior: Point-to-point; kernel maintains two separate buffers per socket (Client->Server, Server->Client)
  zero_copy_shared_memory:
    mechanisms: POSIX (`shm_open`, `mmap`) / System V (`shmget`, `shmat`)
    speed: Fastest IPC method; maps physical RAM directly into virtual spaces of multiple processes (0 kernel copies)
    synchronization: Manual required (Mutexes, RWLocks/LWLocks, Spinlocks, Futexes)
    failure_modes:
      orphaned_locks: Writer process killed by SIGKILL leaves lock byte set to "locked" in RAM
      mitigations:
        posix_robust_mutexes: Kernel detects dead owner and flags `EOWNERDEAD` to next waiting process
        postgres_panic: Main postmaster process detects worker crash, terminates all workers, and wipes/reinitializes `shared_buffers`
  event_notification:
    signals: Asynchronous lifecycle notifications (e.g., `SIGTERM`, `SIGKILL`, `SIGCHLD`)
    eventfd: Lightweight kernel wait/notify counter mechanism
    posix_mq: Priority-based kernel message queues

server_concurrency_architectures:
  multi_process_worker_pool:
    example: PostgreSQL 🐘
    mechanism: Master process accepts client UDS connection, forks dedicated worker backend process per client
    pros: Fault isolation (one client crash doesn't break server), multi-core query execution
    coordination: POSIX shared memory (`shared_buffers`) with Reader-Writer Locks
  single_threaded_event_loop:
    example: Redis ⚡ / Node.js
    mechanism: Single process thread monitors thousands of client socket FDs via `epoll()`
    pros: Zero lock contention, guaranteed atomic operations, low context-switching overhead
    cons: Long $O(N)$ operations block the entire event loop for all clients

debugging_and_diagnostics_toolkit:
  strace:
    command: "strace -e trace=read,write,futex -p <PID>"
    purpose: Inspect system calls of a stuck or running process in real time
  lsof:
    command: "lsof -p <PID>"
    purpose: Map File Descriptors (0, 1, 2, 3...) to physical files, pipes, or sockets
  ss:
    command: "ss -x src /tmp/demo.sock"
    purpose: View Unix socket states and kernel queue buffer sizes (`Recv-Q` / `Send-Q`)
  ipcs:
    command: "ipcs -m"
    purpose: Check active shared memory segments and process attachment count (`nattch`)
  proc_filesystem:
    command: "ls -l /proc/<PID>/fd/"
    purpose: Inspect symbolic links to find raw socket inode numbers

