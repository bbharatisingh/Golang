# Caching Notes

## What is Caching?

Caching is the process of storing frequently accessed data in a faster storage layer so future requests can be served quickly.

Instead of querying the database repeatedly:

Client -> Server -> Database

we use:

Client -> Server -> Cache -> Database (only if cache miss)

---

# Why Caching is Needed

* Reduces database load
* Improves response time
* Improves scalability
* Reduces server cost
* Handles high traffic efficiently

---

# Cache Flow

## Cache Miss

Data is not present in cache.

Flow:

1. Request comes to server
2. Server checks cache
3. Cache miss occurs
4. Server fetches data from DB
5. Store data in cache
6. Return response to user

---

## Cache Hit

Data is already present in cache.

Flow:

1. Request comes to server
2. Server checks cache
3. Data found in cache
4. Return response immediately

---

# Important Terms

## Cache Hit

Requested data found in cache.

Fast response.

---

## Cache Miss

Requested data not found in cache.

Requires DB query.

---

## TTL (Time To Live)

Time duration after which cached data expires automatically.

Example:

* Cache profile data for 10 minutes

Purpose:

* Prevent stale data

---

# Commonly Cached Data

Good candidates:

* User profiles
* Product pages
* Social media feeds
* Trending videos
* Frequently accessed APIs

Bad candidates:

* Real-time bank balances
* Highly dynamic critical data

---

# Popular Cache Tool

## Redis

Redis is an in-memory key-value store commonly used for caching.

Why Redis?

* Extremely fast
* Stored in RAM
* Supports TTL
* Supports distributed caching

---

# Cache Strategies

## 1. Cache Aside (Lazy Loading)

Most common strategy.

Flow:

1. Check cache
2. If miss -> query DB
3. Store result in cache
4. Return response

Advantages:

* Simple
* Efficient

Disadvantages:

* First request slower

---

## 2. Write Through Cache

Whenever DB updates:

* Update cache immediately

Advantages:

* Fresh cache

Disadvantages:

* Slightly slower writes

---

## 3. Write Back Cache

Update cache first.
Database updated asynchronously later.

Advantages:

* Fast writes

Disadvantages:

* Risk of data loss

---

# Cache Invalidation

Problem:
Cached data may become outdated.

Example:

* User updates profile picture
* Cache still contains old picture

Solutions:

* TTL expiry
* Update cache on write
* Delete cache after DB update

---

# Tradeoff

Caching improves performance but may reduce consistency.

Tradeoff:

* Speed vs Freshness

---

# Typical Architecture

Client
↓
Application Server
↓
Redis Cache
↓ (on miss)
Database

---

# Key Interview Insight

Whenever database becomes bottleneck:

Ask:
"Can this data be cached?"

Caching is usually the first scaling optimization.

---

# Real World Examples

## Instagram

* Cache feeds
* Cache user profiles

## YouTube

* Cache trending videos

## Amazon

* Cache product pages

---
# Cache Update vs LRU Notes

## Important Distinction

There are TWO separate concepts:

1. Cache Update Strategy
2. Cache Eviction Policy

People often confuse them.

---

# 1. Cache Update Strategy

This decides:

"How does cache get updated when data changes?"

Examples:

* Cache Aside
* Write Through
* Write Back

These are update/write strategies.

---

# 2. Cache Eviction Policy

This decides:

"When cache becomes full, what data should be removed?"

Examples:

* LRU
* LFU
* FIFO

These are eviction algorithms.

---

# So Does Cache Use LRU?

YES.

But:
LRU is only used for eviction.

NOT for updating cache.

---

# Example

Suppose:
Cache capacity = 3

Cache contains:
[A, B, C]

New item D arrives.

Cache full.

Now eviction policy runs.

If using LRU:

* remove least recently used item
* insert D

---

# How Cache Actually Updates

## Cache Aside Strategy (Most Common)

Flow:

1. Request comes
2. Check cache

If miss:

* fetch from DB
* update cache
* return response

Pseudo Flow:

if key in cache:
return cache[key]

data = DB.get(key)
cache[key] = data
return data

---

# What Happens When DB Data Changes?

Example:
User updates profile picture.

Need to update cache too.

Common approaches:

---

## Option 1: Update Cache Immediately

DB.update(data)
cache[key] = data

Called:
Write Through

---

## Option 2: Delete Cache

DB.update(data)
delete cache[key]

Next request:

* cache miss
* fresh DB fetch

Very common approach.

---

# Where LRU Comes In

Suppose cache memory is limited.

Example:
Cache can store only 1000 items.

New data arrives.

Need to remove something.

LRU decides:
"Remove least recently used item."

---

# Cache Lifecycle

## Step 1

Data requested.

---

## Step 2

If missing:

* fetch from DB
* store in cache

---

## Step 3

Cache becomes full.

---

## Step 4

Eviction policy runs.

Example:

* LRU removes old unused item

---

# Common Eviction Algorithms

## 1. LRU (Least Recently Used)

Remove least recently accessed item.

Most popular.

---

## 2. LFU (Least Frequently Used)

Remove least frequently accessed item.

---

## 3. FIFO

Remove oldest inserted item.

---

# Real World Systems

## Redis

Supports:

* LRU
* LFU
* Random eviction
* TTL based eviction

---

# Key Interview Insight

Cache systems need BOTH:

1. Update strategy
2. Eviction strategy

Example:

Update Strategy:

* Cache Aside

Eviction Strategy:

* LRU

These solve different problems.

---

# Simple Mental Model

Cache Update:
"How data enters/changes in cache"

LRU:
"How data leaves cache"
