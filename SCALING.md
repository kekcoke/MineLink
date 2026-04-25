# Scaling Constraints & Strategies

This MVP operates under a strict resource cap to mimic low-footprint edge hardware.

## The 2-Pod Constraint
* **Implementation:** Every service deployment in the Kubernetes manifests is hardcoded with `replicas: 2`. 
* **Reasoning:** This ensures that if one node/pod crashes at the mine pit, a second instance is already running to pick up the load. We do not scale beyond 2 to maintain predictability of resources in the edge environment.

## Horizontal vs. Vertical Scaling
* When the fleet size grows, vertical scaling (adding CPU/RAM to existing pods) is preferred for the C++ Edge simulator to minimize context switching.
* Horizontal scaling (spinning up new instances of edge brokers) should be handled by bridging brokers rather than increasing pod counts beyond 2 in the primary cluster.