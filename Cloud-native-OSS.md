# Open source, free tools & cloud-native stack (enterprise)

What most enterprises actually run: **OSS core + managed cloud**. Map each row to a paper in `Papers.md`, then to AWS when you interview there.

Primary index: [CNCF Landscape](https://landscape.cncf.io/) · [CNCF graduated projects](https://www.cncf.io/projects/)

---

## How to use this

For every category: **concept paper → OSS you can run → AWS managed cousin → when to pick which**.

Staff interview answer shape: “We’d run X OSS if we need portability / cost / on-prem; on AWS I’d use Y managed because Z ops burden.”

---

## 1. Compute & orchestration

| Role | OSS / free | Paper / origin | AWS managed |
|---|---|---|---|
| Container runtime | containerd, CRI-O | OCI | ECS/EKS/Fargate under the hood |
| Orchestration | **Kubernetes** (CNCF) | Borg, Omega | **EKS**, ECS (AWS-native) |
| Older / niche | Nomad, Mesos/Marathon | Mesos paper | — |
| Serverless frameworks | Knative, OpenFaaS | — | **Lambda**, App Runner |
| VMs / images | Firecracker (AWS open), QEMU | Nitro/Firecracker talks | Lambda/Fargate microVMs |
| Pack / GitOps | Helm, Kustomize, **Argo CD**, Flux | — | — (use with EKS) |
| IaC | Terraform/OpenTofu, Pulumi, Crossplane, CDK | — | CloudFormation, CDK |

**Must know K8s primitives:** Pod, Deployment, Service, Ingress, ConfigMap/Secret, HPA, PDB, StatefulSet, CSI volumes.

**Hands-on lab:** minikube/kind → deploy API + Redis + Ingress; break a node; watch reschedule.

---

## 2. Service discovery, config, coordination

| Role | OSS | Paper / concept | AWS |
|---|---|---|---|
| Coordination / locks | **etcd**, ZooKeeper | Raft, Zab, Chubby | DynamoDB locks, EventBridge, Step Functions (not a ZK SaaS) |
| Service discovery | CoreDNS, Consul, Eureka (legacy) | gossip / registry | Cloud Map, ALB/NLB, ECS Service Connect |
| Feature flags / config | Consul KV, etcd, Unleash | — | AppConfig, SSM Parameter Store, Secrets Manager |
| Secrets | Vault (HashiCorp), SOPS, External Secrets Operator | — | Secrets Manager, ASM |

**Hands-on lab:** 3-node etcd; leader kill; read linearizable key.

---

## 3. Datastores

| Role | OSS / free | Paper lineage | AWS |
|---|---|---|---|
| RDBMS | **PostgreSQL**, MySQL/MariaDB | System R lineage | RDS, Aurora |
| PG HA / shard | Patroni, Citus, Vitess (MySQL) | Vitess @ YouTube | Aurora, RDS Proxy |
| Wide-column | **Cassandra**, ScyllaDB, HBase | Dynamo + Bigtable | Keyspaces (Cassandra API), DynamoDB |
| Distributed SQL | CockroachDB, TiDB/Yugabyte, Spanner-inspired | Spanner | Aurora DSQL / Spanner (GCP) |
| Document | MongoDB, FerretDB | — | DocumentDB |
| Cache | **Redis**/Valkey, KeyDB, Memcached | — | ElastiCache, MemoryDB |
| Embedded engine | RocksDB, Badger, LevelDB | LSM / Bigtable SSTable | used inside many AWS systems |
| Search | Elasticsearch / OpenSearch, Meilisearch, Typesense | Lucene | OpenSearch Service |
| Graph | Neo4j, JanusGraph | — | Neptune |
| Time series | Prometheus (TSDB), InfluxDB, TimescaleDB | — | Timestream, Managed Prometheus |
| OLAP / column | **ClickHouse**, Druid, Pinot | Dremel / C-Store | Redshift, Athena |
| Vector (AI apps) | pgvector, Milvus, Qdrant, Weaviate | — | OpenSearch k-NN, MemoryDB vector |

**Hands-on lab:** docker-compose Postgres + Redis; then Cassandra 3-node; contrast consistency knobs.

---

## 4. Object / file / block storage

| Role | OSS | Paper | AWS |
|---|---|---|---|
| Object | **MinIO**, Ceph RGW, SeaweedFS, Swift | GFS / S3 talks | **S3** |
| Distributed FS | Ceph FS, Gluster, HDFS | GFS, Ceph | EFS, FSx |
| Block | Ceph RBD, Longhorn (K8s) | — | EBS |
| Local engine | RocksDB, SQLite | — | — |

**Hands-on lab:** MinIO S3 API + versioning; compare to S3 consistency post in `Papers.md`.

---

## 5. Messaging & streaming

| Role | OSS | Paper | AWS |
|---|---|---|---|
| Log / event streaming | **Apache Kafka**, Redpanda, WarpStream | Kafka paper, Kora, Kreps Log | **MSK**, Kinesis Data Streams |
| Multi-tenant messaging | **Apache Pulsar** | BookKeeper | — (MSK/Pulsar partners) |
| Classic queues | RabbitMQ, ActiveMQ Artemis | — | **SQS**, MQ |
| Lightweight | NATS/JetStream, Redis Streams | — | SNS+SQS |
| Stream processing | **Flink**, Kafka Streams, Spark Structured Streaming, Faust | MillWheel, Flink papers, Dataflow | Kinesis Analytics / Managed Flink, EMR |
| CDC | Debezium, Maxwell | — | DMS, DynamoDB Streams |
| Schema | Apicurium, Confluent Schema Registry, Buf | — | Glue Schema Registry |

**Enterprise default today:** Kafka (or MSK) + Flink/Kafka Streams + Schema Registry.

**Hands-on lab:** Redpanda or Kafka in Docker → produce/consume → consumer group rebalance; then Flink word-count or Kafka Streams.

---

## 6. API edge, RPC, mesh

| Role | OSS | Notes | AWS |
|---|---|---|---|
| Edge proxy | **Envoy**, NGINX, HAProxy, Caddy, Traefik | Envoy = mesh data plane | ALB, NLB, CloudFront, API Gateway |
| Service mesh | Istio, Linkerd, Consul Connect | mTLS, retries, outlier detection | App Mesh (legacy-ish), VPC Lattice |
| RPC | **gRPC**, Connect-RPC, Twirp | Protobuf contracts | ALB gRPC, App Mesh |
| API gateway OSS | Kong, Apache APISIX, Tyk, KrakenD | — | API Gateway, ALB |

**Hands-on lab:** gRPC service + Envoy sidecar retries/timeouts; chaos latency injection.

---

## 7. Observability & ops

| Role | OSS | Paper | AWS |
|---|---|---|---|
| Metrics | **Prometheus**, Grafana, VictoriaMetrics, Thanos/Mimir | — | CloudWatch, AMP, AMG |
| Logs | Fluent Bit, Loki, OpenSearch | — | CloudWatch Logs, OpenSearch |
| Traces | **OpenTelemetry**, Jaeger, Zipkin | Dapper | X-Ray, OTel Collector on ECS/EKS |
| Profiling | Pyroscope, Parca, async-profiler | — | CodeGuru (limited) |
| Chaos | Chaos Mesh, Litmus, Toxiproxy, Jepsen | Yuan OSDI failures paper | FIS (Fault Injection Service) |
| Progressive delivery | Argo Rollouts, Flagger | — | CodeDeploy, ALB weighted |

**Enterprise default:** OTel → Prometheus/Grafana → Tempo/Jaeger → Loki/ELK.

---

## 8. Workflow, batch, lakehouse

| Role | OSS | Paper / idea | AWS |
|---|---|---|---|
| Durable workflows | **Temporal**, Cadence | Sagas, long-running tx | Step Functions |
| Batch DAG | Airflow, Dagster, Prefect | — | MWAA, Glue workflows |
| K8s jobs | Argo Workflows | — | Batch, Step Functions |
| Table formats | **Iceberg**, Delta Lake, Hudi | ACID on object storage | Athena/Iceberg, EMR |
| Query engines | Trino/Presto, Spark, DuckDB | Dremel | Athena, EMR, Redshift Spectrum |
| Catalog | Hive Metastore, Unity-style OSS, Nessie | — | Glue Data Catalog |

---

## 9. Security & policy (cloud-native)

| Role | OSS | AWS |
|---|---|---|
| Policy | **OPA**/Gatekeeper, Kyverno | IAM, SCP, Cedar (AWS Verified Permissions) |
| Supply chain | Sigstore, Cosign, Trivy, Grype, Syft | ECR scanning, Inspector |
| Runtime | Falco, Tetragon | GuardDuty |
| Identity | Keycloak, Dex, Authentik | Cognito, IAM Identity Center |
| Certs | cert-manager, Smallstep | ACM, Private CA |

---

## 10. “OSS ↔ AWS” cheat sheet (staff interviews)

| If the design needs… | Reach for OSS | On AWS say… |
|---|---|---|
| Linearizable coordination | etcd / ZK | DynamoDB conditional writes / DSQL / careful with Step Functions |
| Durable log bus | Kafka / Redpanda | MSK or Kinesis |
| Tunable AP KV | Cassandra / Scylla | DynamoDB (different: Multi-Paxos partitions) |
| Spanner-like SQL | Cockroach / TiDB | Aurora (+ DSQL where fits) / avoid fake Spanner claims |
| Object store | MinIO / Ceph | S3 |
| Exactly-once stream proc | Flink | Managed Service for Apache Flink |
| Workflows | Temporal | Step Functions |
| K8s platform | kubeadm / EKS Anywhere | EKS |
| Cache | Redis/Valkey | ElastiCache / MemoryDB |

**Trap:** “We’ll use Kafka + Cassandra + etcd + Flink + Temporal on EKS” is a full platform company. Staff answer = **narrow** the blast radius; managed where undifferentiated.

---

## Papers to pair (do not drop classics)

Keep this path from `Papers.md`:

**GFS → Bigtable → Spanner → ZooKeeper → Kafka or Aurora**

Then OSS twins:

| Classic paper | OSS twin |
|---|---|
| GFS | HDFS, Ceph, MinIO (S3 API) |
| Bigtable | HBase, Cassandra (data model mix) |
| Dynamo | Cassandra, Riak, Voldemort (hist.) |
| Spanner | CockroachDB, TiDB, Yugabyte |
| Chubby / ZK | etcd, Consul |
| MapReduce | Spark, Hadoop |
| MillWheel / Dataflow | Flink, Beam |
| Borg | Kubernetes |
| Dapper | Jaeger + OpenTelemetry |
| Kafka paper | Apache Kafka / Redpanda / MSK |

---

## Suggested OSS learning labs (after Hands-on P0–P4)

| Lab | Stack | Locks in |
|---|---|---|
| L1 | kind + Deployment + Service + Ingress | K8s scheduling, probes |
| L2 | etcd 3-node | Raft under your fingers |
| L3 | Kafka/Redpanda + 2 consumer groups | partitions, rebalance, lag |
| L4 | Cassandra `LOCAL_QUORUM` vs `ONE` | tune consistency |
| L5 | Flink or Kafka Streams job | state + checkpoint |
| L6 | MinIO + Iceberg or plain Parquet | lakehouse vs DB |
| L7 | OTel demo (Astronomy Shop) | traces across services |
| L8 | Temporal hello-saga | durable execution vs Step Functions |

Add these as stretch after matching `Hands-on.md` projects.

---

## Free learning resources (not papers)

- [CNCF Landscape](https://landscape.cncf.io/)
- [Kubernetes docs — concepts](https://kubernetes.io/docs/concepts/)
- [etcd docs — learning](https://etcd.io/docs/latest/learning/)
- [Apache Kafka docs](https://kafka.apache.org/documentation/)
- [Jepsen analyses](https://jepsen.io/analyses) — Cassandra, Elasticsearch, Mongo, Kafka, …
- [Awesome Distributed Systems](https://github.com/yeshengm/awesome-distributed-systems)
- AWS skill builder + Well-Architected (free tier of knowledge)

---

## Related vault files

- `Papers.md` — theory + classic systems + OSS paper list + reading order (includes GFS→…→Kafka/Aurora)
- `Hands-on.md` — build-from-scratch projects
- This file — enterprise OSS/CNCF/AWS map
