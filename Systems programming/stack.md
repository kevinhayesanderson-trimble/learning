why should we have a stack mempry ? why not use heap memory for everything?
We should have stack memory because it provides a fast and efficient way to manage function calls and local variables. The stack operates in a last-in, first-out (LIFO) manner, which allows for quick allocation and deallocation of memory as functions are called and return. This makes it ideal for managing temporary data that is only needed during the execution of a function.
easy to deallocate: When a function is called, its local variables are pushed onto the stack, and when the function returns, the stack pointer is simply moved back, effectively deallocating the memory. This is much faster than managing memory on the heap, which requires more complex allocation and deallocation processes.

stack is not fixed size: The stack has a fixed size, which can help prevent memory leaks and fragmentation that can occur with heap memory. Since the stack is managed automatically by the operating system, it reduces the risk of memory management errors.

a single stack frame: Each function call creates a new stack frame, which contains the function's local variables and return address. This allows for easy tracking of function calls and their associated data, making it easier to debug and understand the program's flow. this can also grow and shrink as needed, allowing for efficient use of memory.

