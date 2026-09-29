# The Golden Rules of Rust: Ownership, Borrowing, and Lifetimes

This document outlines the core memory management principles of Rust, mapping practical mental models to the official compiler rules.

---

## 1. Core Ownership

**The Practical Rules:**
* Every value is "owned" by a single variable, argument, struct, vector, etc., at a time.
* Reassigning the value to a variable, passing it to a function, or putting it into a vector *moves* the value. The old owner cannot be used to access the value anymore.
* When an owner goes out of scope, the value owned by it is dropped (cleaned up in memory).

**The Official Rust Rules:**
> * Each value in Rust has an owner.
> * There can only be one owner at a time.
> * When the owner goes out of scope, the value will be dropped.

**The Plain English Version:**
Every piece of data has exactly one "boss." If you hand the data to another variable or function, you are handing over the keys—you can no longer use it. When the boss's job is done (they go out of scope), the data goes in the trash immediately.

**Reference:** [The Rust Book, Ch 4.1: What is Ownership?](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html#ownership-rules)

---

## 2. Borrowing & References

**The Practical Rules:**
* You can create many read-only references to a value that exist at the same time.
* You cannot move a value while a reference to the value exists.
* You can make a writeable (mutable) reference to a value ONLY if there are no read-only references currently in use. One mutable reference to a value can exist at a time.
* You cannot mutate a value through the owner when any reference (mutable or immutable) to the value exists.

**The Official Rust Rules:**
> * At any given time, you can have *either* one mutable reference *or* any number of immutable references.
> * References must always be valid. (The compiler strictly enforces anti-aliasing during mutation).

**The Plain English Version:**
You can either have a room full of people reading a document, OR exactly one person locked in a room editing it. Never both. Furthermore, if you rent out your house (borrow it), you as the owner cannot sell it (move it) or remodel it (mutate it) until the renters leave.

**Reference:** [The Rust Book, Ch 4.2: References and Borrowing](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html#the-rules-of-references)

---

## 3. The `Copy` Trait Exception

**The Practical Rule:**
* Some types of values are copied instead of moved (e.g., numbers, booleans, characters, and arrays/tuples containing only copyable elements).

**The Official Rust Rule:**
> * If a type implements the `Copy` trait, variables that use it do not move, but rather are trivially copied, making them still valid after assignment to another variable.

**The Plain English Version:**
Simple, fixed-size data (like the number `5` or a boolean) is so cheap to photocopy that Rust just hands out identical photocopies instead of dealing with strict ownership transfers.

**Reference:** [The Rust Book, Ch 4.1: Stack-Only Data: Copy](https://doc.rust-lang.org/book/ch04-01-what-is-ownership.html#stack-only-data-copy)

---

## 4. Lifetimes & Validity

**The Practical Rules:**
* There cannot be references to a value when its owner goes out of scope.
* References to a value cannot outlive the value they refer to.

**The Official Rust Rules:**
> * References must always be valid.
> * The lifetime of a reference cannot exceed the lifetime of the value it borrows.

**The Plain English Version:**
You cannot hold a map to a house that has already been demolished. Rust physically will not compile your code if there is any chance a reference points to "dead" memory.

**References:** 
* [The Rust Book, Ch 4.2: Dangling References](https://doc.rust-lang.org/book/ch04-02-references-and-borrowing.html#dangling-references)
* [The Rust Book, Ch 10.3: Validating References with Lifetimes](https://doc.rust-lang.org/book/ch10-03-lifetime-syntax.html)

---

## 5. The Core Philosophy

**The Practical Rules:**
* These rules will dramatically change how you write code (compared to other languages).
* When in doubt, remember that Rust wants to minimize unexpected updates to data.

**The Rust Philosophy:**
Rust's strictness exists to prevent "Data Races" and memory unsafety without needing a garbage collector. A data race happens when two pointers access the same data at the same time, at least one is writing to it, and there is no synchronization.

**The Plain English Version:**
Rust forces you to think about the architecture of your data upfront. By proving to the compiler exactly who is looking at what and when, you eliminate entire classes of bugs (like unexpected mutations, double-frees, or memory leaks) before the code even runs.

**Reference:** [The Rust Book, Ch 16: Fearless Concurrency](https://doc.rust-lang.org/book/ch16-00-concurrency.html)