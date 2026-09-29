# System Design — HLD + LLD (interview + enterprise)

Companion to `Papers.md` (distributed systems theory), `Cloud-native-OSS.md` (tools), `Hands-on.md` (DS builds).

| File | Use |
|---|---|
| **This file** | HLD framework, building blocks, problem catalog, resources |
| `LLD.md` | OOP, SOLID, patterns, machine-coding problems |
| `System-Design-Hands-on.md` | Timed drills + projects (HLD + LLD) |
| `Papers.md` | Depth when interviewer pushes on consistency, logs, geo, etc. |
| `Cloud-native-OSS.md` | What to name as concrete tech (Kafka vs SQS, etc.) |
| `Theory.md` | Volatility-based decomposition (Udi Dahan-style) |
| `API design.md` | REST / RPC / GraphQL notes |

---

## What “system design” means in interviews

| Layer | Asked as | Output | Bar |
|---|---|---|---|
| **HLD** | “Design Twitter / Uber / S3…” | Boxes, data flow, scale math, failures, tradeoffs | Senior+ core |
| **LLD** | “Design Parking Lot / Chess / Rate limiter classes” | Class diagram, APIs, patterns, working code | Many FAANG (esp. Amazon, Uber, Atlassian, some Meta) |
| **DS deep** | “How does your chat survive partition?” | Papers vocabulary from `Papers.md` | Staff / infra teams |

Staff twist: fewer “URL shortener from scratch”; more **ambiguous product**, multi-tenant, cost, multi-region, migration, operability.

---

## Meta resources (HLD)

| Resource | Why |
|---|---|
| [ByteByteGo — system-design-101](https://github.com/ByteByteGoHq/system-design-101) | Visual building blocks |
| Alex Xu — *System Design Interview* Vol 1 & 2 | Problem walkthroughs |
| [Grokking the System Design Interview](https://www.designgurus.io/course/grokking-the-system-design-interview) | Structured course |
| [donnemartin/system-design-primer](https://github.com/donnemartin/system-design-primer) | Free classic GitHub primer |
| [awesome-system-design](https://github.com/madd86/awesome-system-design) / similar lists | Index |
| DDIA — Kleppmann *Designing Data-Intensive Applications* | Best book bridge to `Papers.md` |
| [MIT 6.824](https://pdos.csail.mit.edu/6.824/) | Labs if you want implementation depth |
| Your `Papers.md` + `Cloud-native-OSS.md` | Depth + concrete tech |

---

## HLD interview framework (use every time)

Time-box ~35–45 min.

### 1. Clarify (3–5 min)
- Functional: core use cases, who is user, read vs write heavy
- Non-functional: QPS, latency p99, data size, retention, consistency, availability, cost
- Out of scope: explicit “not building X today”
- Soft vs hard constraints

### 2. Back-of-envelope (2–3 min)
- DAU/MAU → QPS (peak ~2–5× average)
- Storage = objects × size × retention × replication
- Bandwidth, cache hit assumptions
- Say numbers out loud; wrong-but-reasoned > no numbers

### 3. High-level API + entities (3–5 min)
- 3–7 endpoints or events
- Core entities + relationships
- Sync API vs async event where it matters

### 4. First architecture (5 min)
- Clients → edge/CDN → gateway → services → data stores → async
- Draw **one** happy path end-to-end before optimizing

### 5. Deep dives (15–20 min) — pick what interviewer cares about
- Data model / partition key
- Consistency & conflicts
- Cache strategy & invalidation
- Fan-out (push vs pull)
- Failure modes (AZ, region, poison, hot key)
- Scaling bottleneck + next fix

### 6. Wrap (2 min)
- Tradeoffs summary, what you’d monitor, what you’d do at 10×

**Staff add-ons:** multi-tenant isolation, cost, SLO/error budget, migration/rollback, compliance, “what we buy vs build”.

---

## HLD building blocks (must be fluent)

Cross-link `Papers.md` / `Cloud-native-OSS.md` for each.

| Block | Know | Typical choices |
|---|---|---|
| Load balancing | L4 vs L7, sticky, health | NLB/ALB, Envoy, nginx |
| CDN / edge | cache hierarchy, purge | CloudFront, Fastly |
| API gateway | auth, rate limit, routing | APIGW, Kong, Envoy |
| Compute | sync vs async, autoscale | k8s, Lambda, ECS |
| Cache | aside / through / ahead; stampede | Redis, CDN, local |
| DB OLTP | row store, indexes, FA | Postgres, Aurora, DynamoDB |
| DB OLAP | columnar, scan | ClickHouse, Redshift, Athena |
| Search | inverted index | OpenSearch, ES |
| Blob | immutable objects | S3, MinIO |
| Queue / bus | at-least-once, order | SQS, Kafka, Pulsar |
| Stream proc | windows, state | Flink, Kinesis Analytics |
| Coordination | locks, leader, config | etcd, ZK, DynamoDB |
| Idempotency | keys, outbox | see Helland / Hands-on P6 |
| Rate limit | token/leaky bucket, sliding | Redis, gateway |
| Fan-out | push, pull, hybrid | notifications, feeds |
| Geo | active-active vs active-passive | Global Tables, CRR |
| Observability | RED/USE, traces | OTel, Prometheus |

### Capacity heuristics (memorize order of magnitude)

| Thing | Ballpark |
|---|---|
| 1 day | ~10⁵ s |
| Read-heavy cache hit | aim 80–99% for hot keys |
| Single Redis instance | ~100k–1M simple ops/s (order-of) |
| Postgres primary | ~1k–10k simple writes/s before serious design (order-of; measure) |
| S3 | unlimited scale; design for request rate + prefixes |
| Kafka partition | ordered per partition; throughput scales with partitions |

---

## Consistency cheat sheet for HLD (say this calmly)

| Need | Pattern |
|---|---|
| Absolute uniqueness / money / inventory | single writer, serializable/consensus, or conditional write |
| Read-your-writes for one user | sticky session / session guarantee / read from leader |
| Feed / social | eventual OK; timeline fan-out tradeoffs |
| Multi-region low latency | PACELC: usually AP + conflict policy or region pinning |
| Exactly-once *effect* | idempotency + outbox; not “magic queue” |

Details: `Papers.md`, `Consistency.md`, `Replication.md`.

---

## HLD problem catalog (cover these)

Do each with the framework. Star = FAANG favorites.

### Foundations
1. ★ URL shortener  
2. ★ Rate limiter  
3. ★ Unique ID generator (Snowflake-style)  
4. Pastebin / file upload  
5. Key-value store (ties to `Hands-on.md` P4)

### Social / content
6. ★ News feed (FB/Instagram)  
7. ★ Twitter/X timeline + post  
8. YouTube / Netflix streaming (metadata + CDN)  
9. ★ WhatsApp / Slack chat  
10. ★ Notification system (push/email/SMS)

### Marketplace / location
11. ★ Uber / Lyft matching  
12. ★ Ticketmaster / book seats  
13. ★ Yelp / nearby search  
14. E-commerce product + cart + checkout (not full payments deep-dive unless asked)

### Infra / platform
15. ★ Web crawler  
16. ★ Search autocomplete  
17. ★ Distributed cache  
18. ★ Metrics / monitoring system (Prometheus-like)  
19. ★ Logging / analytics pipeline  
20. ★ Job scheduler / cron at scale  
21. ★ S3-like object store  
22. ★ Dropbox / Google Drive  
23. CDN  
24. DNS  

### Transactions / money / integrity
25. ★ Payment / wallet (ledger, idempotency)  
26. Stock exchange / matching engine (advanced)  
27. ★ Ad click aggregator  
28. ★ Flash sale / inventory

### Modern / staff-flavored
29. Multi-tenant SaaS feature flags  
30. Real-time collaboration (presence; OT/CRDT awareness)  
31. ML feature store / embedding retrieval (high level)  
32. IoT ingest + device shadow (your background edge)  
33. Multi-region active-active API  
34. Cost-aware redesign of a chatty microservice system  

For each solved problem, store: requirements, diagram link/ASCII, partition keys, failure table, AWS + OSS options.

---

## Mapping HLD → your DS vault

| Design topic | Read |
|---|---|
| Feed fan-out | Helland, Dynamo, TAO |
| Chat delivery | Kafka/log, at-least-once, idempotency |
| KV / session store | Dynamo, DynamoDB ATC 2022, consistent hashing |
| Object store | GFS, S3 ATD posts |
| Exactly-once pipeline | Kreps Log, MillWheel/Dataflow, outbox Hands-on |
| Global consistency | Spanner, COPS, PNUTS |
| Coordination | Chubby, ZK, etcd, Raft |
| Hot keys | Tail at Scale, DynamoDB GAC |

---

## Suggested HLD learning order

1. Building blocks (ByteByteGo 101 + primer GitHub)  
2. Framework reps on URL shortener + rate limiter + KV  
3. Social: feed + chat + notifications  
4. Infra: metrics, crawler, object store  
5. Hard: Uber, ticket booking, payments  
6. Staff: multi-region + multi-tenant + cost  
7. Parallel: keep coding DSA; do LLD from `LLD.md`

---

## Agent prompt (HLD)

```
Using System-Design.md + Papers.md + Cloud-native-OSS.md:
1) Interview me on designing <PROBLEM> (45 min simulation)
2) Enforce the HLD framework; interrupt if I skip numbers or failures
3) Push twice on consistency and hot keys
4) Ask for AWS and OSS alternatives
5) Score me: clarify / math / API / architecture / deep-dive / tradeoffs / communication
6) Give a remediation plan tied to vault notes
```
