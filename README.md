# MineLink FMS Boilerplate

MineLink is a portfolio project - An MVP for a distributed, mission-critical Fleet Management System designed for industrial mining operations. It bridges real-time vehicle telemetrics at the edge with centralized business logic in the office.
It also demonstrates simulative adaptability akin to a game.

## Project Structure
* `/services/auth`: Go-based identity and access management.
* `/services/simulator`: C++/Python mock telemetry generator.
* `/services/dispatch`: C# service for real-time edge processing.
* `/services/crud`: C# service for fleet configurations and data access.
* `/k8s`: Kubernetes deployment manifests.

## Key Decisions & Architecture
Please refer to [architecture.md](architecture.md) for a detailed breakdown of the stretched cluster architecture, sidecar patterns, and the tradeoffs made regarding edge network latency and high availability.

## Roadmap to Autonomy, Advanced Analytics, and Simulation Play
1. **Phase 1 (Current):** Telemetry ingestion, basic dispatching, and secure edge-to-cloud delivery.
2. **Phase 2 (IoT Sensor Streams & Analytics):** Direct ingestion of high-fidelity vibration, pressure, and acoustic sensors into a Spark/Flink pipeline for predictive maintenance.
3. **Phase 3 (Autonomous Orchestration):** Transitioning from manual dispatch to full equipment orchestration, managing pathing and task assignment for autonomous haulage systems (AHS).
4. **Phase 4 (Orchestration CLI & Dashboard):** Deployment of a management CLI for agent provisioning and goal setting, paired with a unified dashboard for synchronous real-time monitoring and asynchronous historical telemetry analysis. 

## Encryption Standards
* **In-Transit:** All cluster communication is forced through mutual TLS (mTLS) via the service mesh. Public APIs are exposed via HTTPS only.
* **At-Rest:** DB storage volumes use AES-256 block-level encryption (AWS KMS or dm-crypt on bare metal).