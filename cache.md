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

# Important Concepts to Learn Next

1. Redis
2. Load Balancer
3. Database Indexing
4. CDN
5. Replication
