# Netflix + CDN + Load Balancer Flow Notes

# Goal

Serve millions of users globally with:

* low latency
* high availability
* smooth video streaming

Netflix uses:

1. Load Balancers
2. Horizontal Scaling
3. CDN (Open Connect)
4. Distributed Servers

---

# High Level Architecture

```
            Users
               ↓
      Global Load Balancer
               ↓
    Regional Load Balancer
         ↓      ↓      ↓
        S1      S2     S3
         ↓
    Netflix Backend APIs
         ↓
 Redis / Databases / Metadata
         ↓
 CDN Edge Servers (Video Chunks)
         ↓
         User
```

---

# Step-by-Step Flow

# Step 1 — User Opens Netflix

User clicks:
"Play Movie"

Request goes to:
Global Load Balancer

Purpose:

* route user to nearest region
* distribute traffic

---

# Step 2 — Regional Load Balancer

Regional LB chooses one application server.

Example:

* S1 busy
* S2 less loaded

Request routed to:
S2

Algorithms:

* Round Robin
* Least Connections
* Geo Routing

---

# Step 3 — Backend Processing

Application server handles:

* authentication
* subscription verification
* recommendation APIs
* movie metadata

Important:
Application server does NOT stream full video.

---

# Step 4 — Find Nearest CDN

Netflix identifies nearest CDN edge server.

Example:
India user ->
Mumbai CDN server

Purpose:
reduce latency.

---

# Step 5 — Video Streaming from CDN

Video split into small chunks.

CDN streams chunks directly to user.

Flow:

User
↓
Nearest CDN Edge Server
↓
Video Chunks Streamed

Benefits:

* smooth playback
* less buffering
* lower latency

---

# Step 6 — Cache Miss in CDN

If CDN lacks video chunk:

CDN
↓
Origin Storage Server
↓
Fetch chunk
↓
Store locally
↓
Serve user

This is:
CDN Cache Miss

---

# Horizontal Scaling in Netflix

Suppose traffic spikes after new release.

Netflix adds:
S4 S5 S6 ...

Load balancer automatically distributes requests.

This is:
Horizontal Scaling

---

# Failure Handling

Suppose:
S2 crashes.

Health checks detect failure.

Load balancer reroutes traffic:

* S1
* S3

Users continue streaming.

This provides:

* High Availability
* Fault Tolerance

---

# Why CDN is Critical

Without CDN:
all videos stream from central servers.

Problems:

* huge bandwidth cost
* high latency
* buffering globally

CDN solves:

* geographic caching
* faster delivery
* reduced backend load

---

# Key Components Summary

## Load Balancer

Distributes requests.

---

## App Servers

Handle APIs/business logic.

---

## Redis

Stores cache/session data.

---

## CDN

Streams video chunks globally.

---

## Database

Stores metadata/user info.

# Key System Design Insights

1. App servers are stateless
2. CDN handles heavy traffic
3. Load balancer enables scaling
4. Redis reduces DB load
5. Horizontal scaling handles millions of users

---

# Interview Keywords

* Horizontal Scaling
* Stateless Servers
* CDN Edge Servers
* Cache Hit/Miss
* Load Balancer
* Fault Tolerance
* High Availability
* Geo Routing
* Video Chunking
* Distributed Systems
