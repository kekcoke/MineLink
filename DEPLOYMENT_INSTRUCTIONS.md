# Deployment Instructions

## Prerequisites
* Docker & Kubernetes CLI (`kubectl`)
* A running K8s cluster (or Minikube/K3s for local testing)
* Helm (for package management)

## Steps

### 1. Secrets and Encryption Keys
Generate your SSL certs for mTLS and apply them as Kubernetes secrets:
```bash
kubectl create secret tls edge-broker-tls --cert=edge.crt --key=edge.key