# Database Replication Notes

# What is Database Replication?

Replication means:
copying data from one database server to other database servers.

Architecture:

```
            Primary DB
           /    |    \
          ↓     ↓     ↓
     Replica1 Replica2 Replica3
```

Primary:

* handles writes

Replicas:

* handle reads

---

# Why Replication is Needed

Single DB server becomes bottleneck.

Problems:

* too many read requests
* high CPU usage
* slow queries
* server crashes

Replication helps distribute load.

---

# Main Purpose of Replication

## 1. Read Scaling

Most applications have:
more reads than writes.

Example:
Instagram:

* millions of feed reads
* fewer writes/posts

Instead of:
all reads hitting primary DB

Use replicas.

Flow:

App
↓
Primary DB (writes)
↓
Replica DBs (reads)

Result:
reduced DB load.

---

# 2. High Availability

Suppose primary crashes.

Replica can be promoted as new primary.

System continues working.

Improves:
fault tolerance.

---

# 3. Backup / Disaster Recovery

Replicas act as backup copies.

If one DB corrupted:
data still available elsewhere.

---

# 4. Geographic Distribution

Users in India:
read from India replica.

Users in USA:
read from USA replica.

Reduces latency.

---

# Replication Flow

# Step 1

Write happens on Primary.

Example:
INSERT new post

---

# Step 2

Primary updates its logs.

---

# Step 3

Replica copies changes asynchronously/synchronously.

---

# Step 4

Replica data updated.

---

# Types of Replication

# 1. Asynchronous Replication (Most Common)

Primary responds immediately.

Replica updates later.

Advantages:

* fast writes
* low latency

Disadvantages:
Replication lag possible.

---

# Replication Lag

Replica may temporarily contain old data.

Example:
User posts photo.
Immediate read from replica may not show it yet.

This is:
eventual consistency.

---

# 2. Synchronous Replication

Primary waits for replicas before confirming write.

Advantages:

* strong consistency

Disadvantages:

* slower writes
* higher latency

---

# Important Interview Tradeoff

Asynchronous:
faster but slightly stale.

Synchronous:
consistent but slower.

---

# Real Industry Usage

Most companies use:
asynchronous replication.

Because:
availability + performance more important.

---

# Example — Instagram

Writes:

* new post
* new like

go to:
Primary DB

Reads:

* feed loading
* profile viewing

go to:
Replicas

This massively reduces DB load.

---

# Example — Netflix

User playback history:
write -> primary

Movie metadata:
read -> replicas

---

# Benefits of Replication

## 1. Read Scalability

More replicas ->
more read throughput.

---

## 2. High Availability

DB failure recovery.

---

## 3. Fault Tolerance

System survives node crashes.

---

## 4. Lower Latency

Regional replicas closer to users.

---

# Limitations / Bottlenecks

# 1. Write Bottleneck Still Exists

All writes still go to:
Primary DB.

If write traffic huge:
primary becomes bottleneck.

Replication mainly scales READS,
not WRITES.

---

# 2. Replication Lag

Replicas may contain stale data.

Can cause:
inconsistent reads.

---

# 3. Failover Complexity

Promoting replica to primary is complex.

Need:
leader election
consensus systems

---

# 4. Storage Cost

Multiple DB copies required.

More infrastructure cost.

---

# 5. Network Overhead

Replication traffic between DBs.

---

# What Next for Betterment?

When replication not enough,
next step is:

# Database Sharding

Because:
replication solves READ scaling.

Sharding solves:
WRITE scaling + data size scaling.

---

# Replication vs Sharding

## Replication

Copies SAME data.

Purpose:
read scaling + HA.

---

## Sharding

Splits DIFFERENT data across DBs.

Purpose:
write scaling + storage scaling.

---

# Typical Modern Architecture

```
             Load Balancer
                    ↓
              App Servers
                    ↓
            Primary Database
               /    |    \
              ↓     ↓     ↓
         Read Replicas
```

Reads:
replicas

Writes:
primary

---

# Important Interview Keywords

* Primary-Replica
* Read Scaling
* Replication Lag
* Eventual Consistency
* High Availability
* Failover
* Asynchronous Replication
* Synchronous Replication

---

# Key System Design Insight

Replication is mainly used to solve:
READ bottlenecks.

It does NOT fully solve:
WRITE bottlenecks.

For that,
systems move to:
Sharding.
