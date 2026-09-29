# Low-Level Design (LLD) / Machine Coding

Class design, OOP, patterns, extensible APIs — what Amazon/Uber/Atlassian/many FAANG rounds call “LLD” or “machine coding”.

Companion: `System-Design.md` (HLD), `System-Design-Hands-on.md` (drills), `Theory.md` (volatility decomposition).

---

## What interviewers want

| Signal | Evidence |
|---|---|
| Clear domain model | Entities, responsibilities, no god class |
| Extensibility | New requirement doesn’t rewrite everything (OCP) |
| Correct abstractions | Interfaces at volatility boundaries (`Theory.md`) |
| Working code | Compiles/runs; happy path + 1–2 edge cases |
| Concurrency awareness | When asked: thread safety, locking grain |
| Tradeoffs | “I chose X pattern because Y” |

Not required: every Gang of Four pattern by name. Required: apply 3–5 patterns *when they earn their keep*.

---

## Meta resources (LLD)

| Resource | Why |
|---|---|
| Head First Design Patterns | Intuition |
| Gang of Four (Gamma et al.) | Reference |
| *Clean Architecture* / *Clean Code* — Uncle Bob | Boundaries (pair with `Theory.md`) |
| [Grokking the Low Level Design Interview](https://www.designgurus.io/) | Structured LLD problems |
| [prasadgujar/low-level-design-primer](https://github.com/prasadgujar/low-level-design-primer) (and forks) | Problem lists |
| Refactoring.Guru — Design Patterns | Visual catalog |
| Effective Java / Effective Python / similar | Language idioms for machine coding |

---

## LLD interview framework (~60–90 min machine coding; ~45 min whiteboard)

### 1. Requirements (5–10 min)
- Actors, use cases, explicit non-goals
- Ask: single machine vs multi? in-memory vs DB? concurrent users?
- List entities on paper first

### 2. Core objects & relations (5–10 min)
- Nouns → classes; verbs → methods
- UML-ish: associations, inheritance only where “is-a” is stable
- Prefer composition over inheritance

### 3. APIs (5 min)
- Public methods of service/facade
- Error cases (exceptions / Result types)

### 4. Implement vertical slice (rest of time)
- One happy path end-to-end first
- Then strategy/plugins for variation
- Unit tests if environment allows

### 5. Extensions (interviewer will add)
- “Now support X” — show OCP: new class, minimal edits

---

## SOLID (must explain with your code)

| Letter | Meaning | Interview example |
|---|---|---|
| **S**ingle responsibility | One reason to change | `PricingService` ≠ `NotificationService` |
| **O**pen/closed | extend without modify | new `PaymentMethod` via interface |
| **L**iskov | subtypes substitutable | don’t break `Bird.fly()` with Penguin |
| **I**nterface segregation | small interfaces | `Readable` vs fat `Repository` |
| **D**ependency inversion | depend on abstractions | inject `Lock`, don’t `new RedisLock()` |

Also: **DRY**, **KISS**, **YAGNI** — don’t pattern-spam.

---

## Design patterns — interview priority

### Creational
- **Factory / Abstract Factory** — create families (e.g. vehicle types)
- **Builder** — complex immutable config
- **Singleton** — use sparingly; prefer DI; mention thread-safe init if used
- **Prototype** — rare in interviews

### Structural
- **Adapter** — wrap third-party API
- **Decorator** — add behavior (auth, logging) without subclass explosion
- **Facade** — simplify subsystem for client
- **Proxy** — lazy load, remote, access control
- **Composite** — trees (file system, org chart)

### Behavioral
- **Strategy** — interchangeable algorithms (fare, payment, compression)
- **Observer / Pub-Sub** — events (in-process)
- **State** — explicit state machines (vending, order lifecycle)
- **Command** — undo/redo, job queue items
- **Template method** — fixed algorithm skeleton
- **Iterator** — custom collections
- **Chain of responsibility** — middleware / handlers
- **Mediator** — chat room / air traffic-style coupling reduction

### Concurrency (when asked)
- Thread-safe singleton / pool
- Producer-consumer
- Read-write lock
- Actor-ish “single thread owns state”

### Enterprise / application
- Repository, Unit of Work
- DTO vs domain entity
- MVC / Hexagonal / Clean Architecture ports & adapters
- Outbox (ties to DS Hands-on P6)

---

## Volatility & LLD (from `Theory.md`)

Decompose on **what changes**, not functions:
- Managers = use-case / workflow volatility  
- Engines = business rule volatility  
- Don’t mirror org chart or “Controller/Service/Repo” blindly if it fights change  

In interviews: say “I’ll isolate the rule that will change when we add vehicle types” → Strategy/Factory.

---

## LLD problem catalog

### Classic machine coding
1. ★ Parking lot  
2. ★ Elevator / lift  
3. ★ Library management  
4. ★ Movie ticket booking (BookMyShow) — also bridges HLD  
5. ★ Snake & ladder / chess / tic-tac-toe / card game  
6. ★ Vending machine  
7. ★ ATM  
8. ★ Traffic light / intersection  
9. ★ Hotel booking  
10. ★ Splitwise / expense sharing  

### Concurrency / systems-flavored LLD
11. ★ Rate limiter (token bucket classes) — bridge to HLD  
12. ★ Thread-safe cache (LRU) with TTL  
13. ★ Blocking queue / thread pool  
14. ★ Pub-sub in process  
15. ★ Connection pool  
16. ★ Scheduler / cron in-process  

### Extensibility drills
17. ★ Logging framework  
18. ★ Notification system (email/SMS/push plugins)  
19. ★ Payment plugin system  
20. ★ Rule engine (discount coupons)  
21. ★ File system (composite)  
22. ★ Unix find / filter chain  

### Amazon-style favorites
23. ★ Online shopping cart + inventory  
24. ★ Warehouse / inventory allocate  
25. ★ Meeting room / calendar booking  
26. ★ Food order / delivery agent assignment (simplified)

For each: class diagram, key interfaces, sequence for 1 use case, extension point for “new type”.

---

## Language tips (pick one for interviews)

| Language | LLD tips |
|---|---|
| Java | Interfaces, enums, concurrent collections; most LLD content is Java-shaped |
| C# | Similar to Java; good DI story |
| Python | Protocols/ABC, dataclasses; be explicit about threads/GIL if concurrent |
| C++ | ownership, smart pointers; rarer for pure LLD rounds |
| Go | interfaces are implicit; less classic GoF; still do Strategy via interfaces |

Practice **in the language of the interview loop**.

---

## Anti-patterns (call out if you catch yourself)

- God class (`GameManager` does everything)  
- Deep inheritance for variation (use Strategy/composition)  
- Premature singleton everywhere  
- Anemic domain + all logic in one service (sometimes OK — justify)  
- Pattern tourism (Decorator wrapping Decorator wrapping…)  

---

## Suggested LLD learning order

1. SOLID + composition > inheritance (1–2 days of deliberate examples)  
2. Strategy, Factory, Observer, State, Singleton-done-right  
3. Parking lot + LRU cache + rate limiter (code all three)  
4. Elevator or chess (state complexity)  
5. BookMyShow or Splitwise (richer domain)  
6. Concurrent: blocking queue + thread-safe cache  
7. Timed: 90 min parking lot from blank repo  

---

## Agent prompt (LLD)

```
Using LLD.md + Theory.md:
1) Give me requirements for <PROBLEM> as an interviewer
2) I propose classes; critique against SOLID and volatility
3) Then TDD: generate tests first; I implement
4) Midway add a new requirement that should only need a new class
5) Score: model / extensibility / code quality / tests / communication
```
