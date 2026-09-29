# Best papers in distributed systems

Curated from primary sources plus: [Murat Demirbas foundational list](http://muratbuffalo.blogspot.com/2021/02/foundational-distributed-systems-papers.html), [alexprut Essential/Virtuoso](https://github.com/alexprut/distributed-systems), [Papers We Love — distributed_systems](https://github.com/papers-we-love/papers-we-love/tree/main/distributed_systems), [Heidi Howard consensus reading list](https://github.com/heidihoward/distributed-consensus-reading-list), [Paper Trail — theory for engineers](https://www.the-paper-trail.org/post/2014-08-09-distributed-systems-theory-for-the-distributed-systems-engineer/), [dancres reading list](https://dancres.github.io/Pages/), [awesome-distributed-systems](https://github.com/yeshengm/awesome-distributed-systems), [All Things Distributed](https://www.allthingsdistributed.com/articles.html).

---

## First principles (concepts before systems)

Mental model first — these change how you think. Read before diving into Paxos/Dynamo.

### Mindset / why distributed is hard

- **The Fallacies of Distributed Computing** — Deutsch / Rotem-Gal-Oz
  - network is reliable; latency is zero; bandwidth is infinite; network is secure; topology doesn't change; one administrator; transport cost is zero; network is homogeneous
  - https://en.wikipedia.org/wiki/Fallacies_of_distributed_computing
- **A Note on Distributed Computing** — Waldo, Wyant, Wollrath, Kendall (1994)
  - remote ≠ local; you cannot hide the network behind an RPC abstraction forever
  - https://scholar.harvard.edu/files/waldo/files/waldo-94.pdf
- **Notes on Distributed Systems for Young Bloods** — Hodges
  - https://www.somethingsimilar.com/2013/01/14/notes-on-distributed-systems-for-young-bloods/
- **Distributed Systems for Fun and Profit** — Mikito Takada (short book)
  - https://book.mixu.net/distsys/
- **Hints for Computer Systems Design** — Lampson (1983)
  - https://www.microsoft.com/en-us/research/publication/hints-for-computer-system-design/
- **End-to-End Arguments in System Design** — Saltzer, Reed, Clark (1984)
  - https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf
- **Safety vs liveness** — nothing bad ever happens vs something good eventually happens
  - vocabulary for every later impossibility / protocol paper

### Time, causality, global state

- **Time, Clocks, and the Ordering of Events in a Distributed System** — Lamport (1978)
  - happens-before, logical clocks — *the* DS primer
  - https://lamport.azurewebsites.net/pubs/time-clocks.pdf
- **Timestamps in Message-Passing Systems That Preserve the Partial Ordering** — Fidge (1988) / **Virtual Time and Global States** — Mattern (1989)
  - vector clocks
  - http://courses.csail.mit.edu/6.852/01/papers/VirtTime_GlobState.pdf
- **Distributed Snapshots: Determining Global States of a Distributed System** — Chandy & Lamport (1985)
  - consistent cuts without stopping the world
  - https://www.microsoft.com/en-us/research/publication/distributed-snapshots-determining-global-states-of-a-distributed-system/
- **Practical Uses of Synchronized Clocks in Distributed Systems** — Liskov (1991)
- **There is No Now** — Sheehy (ACM Queue 2015) — exposition
- **Why Logical Clocks are Easy** — Baquero & Preguiça (ACM Queue 2016)
- **Hybrid Logical Clocks** — Kulkarni et al. (2014)
  - https://cse.buffalo.edu/~demirbas/publications/hlc.pdf

### Impossibility & models of failure / time

- **Two Generals / Coordinated Attack** — folk theorem
  - https://en.wikipedia.org/wiki/Two_Generals%27_Problem
- **The Byzantine Generals Problem** — Lamport, Shostak, Pease (1982)
  - https://lamport.azurewebsites.net/pubs/byz.pdf
- **Reaching Agreement in the Presence of Faults** — Pease, Shostak, Lamport (1980)
  - https://www.microsoft.com/en-us/research/wp-content/uploads/2016/12/Reaching-Agreement-in-the-Presence-of-Faults.pdf
- **Impossibility of Distributed Consensus with One Faulty Process (FLP)** — Fischer, Lynch, Paterson (1985)
  - async + crash → consensus cannot be both always safe and always terminating
  - https://groups.csail.mit.edu/tds/papers/Lynch/jacm85.pdf
- **Consensus in the Presence of Partial Synchrony** — Dwork, Lynch, Stockmeyer (1988)
  - the model that makes practical consensus possible
  - https://groups.csail.mit.edu/tds/papers/Lynch/jacm88.pdf
- **Unreliable Failure Detectors for Reliable Distributed Systems** — Chandra & Toueg (1996)
  - how to circumvent FLP with oracles (◇P, Ω, …)
  - https://www.cs.utexas.edu/~lorenzo/corsi/cs380d/papers/p225-chandra.pdf
- **Harvest, Yield, and Scalable Tolerant Systems** — Fox & Brewer (1999)
  - practical CAP / graceful degradation
  - https://groups.csail.mit.edu/tds/papers/Brewer/HB1.pdf
- **Brewer's conjecture and the feasibility of CAP** — Gilbert & Lynch (2002)
  - formal CAP
  - https://www.glassbeacon.com/public/html/events/W21CSOC128/pdf/p30-gilbert.pdf
- **Towards robust distributed systems** — Brewer PODC keynote (2000)
  - https://www.cs.berkeley.edu/~brewer/cs262b-2004/PODC-keynote.pdf
- **CAP Twelve Years Later** — Brewer (2012)
  - https://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed/
- **PACELC** — Abadi (2012)
  - if Partition: A vs C; Else: Latency vs Consistency
  - https://www.cs.umd.edu/~abadi/papers/abadi-pacelc.pdf
- **A Hundred Impossibility Proofs for Distributed Computing** — Lynch (survey flavor)
  - https://github.com/papers-we-love/papers-we-love/blob/main/distributed_systems/a-hundred-impossibility-proofs-for-distributed-computing.pdf

### Fault hierarchy (internalize)

- crash-stop → omission → Byzantine (harder up the stack)
- sync / partially sync / async time models
- quorum intersection → single-copy serializability
- atomic broadcast ≡ consensus (same hardness)

---

## Consensus & state-machine replication

- **Implementing Fault-Tolerant Services Using the State Machine Approach** — Schneider (1990)
  - *the* framing: replicate a deterministic state machine
  - https://www.cs.cornell.edu/fbs/publications/SMSurvey.pdf
- **How to Build a Highly Available System Using Consensus** — Lampson (1996)
  - https://www.microsoft.com/en-us/research/publication/how-to-build-a-highly-available-system-using-consensus/
- **Viewstamped Replication** — Oki & Liskov (SOSP 1988)
  - https://www.pmg.csail.mit.edu/papers/vr.pdf
- **Viewstamped Replication Revisited** — Liskov & Cowling (2012)
  - https://pmg.csail.mit.edu/papers/vr-revisited.pdf
- **The Part-Time Parliament (Paxos)** — Lamport (1998 / 2001)
  - https://lamport.azurewebsites.net/pubs/lamport-paxos.pdf
- **Paxos Made Simple** — Lamport (2001)
  - https://lamport.azurewebsites.net/pubs/paxos-simple.pdf
- **Paxos Made Moderately Complex** — van Renesse & Altinbuken (2015)
- **Paxos Made Live** — Chandra, Griesemer, Redstone (2007)
  - https://www.cs.utexas.edu/users/lorenzo/corsi/cs380d/papers/paper2-1.pdf
- **Practical Byzantine Fault Tolerance (PBFT)** — Castro & Liskov (OSDI 1999)
  - https://www.microsoft.com/en-us/research/wp-content/uploads/2017/01/p398-castro-bft-tocs.pdf
- **In Search of an Understandable Consensus Algorithm (Raft)** — Ongaro & Ousterhout (2014)
  - https://raft.github.io/raft.pdf
- **There Is More Consensus in Egalitarian Parliaments (EPaxos)** — Moraru, Andersen, Kaminsky (SOSP 2013)
  - https://www.cs.cmu.edu/~dga/papers/epaxos-sosp2013.pdf
- **Flexible Paxos** — Howard, Malkhi, Spiegelman (2016)
  - quorum intersection revisited
  - https://arxiv.org/abs/1608.06696
- **The Chubby Lock Service** — Burrows (2006)
  - https://research.google/pubs/pub27897/
- **ZooKeeper: Wait-free coordination for Internet-scale systems** — Hunt et al. (2010)
  - https://www.usenix.org/legacy/event/usenix10/tech/full_papers/Hunt.pdf
- **Zab: High-performance broadcast for primary-backup systems** / A simple totally ordered broadcast protocol
  - https://github.com/papers-we-love/papers-we-love/blob/main/distributed_systems/zab-high-performance-broadcast-for-primary-backup-systems.pdf
- **Chain Replication for Supporting High Throughput and Availability** — van Renesse & Schneider (OSDI 2004)
  - https://www.usenix.org/legacy/events/osdi04/tech/full_papers/renesse/renesse.pdf
- Virtual Synchrony / ISIS — Birman et al. (group membership + reliable multicast)

---

## Consistency models & weak replication

- **Linearizability** — Herlihy & Wing (1990)
  - https://cs.brown.edu/~mph/HerlihyW90/p463-herlihy.pdf
- **Session Guarantees for Weakly Consistent Replicated Data** — Terry et al. (1994)
  - read-your-writes, monotonic reads/writes, causal
  - https://www.cs.utexas.edu/~dahlin/Classes/GradOS/papers/SessionGuaranteesPDIS.pdf
- **Managing Update Conflicts in Bayou** — Terry et al. (SOSP 1995)
  - https://www.cs.utexas.edu/~lorenzo/corsi/cs380d/papers/p172-terry.pdf
- **Optimistic Replication** — Saito & Shapiro (2005 survey)
  - https://pages.cs.wisc.edu/~remzi/Classes/739/Spring2004/Papers/optimistic-survey.pdf
- **Eventually Consistent** — Vogels (2008)
  - https://www.allthingsdistributed.com/2008/12/eventually_consistent.html
- **Conflict-free Replicated Data Types (CRDTs)** — Shapiro, Preguiça, Baquero, Zawirski (2011)
  - https://hal.inria.fr/inria-00555588/document
- **Consistency Analysis in Bloom: a CALM and Collected Approach** — Alvaro et al. (CIDR 2011)
  - when coordination is unnecessary
- **Life Beyond Distributed Transactions: an Apostate's Opinion** — Helland (CIDR 2007)
  - https://queue.acm.org/detail.cfm?id=3025012
- **Data on the Outside versus Data on the Inside** — Helland
- **Highly Available Transactions (HAT)** — Bailis et al.
  - https://www.vldb.org/pvldb/vol7/p181-bailis.pdf
- **Consistency Without Borders** / baseball examples — Whittaker
  - https://mwhittaker.github.io/consistency_in_distributed_systems/
- **The Dangers of Replication and a Solution** — Gray et al.
  - eager vs lazy replication tradeoffs

---

## Placement, membership, gossip

- **Consistent Hashing and Random Trees** — Karger et al. (1997)
  - foundation of Dynamo/Cassandra ring
  - https://www.cs.princeton.edu/courses/archive/fall09/cos518/papers/chash.pdf
- **Chord: A Scalable Peer-to-peer Lookup Service** — Stoica et al. (SIGCOMM 2001)
  - https://pdos.csail.mit.edu/papers/chord:sigcomm01/chord_sigcomm.pdf
- **Pastry** / **Kademlia** — DHT lineage
- **Epidemic Algorithms for Replicated Database Maintenance** — Demers et al. (1987)
- **SWIM: Scalable Weakly-consistent Infection-style Process Group Membership** — Das, Gupta, Motivala
  - https://research.cs.cornell.edu/projects/Quicksilver/public_pdfs/SWIM.pdf
- **Self-stabilizing Systems in Spite of Distributed Control** — Dijkstra (1974)

---

## Replication & storage systems (industry classics)

- **The Google File System (GFS)** — Ghemawat, Gobioff, Leung (2003)
  - https://research.google/pubs/pub51/
- **Bigtable** — Chang et al. (2006)
  - https://research.google/pubs/pub27898/
- **Dynamo** — DeCandia et al. (2007)
  - https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf
  - announcement: https://www.allthingsdistributed.com/2007/10/amazons_dynamo.html
- **Cassandra** — Lakshman & Malik (2009)
  - https://www.cs.cornell.edu/projects/ladis2009/papers/lakshman-ladis2009.pdf
- **PNUTS** — Cooper et al. (Yahoo, 2008)
- **Megastore** — Baker et al. (2011)
- **Spanner** — Corbett et al. (2012)
  - https://research.google/pubs/pub39966/
- **Spanner, TrueTime and the CAP Theorem** — Brewer (2017)
  - https://cloud.google.com/spanner/docs/whitepapers/SpannerAndCap
- **F1** — Shute et al. (2013)
- **Amazon Aurora** (SIGMOD 2017) + consensus-avoidance paper (SIGMOD 2018)
  - https://www.allthingsdistributed.com/files/p1041-verbitski.pdf
  - https://dl.acm.org/citation.cfm?id=3183713.3196937
- **Amazon Distributed Computing Manifesto** (1998 / published 2022)
  - https://www.allthingsdistributed.com/files/amazon-distributed-computing-manifesto-1998.pdf
- **TAO** — Bronson et al. (Facebook, 2013)
- **COPS** — Don't Settle for Eventual (causal geo) — Lloyd et al. (SOSP 2011)
  - https://www.cs.cmu.edu/~dga/papers/cops-sosp2011.pdf

---

## Transactions & concurrency

- **On Optimistic Methods for Concurrency Control** — Kung & Robinson (1981)
- **Concurrency Control and Recovery in Database Systems** — Bernstein, Hadzilacos, Goodman (1987 book)
- **Notes on Distributed Databases** / 2PC lineage — Gray, Lampson & Sturgis
- **Sagas** — Garcia-Molina & Salem (1987)
  - https://www.cs.cornell.edu/andru/cs711/2002fa/reading/sagas.pdf
- **Percolator** — Peng & Dabek (2010)
  - https://research.google/pubs/pub36726/
- **Calvin** — Thomson et al. (2012)
  - https://cs.yale.edu/homes/thomson/publications/calvin-sigmod12.pdf

---

## Compute / logs / streaming

- **MapReduce** — Dean & Ghemawat (2004)
  - https://research.google/pubs/pub62/
- **Dryad** — Isard et al. (2007)
- **Dremel** — Melnik et al. (2010)
- **Kafka: a Distributed Messaging System for Log Processing** — Kreps et al. (2011)
  - https://notes.stephenholiday.com/Kafka.pdf
- **MillWheel** → **Dataflow / Beam model** — Akidau et al.
  - https://research.google/pubs/pub43864/
- **Resilient Distributed Datasets (RDD / Spark)** — Zaharia et al. (NSDI 2012)
- **The Tail at Scale** — Dean & Barroso (2013)
  - https://research.google/pubs/pub40801/
- **Lessons from Giant-Scale Services** — Brewer (2001)
- **On Designing and Deploying Internet-Scale Services** — Hamilton (LISA 2007)
  - https://mvdirona.com/jrh/talksandpapers/jamesrh_lisa.pdf
- **Dapper** — large-scale distributed tracing (Google)
  - https://research.google/pubs/pub36356/

---

## Failure, complexity, operations

- **Why Do Computers Stop and What Can Be Done About It?** — Gray (1985)
  - http://www.hpl.hp.com/techreports/tandem/TR-85.7.pdf
- **Crash-Only Software** — Candea & Fox (HotOS 2003)
- **How Complex Systems Fail** — Cook
  - https://how.complexsystems.fail/
- **Simple Testing Can Prevent Most Critical Failures** — Yuan et al. (OSDI 2014)
  - https://www.usenix.org/system/files/conference/osdi14/osdi14-paper-yuan.pdf
- Post-mortems: https://github.com/danluu/post-mortems
- **aphyr / Jepsen** — https://github.com/aphyr/distsys-class + jepsen.io writeups

---

## Gap fill — staff interviews + AWS cloud (papers & primary docs)

These close the gaps between classic theory and what a staff AWS loop actually probes. Prefer these after the first-principles core. Hands-on builds: see `Hands-on.md`.

### Quorums, voting, leases (often assumed, rarely named)

- **Weighted Voting for Replicated Data** — Gifford (1979)
  - quorum R+W > N; Dynamo/Cassandra N/R/W descends from this
  - http://www.cs.cmu.edu/~15-610/READINGS/required/availability/gifford79.pdf
- **Leases: An Efficient Fault-Tolerant Mechanism for Distributed File Cache Consistency** — Gray & Cheriton (1989)
  - time-bound ownership; Chubby / DynamoDB locks / leader leases
  - http://www.stanford.edu/class/cs240/readings/89-leases.pdf

### Messaging, exactly-once, logs

- **The Log: What Every Software Engineer Should Know About Real-time Data's Unifying Abstraction** — Jay Kreps
  - https://engineering.linkedin.com/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying
- **Kafka** — Kreps et al. (already above) — ordered log + consumer groups
- **Transactional messaging / outbox pattern** (concept paper + industry)
  - **Idempotence Is Not a Strategy** / outbox discussions — pair with Helland
  - **Building on Quicksand** — Helland (CIDR 2009)
    - https://blueprint.csail.mit.edu/wp-content/uploads/2016/10/Building-on-Quicksand.pdf
- **Exactly-once semantics in practice** — MillWheel / Dataflow (already above); Kafka EOS blogs (Confluent)
- **Starbucks Does Not Use Two-Phase Commit** — Alpert / Fowler essay
  - asynchronous sagas in the real world
  - https://www.enterpriseintegrationpatterns.com/ramblings/18_starbucks.html

### Multi-region / geo consistency

- **PNUTS: Yahoo!'s Hosted Data Serving Platform** — Cooper et al. (VLDB 2008)
  - per-record timeline consistency; proto–Global Tables thinking
  - https://www.cs.ucsb.edu/~agrawal/fall2009/PNUTS.pdf
- **MDCC: Multi-Data Center Consistency** — Kraska et al. (EuroSys 2013)
  - https://amplab.cs.berkeley.edu/wp-content/uploads/2013/03/mdcc-eurosys13.pdf
- **Don't Settle for Eventual: COPS** — Lloyd et al. (SOSP 2011)
  - causal+ geo
  - https://www.cs.cmu.edu/~dga/papers/cops-sosp2011.pdf
- **Spanner, TrueTime and the CAP Theorem** — Brewer
  - https://cloud.google.com/spanner/docs/whitepapers/SpannerAndCap
- **Stronger Semantics for Low-Latency Geo-Replicated Storage (Eiger)** — Lloyd et al.

### Hot keys, admission control, predictable performance

- **The Tail at Scale** — Dean & Barroso (already above) — hedged requests, latency variance
- **Amazon DynamoDB: A Scalable, Predictably Performant, and Fully Managed NoSQL Database Service** — Elhemali et al. (USENIX ATC 2022)
  - *must read for AWS staff* — Multi-Paxos per partition, GAC/token buckets, adaptive capacity, WAL→S3
  - https://www.usenix.org/system/files/atc22-elhemali.pdf
- **Lessons learned from 10 years of DynamoDB** — Amazon Science
  - https://www.amazon.science/blog/lessons-learned-from-10-years-of-dynamodb
- **DynamoDB, Ten Years Later** (summary)
  - https://www.mydistributed.systems/2022/10/dynamodb-ten-years-later.html

### AWS primary systems papers / deep dives (map theory → service)

| Service | Paper / primary | Maps from |
|---|---|---|
| S3 | **Building and operating a pretty big storage system called S3** — Vogels / re:Invent talk + ATD | GFS durability thinking, erasure, request patterns |
| S3 | [Building and operating… S3](https://www.allthingsdistributed.com/2023/07/building-and-operating-a-pretty-big-storage-system.html) | |
| S3 consistency | [Diving Deep on S3 Consistency](https://www.allthingsdistributed.com/2021/04/s3-strong-consistency.html) | CAP / read-after-write |
| DynamoDB | ATC 2022 paper above | Dynamo → managed Multi-Paxos |
| Aurora | Verbitski SIGMOD 2017 + 2018 (already above) | log is DB; quorum storage |
| Lambda networking | [The invisible engineering behind Lambda’s network](https://www.allthingsdistributed.com/2026/04/the-invisible-engineering-behind-lambdas-network.html) | multi-tenant isolation |
| Nitro | [Reinventing virtualization with the AWS Nitro System](https://www.allthingsdistributed.com/2020/09/reinventing-virtualization-with-nitro.html) | offload / blast radius |
| Constant work | [Standing on the shoulders of giants: Colm on constant work](https://www.allthingsdistributed.com/2023/11/constant-work.html) | staff ops principle |
| Control planes | [On building scalable control planes](https://www.allthingsdistributed.com/2026/08/on-building-scalable-control-planes.html) | API/control vs data plane |
| Frugal Architect | [The Frugal Architect](https://www.thefrugalarchitect.com/) — Vogels | cost as architecture axis |
| Distributed Computing Manifesto | ATD PDF (already above) | Amazon SOA DNA |

### Idempotency, APIs, distributed transactions (staff design vocabulary)

- **Making Reliable Distributed Systems in the Presence of Software Errors** — Armstrong (Erlang thesis)
  - supervision trees; “let it crash”
  - http://www.erlang.org/download/armstrong_thesis_2003.pdf
- **Life Beyond Distributed Transactions** + **Building on Quicksand** — Helland (already / above)
- **Sagas** (already above) — long-running business tx
- **Calvin** / **Percolator** (already above) — two ends of distributed tx design space
- **F1 / Spanner SQL** — externally consistent distributed SQL (already above)

### Observability & testing distributed systems

- **Dapper** (already above) → AWS X-Ray / OpenTelemetry mental model
- **Jepsen analyses** — Kyle Kingsbury (read 2–3: Kafka, Elasticsearch, Mongo, or DynamoDB-era notes)
  - https://jepsen.io/analyses
- **IronFleet: Proving Practical Distributed Systems Correct** — Hawblitzel et al. (optional depth)
- **SEDA** — Welsh, Culler, Brewer (staged event-driven; backpressure ancestor)
  - https://www.usenix.org/legacy/events/osdi02/tech/full_papers/welsh/welsh.pdf

### Cost, economics, giant-scale ops

- **Distributed Computing Economics** — Jim Gray
  - https://www.microsoft.com/en-us/research/publication/distributed-computing-economics/
- **Rules of Thumb in Data Engineering** — Gray & Shenoy
- **On Designing and Deploying Internet-Scale Services** — Hamilton (already above)
- **Lessons from Giant-Scale Services** — Brewer (already above)

---

## Meta reading lists (keep these open)

| Source | Why |
|---|---|
| [Murat — Foundational DS papers](http://muratbuffalo.blogspot.com/2021/02/foundational-distributed-systems-papers.html) | Best chronological first-principles taxonomy |
| [alexprut/distributed-systems](https://github.com/alexprut/distributed-systems) | Essential / Virtuoso / Maestro PDF bundles + citation order |
| [papers-we-love/distributed_systems](https://github.com/papers-we-love/papers-we-love/tree/main/distributed_systems) | Hosted PDFs for many classics |
| [heidihoward/distributed-consensus-reading-list](https://github.com/heidihoward/distributed-consensus-reading-list) | Deep consensus / BFT / quorums |
| [Paper Trail — theory for engineers](https://www.the-paper-trail.org/post/2014-08-09-distributed-systems-theory-for-the-distributed-systems-engineer/) | What an engineer must internalize (not a PhD laundry list) |
| [dancres reading list](https://dancres.github.io/Pages/) | Helland essays + Google/Amazon systems + gossip/P2P |
| [yeshengm/awesome-distributed-systems](https://github.com/yeshengm/awesome-distributed-systems) | Index of indexes |
| [MIT 6.824](https://pdos.csail.mit.edu/6.824/) | Course + lab papers (Raft, GFS, …) |
| [CMU 15-749 readings](http://www.andrew.cmu.edu/course/15-749/READINGS/required/) | Required classic set |

## From [All Things Distributed](https://www.allthingsdistributed.com/articles.html) (Werner Vogels)

Crawled the articles index (444 posts); extracted papers from weekend-reading / Dynamo / Aurora / manifesto / reading-list posts. Dead `wv.ly` shortlinks replaced with canonical URLs where known; otherwise the ATD post is the pointer.

Index of the series: [Back-to-Basics Readings of 2012](https://www.allthingsdistributed.com/2012/12/paper-readings-2012.html)

### Fault tolerance, agreement, consistency

- **The Byzantine Generals Problem** — Lamport, Shostak, Pease
  - https://www.microsoft.com/en-us/research/wp-content/uploads/2016/12/The-Byzantine-Generals-Problem.pdf
- **Reaching Agreement in the Presence of Faults** — Pease, Shostak, Lamport
  - https://www.microsoft.com/en-us/research/wp-content/uploads/2016/12/Reaching-Agreement-in-the-Presence-of-Faults.pdf
- **SIFT: Design and Analysis of a Fault-Tolerant Computer for Aircraft Control**
  - https://www.microsoft.com/en-us/research/wp-content/uploads/2016/12/Design-and-Analysis-of-a-Fault-Tolerant-Computer-for-Aircraft-Control.pdf
- **Why Do Computers Stop and What Can Be Done About It?** — Gray (1985)
  - http://www.hpl.hp.com/techreports/tandem/TR-85.7.pdf
- **A Survey of Rollback-Recovery Protocols in Message-Passing Systems** — Elnozahy, Alvisi, Wang, Johnson
  - http://www.cs.utexas.edu/users/lorenzo/papers/SurveyFinal.pdf
- **Distributed Snapshots: Determining Global States of a Distributed System** — Chandy & Lamport
  - http://research.microsoft.com/en-us/um/people/lamport/pubs/chandy.pdf
- **Virtual Time and Global States of Distributed Systems** — Mattern (1989)
  - http://courses.csail.mit.edu/6.852/01/papers/VirtTime_GlobState.pdf
- **Weighted Voting for Replicated Data** — Gifford (1979)
  - http://www.cs.cmu.edu/~15-610/READINGS/required/availability/gifford79.pdf
- **Leases: An Efficient Fault-Tolerant Mechanism for Distributed File Cache Consistency** — Gray & Cheriton (1989)
  - http://www.stanford.edu/class/cs240/readings/89-leases.pdf
- **Epidemic algorithms for replicated database maintenance** — Demers et al. (Xerox PARC)
  - http://ftp.se.scene.org/pub/bitsavers.org/pdf/xerox/parc/techReports/CSL-89-1_Epidemic_Algorithms_for_Replicated_Database_Maintenance.pdf
- **Epidemic algorithms in replicated databases**
  - http://www.cse.scu.edu/~jholliday/112609-2.pdf
- **Gossips and Telephones**
  - https://www.allthingsdistributed.com/files/gossips-telephones.pdf
- **Group membership in the epidemic style**
  - http://www.soe.ucsc.edu/share/technical-reports/1992/ucsc-crl-92-13.pdf
- **Randomized rumor spreading**
  - http://archive.cone.informatik.uni-freiburg.de/pubs/rumor.pdf

### Distributed systems & OS classics

- **Grapevine: An Exercise in Distributed Computing** — Birrell et al.
  - https://birrell.org/andrew/papers/Grapevine.pdf
- **Experience with Grapevine: The Growth of a Distributed System**
  - http://birrell.org/andrew/papers/ExperienceWithGrapevine.pdf
- **Scale and Performance in a Distributed File System (AFS)** — Howard et al. (1988)
  - http://inst.cs.berkeley.edu/~cs262/sp02/Papers/afs.pdf
- **Automatic Reconfiguration in Autonet** — Rodeheffer & Schroeder
  - http://birrell.org/andrew/papers/059-Autonet.pdf
- **End-to-End Arguments in System Design** — Saltzer, Reed, Clark (1984)
  - https://web.mit.edu/Saltzer/www/publications/endtoend/endtoend.pdf
- **Hints for Computer Systems Design** — Lampson (1983)
  - https://www.microsoft.com/en-us/research/publication/hints-for-computer-system-design/
- **On the Naming and Binding of Network Destinations** — Saltzer (RFC 1498)
  - https://www.rfc-editor.org/rfc/rfc1498
- **SEDA: An Architecture for Well-Conditioned, Scalable Internet Services** — Welsh, Culler, Brewer (SOSP 2001)
  - https://www.usenix.org/legacy/events/osdi02/tech/full_papers/welsh/welsh.pdf
  - ATD: https://www.allthingsdistributed.com/2012/08/staged-event-driven-architecture.html
- **Adaptive Load Sharing in Homogeneous Distributed Systems** — Eager, Lazowska, Zahorjan (1986)
  - https://homes.cs.washington.edu/~lazowska/qsp/Images/Papers/PDF/ALS.pdf
- **Using Encryption for Authentication in Large Networks of Computers** — Needham & Schroeder (1978)
  - http://jmiller.uaa.alaska.edu/cse465-fall2012/papers/needham1978.pdf
- **Disco: Running Commodity Operating Systems on Scalable Multiprocessors** — Bugnion et al.
  - http://www.stanford.edu/class/cs240/readings/disco.pdf
- **Xen and the Art of Virtualization** — Barham et al. (SOSP 2003)
  - http://www.cl.cam.ac.uk/Research/SRG/netos/papers/2003-xensosp.pdf
- **U-Net: A User-Level Network Interface for Parallel and Distributed Computing**
  - https://www.allthingsdistributed.com/files/u-net.pdf
- **Limits to Low-Latency Communication on High-Speed Networks** — Thekkath
  - http://www.thekkath.org/Documents/lowlatency.pdf
- **Sparse Partitions** — Awerbuch & Peleg (FOCS 1990)
  - http://courses.csail.mit.edu/6.885/spring06/papers/AwerbuchPeleg-focs.pdf
- **The Structure of the THE Multiprogramming System** — Dijkstra
  - http://uosis.mif.vu.lt/~liutauras/books/Dijkstra%20-%20The%20structure%20of%20the%20THE%20multiprogramming%20system.pdf
- **The Working Set Model for Program Behavior** — Denning (1968)
  - http://denninginstitute.com/pjd/PUBS/WSModel_1968.pdf
- **The Emperor's Old Clothes** — Hoare
  - https://www.cs.ucsb.edu/~cs170/documents/oldclothes.pdf
- **The Rise of Worse is Better** — Gabriel
  - https://www.dreamsongs.com/RiseOfWorseIsBetter.html
- OS books list (not papers): [The OS Classics](https://www.allthingsdistributed.com/2020/07/the-os-classics.html)

### Storage, DB recovery, columns, caches

- **A Case for Redundant Arrays of Inexpensive Disks (RAID)** — Patterson, Gibson, Katz
  - http://www.cs.cmu.edu/~garth/RAIDpaper/Patterson88.pdf
- **RAID: High-Performance, Reliable Secondary Storage** — Chen et al.
  - https://web.eecs.umich.edu/~pmchen/papers/chen94_1.pdf
- **The Design and Implementation of a Log-Structured File System** — Rosenblum & Ousterhout
  - https://people.eecs.berkeley.edu/~brewer/cs262/LFS.pdf
- **An Implementation of a Log-Structured File System for UNIX** — Seltzer et al.
  - https://www.usenix.org/legacy/publications/library/proceedings/sd93/seltzer.pdf
- **The 5 Minute Rule…** — Gray & Putzolu (1987)
  - http://www.hpl.hp.com/techreports/tandem/TR-86.1.pdf
- **The Five-Minute Rule Ten Years Later** — Gray & Graefe (1997)
  - ftp://ftp.research.microsoft.com/pub/tr/tr-97-33.pdf
- **The Five-Minute Rule 20 Years Later** — Graefe (2008)
  - http://cacm.acm.org/magazines/2009/7/32091-the-five-minute-rule-20-years-later/fulltext
- **Principles of Transaction-Oriented Database Recovery** — Härder & Reuter (1983)
  - http://www.minet.uni-jena.de/dbis/lehre/ws2005/dbs1/HaerderReuter83.pdf
- **Granularity of Locks and Degrees of Consistency** — Gray, Lorie, Putzolu, Traiger (1976)
  - https://jimgray.azurewebsites.net/papers/GranularityOfLocks.pdf
  - ATD: https://www.allthingsdistributed.com/2012/08/granularity-of-locks.html
- **A Decomposition Storage Model** (column stores precursor)
  - http://www3.in.tum.de/teaching/ws0506/MMDBMS/download/decomposition-storage-model.pdf
- **C-Store** — Stonebraker et al.
  - http://db.csail.mit.edu/projects/cstore/vldb.pdf
- **Column-Stores vs. Row-Stores (VLDB 2009 tutorial)** — Harizopoulos, Abadi, Boncz
  - http://www.cs.yale.edu/homes/dna/talks/Column_Store_Tutorial_VLDB09.pdf
- **Space/Time Trade-offs in Hash Coding with Allowable Errors (Bloom filters)** — Bloom
  - http://dmod.eu/deca/ft_gateway.cfm.pdf
- **Cache-, Hash- and Space-Efficient Bloom Filters**
  - http://algo2.iti.kit.edu/singler/publications/cacheefficientbloomfilters-wea2007.pdf
- **Summary Cache: A Scalable Wide-Area Web Cache Sharing Protocol** — Fan, Cao, Almeida, Broder
  - http://pages.cs.wisc.edu/~jussara/papers/00ton.pdf
- **Join Processing in Relational Databases** — Mishra & Eich (survey)
  - ATD: https://www.allthingsdistributed.com/2013/04/join-processing-relational-databases.html
- **Practical Applications of Triggers and Constraints**
  - ATD: https://www.allthingsdistributed.com/2013/04/practical-applications-triggers-constraints.html
- **An Introduction to Spatial Database Systems**
  - ATD: https://www.allthingsdistributed.com/2013/08/spatial-databases.html

### Security, anonymity, encrypted DB

- **Tor: The Second-Generation Onion Router**
  - http://www.onion-router.net/Publications/tor-design.pdf
- **Hiding Routing Information** (onion routing precursor)
  - http://www.onion-router.net/Publications/IH-1996.pdf
- **CryptDB: Protecting Confidentiality with Encrypted Query Processing** — Popa et al. (SOSP 2011)
  - http://nms.lcs.mit.edu/papers/cryptdb-sosp11.pdf
- **How to Time-Stamp a Digital Document** — Haber & Stornetta
  - https://www.anf.es/pdf/Haber_Stornetta.pdf
- **Cryptographic Support for Secure Logs on Untrusted Machines** — Schneier & Kelsey
  - https://www.schneier.com/academic/paperfiles/paper-secure-logs.pdf
- **The Eternity Service** — Anderson
  - https://www.cl.cam.ac.uk/~rja14/Papers/eternity.pdf

### Graphs / ML / misc. from the same ATD series

- **Distributed GraphLab** — Low et al.
  - https://select.cs.cmu.edu/publication/selectpubs/preprints/vldb2012low.pdf
- **PowerGraph: Distributed Graph-Parallel Computation on Natural Graphs** — Gonzalez et al. (OSDI 2012)
  - https://www.usenix.org/system/files/conference/osdi12/osdi12-final-167.pdf
- **Survey of Local Algorithms** — Suomela
  - http://users.ics.aalto.fi/suomela/doc/local-survey.pdf
- **Exploring Complex Networks** — Strogatz (Nature)
  - http://www.nature.com/nature/journal/v410/n6825/pdf/410268a0.pdf
- **A Few Useful Things to Know about Machine Learning** — Domingos
  - http://homes.cs.washington.edu/~pedrod/papers/cacm12.pdf
- **Deep Learning in Neural Networks: An Overview** — Schmidhuber
  - http://people.idsia.ch/~juergen/DeepLearning8Oct2014.pdf
- **Data Compression** — Hirschberg survey
  - http://www.ics.uci.edu/~dan/pubs/DataCompression.pdf
- **The Monte Carlo Method** — Metropolis & Ulam
  - http://homepages.rpi.edu/~angel/MULTISCALE/metropolis_Ulam_1949.pdf
- **Using Continuations to Implement Thread Management…**
  - ATD: https://www.allthingsdistributed.com/2013/05/continuations-in-operating-system.html
- **Auctions and Bidding: A Guide for Computer Scientists**
  - ATD: https://www.allthingsdistributed.com/2013/06/auctions-and-bidding.html
- **From Push to Pull** — John Seely Brown
  - http://www.johnseelybrown.com/pushmepullyou4.72.pdf
- **The Machine Stops** — Forster (from 2025 reading list)
  - https://www.cs.ucdavis.edu/~koehl/Teaching/ECS188/PDF_files/Machine_stops.pdf

## Open source & cloud-native systems papers (enterprise)

Industry runs **OSS + managed wrappers**. Read the paper, then touch the project. Full tool map + AWS equivalents: `Cloud-native-OSS.md`. Landscape: [CNCF Landscape](https://landscape.cncf.io/).

### Orchestration & cluster management

- **Large-scale cluster management at Google with Borg** — Verma et al. (EuroSys 2015)
  - https://research.google/pubs/pub43438/
- **Omega: flexible, scalable schedulers for large compute clusters** — Schwarzkopf et al.
  - https://research.google/pubs/pub41684/
- **Kubernetes** design docs / SIG architecture (Borg’s open heir)
  - https://kubernetes.io/docs/concepts/overview/
- **Mesos: A Platform for Fine-Grained Resource Sharing** — Hindman et al. (NSDI 2011)

### Coordination / KV / consensus in production OSS

- **ZooKeeper** — Hunt et al. (already above)
- **etcd** — Raft-based; Kubernetes control plane backbone
  - https://etcd.io/docs/latest/learning/why/
- **Consul** — service discovery + Raft sessions (HashiCorp docs + gossip)
- **Chubby** (already above) — conceptual parent of lock/lease services

### Datastores (OSS enterprises actually run)

- **Cassandra** — Lakshman & Malik (already above) → ScyllaDB (drop-in perf fork)
- **HBase** — Bigtable open clone; HDFS papers
- **Riak** / Bitcask — Dynamo descendant
- **Redis** — in-memory structures; Redis Cluster spec
  - https://redis.io/docs/latest/operate/oss_and_stack/reference/cluster-spec/
- **RocksDB** — Facebook embedded storage engine (under many systems)
  - https://rocksdb.org/
- **CockroachDB** — Spanner-inspired; [CockroachDB design](https://www.cockroachlabs.com/docs/stable/architecture/overview)
- **TiDB / TiKV** — Raft regions; Spanner-like
- **Vitess** — YouTube; MySQL sharding at scale (CNCF)
  - https://vitess.io/docs/overview/whatisvitess/
- **ClickHouse** — columnar OLAP (paper/blog lineage)
- **Elasticsearch** — Lucene distributed search
- **PostgreSQL** + **Patroni** / **Citus** — enterprise default RDBMS + HA/sharding
- **Apache Cassandra / Scylla / YugabyteDB** — wide-column / distributed SQL variants

### Logs, streaming, queues

- **Kafka** — Kreps et al. (already above)
- **Kora: A Cloud-Native Event Streaming Platform for Kafka** — Confluent (VLDB 2023)
  - https://www.vldb.org/pvldb/vol16/p3822-povzner.pdf
- **The Cloud-Native Chasm** — reinventing Kafka as managed service
  - https://assets.confluent.io/m/3f006525da43c842/original/20211208-WP-Cloud_Native_Chasm_Lessons_Learned_from_Reinventing_Apache_Kafka_as_Cloud_Native_Online_Service.pdf
- **Apache Pulsar** — BookKeeper + brokers; multi-tenant messaging
- **Apache Flink** — stateful stream processing; exactly-once
  - Disaggregated state Flink 2.0 (VLDB): https://www.vldb.org/pvldb/vol18/p4846-mei.pdf
- **Apache Spark / RDD** (already above)
- **NATS / JetStream**, **RabbitMQ**, **Apache ActiveMQ / Artemis**
- **Redpanda** — Kafka API, no JVM/ZK historically

### Object / file / blob OSS

- **Ceph** — reliable, scalable distributed storage
- **MinIO** — S3-compatible object store
- **SeaweedFS**, **OpenStack Swift**
- **HDFS** — GFS descendant (MapReduce ecosystem)

### Service mesh, proxy, RPC

- **Envoy** — Lyft; universal data-plane proxy (CNCF)
- **Linkerd**, **Istio** — service mesh
- **gRPC** + **Protobuf** — Google RPC stack (open)
- **Cilium** — eBPF networking/security (CNCF)

### Observability (OSS default stack)

- **Dapper** (already above) → **Jaeger** / **Zipkin** / **OpenTelemetry**
- **Prometheus** + **Grafana** + **Thanos** / **Cortex** / **Mimir**
- **Fluentd** / **Fluent Bit** / **Loki** / **OpenSearch**
- **Falco** — runtime security

### Workflow / durable execution

- **Temporal** (Cadence fork) — durable workflows
- **Apache Airflow** — batch DAGs
- **Argo Workflows / Events / CD** — K8s-native (CNCF)
- **Cadence** — Uber

### Lakehouse / analytics table formats

- **Apache Iceberg**, **Delta Lake**, **Apache Hudi** — ACID tables on object storage
- **Apache Parquet** / **ORC** — columnar files (Dremel lineage)

---

## Suggested reading order (first principles → practice)

Do **not** skip classic systems for AWS-only reading. Paths are parallel after step 5.

1. **Mindset:** Fallacies → Note on Distributed Computing → Young Bloods → Fun and Profit
2. **Time/state:** Lamport clocks → vector clocks / Mattern → Chandy–Lamport snapshots
3. **Limits:** Two Generals → Byzantine Generals → FLP → failure detectors (Chandra/Toueg) → CAP/PACELC (Gilbert–Lynch + Brewer + Abadi)
4. **Primitives:** Schneider state-machine tutorial → Gifford voting → Paxos Made Simple → Raft
5. **Weak consistency:** Session guarantees → Dynamo → Helland Life Beyond DT / Building on Quicksand → CRDTs
6. **Systems (classic — keep this path):** GFS → Bigtable → Spanner → ZooKeeper; then Kafka or Aurora
7. **OSS twins (same concepts, runnable):** HDFS/HBase or Ceph/MinIO → Cassandra/Scylla → etcd → Kafka (+ optional Flink/Pulsar) → see `Cloud-native-OSS.md`
8. **AWS bridge:** DynamoDB ATC 2022 → Aurora SIGMOD → S3 consistency + “pretty big storage” → Kreps Log essay
9. **Geo:** PNUTS or COPS → Spanner TrueTime/CAP note
10. **Cloud-native platform:** Borg/Omega → Kubernetes concepts → Envoy/OTel/Prometheus
11. **Ops:** Gray failures → Tail at Scale → Hamilton → constant work / Frugal Architect
12. **Code:** execute `Hands-on.md` (+ OSS labs in `Cloud-native-OSS.md`)

alexprut **Essential** PDF order is a good parallel track: clocks → snapshots → Byzantine → Paxos → CAP → consistent hashing → MapReduce → Paxos Made Live → Bitcoin.

## Already cited in this vault

- Dynamo — see `distributed systems.md`
- Chain replication, Whittaker consistency notes, Bailis HAT — see `Consistency.md`
- PACELC — see `Replication.md`
- Single-node KV starter — see `distributed-key-value-store.md`
- DS build sequence — see `Hands-on.md`
- Enterprise OSS / CNCF / AWS map — see `Cloud-native-OSS.md`
- High-level design (HLD) — see `System-Design.md`
- Low-level design (LLD) — see `LLD.md`
- HLD + LLD drills — see `System-Design-Hands-on.md`
