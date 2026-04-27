# MineLink: Product Roadmap & Implementation Strategy

This roadmap outlines the path from architectural specification to a production-ready, hybrid-cloud Fleet Management System (FMS). It prioritizes **Test-Driven Development (TDD)**, **Zero-Trust Security**, and **High-Availability Edge Operations**.

---

## Phase 1: Specification & Contract-First Development
*Goal: Eliminate integration friction by defining the "Rules of Engagement" before a single line of service code is written.*

1.  **Schema Definition (AsyncAPI/Protobuf):**
    *   Define all MQTT and gRPC contracts in a central repository.
    *   Ensure strong typing for telemetry (`truck/status`) and C2 commands (`c2/tactical/assignment`).
2.  **TDD: Mocking the Mesh:**
    *   Develop a "System Mock" suite that simulates the Edge Broker and Redis.
    *   **Success Metric:** A test runner can simulate a full shift rotation (8 hours compressed) and verify that Redis state matches the expected MQTT command sequence.
3.  **Security Baseline:**
    *   Define the PKI (Public Key Infrastructure) for mTLS.
    *   Draft the `NetworkPolicy` manifests to isolate the Edge Redis from public ingress.

---

## Phase 2: Hybrid Infrastructure Provisioning (IaC)
*Goal: Automate the deployment of the "Stretched Cluster" using Infrastructure as Code.*

1.  **Terraform: Cloud & Edge Foundation:**
    *   Provision the Cloud EKS/GKE cluster and the Edge K3s nodes.
    *   Configure the VPN/SD-WAN tunnel between the Pit and the Office HQ.
2.  **State Layer Provisioning:**
    *   **Edge:** Deploy Redis with AOF (Append Only File) persistence enabled for crash recovery.
    *   **Cloud:** Provision TimescaleDB for high-ingestion telemetry storage.
3.  **Scaling & Security Hardening:**
    *   Apply the **2-Pod Constraint** via K8s Affinity/Anti-Affinity rules.
    *   Implement **Vault/External Secrets** to manage mTLS certificates and DB credentials.

---

## Phase 3: Service Implementation (The Agentic Core)
*Goal: Build the hierarchical C2 services using the specified language stack.*

1.  **The Supervisor Agent (Go - Edge):**
    *   **Focus:** Concurrency and Orchestration.
    *   Implement the Redis-backed state machine for shift management.
    *   **TDD:** Unit tests for "Shift Hand-off" logic with zero data loss.
2.  **The Operator Agent (C++ - Edge):**
    *   **Focus:** High-performance load generation and Actor Model.
    *   Implement the telemetry loop and resource depletion logic (fuel/tire wear).
    *   **TDD:** Memory-leak profiling and stress-testing the MQTT bridge at 1000+ msgs/sec.
3.  **The Dispatch Agent (C# - Cloud):**
    *   **Focus:** Optimization and Global Strategy.
    *   Implement the goal-seeking algorithm to balance fleet throughput.
    *   **TDD:** Integration tests verifying that "Global Intents" from C# are correctly received as "Tactical Tasks" by Go.

---

## Phase 4: Validation & Resilience Testing
*Goal: Prove the system can survive the harsh environment of an industrial mine.*

1.  **Chaos Engineering (The "Cut the Fiber" Test):**
    *   Simulate a network partition between Edge and Cloud.
    *   **Requirement:** The Pit (Supervisor + Operators) must continue 8-hour shifts autonomously using Edge Redis state.
2.  **Synthetic Load Stress:**
    *   Scale the Operator Agent threads to simulate a 200-vehicle fleet.
    *   Monitor the TimescaleDB ingestion lag and CPU pressure on the Edge K3s nodes.
3.  **Security Audit:**
    *   Verify that unencrypted MQTT traffic is rejected by the broker.
    *   Confirm that a compromised Operator thread cannot access the Cloud Postgres credentials.

---

## Hybrid Cloud Best Practices (Scaling & Security)

| Dimension | Edge Strategy (The Pit) | Cloud Strategy (The Office) |
| :--- | :--- | :--- |
| **Scaling** | **Vertical:** Increase CPU/RAM for the 2 fixed pods. | **Horizontal:** Auto-scale pods based on API demand. |
| **State** | **Redis (Volatile/Tactical):** Low latency, local. | **Postgres (Persistent/Strategic):** Long-term analytics. |
| **Security** | **Network Isolation:** No direct public access. | **WAF & OAuth2:** Secure public management API. |
| **Resilience** | **Autonomy:** Operates during Cloud downtime. | **Global View:** Reconciles state after Edge reconnects. |
