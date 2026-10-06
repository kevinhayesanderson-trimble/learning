
_cache write-through_: when a thread updates the cache contents, we mirror the update back to the main memory

_cache-coherency protocols_:
- mechanism for dealing with reads and writes on memory and caches in a multiprocessor system
- write-back to the main memory with cache invalidation in other threads(listening to bus update messages) is one of the protocol

_coherency wall_:

_Escape analysis_:
Escape analysis consists of the compiler algorithms that decide whether a variable should be allocated on the heap instead of the stack.
the local variable shared by goroutines escaped to the heap
cost - go garbage collector collects in heap


