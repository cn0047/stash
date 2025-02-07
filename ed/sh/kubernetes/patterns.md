K8S Patterns
-

Foundational patterns:
* Predictable demands (kind: PriorityClass, ResourceQuota, LimitRange).
* Declarative deployment (kind: Deployment).
* Health probe (HTTP, TCP, GRPC, exec).
* Managed lifecycle (lifecycle.postStart, lifecycle.preStop, initContainers).
* Automated placement (kind: KubeSchedulerConfiguration; pod affinity).

Behavioral Patterns:
* Batch job (kind: Job).
* Periodic job (kind: CronJob).
* Daemon service (kind: DaemonSet).
* Singleton service (kind: PodDisruptionBudget).
* Stateless service (offload the state to some other stateful system or data store).
* Stateful service.
* Service discovery.
* Self awareness.
