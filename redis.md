# Redis Notes for System Design Interviews

# What is Redis?

Redis is an in-memory key-value datastore.

Redis = Remote Dictionary Server

It is commonly used for:

* caching
* session storage
* rate limiting
* queues
* pub/sub
* counters
* real-time analytics

---

# Why Redis is Fast

Redis stores data in RAM instead of disk.

RAM is much faster than databases reading from storage.

Typical latency:

* Redis -> microseconds
* DB -> milliseconds

Huge performance difference at scale.

---

# Core Idea

Instead of:

Application -> Database

we use:

Application -> Redis Cache -> Database

Redis serves frequently accessed data quickly.

---

# Why Companies Use Redis

Benefits:

* extremely fast
* reduces DB load
* supports expiration (TTL)
* supports distributed caching
* supports many data structures

Used by:

* Instagram
* Twitter/X
* Netflix
* Uber
* Amazon

---

# Redis as a Cache

## Example Flow

### First Request

1. User requests profile
2. Redis checked
3. Cache miss
4. Fetch from DB
5. Store in Redis
6. Return response

---

### Second Request

1. Redis checked
2. Cache hit
3. Return instantly

No DB query needed.

---

# Redis Data Structures

Very important interview topic.

---

# 1. Strings

Most common.

Example:

* user profile JSON
* token
* counters

Commands:
SET key value
GET key

Example:
SET user:1 "Bhavana"

---

# 2. Hashes

Store object-like data.

Example:
HSET user:1 name Bhavana age 31

Useful for:

* profiles
* metadata

---

# 3. Lists

Ordered collection.

Useful for:

* feeds
* queues
* recent items

Commands:
LPUSH
RPUSH
LPOP

---

# 4. Sets

Unique unordered elements.

Useful for:

* followers
* tags
* unique visitors

Commands:
SADD
SMEMBERS

---

# 5. Sorted Sets (VERY IMPORTANT)

Stores:
(value, score)

Sorted automatically by score.

Used for:

* leaderboards
* rankings
* trending systems

Example:
(score = likes/views/time)

Commands:
ZADD
ZRANGE

---

# TTL (Time To Live)

Redis supports automatic expiration.

Example:
SETEX user:1 600 data

Meaning:

* expire after 600 seconds

Very important for caching.

---

# Cache Eviction Policies

When memory becomes full:
Redis removes data.

Popular policies:

## LRU

Least Recently Used

---

## LFU

Least Frequently Used

---

## TTL Based

Remove keys nearing expiration.

---

# Persistence in Redis

Redis is in-memory but can also save data to disk.

Two mechanisms:

---

## 1. RDB Snapshot

Takes periodic snapshots.

Good:

* compact
* faster recovery

Bad:

* some data loss possible

---

## 2. AOF (Append Only File)

Logs every write operation.

Good:

* better durability

Bad:

* larger files

---

# Redis Replication

Master -> Replica architecture

Purpose:

* high availability
* read scaling

Flow:

Application
↓
Master Redis
↓
Replica Redis

Writes:

* master

Reads:

* replicas

---

# Redis Sentinel

Used for:

* automatic failover
* monitoring

If master crashes:

* replica promoted automatically

Important interview term.

---

# Redis Cluster

Used for horizontal scaling.

Data split across multiple Redis nodes.

Technique:

* sharding

Important because:
Single Redis server cannot scale infinitely.

---

# Common Redis Interview Use Cases

---

# 1. Caching

Most common.

Example:

* user profiles
* feeds
* product pages

---

# 2. Rate Limiter

Track request counts.

Example:
"Allow only 100 requests/minute"

Redis ideal because:

* atomic counters
* expiration support

---

# 3. Session Store

Store user login sessions.

Example:
session:user123

---

# 4. Leaderboards

Use Sorted Sets.

Example:
Top players by score.

---

# 5. Real-Time Analytics

Track:

* views
* clicks
* likes

---

# 6. Message Queues

Redis lists used as lightweight queues.

---

# Typical System Design Architecture

Client
↓
Application Server
↓
Redis Cache
↓ (cache miss)
Database

---

# Important Tradeoffs

## Advantages

* very fast
* simple
* scalable
* supports TTL
* reduces DB load

---

## Disadvantages

* RAM expensive
* limited memory
* stale cache possible
* persistence weaker than DB

---

# Common Interview Questions

## Why use Redis over DB?

Because Redis is:

* memory-based
* much faster

DB optimized for:

* persistence
* complex queries

Redis optimized for:

* speed

---

# Why not store everything in Redis?

Because:

* RAM expensive
* limited memory
* data durability concerns

---

# Redis vs Memcached

## Redis

* richer data structures
* persistence
* replication
* pub/sub

## Memcached

* simpler
* pure cache

Redis generally preferred nowadays.

---

# High-Level Redis Interview Summary

Redis is:

* an in-memory datastore
* mainly used for caching
* extremely fast
* supports TTL
* supports replication and clustering
* commonly used to reduce database load

-
