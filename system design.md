# 🗺️ DSA + System Design Recovery Roadmap (Updated)
### May 8 – May 24, 2026
> ⚕️ **Post-Op Note:** Optimized for C5–C6 recovery. Focus on eye-level monitors and lumbar support.

---

## 📋 Daily Protocol
- 🕐 **Total Screen Time**: 1.5 – 2 Hours.
- 🧘 **Neck Check**: Reset posture every 20 minutes.
- 🎧 **Audio First**: If neck strain occurs, switch to System Design podcasts/videos and listen while lying flat.

---

## 🗓️ Week 1: Foundation & High-Level Components

### 📅 Friday, May 8 (Today)
**⚕️ STATUS: ACTIVE RECOVERY**
*   **DSA**: None (Rest Day).
*   **System Design Mini-Practice (20 min)**: Audio-only immersion.
    *   **Topic**: The 4-Step System Design Interview Framework.
    *   **Resource**: [ByteByteGo: How to Answer SD Questions](https://blog.bytebytego.com/p/how-to-answer-system-design-interview)
    *   **Goal**: Internalize the flow: Requirements → API Design → Data Schema → High-Level Design.

### 📅 Saturday, May 9
**Topic**: Binary Search + Load Balancing
*   **DSA (45 min)**: LeetCode 704 (Binary Search), 153 (Min in Rotated Array).
*   **System Design (30 min)**: 
    *   **Topic**: Load Balancers & Reverse Proxies.
    *   **Resource**: [Sam Who: Load Balancing (Interactive)](https://samwho.dev/load-balancing/)
    *   **Concept**: L4 vs L7 balancing, sticky sessions, health checks.

### 📅 Sunday, May 10
**⚕️ FULL REST DAY**
*   No screen time. 
*   Mental Exercise: Visualize the traffic flow from a mobile app to a database.

---

## 🗓️ Week 2: Scaling & Data Patterns

### 📅 Monday, May 11
**Topic**: Linked Lists + Database Fundamentals
*   **DSA (45 min)**: LeetCode 206 (Reverse LL), 21 (Merge Sorted LL).
*   **System Design (30 min)**:
    *   **Topic**: SQL vs NoSQL & ACID vs BASE.
    *   **Resource**: [SQL vs NoSQL Deep Dive](https://www.prisma.io/dataguide/types/relational/comparing-sql-and-nosql)
    *   **Sketch**: Draw a simple Relational Schema for an E-commerce user.

### 📅 Tuesday, May 12
**Topic**: Binary Trees + Sharding & Replication
*   **DSA (45 min)**: LeetCode 104 (Max Depth), 226 (Invert Tree).
*   **System Design (30 min)**:
    *   **Topic**: Horizontal Scaling & Data Partitioning.
    *   **Resource**: [Database Sharding - DigitalOcean](https://www.digitalocean.com/community/tutorials/understanding-database-sharding)
    *   **Concept**: Consistent Hashing, Range Sharding, Hotspots.

### 📅 Wednesday, May 13
**Topic**: BST/Tries + Caching
*   **DSA (45 min)**: LeetCode 98 (Validate BST), 208 (Implement Trie).
*   **System Design (30 min)**:
    *   **Topic**: Caching Strategies & CDNs.
    *   **Resource**: [Caching Patterns](https://aws.amazon.com/caching/)
    *   **Concept**: Cache-aside, Write-through, Eviction (LRU/LFU).

### 📅 Thursday, May 14
**⚕️ LIGHT REVIEW DAY**
*   **Activity**: Re-read notes from Mon-Wed.
*   **SD Focus**: API Protocols (REST vs gRPC vs GraphQL).

---

## 🗓️ Week 3: Deep Dives & Full Designs

### 📅 Friday, May 15
**Topic**: Stacks/Queues + Rate Limiting
*   **DSA**: LeetCode 20 (Valid Parentheses), 739 (Daily Temps).
*   **System Design**: **Design a Rate Limiter**.
    *   **Concept**: Token Bucket vs Leaking Bucket. [Resource](https://stripe.com/blog/rate-limiters).

### 📅 Saturday, May 16
**Topic**: Graphs + URL Shortener
*   **DSA**: LeetCode 200 (Number of Islands).
*   **System Design**: **Design TinyURL**.
    *   **Concept**: Key Generation Service (KGS), Base62 encoding.

### 📅 Sunday, May 17
**⚕️ FULL REST DAY**

### 📅 Monday, May 18
**Topic**: Advanced Graphs + Chat Systems
*   **DSA**: LeetCode 743 (Network Delay).
*   **System Design**: **Design WhatsApp**.
    *   **Concept**: WebSockets, Message Queues, Presence Servers.

### 📅 Tuesday, May 19
**Topic**: DP 1D + Proximity Services
*   **DSA**: LeetCode 70 (Climbing Stairs), 322 (Coin Change).
*   **System Design**: **Design Uber/Yelp**.
    *   **Concept**: Geospatial Indexing (Quadtrees, Geohashes).

### 📅 Wednesday, May 20
**Topic**: DP 2D + Video Streaming
*   **DSA**: LeetCode 1143 (LCS).
*   **System Design**: **Design Netflix/YouTube**.
    *   **Concept**: Video Transcoding, Adaptive Bitrate Streaming (DASH/HLS).

### 📅 Thursday, May 21
**Topic**: Web Crawlers
*   **System Design**: **Design a Web Crawler**.
    *   **Concept**: Politeness (Robots.txt), URL Frontier, Deduplication.

### 📅 Friday, May 22
**Topic**: Social Media Feeds
*   **System Design**: **Design Twitter Newsfeed**.
    *   **Concept**: Fan-out on Write vs. Fan-out on Read.

### 📅 Saturday, May 23
**Topic**: Distributed Systems Special
*   **SD Focus**: CAP Theorem, Distributed Transactions (2PC/Sagas).

### 📅 Sunday, May 24
**🏁 THE FINISH LINE**
*   **Mock**: Full end-to-end design of **Instagram** (on paper/whiteboard).
*   **Review**: Top 10 most missed DSA problems.
