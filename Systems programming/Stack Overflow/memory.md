virtual memory:
- sub-system
- Address Space Layout Randomization (ASLR) is a computer security feature that randomly arranges the memory locations of key program data—such as the stack, heap, shared libraries, and the main executable base—every time a program runs
- 

Stack overflow:
- Stack overflow is a common software bug that occurs when a program writes more data to the stack
stack memory address space than it was allocated for. This can happen due to deep or infinite recursion, excessive memory allocation on the stack, or other programming errors. When a stack overflow occurs, it can lead to program crashes, unexpected behavior, or security vulnerabilities.

stack memory address space:
    - start address - 0x7fffe0000000
    - end address - 0x7fffffffffff
    why is it decreasing in decreasing order? - The stack grows downwards in memory, meaning that as more data is pushed onto the stack, the stack pointer moves to lower memory addresses. This is a common design choice in many computer architectures, allowing for efficient use of memory and enabling the stack to grow towards the heap, which typically grows upwards in memory.

    