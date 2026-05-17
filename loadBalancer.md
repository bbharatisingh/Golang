# Load Balancer Notes

# What is a Load Balancer?

A Load Balancer distributes incoming traffic across multiple servers.

Purpose:

* prevent server overload
* improve scalability
* improve availability
* avoid single point of failure

---

# Why Load Balancer is Needed

Suppose:

1 server can handle:
10,000 requests/sec

Traffic becomes:
100,000 requests/sec

Single server crashes or becomes slow.

Solution:
Add multiple servers.

Need something to distribute traffic.

That component is:
Load Balancer

---

# Basic Architecture

Users
↓
Load Balancer
↓   ↓   ↓
S1  S2  S3

Flow:

1. User sends request
2. Load balancer receives request
3. LB forwards request to one server

---

# Benefits of Load Balancer

## 1. Scalability

Can add more servers dynamically.

---

## 2. High Availability

If one server dies:

* LB routes traffic to healthy servers

---

## 3. Fault Tolerance

System continues working despite failures.

---

## 4. Better Performance

Traffic spread evenly.

No server overloaded.

---

# Common Algorithms

# 1. Round Robin

Requests distributed sequentially.

Example:

Req1 -> S1
Req2 -> S2
Req3 -> S3
Req4 -> S1

Simple and common.

---

# 2. Least Connections

Send request to server with fewest active connections.

Useful when:

* request durations vary

---

# 3. IP Hash

Hash client IP.

Same user routed to same server.

Useful for:

* sticky sessions

---

# Health Checks

Load balancer periodically checks:

"Is server alive?"

If server unhealthy:

* stop sending traffic to it

Very important concept.

---

# Sticky Sessions

Same user routed to same server.

Needed when:
session stored locally on server.

Problem:
If server crashes:
session lost.

Modern systems avoid sticky sessions using:

* Redis session storage
* stateless servers

---

# Stateless Servers

Preferred design.

Any server can handle any request.

Session/data stored externally:

* Redis
* DB

Benefits:

* easier scaling
* better fault tolerance

---

# Reverse Proxy

Load balancer acts as reverse proxy.

Meaning:
Client talks only to LB.

Servers hidden internally.

Benefits:

* security
* centralized routing
* SSL termination

---

# Types of Load Balancer

## 1. Layer 4 (Transport Layer)

Routes using:

* IP
* TCP/UDP

Faster.

---

## 2. Layer 7 (Application Layer)

Routes using:

* URL
* headers
* cookies

Smarter routing.

Example:

* /images -> image servers
* /api -> backend servers

---

# Software Load Balancers

Examples:

* Nginx
* HAProxy

Cloud:

* AWS ELB
* Google Cloud Load Balancer

---

# Horizontal Scaling

Load balancer enables:

Horizontal Scaling

Meaning:
add more servers instead of increasing server power.

Architecture:

Users
↓
LB
↓ ↓ ↓ ↓ ↓
Many Servers

---

# Important Interview Problems Solved by LB

1. Single server overload
2. High traffic spikes
3. Server failures
4. Scalability issues

---

# Typical System Design Flow

Client
↓
Load Balancer
↓
Application Servers
↓
Redis Cache
↓
Database

---

# Important Tradeoffs

Advantages:

* scalability
* high availability
* fault tolerance

Disadvantages:

* extra infrastructure
* LB itself can become bottleneck

Solution:
Use multiple load balancers.

---

# Most Important Interview Concepts

1. Round Robin
2. Health Checks
3. Stateless Servers
4. Sticky Sessions
5. Reverse Proxy
6. Horizontal Scaling
7. Fault Tolerance

---

# Key Interview Insight

Load Balancer enables:

* horizontal scaling
* high availability

Without LB:
large-scale systems are impossible.
