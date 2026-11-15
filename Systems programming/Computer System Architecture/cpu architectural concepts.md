
---

## 1. The Processor State: Registers, Addressing, and ABI Conventions

RAM is orders of magnitude slower than a CPU core. Accessing main memory takes roughly 100 to 200 clock cycles, whereas accessing an internal CPU register takes less than 1 cycle. Therefore, **registers are the fundamental working set of the CPU.**

* **The Accumulator (`RAX`):** While modern ISAs allow arithmetic in any general-purpose register, legacy design persists. `RAX` acts as the universal return-value register for function calls.
* **Control Registers (`RIP` / `PC`):** The Instruction Pointer (`RIP` on x86-64, `PC` on ARM) holds the memory address of the next instruction. It cannot be modified via standard ALU math (`add rip, 4` is illegal); control-flow transfers (`jmp`, `call`, `ret`) explicitly overwrite `RIP`.
* **Stack Management (`RSP` and `RBP`):** The Stack Pointer (`RSP`) tracks the top of the call stack (which grows downward in memory). The Frame Pointer (`RBP`) anchors the current stack frame, providing a stable base offset (e.g., `[rbp - 24]`) to access local variables even as `RSP` fluctuates.
* **ABI Register Passing:** Application Binary Interfaces (ABIs)—such as the System V AMD64 ABI used on Linux/macOS—mandate that the first integer/pointer arguments to a function are passed via specific registers (`RDI`, `RSI`, `RDX`, `RCX`, `R8`, `R9`) rather than pushed onto the stack, vastly reducing memory traffic.

---

## 2. Memory Subsystem Physics: Endianness and Alignment

Hardware does not view memory as a formless stream; it imposes strict organizational rules that directly impact performance and data integrity.

### Endianness

Both x86-64 and ARM64 are **Little-Endian**, meaning the Least Significant Byte (LSB) is stored at the lowest memory address.

* *Engineering Impact:* When inspecting memory buffers or dealing with network protocols (which enforce **Big-Endian** Network Byte Order), failing to byte-swap values using instructions like `bswap` or conversion functions (`ntohs`) results in corrupted protocol headers and port numbers.

### Word Alignment and Padding

A memory access is *aligned* when an $N$-byte data type is stored at an address divisible by $N$.

* **Struct Padding:** Compilers automatically insert padding bytes into data structures to ensure fields land on natural alignment boundaries. Reordering struct fields from largest to smallest eliminates dead padding bytes, shrinking memory footprints and improving cache density.
* **Cache Line Splitting:** CPUs fetch RAM into L1 cache in aligned 64-byte chunks. Data that crosses a 64-byte cache line boundary forces the memory controller to issue two separate cache fetches instead of one, stalling execution.
* **Atomics and SIMD:** Unaligned memory accesses break hardware atomic operations (causing bus locks or faults) and trigger segmentation faults when executing SIMD vector instructions (`AVX`/`SSE`) that demand strict 16- or 32-byte alignment.

---

## 3. Microarchitecture: How CISC Meets RISC ($\mu$ops)

An Instruction Set Architecture (ISA) is merely the software contract. The actual microarchitecture determines how the silicon fulfills that contract.

* **x86-64 (CISC) vs. ARM64/RISC-V (RISC):** CISC architectures feature variable-length instructions (1 to 15 bytes) and memory-operand instructions (`add [rax], rbx`). RISC architectures enforce fixed-length instructions (e.g., 4 bytes) and a strict **Load/Store model**—you cannot perform math directly on RAM; data must be explicitly loaded into a register first.
* **Micro-operations ($\mu$ops):** Modern high-performance x86-64 CPUs do not execute complex x86 instructions natively. Instead, hardware decoders split incoming x86 instructions into simple, RISC-like **micro-ops ($\mu$ops)**.
* *Fast-Path:* ~90% of instructions (`mov`, `add`, `cmp`) decode directly into 1–4 $\mu$ops in hardware.
* *Slow-Path (Microcode ROM):* Complex or rare instructions (`cpuid`, `rep movsb`) trap into an internal microcode ROM sequencer that outputs a long sequence of $\mu$ops.



---

## 4. Pipelining and the Iron Law of Performance

CPU performance is governed by the Iron Law:

$$\text{Time per Program} = \text{Instructions per Program} \times \text{CPI} \times \text{Clock Cycle Time}$$

To maximize clock frequency (minimizing Cycle Time), processors break instruction execution into a multi-stage pipeline (typically Fetch, Decode, Execute, Memory Access, Write-Back). Instead of waiting for an instruction to completely finish before starting the next, instructions overlap. In an ideal pipeline, a completed instruction finishes every clock cycle ($\text{CPI} = 1$).

However, **pipeline hazards** inject `nop` bubbles that drive $\text{CPI} > 1$.

### Harvard Architecture at the L1 Cache Level

To prevent instructions (`IF` stage) and data (`MA` stage) from colliding on a single memory bus during the same clock cycle, CPU cores use a **Harvard architecture** at the L1 level, providing entirely separate, parallel paths for **L1 Instruction Cache (`L1i`)** and **L1 Data Cache (`L1d`)**.

---

## 5. Pipeline Hazards: Stalls, Bypassing, and Branches

When instructions interact in flight, hardware must resolve three classes of hazards:

### Data Hazards (RAW: Read-After-Write)

When Instruction 2 needs a register value being produced by Instruction 1, a dependency forms.

* **Stalling (Interlocks):** If hardware detects a dependency, it can freeze the pipeline stages, injecting idle bubbles. This destroys throughput.
* **Bypassing (Forwarding):** To eliminate bubbles, hardware routes ALU output data *directly* from the execution stage back into the input of the ALU for the next instruction, bypassing the register file entirely. Back-to-back math operations execute with zero stall cycles.
* **The Load-Use Latency Wall:** Bypassing cannot save a `Load` instruction (`lw` / `mov`) because data fetched from the L1 data cache does not arrive until the end of the Memory stage. An instruction immediately consuming that loaded data **must stall for at least one cycle**. This hardware reality is why compilers aggressively reorder code and insert independent work between memory loads and their usage.

### Control Hazards (Branches and Jumps)

Conditional branches (`if`/`else`, loops) alter the control flow. Because the CPU fetches sequential instructions speculatively before the branch condition is evaluated in hardware, a **taken branch** invalidates the instructions currently sitting in the pipeline.

* **Branch Flushing:** When a branch resolves as taken, the CPU must **flush (kill)** the speculatively fetched instructions, converting them into dead bubbles.
* **Mitigation:** Modern CPUs utilize dedicated early-stage hardware comparators to resolve branches faster, alongside sophisticated **Branch Predictors** that guess branch outcomes dynamically to keep deep pipelines full.