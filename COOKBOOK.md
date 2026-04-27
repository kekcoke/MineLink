# MineLink Cookbook: Setup, Troubleshooting & Future Iterations

This guide provides the necessary steps to run MineLink locally, solve common issues, and roadmap the next evolution of the platform.

---

## 1. Local Development Setup

To run the full Hybrid C2 simulation on your machine, you need a local message bus and state store.

### Prerequisites
*   **Go 1.21+** (Supervisor)
*   **dotnet 8.0** (Dispatch)
*   **CMake & C++17 Compiler** (Simulator)
*   **Docker & Docker Compose** (Infrastructure)

### Step A: Infrastructure (The "Mesh")
Start a local MQTT broker (Mosquitto) and Redis instance:
```bash
# Using Docker for quick setup
docker run -d --name minelink-mqtt -p 1883:1883 eclipse-mosquitto
docker run -d --name minelink-redis -p 6379:6379 redis:alpine
```

### Step B: Building & Running Services
Run each layer in a separate terminal:

1.  **Dispatch Agent (Cloud):**
    ```bash
    cd services/dispatch
    dotnet run
    ```
2.  **Supervisor Agent (Edge):**
    ```bash
    cd services/supervisor
    SUPERVISOR_ID=pit-a-alpha go run cmd/supervisor/main.go
    ```
3.  **Operator Simulator (Edge):**
    ```bash
    # Build the C++ simulator
    cd services/simulator
    mkdir build && cd build
    cmake ..
    make
    FLEET_SIZE=10 ./operator_simulator
    ```

---

## 2. Troubleshooting

| Symptom | Probable Cause | Solution |
| :--- | :--- | :--- |
| `Failed to connect to Edge Redis` | Redis is not running on 6379. | Run `docker start minelink-redis`. |
| `MQTT Connection failed` | Broker is down or blocked. | Verify `localhost:1883` is reachable. |
| `Error unmarshalling assignment` | Schema mismatch. | Ensure `asyncapi.yaml` matches the C# and Go models. |
| Supervisor doesn't react to C# intents. | Wrong MQTT topic. | Check `SUPERVISOR_ID` env var matches the Dispatch intent target. |

---

## 3. Next Iterations & Feature Enrichment

### A. The "SIMS" God View (Desktop GUI)
The current CLI control plane should be complemented by a high-fidelity Desktop Application (Phase 4).
*   **Technology:** Electron (React/TypeScript) or Unity (for 3D visualization).
*   **Features:**
    *   **Drag-and-Drop Provisioning:** Visually drag a "Supervisor" into a pit to start a deployment.
    *   **Live Telemetry Map:** Real-time icons for `TRK-XXX` units moving based on GPS payloads.
    *   **God Mode Overrides:** Button to "Force Refuel" or "E-Stop" all agents in a sector.

### B. Predictive Health (ML Integration)
Upgrade the **Predictive Health Agent** (Python/TensorFlow) to consume the `telemetry/truck/+/status` stream.
*   **Goal:** Detect anomalies in `engineTemp` vs. `fuelLevel` to predict mechanical failure 30 minutes before it occurs.
*   **Feedback Loop:** Automatically trigger a `REFUEL_NOW` or `MAINTENANCE` command via the Supervisor.

### C. Advanced Autonomous Pathing
Replace the "Direct Move" logic in the C++ Operators with **A* Pathfinding**.
*   **Complexity:** Operators must navigate around static obstacles (pit walls) and dynamic ones (other trucks) to optimize cycle times.

### D. Multi-Tenancy & Global Scale
Transition from a single pit to a global "Mine Portfolio" view.
*   **Tech:** Use **MQTT Bridging** to link local Edge Brokers in Australia, Chile, and Canada to a single global Cloud Dispatcher.
