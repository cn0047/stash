K8S Patterns
-

Foundational patterns:
* Predictable demands (kind: PriorityClass, ResourceQuota, LimitRange).
* Declarative deployment (kind: Deployment).
* Health probe (HTTP, TCP, GRPC, exec).
* Managed lifecycle (lifecycle.postStart, lifecycle.preStop, initContainers).
* Automated placement (kind: KubeSchedulerConfiguration; pod affinity).
