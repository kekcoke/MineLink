# System Architecture & Tradeoffs

## Architecture Diagram (Mermaid)

```mermaid
graph TD
    %% User Layer
    CLI[MineLink CLI] -->|Provision/Set Goals| Supervisor
    Dashboard[Monitoring Dashboard] -->|Query Sync/Async| CrudSvc
    Dashboard -->|Real-time Stream| OfficeBroker

    %% Agent Hierarchy (C2)
    DispatchAgent[Dispatch Optimization Agent] -->|Strategic Commands| Supervisor
    Supervisor[Off-Site Supervisor Agent] -->|Tactical Assignments| Operators[Off-Site Operator Agents]
    Operators -->|Telemetry Load| EdgeBroker1

    %% Edge 1 (Pit A)
    subgraph Edge_Site_1 [Edge Deployment 1: Pit A]
        TruckSim1[Truck Simulator] -->|mTLS MQTT| EdgeBroker1[Edge MQTT Broker]
        EdgeBroker1 -->|Bridged Stream| OfficeBroker
        HealthCheck1[Dispatch & Health API - C#] -->|Subscribe| EdgeBroker1
    end

    %% Office / Static Deployment
    subgraph Static_Office [Static Deployment: Cloud/HQ]
        OfficeBroker[Central Message Bus]
        AuthSvc[Auth Service - Go]
        CrudSvc[CRUD API - C#]
        DB[(PostgreSQL + TimescaleDB)]
        
        CrudSvc --> DB
        AuthSvc --> DB
        OfficeBroker --> CrudSvc
    end

    %% Connections
    EdgeBroker1 -.->|In-Transit Encryption| OfficeBroker
    CrudSvc -.->|Token Verification| AuthSvc
```

## 1. Hierarchical Command & Control (C2)
MineLink utilizes a three-tier agent hierarchy to manage complex mining operations:
*   **Strategic Layer (Dispatch):** Global optimization and high-level rerouting.
*   **Orchestration Layer (Supervisor):** Translates strategy into tactical shifts and agent provisioning.
*   **Execution Layer (Operators):** Direct equipment interaction and telemetry generation.

## 2. User Interface: CLI & Dashboard
The system replaces a traditional GUI with a two-pronged control plane:
*   **MineLink CLI:** A high-performance interface for developers and operators to provision agents, set production goals via arguments, and trigger shift rotations.
*   **Unified Dashboard:** A hybrid visualization tool that supports:
    *   **Synchronous Monitoring:** Real-time telemetry streams via WebSockets/MQTT.
    *   **Asynchronous Analytics:** Querying historical data and performance metrics from TimescaleDB.

## 3. Temporal Simulation & State Persistence
To simulate real-world mining cycles, the system implements:
*   **Shift-Based Logic:** Operators work 8-hour shifts with mandatory hand-offs and state transitions.
*   **Resource Depletion:** Real-time tracking of fuel, tire pressure, and mechanical health, requiring agents to proactively seek maintenance or refueling.
*   **Load Generation:** Simulated operators provide stress testing for the ingestion pipeline, ensuring the office layer can handle high-frequency telemetry bursts.

## 4. Encryption & Security
*   **In-Transit:** All cluster communication is forced through mutual TLS (mTLS) via the service mesh.
*   **At-Rest:** DB storage volumes use AES-256 block-level encryption.