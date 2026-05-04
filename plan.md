# 🗺️ DSA + System Design Recovery Roadmap
### May 5 – May 24, 2026
> ⚕️ **Note:** You are recovering from **cervical surgery (C5–C6)**. This plan is crafted with **generous rest windows**, short focused sessions (max 1.5–2 hrs/day), and zero screen-strain pressure. Listen to your body. Skip or swap days freely.

---

## 📋 General Daily Rules
- 🕐 **Max screen time**: 1.5–2 hours per day (split into 2 sessions of 45–60 min)
- 🛏️ **Rest between sessions**: minimum 2 hours
- 🧘 **Neck stretches** before/after every session (as approved by physiotherapist)
- 💧 Stay hydrated, maintain posture with neck support
- 📵 No late-night studying
- 🏗️ **Every active day includes a 15–20 min System Design mini-practice** (concept reading or diagram sketching — no heavy typing)

### 🗂️ Daily Session Structure (Active Days)
| Time Slot | Activity | Duration |
|-----------|----------|----------|
| 9:00 – 9:45 AM | DSA Session 1 (2–3 problems) | 45 min |
| 9:45 AM – 12:00 PM | Rest + Physio | ~2 hrs |
| 4:00 – 4:45 PM | DSA Session 2 (1–2 problems) | 45 min |
| 5:00 – 5:20 PM | System Design Mini-Practice | 20 min |
| 5:20 PM onwards | Full rest | — |

---

## 🗓️ Week 1: May 5 – May 10 | Foundation & Arrays

> **Theme**: Warm-up. Get back into the groove gently.

---

### 📅 Day 1 — Monday, May 5
**Topic**: Arrays & Two Pointers
**Time**: ~1.5 hrs total (2 sessions × 45 min)

#### 📖 Resources
- [Two Pointers - NeetCode](https://neetcode.io/roadmap)
- [Array Patterns - LeetCode Explore](https://leetcode.com/explore/learn/card/array-and-string/)

#### ❓ Questions (Easy → Medium)
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Two Sum | [LeetCode 1](https://leetcode.com/problems/two-sum/) | Easy |
| 2 | Best Time to Buy and Sell Stock | [LeetCode 121](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy |
| 3 | Container With Most Water | [LeetCode 11](https://leetcode.com/problems/container-with-most-water/) | Medium |
| 4 | 3Sum | [LeetCode 15](https://leetcode.com/problems/3sum/) | Medium |
| 5 | Trapping Rain Water | [LeetCode 42](https://leetcode.com/problems/trapping-rain-water/) | Hard |
| 6 | Product of Array Except Self | [LeetCode 238](https://leetcode.com/problems/product-of-array-except-self/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: What is a System Design Interview? Framework Overview
- Read: [How to approach SD interviews - ByteByteGo](https://blog.bytebytego.com/p/how-to-answer-system-design-interview)
- Practice: Write down the 4-step framework (Requirements → Estimation → Design → Deep Dive) in your own words
- Sketch: A simple client–server diagram with a load balancer

#### 🛏️ Rest Plan
- Morning session: 9:00 AM – 9:45 AM
- Rest + Physio: 9:45 AM – 12:00 PM
- Evening session: 4:00 PM – 4:45 PM
- SD Mini-Practice: 5:00 PM – 5:20 PM
- Full rest after 5:20 PM

---

### 📅 Day 2 — Tuesday, May 6
**Topic**: Sliding Window
**Time**: ~1.5 hrs

#### 📖 Resources
- [Sliding Window Technique - GeeksforGeeks](https://www.geeksforgeeks.org/window-sliding-technique/)
- [NeetCode Sliding Window Playlist](https://www.youtube.com/watch?v=jM2dhDPYMQM)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Maximum Subarray | [LeetCode 53](https://leetcode.com/problems/maximum-subarray/) | Medium |
| 2 | Longest Substring Without Repeating Characters | [LeetCode 3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | Medium |
| 3 | Minimum Window Substring | [LeetCode 76](https://leetcode.com/problems/minimum-window-substring/) | Hard |
| 4 | Permutation in String | [LeetCode 567](https://leetcode.com/problems/permutation-in-string/) | Medium |
| 5 | Fruit Into Baskets | [LeetCode 904](https://leetcode.com/problems/fruit-into-baskets/) | Medium |
| 6 | Sliding Window Maximum | [LeetCode 239](https://leetcode.com/problems/sliding-window-maximum/) | Hard |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: DNS & How the Web Works
- Read: [What happens when you type google.com - GitHub](https://github.com/alex/what-happens-when)
- Concept: Understand DNS resolution, TCP handshake, HTTP request lifecycle
- Sketch: DNS lookup flow → CDN → Load Balancer → App Server → DB

#### 🛏️ Rest Plan
- Same schedule as Day 1

---

### 📅 Day 3 — Wednesday, May 7
**Topic**: Hashing & HashMaps
**Time**: ~1.5 hrs

#### 📖 Resources
- [HashMap Deep Dive - CS Dojo](https://www.youtube.com/watch?v=sfWyugl4JWA)
- [Hashing - InterviewBit](https://www.interviewbit.com/courses/programming/topics/hashing/)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Valid Anagram | [LeetCode 242](https://leetcode.com/problems/valid-anagram/) | Easy |
| 2 | Group Anagrams | [LeetCode 49](https://leetcode.com/problems/group-anagrams/) | Medium |
| 3 | Longest Consecutive Sequence | [LeetCode 128](https://leetcode.com/problems/longest-consecutive-sequence/) | Medium |
| 4 | Top K Frequent Elements | [LeetCode 347](https://leetcode.com/problems/top-k-frequent-elements/) | Medium |
| 5 | Subarray Sum Equals K | [LeetCode 560](https://leetcode.com/problems/subarray-sum-equals-k/) | Medium |
| 6 | Design HashMap | [LeetCode 706](https://leetcode.com/problems/design-hashmap/) | Easy |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Caching Fundamentals
- Read: [Caching Strategies - AWS](https://aws.amazon.com/caching/) and [Redis vs Memcached - ByteByteGo](https://www.youtube.com/watch?v=a9__D53WsMs)
- Concepts: Cache-aside, Write-through, Write-behind, TTL, Eviction policies (LRU, LFU)
- Think: Where would you add a cache in a simple e-commerce app?

---

### 📅 Day 4 — Thursday, May 8
**⚕️ REST DAY — Light Review Only**

- No new problems
- Re-read solutions from Day 1–3 (reading only, 30 min max)
- Optional: Watch 1 explainer video lying down with audio focus
- Physio + neck rest priority

---

### 📅 Day 5 — Friday, May 9
**Topic**: Stacks & Queues
**Time**: ~1.5 hrs

#### 📖 Resources
- [Stack & Queue - Striver A2Z Sheet](https://takeuforward.org/strivers-a2z-dsa-course/strivers-a2z-dsa-course-sheet-2/)
- [Monotonic Stack - NeetCode](https://www.youtube.com/watch?v=Dq_ObZwTY_Q)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Valid Parentheses | [LeetCode 20](https://leetcode.com/problems/valid-parentheses/) | Easy |
| 2 | Daily Temperatures | [LeetCode 739](https://leetcode.com/problems/daily-temperatures/) | Medium |
| 3 | Largest Rectangle in Histogram | [LeetCode 84](https://leetcode.com/problems/largest-rectangle-in-histogram/) | Hard |
| 4 | Min Stack | [LeetCode 155](https://leetcode.com/problems/min-stack/) | Medium |
| 5 | Evaluate Reverse Polish Notation | [LeetCode 150](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | Medium |
| 6 | Car Fleet | [LeetCode 853](https://leetcode.com/problems/car-fleet/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Message Queues & Async Processing
- Read: [Message Queues - ByteByteGo](https://www.youtube.com/watch?v=oUJbuFMyFDo)
- Concepts: Kafka vs RabbitMQ, Producer-Consumer model, Dead Letter Queue, At-least-once delivery
- Think: How would you use a queue to handle 1M image upload notifications per day?

---

### 📅 Day 6 — Saturday, May 10
**Topic**: Binary Search
**Time**: ~1.5 hrs

#### 📖 Resources
- [Binary Search Patterns - LeetCode](https://leetcode.com/discuss/general-discussion/786126/python-powerful-ultimate-binary-search-template)
- [Binary Search - NeetCode Video](https://www.youtube.com/watch?v=s4DPM8ct1pI)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Binary Search | [LeetCode 704](https://leetcode.com/problems/binary-search/) | Easy |
| 2 | Find Minimum in Rotated Sorted Array | [LeetCode 153](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | Medium |
| 3 | Search in Rotated Sorted Array | [LeetCode 33](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Medium |
| 4 | Median of Two Sorted Arrays | [LeetCode 4](https://leetcode.com/problems/median-of-two-sorted-arrays/) | Hard |
| 5 | Search a 2D Matrix | [LeetCode 74](https://leetcode.com/problems/search-a-2d-matrix/) | Medium |
| 6 | Koko Eating Bananas | [LeetCode 875](https://leetcode.com/problems/koko-eating-bananas/) | Medium |
| 7 | Time Based Key-Value Store | [LeetCode 981](https://leetcode.com/problems/time-based-key-value-store/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Database Fundamentals — Sharding & Replication
- Read: [DB Sharding - ByteByteGo](https://www.youtube.com/watch?v=5faMjKuB9bc)
- Concepts: Horizontal vs Vertical sharding, Master-Slave replication, Read replicas, Indexing
- Think: How do you shard a Users table with 500M rows?

---

### 📅 Day 7 — Sunday, May 11
**⚕️ FULL REST DAY**
- No screen time for DSA
- Light walk if physio approves
- Mental reset

---

## 🗓️ Week 2: May 12 – May 17 | Trees, Graphs & Recursion

> **Theme**: Core data structures. These are the backbone of interviews.

---

### 📅 Day 8 — Monday, May 12
**Topic**: Linked Lists
**Time**: ~1.5 hrs

#### 📖 Resources
- [Linked List Masterclass - Striver](https://takeuforward.org/data-structure/top-linkedlist-interview-questions-structured-path/)
- [Linked List Patterns - NeetCode](https://neetcode.io/roadmap)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Reverse Linked List | [LeetCode 206](https://leetcode.com/problems/reverse-linked-list/) | Easy |
| 2 | Merge Two Sorted Lists | [LeetCode 21](https://leetcode.com/problems/merge-two-sorted-lists/) | Easy |
| 3 | Reorder List | [LeetCode 143](https://leetcode.com/problems/reorder-list/) | Medium |
| 4 | LRU Cache | [LeetCode 146](https://leetcode.com/problems/lru-cache/) | Medium |
| 5 | Linked List Cycle II | [LeetCode 142](https://leetcode.com/problems/linked-list-cycle-ii/) | Medium |
| 6 | Remove Nth Node From End | [LeetCode 19](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | Medium |
| 7 | Merge K Sorted Lists | [LeetCode 23](https://leetcode.com/problems/merge-k-sorted-lists/) | Hard |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Design a Key-Value Store (like Redis)
- Read: [Designing a KV Store - System Design Primer](https://github.com/donnemartin/system-design-primer#design-a-key-value-store-for-a-search-engine)
- Concepts: In-memory storage, persistence (RDB/AOF), replication, partitioning
- Sketch: Basic architecture of a distributed KV store with 3 nodes

---

### 📅 Day 9 — Tuesday, May 13
**Topic**: Binary Trees — Traversals & Basic Problems
**Time**: ~1.5 hrs

#### 📖 Resources
- [Binary Tree Playlist - NeetCode](https://www.youtube.com/watch?v=OnSn2XEQ4MY)
- [Tree Traversals - GeeksforGeeks](https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Invert Binary Tree | [LeetCode 226](https://leetcode.com/problems/invert-binary-tree/) | Easy |
| 2 | Maximum Depth of Binary Tree | [LeetCode 104](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | Easy |
| 3 | Binary Tree Level Order Traversal | [LeetCode 102](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Medium |
| 4 | Serialize and Deserialize Binary Tree | [LeetCode 297](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) | Hard |
| 5 | Diameter of Binary Tree | [LeetCode 543](https://leetcode.com/problems/diameter-of-binary-tree/) | Easy |
| 6 | Binary Tree Right Side View | [LeetCode 199](https://leetcode.com/problems/binary-tree-right-side-view/) | Medium |
| 7 | Path Sum II | [LeetCode 113](https://leetcode.com/problems/path-sum-ii/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: CDN & Content Delivery
- Read: [How CDN works - Cloudflare](https://www.cloudflare.com/learning/cdn/what-is-a-cdn/)
- Watch: [CDN explained - ByteByteGo](https://www.youtube.com/watch?v=RI9np1LWzqw)
- Concepts: Edge servers, Origin server, Cache invalidation, Push vs Pull CDN
- Think: When designing Netflix, where does CDN fit?

---

### 📅 Day 10 — Wednesday, May 14
**Topic**: BST + Tries
**Time**: ~1.5 hrs

#### 📖 Resources
- [BST Operations - Striver](https://takeuforward.org/data-structure/binary-search-tree/)
- [Trie Implementation - NeetCode](https://www.youtube.com/watch?v=oobqoCJlHA0)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Validate Binary Search Tree | [LeetCode 98](https://leetcode.com/problems/validate-binary-search-tree/) | Medium |
| 2 | Kth Smallest Element in a BST | [LeetCode 230](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) | Medium |
| 3 | Implement Trie (Prefix Tree) | [LeetCode 208](https://leetcode.com/problems/implement-trie-prefix-tree/) | Medium |
| 4 | Word Search II | [LeetCode 212](https://leetcode.com/problems/word-search-ii/) | Hard |
| 5 | Lowest Common Ancestor of BST | [LeetCode 235](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | Medium |
| 6 | Design Add and Search Words | [LeetCode 211](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Search Autocomplete System
- Watch: [Design Typeahead - ByteByteGo](https://www.youtube.com/watch?v=9HrnEMkMGeU)
- Concepts: Trie-based prefix search, top-K suggestions, caching popular queries
- Sketch: How a Trie powers autocomplete in a search bar with ranking

---

### 📅 Day 11 — Thursday, May 15
**⚕️ REST DAY — Light Review**
- Review tree problems solved so far (reading only)
- Watch 1 video max (30 min lying back)
- Neck exercises + rest

---

### 📅 Day 12 — Friday, May 16
**Topic**: Graphs — BFS / DFS
**Time**: ~1.5 hrs

#### 📖 Resources
- [Graph Series - Striver](https://takeuforward.org/graph/striver-graph-series-by-takeUforward/)
- [BFS/DFS Visual - VisuAlgo](https://visualgo.net/en/dfsbfs)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Number of Islands | [LeetCode 200](https://leetcode.com/problems/number-of-islands/) | Medium |
| 2 | Clone Graph | [LeetCode 133](https://leetcode.com/problems/clone-graph/) | Medium |
| 3 | Pacific Atlantic Water Flow | [LeetCode 417](https://leetcode.com/problems/pacific-atlantic-water-flow/) | Medium |
| 4 | Course Schedule (Topological Sort) | [LeetCode 207](https://leetcode.com/problems/course-schedule/) | Medium |
| 5 | Surrounded Regions | [LeetCode 130](https://leetcode.com/problems/surrounded-regions/) | Medium |
| 6 | Rotting Oranges | [LeetCode 994](https://leetcode.com/problems/rotting-oranges/) | Medium |
| 7 | Course Schedule II (All orderings) | [LeetCode 210](https://leetcode.com/problems/course-schedule-ii/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Design a Web Crawler
- Watch: [Web Crawler Design - ByteByteGo](https://www.youtube.com/watch?v=BKZxZwUgL3Y)
- Concepts: BFS traversal on URLs, URL deduplication, politeness policy, robots.txt, distributed crawling
- Sketch: Crawler pipeline — URL Frontier → Fetcher → Parser → URL Extractor → Storage

---

### 📅 Day 13 — Saturday, May 17
**Topic**: Graph Advanced — Dijkstra, Union-Find
**Time**: ~1.5 hrs

#### 📖 Resources
- [Dijkstra - Striver](https://www.youtube.com/watch?v=V6H1qAeB-l4)
- [Union Find - NeetCode](https://www.youtube.com/watch?v=ayW5B2W9hfo)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Network Delay Time (Dijkstra) | [LeetCode 743](https://leetcode.com/problems/network-delay-time/) | Medium |
| 2 | Redundant Connection (Union-Find) | [LeetCode 684](https://leetcode.com/problems/redundant-connection/) | Medium |
| 3 | Word Ladder | [LeetCode 127](https://leetcode.com/problems/word-ladder/) | Hard |
| 4 | Swim in Rising Water | [LeetCode 778](https://leetcode.com/problems/swim-in-rising-water/) | Hard |
| 5 | Alien Dictionary | [LeetCode 269](https://leetcode.com/problems/alien-dictionary/) | Hard |
| 6 | Min Cost to Connect All Points (MST) | [LeetCode 1584](https://leetcode.com/problems/min-cost-to-connect-all-points/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Design Google Maps / Proximity Service
- Watch: [Proximity Service - ByteByteGo](https://www.youtube.com/watch?v=M4lR_Va97cQ)
- Concepts: Geohashing, Quadtree, PostGIS, nearest-neighbor search
- Sketch: How would you find all restaurants within 5 km of a user in real time?

---

### 📅 Day 14 — Sunday, May 18
**⚕️ FULL REST DAY**
- No DSA
- Recovery priority
- Optional: Listen to a system design podcast/talk (audio only, no screen)

---

## 🗓️ Week 3: May 19 – May 24 | DP + System Design

> **Theme**: Dynamic Programming & System Design. The power topics. Keep sessions short and focused.

---

### 📅 Day 15 — Monday, May 19
**Topic**: Dynamic Programming — 1D
**Time**: ~1.5 hrs

#### 📖 Resources
- [DP Patterns - LeetCode Discuss](https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns)
- [DP Series - NeetCode](https://www.youtube.com/watch?v=73r3KWiEvyk)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Climbing Stairs | [LeetCode 70](https://leetcode.com/problems/climbing-stairs/) | Easy |
| 2 | House Robber | [LeetCode 198](https://leetcode.com/problems/house-robber/) | Medium |
| 3 | Longest Palindromic Substring | [LeetCode 5](https://leetcode.com/problems/longest-palindromic-substring/) | Medium |
| 4 | Decode Ways | [LeetCode 91](https://leetcode.com/problems/decode-ways/) | Medium |
| 5 | House Robber II | [LeetCode 213](https://leetcode.com/problems/house-robber-ii/) | Medium |
| 6 | Palindromic Substrings | [LeetCode 647](https://leetcode.com/problems/palindromic-substrings/) | Medium |
| 7 | Coin Change | [LeetCode 322](https://leetcode.com/problems/coin-change/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Design a Payment System
- Watch: [Payment System Design - ByteByteGo](https://www.youtube.com/watch?v=olfaBgJrUBI)
- Concepts: Idempotency keys, exactly-once delivery, double-entry ledger, PCI-DSS basics
- Think: How do you prevent double charges if the network times out after payment request?

---

### 📅 Day 16 — Tuesday, May 20
**Topic**: Dynamic Programming — 2D (Grids & Strings)
**Time**: ~1.5 hrs

#### 📖 Resources
- [Grid DP - Striver](https://takeuforward.org/data-structure/grid-unique-paths/)
- [Knapsack & LCS - Aditya Verma](https://www.youtube.com/playlist?list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go)

#### ❓ Questions
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Unique Paths | [LeetCode 62](https://leetcode.com/problems/unique-paths/) | Medium |
| 2 | Longest Common Subsequence | [LeetCode 1143](https://leetcode.com/problems/longest-common-subsequence/) | Medium |
| 3 | Edit Distance | [LeetCode 72](https://leetcode.com/problems/edit-distance/) | Hard |
| 4 | Burst Balloons | [LeetCode 312](https://leetcode.com/problems/burst-balloons/) | Hard |
| 5 | 0/1 Knapsack | [GFG](https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/) | Medium |
| 6 | Distinct Subsequences | [LeetCode 115](https://leetcode.com/problems/distinct-subsequences/) | Hard |
| 7 | Interleaving String | [LeetCode 97](https://leetcode.com/problems/interleaving-string/) | Medium |

#### 🏗️ System Design Mini-Practice (20 min)
**Topic**: Design a Distributed File Storage (like S3/Google Drive)
- Watch: [Design Google Drive - ByteByteGo](https://www.youtube.com/watch?v=jLM1nGiQjXT)
- Concepts: Object storage vs block storage, chunking large files, deduplication, metadata service
- Sketch: Upload flow — Client → API Gateway → Metadata DB → Block Storage → CDN

---

### 📅 Day 17 — Wednesday, May 21
**Topic**: System Design — Fundamentals
**Time**: ~1.5 hrs (reading + diagrams, no code)

#### 📖 Resources
- [System Design Primer - GitHub](https://github.com/donnemartin/system-design-primer)
- [Grokking System Design - educative.io](https://www.educative.io/courses/grokking-the-system-design-interview)
- [ByteByteGo Newsletter](https://blog.bytebytego.com/)

#### 📝 Topics to Cover
| # | Topic | Resource |
|---|-------|----------|
| 1 | CAP Theorem | [CAP Explained - ByteByteGo](https://www.youtube.com/watch?v=_RbsFXWRZ10) |
| 2 | Load Balancing | [Load Balancer - Cloudflare](https://www.cloudflare.com/learning/performance/what-is-load-balancing/) |
| 3 | Caching (Redis, CDN) | [Caching Strategies - AWS](https://aws.amazon.com/caching/) |
| 4 | SQL vs NoSQL | [SQL vs NoSQL - ByteByteGo](https://www.youtube.com/watch?v=W2Z7fbCLSTw) |
| 5 | Consistent Hashing | [Consistent Hashing - Tom Scott](https://www.youtube.com/watch?v=tHEyzVbl4bg) |
| 6 | API Gateway & REST vs gRPC | [API Gateway - ByteByteGo](https://www.youtube.com/watch?v=vgQEGTy58Ow) |
| 7 | Microservices vs Monolith | [Martin Fowler](https://martinfowler.com/articles/microservices.html) |
| 8 | Event-Driven Architecture | [EDA - AWS](https://aws.amazon.com/event-driven-architecture/) |

#### 🧠 SD Concept Practice Questions (answer in writing, ~5 min each)
1. You have a monolithic app with 100K daily users. Traffic is doubling every 3 months. What steps do you take?
2. Explain the trade-offs between consistency and availability using a real-world example.
3. How would you implement rate limiting at 10K requests/sec across 5 servers?
4. Why is consistent hashing better than modulo-based hashing in distributed systems?

---

### 📅 Day 18 — Thursday, May 22
**Topic**: System Design — Design Problems (Part 1)
**Time**: ~1.5 hrs

#### 📖 Resources
- [Design URL Shortener - NeetCode](https://www.youtube.com/watch?v=fMZMm_0ZhK4)
- [Design Pastebin - System Design Primer](https://github.com/donnemartin/system-design-primer/blob/master/solutions/system_design/pastebin/README.md)
- [High Scalability Blog](http://highscalability.com/)

#### 🏗️ Systems to Design (Practice on paper / whiteboard)
| # | System | Key Concepts | Resource |
|---|--------|-------------|----------|
| 1 | URL Shortener (TinyURL) | Hashing, KV Store, Redirection | [Video](https://www.youtube.com/watch?v=fMZMm_0ZhK4) |
| 2 | Rate Limiter | Token Bucket, Sliding Window Log | [ByteByteGo](https://www.youtube.com/watch?v=FU4WlwfS3G0) |
| 3 | Notification System | Message Queue, Fan-out, Push/Pull | [ByteByteGo](https://www.youtube.com/watch?v=bBTPZ9NdSk8) |
| 4 | Pastebin / Code Snippet Store | Object Storage, Expiry, Unique ID gen | [System Design Primer](https://github.com/donnemartin/system-design-primer/blob/master/solutions/system_design/pastebin/README.md) |
| 5 | Job Scheduler | Priority Queue, Cron, Worker Pool | [ByteByteGo](https://www.youtube.com/watch?v=ta5x62cDZx0) |

#### 🧠 Deep-Dive Questions for Day 18
1. Design the database schema for TinyURL — what tables, indexes, and data types would you use?
2. How would you make the URL shortener handle 100K redirects/sec with < 10ms latency?
3. How do you prevent a single bad actor from abusing the rate limiter across multiple IPs?
4. What happens if the notification service goes down — how do you ensure delivery?

---

### 📅 Day 19 — Friday, May 23
**Topic**: System Design — Design Problems (Part 2)
**Time**: ~1.5 hrs

#### 📖 Resources
- [Designing Netflix - NeetCode](https://www.youtube.com/watch?v=psQzyFfsUGU)
- [Designing Twitter - ByteByteGo](https://www.youtube.com/watch?v=wYk0xPP_P_8)
- [Designing WhatsApp - Gaurav Sen](https://www.youtube.com/watch?v=vvhC64hQZMk)

#### 🏗️ Systems to Design
| # | System | Key Concepts | Resource |
|---|--------|-------------|----------|
| 1 | News Feed / Twitter | Fan-out, Timeline, Sharding | [Video](https://www.youtube.com/watch?v=wYk0xPP_P_8) |
| 2 | Chat System (WhatsApp) | WebSocket, Message Queue, Presence | [Video](https://www.youtube.com/watch?v=vvhC64hQZMk) |
| 3 | Distributed Cache (Redis) | Eviction Policies, Replication | [Redis Docs](https://redis.io/docs/manual/replication/) |
| 4 | Video Streaming (YouTube/Netflix) | CDN, Adaptive Bitrate, Metadata DB | [NeetCode](https://www.youtube.com/watch?v=psQzyFfsUGU) |
| 5 | Ride-Sharing (Uber) | Geo-indexing, Matching, Surge Pricing | [ByteByteGo](https://www.youtube.com/watch?v=lsKU38RKQSo) |

#### 🧠 Deep-Dive Questions for Day 19
1. Fan-out on write vs fan-out on read for a Twitter timeline — which do you choose at 10M users vs 1B users?
2. How do you show message delivery receipts (single tick / double tick / blue tick) in a chat system?
3. How would you handle a celebrity user (50M followers) posting a tweet — what breaks and how do you fix it?
4. How does YouTube decide which video quality to serve and when to switch?

---

### 📅 Day 20 — Saturday, May 24
**🏁 Revision + Mock Interview Day**
**Time**: ~2 hrs (special final day)

#### 📝 Revision Checklist
- [ ] Review all DSA patterns covered (Arrays, Sliding Window, Hashing, Stacks, Binary Search, Linked Lists, Trees, Graphs, DP)
- [ ] Re-attempt 2 hard problems from earlier weeks without looking at solutions
- [ ] Whiteboard 1 full system design end-to-end (pick one from the list below)
- [ ] Write down your own cheat-sheet for patterns
- [ ] Review all SD deep-dive questions from Days 17–19
- [ ] Write a 1-page summary of your go-to SD framework

#### 🎯 Final Day System Design Choices (pick 1 to do fully)
| Option | System | Key Areas |
|--------|--------|-----------|
| A | Design Instagram | Media storage, Feed generation, Follower graph |
| B | Design Booking.com | Search, Inventory locking, Double-booking prevention |
| C | Design a Stock Trading Platform | Low latency, Order matching engine, Real-time feeds |
| D | Design an Ad Click Aggregator | High write throughput, Aggregation, Fraud detection |

#### 🧠 Final DSA Bonus Problems (attempt any 2–3)
| # | Problem | Link | Difficulty |
|---|---------|------|------------|
| 1 | Regular Expression Matching | [LeetCode 10](https://leetcode.com/problems/regular-expression-matching/) | Hard |
| 2 | Largest Rectangle in Histogram | [LeetCode 84](https://leetcode.com/problems/largest-rectangle-in-histogram/) | Hard |
| 3 | N-Queens | [LeetCode 51](https://leetcode.com/problems/n-queens/) | Hard |
| 4 | Sudoku Solver | [LeetCode 37](https://leetcode.com/problems/sudoku-solver/) | Hard |
| 5 | Minimum Cost to Cut a Stick | [LeetCode 1547](https://leetcode.com/problems/minimum-cost-to-cut-a-stick/) | Hard |

#### 🎯 Suggested Mock Resources
| Resource | Link |
|----------|------|
| NeetCode Practice | [neetcode.io/practice](https://neetcode.io/practice) |
| Pramp (Free Peer Mock) | [pramp.com](https://www.pramp.com/) |
| interviewing.io | [interviewing.io](https://interviewing.io/) |
| LeetCode Mock Assessment | [LeetCode Assessments](https://leetcode.com/assessment/) |

#### 🛏️ End-of-Plan Rest
- Take the evening completely off
- Celebrate your recovery and consistency! 🎉

---

## 📚 Master Resource List

### DSA
| Resource | Type | Link |
|----------|------|------|
| NeetCode Roadmap | Video + Problems | [neetcode.io](https://neetcode.io/roadmap) |
| Striver A2Z Sheet | Structured Sheet | [takeuforward.org](https://takeuforward.org/strivers-a2z-dsa-course/strivers-a2z-dsa-course-sheet-2/) |
| LeetCode Explore Cards | Interactive | [leetcode.com/explore](https://leetcode.com/explore/) |
| CP-Algorithms | Deep Dive Theory | [cp-algorithms.com](https://cp-algorithms.com/) |
| Visualgo | Visualizations | [visualgo.net](https://visualgo.net/) |

### System Design
| Resource | Type | Link |
|----------|------|------|
| System Design Primer | GitHub Repo | [github.com/donnemartin](https://github.com/donnemartin/system-design-primer) |
| ByteByteGo Blog | Blog + Videos | [blog.bytebytego.com](https://blog.bytebytego.com/) |
| Grokking SD Interview | Course | [educative.io](https://www.educative.io/courses/grokking-the-system-design-interview) |
| Gaurav Sen YouTube | Videos | [youtube.com/c/GauravSensei](https://www.youtube.com/c/GauravSensei) |
| High Scalability | Case Studies | [highscalability.com](http://highscalability.com/) |
| AWS Architecture Center | Real World | [aws.amazon.com/architecture](https://aws.amazon.com/architecture/) |

---

## 📊 Weekly Summary

| Week | Dates | Topics | SD Mini-Practice | Rest Days |
|------|-------|--------|------------------|-----------|
| Week 1 | May 5–11 | Arrays, Sliding Window, Hashing, Stacks, Binary Search (30+ problems) | SD Framework, DNS, Caching, Message Queues, DB Sharding | Thu, Sun |
| Week 2 | May 12–18 | Linked Lists, Trees, BST, Tries, Graphs (35+ problems) | KV Store, CDN, Autocomplete, Web Crawler, Proximity Service | Thu, Sun |
| Week 3 | May 19–24 | DP 1D & 2D (30+ problems), SD Fundamentals + Full System Designs | Payment System, File Storage, Twitter/WhatsApp/Netflix/Uber | Sat eve |

---

## ⚕️ Recovery Guidelines (C5-C6 Post-Surgery)

> ⚠️ This is a study plan, not medical advice. Always follow your neurosurgeon's and physiotherapist's instructions.

- **Screen distance**: Keep monitor/screen at eye level to avoid neck flexion
- **Posture**: Use a cervical pillow/neck roll while sitting
- **Max sitting duration**: 20–30 min at a stretch, then move/lie down
- **Reading position**: Prefer lying at 30–45° incline with screen elevated
- **Breaks**: Every 20 min — look away (20-20-20 rule)
- **Signs to stop**: Any tingling in arms/hands, worsening headache, neck stiffness → stop immediately and rest

---

*🗓️ Plan Created: May 4, 2026 | Recovery Roadmap v1.0*
*💪 Wishing you a smooth recovery and great interviews ahead!*
