we used to have co-operative multitasking, where one process would yield control to another process, allowing multiple processes to run seemingly simultaneously. However, this approach had limitations, as a single misbehaving process could monopolize the CPU and prevent other processes from running.

now we have preemptive multitasking, where the operating system can interrupt a running process and switch to another process, ensuring that all processes get a fair share of CPU time. This is achieved through the use of timers and interrupts, allowing the OS to regain control and manage process scheduling effectively.

timer interrupt - A timer interrupt is a signal sent to the CPU at regular intervals by a hardware timer. When the timer interrupt occurs, the operating system can pause the currently running process and switch to another process, allowing for better multitasking and responsiveness. This mechanism ensures that no single process can dominate the CPU, promoting fairness and efficiency in process scheduling.

flavours of unix:
    - GNU: A widely used open-source operating system kernel that is the foundation for many Linux distributions. It is known for its stability, security, and flexibility, making it popular for servers, desktops, and embedded systems.
    - BSD (Berkeley Software Distribution): A Unix-like operating system that originated from the University of California, Berkeley. It has several variants, including FreeBSD, OpenBSD, and NetBSD, each with its own focus on performance, security, and portability.
    - Solaris: A Unix operating system developed by Sun Microsystems (now owned by Oracle). It is known for its scalability, reliability, and advanced features, making it suitable for enterprise environments and high-performance computing.
