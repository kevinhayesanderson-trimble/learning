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


why kernel mediates ?:
- interface to control permissions
- mediation of I/O driver code

inode - 
inode is a data structure used in many file systems to represent a file or directory. It contains metadata about the file, such as its size, ownership, permissions, timestamps, and pointers to the actual data blocks on the storage device. Each file or directory has a unique inode number that allows the file system to locate and manage it efficiently.

inode structure:
- File type (regular file, directory, symbolic link, etc.)
- Permissions (read, write, execute)
- Owner and group information
- File size
- Timestamps (creation, modification, access)
- Pointers to data blocks (direct, indirect, double indirect, etc.)	

buffering in userspace - 
buffering in userspace refers to the practice of temporarily storing data in memory (buffers) within user-level applications before it is processed or sent to the kernel for further handling. This technique helps improve performance by reducing the number of system calls and context switches between user space and kernel space. By accumulating data in buffers, applications can perform I/O operations more efficiently, as they can read or write larger chunks of data at once rather than making frequent small requests to the kernel.

