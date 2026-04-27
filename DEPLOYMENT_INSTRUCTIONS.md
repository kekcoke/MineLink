# Deployment Instructions

## Prerequisites
* Docker & Kubernetes CLI (`kubectl`)
* A running K8s cluster (or Minikube/K3s for local testing)
* Helm (for package management)
* Redis CLI (for local edge state verification)

## 1. Secrets and Encryption Keys
Generate your SSL certs for mTLS and apply them as Kubernetes secrets. All communication between the Edge and Cloud is encrypted.
```bash
kubectl create secret tls edge-broker-tls --cert=edge.crt --key=edge.key
```

## 2. Edge Deployment (The Pit)
The edge deployment mimics a low-footprint industrial environment.

### A. Edge Redis (State Store)
Deploy a lightweight Redis instance to handle local agent state:
```bash
helm install edge-state bitnami/redis --set architecture=standalone,auth.enabled=false
```

### B. Supervisor Agent (Go) & Simulator Pods (C++)
Apply the edge manifests. Note the strict `replicas: 2` constraint in the manifests to ensure high availability without exceeding hardware caps.
```bash
kubectl apply -f k8s/edge/supervisor-deployment.yaml
kubectl apply -f k8s/edge/simulator-deployment.yaml
```

## 3. Cloud Deployment (The Office)
The cloud layer manages global optimization and heavy data storage.

### A. Dispatch Optimization Agent (C#)
Deploy the central optimization loop:
```bash
kubectl apply -f k8s/cloud/dispatch-deployment.yaml
```

### B. Data & Analytics
Deploy the primary persistence layer (PostgreSQL + TimescaleDB):
```bash
kubectl apply -f k8s/cloud/database-deployment.yaml
```

## 4. Verification
Verify that the Supervisor has successfully registered with the Cloud Dispatcher:
```bash
# Check Supervisor logs for MQTT handshake
kubectl logs -l app=supervisor

# Check local state in Edge Redis
redis-cli -h edge-state-master HGETALL operator:active:shifts
```
