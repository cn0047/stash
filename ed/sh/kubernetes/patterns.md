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
* Init container (initContainers executed first, if it finished successfully then executed application container).
* Sidecar (helper container that runs alongside: logging, monitoring, service mesh, configuration management, etc.).
* Adapter (sidecar which helps with incompatible interface of main container).
* Ambassador (sidecar, which is like smart proxy to the outside world).

Configuration Patterns:
* EnvVar configuration (env, valueFrom.configMapKeyRef, valueFrom.secretKeyRef).
* Configuration resource (kind: ConfigMap; envFrom.configMapRef, configMap).
* Immutable configuration (initContainers write into volume (emptyDir), and containers read from immutable volume).
* Configuration template (volume with params -> initContainers + template -> emptyDir volume -> read by main container).

Security Patterns:
* Process containment (securityContext.runAsUser, securityContext.runAsGroup, securityContext.runAsNonRoot).
* Network segmentation (kind: NetworkPolicy, AuthorizationPolicy).
* Secure configuration (kind: SealedSecret, SecretStore, ExternalSecret, SecretProviderClass).
* Access control (kind: ServiceAccount, Secret, Role, ClusterRole).

Advanced Patterns:
* Controller (controller actively monitors and maintains set of k8s resources in desired state, built-in controllers: ReplicaSet, DaemonSet, StatefulSet, Deployment, Service).
* Operator (operator - controller that uses CRD to encapsulate operational knowledge, kind: CustomResourceDefinition).
* Elastic scale.
* Image builder.
