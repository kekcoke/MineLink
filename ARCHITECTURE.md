# System Architecture & Tradeoffs

## Architecture Diagram (Mermaid)

```mermaid
graph TD
    %% Edge 1 (Pit A)
    subgraph Edge_Site_1 [Edge Deployment 1: Pit A]
        TruckSim1[Truck Simulator - Python/C++] -->|mTLS MQTT| EdgeBroker1[Edge MQTT Broker]
        EdgeBroker1 -->|Bridged Stream| OfficeBroker
        HealthCheck1[Dispatch & Health API - C#] -->|Subscribe| EdgeBroker1
    end

    %% Edge 2 (Pit B)
    subgraph Edge_Site_2 [Edge Deployment 2: Pit B]
        TruckSim2[Truck Simulator - Python/C++] -->|mTLS MQTT| EdgeBroker2[Edge MQTT Broker]
        EdgeBroker2 -->|Bridged Stream| OfficeBroker
        HealthCheck2[Dispatch & Health API - C#] -->|Subscribe| EdgeBroker2
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
    EdgeBroker2 -.->|In-Transit Encryption| OfficeBroker
    Users((Office & Field Users)) -->|HTTPS| CrudSvc
    CrudSvc -.->|Token Verification| AuthSvc