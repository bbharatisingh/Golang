# Sharding Example — Simplified

# Goal of Sharding

Split huge database across multiple servers.

Reason:
one DB cannot handle:

* massive users
* huge writes
* huge storage

---

# Example Setup

Suppose:
10 million users.

One DB server becomes overloaded.

So we split data across:

Shard A
Shard B
Shard C
Shard D

Each shard stores only PART of users.

---

# Using userId as Shard Key

Shard Key:
field used to decide where data goes.

Here:
userId

---

# Simple Idea

Suppose:

Users:
0–2.5M  -> Shard A
2.5–5M  -> Shard B
5–7.5M  -> Shard C
7.5–10M -> Shard D

Now:
each DB handles fewer users.

Load distributed.

---

# Query Example

Query:

"Get profile of user 8421"

System does:

1. Check shard key:
   userId = 8421

2. Routing logic decides:
   8421 belongs to Shard A

3. Query sent only to:
   Shard A

Fast lookup.

Instead of searching all DBs.

---

# What is Hashing?

Instead of manually assigning ranges,
systems often use:

hash(userId)

Reason:
better distribution.

Example:

hash(8421) -> Shard C

Router automatically determines shard.

---

# Problem — Hot Shard

Suppose:
Shard C gets too much traffic.

Maybe:
many active users mapped there.

Now:
Shard C overloaded.

Called:
Hot Shard Problem.

---

# Solution — Add New Shard

Add:
Shard E

Now redistribute some data.

Before:

Shard C handles:
40% traffic

After:

Shard C -> 20%
Shard E -> 20%

Load balanced.

---

# What are Virtual Nodes?

Instead of:
one huge fixed range per shard,

system creates:
many small partitions called virtual nodes.

Example:

Shard C may own:
V1 V2 V3 V4

When adding Shard E:

Move:
V3 V4 -> Shard E

Only partial data moves.

Benefits:

* less migration
* smoother scaling

---

# What is Consistent Hashing?

Technique to:
minimize data movement when adding/removing shards.

Without consistent hashing:
adding one shard may require moving MOST data.

With consistent hashing:
only small portion moves.

Very important distributed systems concept.

---

# Final Architecture

```
            Router
               ↓
  --------------------------------
  ↓         ↓        ↓         ↓
```

Shard A   Shard B  Shard C   Shard D

Query:
hash(userId)
decides shard.

---

# Key Insight

Replication:
copies SAME data.

Sharding:
splits DIFFERENT data across DBs.

Purpose of sharding:

* write scaling
* storage scaling
* distribute load
# Sharding — Common Interview Questions

# 1. When Should I Shard?

Shard when:

* write traffic becomes too high
* storage too large
* DB latency increases
* single DB cannot scale further

Usually after:

1. Vertical Scaling
2. Replication

are no longer enough.

---

# 2. What Makes a Good Shard Key?

A good shard key should be:

## 1. High Cardinality

Many unique values.

Good:

* user_id

Bad:

* gender

---

## 2. Evenly Distributed

Traffic spread uniformly.

Avoid:
all traffic going to one shard.

---

## 3. Frequently Used in Queries

Most queries should contain shard key.

Example:

SELECT * FROM users WHERE user_id=123

Easy routing.

---

# 3. Why are Cross-Shard Joins Expensive?

Suppose:

User data -> Shard A
Order data -> Shard D

Now JOIN requires:

* querying multiple DBs
* network communication
* data aggregation

Very slow.

Preferred:

* colocate related data
  OR
* perform joins in application layer

---

# 4. How is Uniqueness Maintained Across Shards?

Problem:
Auto-increment IDs may clash.

Solutions:

## 1. UUIDs

Globally unique.

---

## 2. Snowflake IDs

Distributed unique ID generation.

Used by:

* Twitter/X

---

## 3. Central ID Service

One service generates unique IDs.

---

# 5. How Do You Add New Shards Later?

Without proper strategy:
adding shard requires moving huge data.

Solution:

## Consistent Hashing

Only small portion of data moves.

OR

## Mapping Table

Router stores:
which shard owns which range.

---

# 6. What is Consistent Hashing?

Technique for:
minimal data movement during scaling.

Benefits:

* easier shard addition/removal
* smoother rebalancing

Very important distributed systems concept.

---

# 7. What is a Hot Shard?

One shard receives disproportionate traffic.

Example:
celebrity users mapped to same shard.

Problems:

* overload
* latency
* uneven scaling

Solutions:

* better shard key
* resharding
* virtual nodes

---

# 8. Replication vs Sharding

## Replication

Copies SAME data.

Goal:
read scaling.

---

## Sharding

Splits DIFFERENT data.

Goal:
write/storage scaling.

---

# 9. Can Shards Have Replicas?

YES.

Very common architecture.

Example:

Shard A Primary
↓
Replicas

Shard B Primary
↓
Replicas

Purpose:
both read scaling and write scaling.

---

# 10. Typical Industry Flow

1. Start with single DB
2. Vertical scaling
3. Add replicas
4. Introduce sharding
5. Add replicas per shard

This is common scaling evolution.

---

# Key Interview Insight

Replication solves:
READ bottlenecks.

Sharding solves:
WRITE + storage bottlenecks.
