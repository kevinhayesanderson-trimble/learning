commands:
- apropos pid
	- to get all the systemcall wrapper functions 
- man function
	- to get the definition of the syscall wrapper function
- man pages sections
	- 1 - shell commands
	- 2 - System Calls - invoked by c code 
	- 3 - std lib subroutines - executed purely in user space
	- 4 - Special files - usually found in /dev
	- 5 - File formats and conventions - e.g. /etc/passwd
	- 6 - Games and screensavers
	- 7 - Miscellaneous
	- 8 - System administration commands
- strace -p pid
	- to trace the system calls made by a process with the given pid
	- -e option to specify which system calls to trace, e.g. strace -e open,read,write -p pid
	- -f option to follow child processes created by the traced process

- stat
	- to get the file status information, including the inode number, file size, permissions, and timestamps

- ltrace , uftrace , bpftrace :
	- to trace the library calls made by a process with the given pid
	- ltrace -p pid
	- uftrace -p pid
	- bpftrace -e 'tracepoint:syscalls:sys_enter_openat { printf("%s\n", str(args->filename)); }' -p pid


POSIX:
- POSIX is a family of standards specified by the IEEE for maintaining compatibility between operating systems.

	