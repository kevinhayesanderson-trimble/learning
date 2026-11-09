[Computation Structures](https://learn.mit.edu/video-playlist/7859/mit-6-004-computation-structures-spring-2017)

https://ocw.mit.edu/courses/6-004-computation-structures-spring-2017/

File system is abstraction over the physical storage devices. It provides a way to organize and manage files and directories on a storage medium, such as a hard drive, SSD, or USB drive. The file system defines how data is stored, accessed, and manipulated, allowing users and applications to interact with files in a consistent manner.

Kernel- 
The kernel is the core component of an operating system that manages system resources and provides essential services to applications. It acts as a bridge between hardware and software, handling tasks such as process management, memory management, device drivers, and system calls. The kernel operates in a privileged mode, allowing it to control access to hardware and ensure the stability and security of the system.

Operating System (OS) -
An operating system is a software that manages computer hardware and software resources and provides common services for computer programs. It acts as an intermediary between users and the computer hardware, enabling users to interact with the system and run applications. The OS handles tasks such as process scheduling, memory management, file system management, device management, and user interface.

kernel vs OS -
The kernel is a part of the operating system, responsible for managing system resources and providing essential services. The operating system encompasses the kernel along with other components, such as user interfaces, libraries, and utilities, that provide a complete environment for running applications. In summary, the kernel is the core of the OS, while the OS includes the kernel and additional software that enables users to interact with the system.



Kernel responsibilities:
-  I/O devices
	- Disk
	- Networking
	- Keyboard
	- Display
	- Mouse
	- ....
- Process Management
	- Origination/launching
	- Lifecycle management
	- Signalling/IPC
	- Scheduling
		- Fairness
		- interactive process prioritized
		- Multicore
		- 
- Access to RAM ?
	Access to RAM is managed by the kernel through memory management techniques such as paging, segmentation, and virtual memory. The kernel allocates and deallocates memory for processes, ensuring that each process has its own isolated address space and preventing unauthorized access to other processes' memory.


why kernel mediates ?:
- interface to control permissions
- mediation of I/O driver code

inode - 
inode is a data structure used in many file systems to represent a file or directory. It contains metadata about the file, such as its size, ownership, permissions, timestamps, and pointers to the actual data blocks on the storage device. Each file or directory has a unique inode number that allows the file system to locate and manage it efficiently.

how to find inode number of a file:
To find the inode number of a file in a Unix-like operating system, you can use the `ls -i` command in the terminal. For example, running `ls -i filename` will display the inode number of the specified file. Alternatively, you can use the `stat` command, which provides detailed information about the file, including its inode number. For example, `stat filename` will show the inode number along with other file attributes.



inode structure:
- File type (regular file, directory, symbolic link, etc.)
- Permissions (read, write, execute)
- Owner and group information
- File size
- Timestamps (creation, modification, access)
- Pointers to data blocks (direct, indirect, double indirect, etc.)	

inode is used to directly seek to the data blocks of a file without having to traverse the directory structure, allowing for efficient file access and manipulation. The inode structure provides a way to manage and organize files on the storage medium, enabling the file system to locate and retrieve data quickly.


buffering in userspace - 
buffering in userspace refers to the practice of temporarily storing data in memory (buffers) within user-level applications before it is processed or sent to the kernel for further handling. This technique helps improve performance by reducing the number of system calls and context switches between user space and kernel space. By accumulating data in buffers, applications can perform I/O operations more efficiently, as they can read or write larger chunks of data at once rather than making frequent small requests to the kernel.

