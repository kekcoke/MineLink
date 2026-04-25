# AI & Automation Agents

To evolve from a passive monitoring system to an active FMS, the platform utilizes specialized service agents:

## 1. The Dispatch Optimization Agent
* **Role:** Analyzes real-time traffic in the mine pit, queue sizes at shovel sites, and vehicle fuel levels.
* **Skill:** Operates an optimization loop to issue direct rerouting commands to trucks via MQTT to prevent idle time.

## 2. The Predictive Health Agent
* **Role:** Monitors telemetry streams from the Simulator.
* **Skill:** Utilizes lightweight ML models on the edge to detect anomalies (e.g., rising engine temps combined with drop in pressure) and automatically flags equipment for maintenance before a catastrophic failure occurs.

## 3. The Off-Site Operator Agents
* **Role:** Simulate and act as off-site workers directly managing vehicle and other related equipment.
* **Skill:** Work and produce output which are mining and consuming vehicle levels on 8 hours shifts. The output will generate real-time traffic for data indigestion.

## 4. The Off-Site Supervisor Agent
* **Role:** Orchestrate and delete offsite orchestrators, under direction of Dispatch Optimization Agent
* **Skill:** Delegates tasks commands to workers in accordance to Dispatch and ensures outputs are met as directed of the Optimization Agent.
