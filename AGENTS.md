# AI & Automation Agents

To evolve from a passive monitoring system to an active FMS, the platform utilizes specialized service agents organized in a hierarchical Command & Control (C2) structure:

## 1. The Dispatch Optimization Agent (Strategic Layer)
* **Role:** Analyzes global mine state, including real-time traffic in the mine pit, queue sizes at shovel sites, and vehicle fuel levels.
* **Skill:** Operates an optimization loop to issue high-level rerouting and task priority commands to the Supervisor layer via MQTT.

## 2. The Off-Site Supervisor Agent (Orchestration Layer)
* **Role:** Translates strategic goals from the Dispatch Optimization Agent into tactical task assignments.
* **Skill:** Orchestrates the lifecycle of Off-Site Operator Agents. It provisions, monitors, and rotates workers based on 8-hour shift requirements and ensures production outputs align with the Dispatch Agent's directives.

## 3. The Off-Site Operator Agents (Execution Layer)
* **Role:** Simulated digital workers directly managing vehicles and heavy equipment in 8-hour shifts.
* **Skill:** Executes mining tasks, consumes vehicle resources (fuel/tire life), and generates high-fidelity telemetry. This layer provides the primary "load generation" for testing the system's ingestion limits.

## 4. The Predictive Health Agent (Monitoring Layer)
* **Role:** Parallel monitoring of telemetry streams from the Simulator and Operator layers.
* **Skill:** Utilizes lightweight ML models on the edge to detect anomalies (e.g., rising engine temps combined with drop in pressure) and automatically flags equipment for maintenance, providing feedback to the Dispatch Agent for rerouting.

