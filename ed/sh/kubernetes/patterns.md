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
* Stateless service (example: nginx, authentication microservice, etc.).
* Stateful service (example: mysql, rabbitmq, etc.).
* Service discovery (CoreDNS).
* Self awareness (fieldRef.fieldPath, resourceFieldRef.resource).

Structural Patterns:
* Init container.
* Sidecar.
* Adapter.
* Ambassador.

Configuration Patterns:
* EnvVar configuration.
* Configuration resource.
* Immutable configuration.
* Configuration template.

Security Patterns:
* Process containment.
* Network segmentation.
* Secure configuration.
* Access control.

Advanced Patterns:
* Controller.
* Operator.
* Elastic scale.
* Image builder.
