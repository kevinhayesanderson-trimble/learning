# Hands-on projects — grasp DS concepts by building

Feed this file + `Papers.md` to an agent: extract concepts, then implement projects **in order**. Each project lists papers, acceptance tests, and the interview line you should be able to say after finishing.

**Stack defaults:** Python 3.12+ (you already have `server.py` / `client.py` in this vault) or Go. AWS projects use boto3 + LocalStack *or* a real AWS sandbox account.

**Rule:** finish acceptance criteria before stretch. Write a 10-line “what I would tell an interviewer” note at the end of each project.

---

## Track map

| # | Project | Concepts locked in | Papers / docs |
|---|---|---|---|
| 0 | Clocks lab | happens-before, Lamport, vector clocks | Lamport 1978, Mattern/Fidge |
| 1 | Single-node KV | API, framing, in-memory store | continue `distributed-key-value-store.md` |
| 2 | Persistent KV + WAL | durability, crash recovery | Gray recovery, LFS intuition |
| 3 | Consistent hashing ring | partition, rebalance, virtual nodes | Karger consistent hashing, Dynamo § |
| 4 | Mini-Dynamo (quorum KV) | N/R/W, sloppy quorum, vector clocks, hinted handoff | Dynamo, Gifford voting |
| 5 | Mini-Raft | leader election, log replication, commit index | Raft paper + student guide |
| 6 | Outbox + at-least-once | idempotency, exactly-once *effect* | Helland, Sagas, Kreps Log |
| 7 | CRDT counter/set | conflict-free merge | Shapiro CRDTs |
| 8 | AWS DynamoDB patterns | PK/SK, conditional writes, streams | DynamoDB ATC 2022 |
| 9 | AWS messaging pipeline | SQS/SNS, DLQ, retries | Kreps Log, Starbucks≠2PC |
| 10 | Capstone: multi-tenant event platform | staff narrative end-to-end | compose 4–9 |

---

## P0 — Clocks lab (1–2 evenings)

**Build:** library + CLI demo with 3 simulated nodes exchanging messages.

**Must implement**
- Lamport clock on send/recv/local event
- Vector clock; detect concurrent vs causal
- Print happens-before matrix for a scripted scenario

**Acceptance**
- Given a fixed message script, assert which events are concurrent
- Unit tests for tick / merge / compare

**Interview line:** “Without a shared clock I only have partial order; vector clocks tell me when two updates conflict vs one happened-before the other.”

---

## P1 — Single-node KV (extend existing)

**Build:** finish `distributed-key-value-store.md` Step 1+ 

**Must implement**
- GET/SET over UDP or TCP localhost
- Concurrent clients (thread or async)
- Clear wire protocol

**Acceptance**
- 2 clients SET/GET without corruption under sequential load
- Document bottleneck hypothesis (CPU vs lock vs copy)

**Interview line:** “Before distribution I know my single-node bottleneck and API semantics.”

---

## P2 — WAL + crash recovery

**Build:** append-only write-ahead log; replay on boot; optional snapshot.

**Must implement**
- Every SET appends log record then updates memtable
- Kill -9 mid-write; restart recovers last durable prefix
- fsync policy flag: every write vs batch

**Acceptance**
- Property test: random SET sequence + crash → recovered state ⊆ durable prefix
- Measure latency: fsync every write vs every 50ms

**Interview line:** “Durability is a fsync policy; availability vs durability is an explicit tradeoff.”

---

## P3 — Consistent hashing ring

**Build:** library that maps keys → nodes; add/remove node with minimal remaps.

**Must implement**
- Hash ring + virtual nodes (e.g. 100 vnodes/node)
- `get_preference_list(key, n=3)`
- Simulate join/leave; count keys that move

**Acceptance**
- Adding 4th node moves ~1/4 of keys (±10% with vnodes)
- Unit tests for preference list stability when ring unchanged

**Interview line:** “Consistent hashing + vnodes keeps rebalance proportional and avoids hot single-hash spots.”

---

## P4 — Mini-Dynamo (the money project)

**Build:** 3–5 process cluster; client talks to any coordinator.

**Must implement**
- Consistent hashing preference list (from P3)
- Quorum GET/PUT with configurable N, R, W
- Vector clocks on values; return siblings on conflict (or LWW flag)
- Hinted handoff *or* read repair (pick one, document the other)
- Failure injection: kill one replica mid-write

**Acceptance**
- With N=3,W=2,R=2: survive 1 node loss; read returns written value after quorum
- Conflict test: concurrent writes → siblings observable
- Chaos script: random kill/restart for 2 minutes; no silent data loss when W+R>N

**Interview line:** “Dynamo chooses AP: tunable quorums, version vectors, and client-side or coordinator merge — DynamoDB later moved partitions to Multi-Paxos for predictable strong consistency.”

**Stretch:** gossip membership (SWIM-lite).

---

## P5 — Mini-Raft (or MIT 6.824 labs)

**Build:** 3-node Raft; replicated log; state machine = KV.

**Options**
- From scratch (hard, high learning) — follow Raft paper Figure 2 strictly
- Or complete [MIT 6.824](https://pdos.csail.mit.edu/6.824/) Labs 2–3 in Go

**Must implement**
- Leader election + heartbeats
- Log replication; commit only when majority persisted
- Step down on higher term

**Acceptance**
- Jepsen-style linearizability checker on KV *or* published Raft test suite
- Partition leader → new leader elects; no dual commit for same index/term

**Interview line:** “Consensus gives a single linearizable log; I pay latency and availability during partitions — opposite of Dynamo’s default tradeoff.”

---

## P6 — Transactional outbox → at-least-once bus

**Build:** local DB (SQLite/Postgres) + worker publisher + consumer; bus can be in-proc queue first, then SQS.

**Must implement**
- Business write + outbox row in **one local transaction**
- Publisher polls outbox, publishes, marks sent
- Consumer idempotency key store (dedupe table)
- Simulate duplicate delivery

**Acceptance**
- Kill publisher after DB commit before publish → restart delivers exactly once *effect*
- Duplicate messages do not double-apply

**Interview line:** “Exactly-once is an application effect via idempotency + outbox; the network is at-least-once.”

---

## P7 — CRDT grow-only counter + OR-set

**Build:** 3 replicas; sync via anti-entropy (push full state or deltas).

**Must implement**
- G-Counter merge
- OR-Set add/remove with unique tags
- Concurrent add/remove scenario tests

**Acceptance**
- Any merge order converges to same state
- Show a case CRDT cannot express (e.g. global uniqueness) and what you’d use instead (consensus)

**Interview line:** “CRDTs buy coordination-free availability for mergeable types; uniqueness and money still need consensus or careful single-writer.”

---

## P8 — AWS DynamoDB patterns (sandbox)

**Build:** small service (Lambda or local script) against DynamoDB.

**Must implement**
- Single-table or clear PK/SK design for one domain (e.g. device → telemetry pointer)
- Conditional write for optimistic concurrency (`attribute_not_exists` / version)
- DynamoDB Streams → Lambda consumer with idempotency
- Demonstrate hot-key failure (write hammer one PK) and mitigation (write sharding / salt)

**Acceptance**
- Load test: document RCU/WCU, throttles, and fix
- Explain strong vs eventual consistent read in your API

**Papers:** DynamoDB ATC 2022 + original Dynamo (contrast leaderless vs Multi-Paxos partitions)

**Interview line:** “DynamoDB is not open Dynamo: per-partition Multi-Paxos, admission control, and predictable single-digit ms — I design keys so partitions stay balanced.”

---

## P9 — AWS messaging pipeline

**Build:** API → SNS → SQS (standard + FIFO variant) → worker → DLQ.

**Must implement**
- Retry with backoff; poison message to DLQ after N
- Idempotent handler (DynamoDB or Redis dedupe)
- Ordering experiment: show standard queue reordering; FIFO tradeoffs
- Optional: EventBridge Pipes or Step Functions for saga compensation

**Acceptance**
- Chaos: handler fails 50% → no loss, DLQ captures poison, metrics exported
- Cost sketch: msgs/day → SQS + Lambda $ estimate

**Interview line:** “I pick bus semantics deliberately: standard SQS for throughput, FIFO for partition-key order, Step Functions when the saga must be visible and compensatable.”

---

## P10 — Capstone (staff narrative)

**Build:** multi-tenant “ingest + process + query latest” platform on AWS.

**Example domain (pick one):**
- IoT device shadow + event log (fits your background)
- Multi-tenant webhooks delivery with retries
- Collaborative document metadata (not OT/CRDT full editor)

**Architecture minimum**
- API Gateway / ALB + compute (Lambda or ECS)
- DynamoDB for state + S3 for payloads
- SQS/Kinesis for async
- Multi-AZ by default; write a 1-page multi-Region *decision* (even if you don’t build it): RPO/RTO, conflict model, cost

**Deliverables**
1. Architecture diagram + threat/failure table (AZ loss, throttle, poison, duplicate, clock skew)
2. Load test numbers
3. 5-minute verbal design as if to interviewers
4. Explicit map: which paper concept → which AWS knob

**Acceptance**
- Another engineer can run `cdk deploy` / `terraform apply` / compose file and hit happy path
- You can answer: consistency choice, idempotency, hot key, DLQ, cost, what breaks at 10×

---

## How to use this with an agent

Prompt pattern:

```
Using Papers.md + Hands-on.md, teach me project P{N}:
1) Extract the concepts and paper claims I must internalize
2) Give a stepwise implementation plan for my stack
3) Generate acceptance tests first (TDD)
4) Review my code against acceptance criteria
5) Quiz me with staff interview follow-ups
```

Do **one** project per chat/session; store learnings back into topic notes (`Consistency.md`, `Replication.md`, etc.).

---

## Weekly cadence (staff prep)

| Week | Focus |
|---|---|
| 1 | P0–P2 + read mindset/time/CAP pages |
| 2 | P3–P4 + Dynamo paper |
| 3 | P5 (Raft) + Schneider SMR |
| 4 | P6–P7 + Helland / CRDT |
| 5 | P8–P9 + DynamoDB ATC 2022 + S3 posts |
| 6 | P10 + mock interviews (2 designs timed) |

---

## Related vault notes

- `Papers.md` — theory + classic systems path (GFS→Bigtable→Spanner→ZooKeeper→Kafka/Aurora) + AWS + OSS papers
- `Cloud-native-OSS.md` — enterprise open source / CNCF / AWS map + labs L1–L8
- `System-Design.md` / `LLD.md` / `System-Design-Hands-on.md` — HLD + LLD interview track
- `distributed-key-value-store.md` — P1 starter
- `Consistency.md` / `Replication.md` / `Partitioning.md` / `Transactions.md` — dump interview talking points here after each project
