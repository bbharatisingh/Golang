# 🏛️ System Design Master Notes & Resources
> **Note:** These notes are optimized for high-level understanding. Focus on the "why" and "trade-offs" rather than just the "how."

---

## 🏗️ 1. Core Architectural Concepts

### A. The CAP Theorem
In a distributed system, when a network partition occurs, you must choose between Consistency or Availability.
*   **Consistency (C):** Every node sees the same data at the same time.
*   **Availability (A):** Every request gets a response (success/failure) but no guarantee of latest data.
*   **Partition Tolerance (P):** The system continues to work despite dropped messages between nodes.
*   **Decision Rule:** Since network failures are inevitable, you usually choose **CP** (Consistency) for financial apps or **AP** (Availability) for social media.

### B. Load Balancing
*   **L4 Load Balancer:** Works at the Transport layer (TCP/UDP). Fast, but can't see "inside" the HTTP request.
*   **L7 Load Balancer:** Works at the Application layer. Can route based on URLs, cookies, or headers.
*   **Algorithms:** 
    *   **Round Robin:** Simple rotation.
    *   **Least Connections:** Best for long-lived tasks.
    *   **Consistent Hashing:** Minimizes reorganization when servers are added/removed (crucial for caches).



---

## 💾 2. Data Management & Scaling

### A. Database Scaling
1.  **Vertical Scaling:** Bigger machine (CPU/RAM). Has an upper limit and single point of failure.
2.  **Horizontal Scaling (Sharding):** Splitting data across multiple machines.
    *   *Key-Based Sharding:* `ShardID = UserID % NumberOfShards`.
    *   *Directory-Based:* A lookup table tells you where the data lives.
3.  **Replication:**
    *   *Leader-Follower:* One master for writes, multiple slaves for reads. Great for Read-Heavy systems.
    *   *Multi-Leader:* Good for multi-region setups but complex conflict resolution.

### B. Caching
*   **Cache-Aside:** Application checks cache first. On miss, it reads from DB and updates cache.
*   **Write-Through:** Data is written to cache and DB at the same time.
*   **Write-Back:** Data is written to cache only; DB is updated later in batches (high risk of data loss).



---

## 🌐 3. Communication & APIs

### A. API Protocols
*   **REST:** Standard, stateless, uses HTTP verbs. Best for public-facing APIs.
*   **gRPC:** Uses Protocol Buffers (binary). Faster and smaller than JSON. Best for internal microservices.
*   **WebSockets:** Bi-directional, persistent connection. Essential for real-time apps (Chat, Stock Tickers).

### B. Messaging & Queues
*   **Asynchronous Processing:** Using a Message Queue (Kafka, RabbitMQ) to handle heavy tasks (like image processing) without blocking the user.
*   **Fan-out:** Sending a single message to multiple subscribers (e.g., a tweet going to all followers' feeds).

---

## 🛠️ 4. System Design Question Templates

### 📝 Design a URL Shortener (TinyURL)
*   **API:** `POST /v1/shorten {url: "..."}` -> returns `tinyurl.com/xyz123`
*   **Storage:** NoSQL (KV Store like DynamoDB or Cassandra) is enough.
*   **Key Logic:** Use **Base62 encoding** (a-z, A-Z, 0-9) to generate short keys.
*   **Scaling:** Use a **Key Generation Service (KGS)** to pre-generate unique IDs so there are no collisions during high traffic.

### 📝 Design a Rate Limiter
*   **Algorithm:** **Token Bucket**. Each user has a bucket that fills with "tokens" at a fixed rate.
*   **Storage:** Use **Redis** for fast incrementing/expiring of counts.
*   **Level:** Can be at the API Gateway level or the Application level.

---

## 📚 5. Curated Resource List

### 🎥 High-Quality Video Channels
*   [ByteByteGo](https://www.youtube.com/@ByteByteGo) - Best visual explanations.
*   [Gaurav Sen](https://www.youtube.com/@gkcs) - Excellent for beginners and fundamental trade-offs.
*   [NeetCode IO](https://www.youtube.com/@NeetCode) - Great for the intersection of DSA and System Design.

### 📖 Essential Reading (Text)
*   [The System Design Primer (GitHub)](https://github.com/donnemartin/system-design-primer) - The gold standard of free resources.
*   [High Scalability Blog](http://highscalability.com/) - Real-world case studies from Netflix, Twitter, and Pinterest.
*   [Designing Data-Intensive Applications (DDIA)](https://www.oreilly.com/library/view/designing-data-intensive-applications/9781491903063/) - The "Bible" of System Design (Advanced).

### ✍️ Interactive Tools
*   [Excalidraw](https://excalidraw.com/) - Best tool for practicing your architecture sketches.
*   [Sam Who?](https://samwho.dev/) - Visual, interactive guides on Load Balancers and Hashing.

---
