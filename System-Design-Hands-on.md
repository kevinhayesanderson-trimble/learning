# System Design hands-on — HLD + LLD drills

Practice reps. Theory lives in `System-Design.md`, `LLD.md`, `Papers.md`, `Cloud-native-OSS.md`.

**Rule:** timed, out loud (or with agent). Write a short postmortem after each drill.

---

## Weekly cadence (combine with DS Hands-on)

| Day | Focus |
|---|---|
| Mon | HLD timed problem (45 min) |
| Tue | LLD machine coding (90 min) or DSA |
| Wed | HLD hard problem + push on failures |
| Thu | LLD concurrency or extensibility drill |
| Fri | Mock (agent/human): full HLD or LLD |
| Sat | Deep dive: map one HLD to papers/OSS/AWS |
| Sun | Rest or light ByteByteGo / DDIA chapter |

---

## HLD drills

### H0 — Building-block flashcards (ongoing)
For each block in `System-Design.md`, answer in 60s:
- What problem it solves  
- Failure mode  
- One OSS + one AWS option  
- When *not* to use it  

### H1 — Framework muscle (week 1)
Timed 45 min each, no notes first pass:
1. URL shortener  
2. Rate limiter  
3. Pastebin  

**Acceptance:** used full framework; stated QPS + storage; named cache + DB; one failure scenario.

### H2 — Social / chat (week 2)
1. News feed  
2. Chat (1:1 + groups)  
3. Notification system  

**Acceptance:** fan-out decision explicit; delivery guarantees stated; online presence approach.

### H3 — Marketplace / booking (week 3)
1. Ticket booking (avoid double sell)  
2. Uber matching (simplified)  
3. Nearby places search  

**Acceptance:** inventory consistency strategy; geo index approach; hotspot handling.

### H4 — Infra platform (week 4)
1. Metrics system  
2. Distributed job scheduler  
3. S3-like object store (high level)  

**Acceptance:** cardinality / hot partition called out; durability story; multi-tenant or ACL mention.

### H5 — Money & integrity (week 5)
1. Payment ledger / wallet  
2. Flash sale  
3. Ad click aggregator  

**Acceptance:** idempotency keys; exactly-once *effect*; ledger immutability.

### H6 — Staff scenarios (week 6+)
1. Multi-region active-active API for session store  
2. Multi-tenant SaaS (noisy neighbor)  
3. Migrate monolith checkout → services without downtime  
4. Cut cost 40% on a chatty event pipeline  

**Acceptance:** RPO/RTO; cost dimension; migration/rollback; what you’d **not** build.

### H7 — “Interviewer push” checklist
After any HLD, answer all:
- [ ] Hot key?  
- [ ] AZ failure?  
- [ ] Duplicate message?  
- [ ] Clock skew?  
- [ ] Schema change?  
- [ ] p99 latency regression?  
- [ ] Cost at 10×?  

---

## HLD deliverable template (paste into notes after each)

```md
## Design: <name> — <date>
### Requirements
### Numbers
### API
### Diagram (ASCII)
### Data model / keys
### Deep dives
### Failures
### OSS options
### AWS options
### Tradeoffs
### What I'd monitor
### Links to Papers.md concepts
```

---

## LLD drills

### L0 — Pattern kata (do once each)
Implement in your interview language (small, tested):
1. Strategy (payment methods)  
2. Factory (vehicle/document types)  
3. Observer (stock ticker / event bus)  
4. State (vending or order)  
5. Decorator (stream/reader or HTTP middleware)  
6. Singleton (thread-safe) — then rewrite with DI and delete singleton  

### L1 — Classics (90 min blank repo each)
1. Parking lot  
2. LRU cache (thread-safe optional stretch)  
3. Rate limiter (token bucket + per-key)  

**Acceptance:** README with class diagram; `test` passes; add “EV spots” or “new vehicle type” with ≤1 file of core edits.

### L2 — State-heavy
1. Elevator  
2. Snake & ladder **or** tic-tac-toe + AI stub  
3. Traffic light intersection  

**Acceptance:** explicit state machine; no boolean soup.

### L3 — Domain-rich
1. BookMyShow (in-memory seats + booking)  
2. Splitwise  
3. Library management  

**Acceptance:** prevent double-booking; clear service boundaries.

### L4 — Concurrent machine coding
1. Blocking queue  
2. Thread pool  
3. In-memory pub-sub  

**Acceptance:** stress test with N producers/consumers; no lost messages under happy path; document locking.

### L5 — Bridge HLD↔LLD
Implement **classes** for a piece of an HLD you already did:
1. URL shortener encoder + storage interface  
2. Notification dispatcher with channel plugins  
3. Idempotent payment command handler + outbox interface  

**Acceptance:** can swap Redis vs memory repository via interface.

---

## Combined capstone (staff narrative)

**C1 — Design + code a slice**
1. HLD: multi-tenant webhook delivery (45 min)  
2. LLD: implement dispatcher + retry policy + idempotency store (half day)  
3. Map to AWS: API GW + SQS + Lambda + DynamoDB (`Cloud-native-OSS.md`)  
4. Write failure table + SLO  

**C2 — Design + code parking→booking bridge**
1. LLD parking lot  
2. HLD “parking SaaS” multi-city  
3. Explain what stays in monolith vs extracts first  

---

## Scoring rubric (self / agent)

### HLD (score 1–5 each)
- Clarifying questions  
- Numbers  
- Clean API  
- End-to-end architecture  
- Data model / partitioning  
- Consistency & failures  
- Tech justification (OSS + AWS)  
- Communication / time management  

**Target senior:** ≥4 average. **Staff:** ≥4 and strong on failures/cost/migration.

### LLD (score 1–5 each)
- Requirements  
- Model clarity  
- SOLID / extensibility  
- Working code  
- Tests  
- Handling new requirement  
- Concurrency (if in scope)  

---

## Agent prompts

### Full HLD mock
```
You are a FAANG interviewer. Using System-Design.md, run a 45-minute design of <PROBLEM>.
Be adversarial on numbers, consistency, and hot keys.
I am the candidate — wait for my answers.
End with rubric scores and a 3-item remediation plan pointing to vault files.
```

### Full LLD mock
```
You are an interviewer for a 90-minute machine coding round.
Problem: <PROBLEM> from LLD.md.
Give requirements in stages. After my class design, require tests then implementation.
Inject one new requirement at minute 60 that should be an OCP win.
Score per System-Design-Hands-on.md rubric.
```

### Teach then drill
```
Using System-Design.md + LLD.md + Papers.md:
Teach the minimum for <TOPIC>, then give a 30-min drill, then quiz me.
```

---

## Mapping to other vault tracks

| After this drill… | Also do |
|---|---|
| HLD KV / shortener | `Hands-on.md` P3–P4 |
| HLD chat / pipeline | `Hands-on.md` P6, P9 |
| HLD object store | Papers GFS/S3 + MinIO lab in `Cloud-native-OSS.md` |
| LLD rate limiter | HLD rate limiter same week |
| LLD + outbox | DS Hands-on P6 |

---

## Related files

- `System-Design.md` — HLD framework + problem catalog  
- `LLD.md` — patterns + LLD catalog  
- `Papers.md` — DS depth  
- `Cloud-native-OSS.md` — concrete tech  
- `Hands-on.md` — implement DS primitives  
- `Theory.md` — volatility decomposition  
- `API design.md` — API style notes  
