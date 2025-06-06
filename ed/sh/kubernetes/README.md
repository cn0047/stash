Kubernetes (k8s)
-
<br>1.10

[docs](https://kubernetes.io/docs/)

````bash
# osx
brew install kubernetes-cli
brew install minikube
brew install docker-machine-driver-xhyve
open https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver

# ${HOME}/.kube
````

Kubernetes - orchestrator for microservice apps (containers) (desired state management tool).
<br>k8s has three namespaces by default: default, kube-system, kube-public.
<br>Minikube - tool to use k8s on local env.

Advantages:
* Scalable (nodes).
* Horizontal scaling (pods).
* Has self-healing.
* Highly-available.
* Automated rollbacks.
* It's portable (gcp, aws, etc.).

Drawbacks:
* Complexity.
* Upfront cost is high.
* Cost (it may be overkill for small orgs).

Manifest file - describe desired state.
Operator - software extensions that make use of custom resources.
Workload - application running on k8s (Deployment, ReplicaSet, StatefulSet, DaemonSet, Job, CronJob).
Cluster - control plane plus set of worker machines (nodes), that run containerized applications.

Cluster limitations: max 110 pods per node, 5000 nodes, 150K pods in total, 300K containers in total.

Affinity - way to control how pods are scheduled onto nodes based on rules and conditions.
* Node Affinity (assign pods to specific nodes).
* Pod Affinity and Anti-Affinity (rules about how pods are scheduled relative to each other (keep pods together/apart)).

**CoreDNS** - used for internal service discovery,
it automatically assigns DNS name to each service in format: `<svc-name>.<namespace>.svc.cluster.local`.

## Master (Control plane)

Control plane - set of components that manage and control the overall system.

`kube-controller-manager` - controls lots of different things
in the cluster. Replication Controller Manager ensures all the Replication Controllers
run on the desired container amount.

`kube-scheduler` - decides which node is suitable for pods to run on,
according to the resource capacity or the balance of the resource utilization on the node.

`etcd` (cluster data store) is an open source distributed key-value store (https://coreos.com/etcd).
k8s stores all the RESTful API objects here.

`kube-apiserver` (API server) - provides an HTTP/HTTPS server,
which provides a RESTful API for all the components in the k8s master.
All k8s commands come here.

## Node (minion)

A node is a worker machine in k8s, previously known as a minion.

`kubelet` is a major process (agent) in the nodes, which reports node activities back
to kubeapiserver periodically, such as pod health, node health, and liveness probe.

`container engine(runtime)` - docker or rkt.

`kube-proxy` - k8s networking, handles the routing between pod load balancer (a.k.a. service)
and pods, it also provides the routing from outside to service.

`/spec /healthz /pods` - contains the services
necessary to run pods and is managed by the master components.

#### Pod

Pod - (like container in docker) the smallest deployable unit in k8s,
description of a set of container(s) that need to run together.
It can contain one or more containers.
Most of the time, we just need one container per pod.
Pod is also designed as mortal.
Containers in pod shares same IP address, memory, mounted volumes, etc.

#### Kinds

**Service** - object that describes a set of pods that provide a useful service.
Types:
* ClusterIP (default)
  Gives service inside cluster that other apps inside cluster can access.
  There is no external access.
* NodePort
  Primitive way to get external traffic directly to service.
  Opens a specific port on all the nodes.
* LoadBalancer (externally available).
* Ingress - Something like smart router.

**ReplicaSet** controls that pods count matches desired count.

**Deployment** controls the creation and destruction of pods,
is the best primitive to manage and deploy our software in k8s.
It supports gracefully deploying, rolling updating,
and rolling back pods and ReplicaSets.
Main feature - zero downtime deployment.

Deployment strategies `StrategyType`:
* RollingUpdate (default) - create new set of pods & gradually replace old pods.
* Recreate - just drop everything and create new (with down time).

Deployment creates ReplicaSet under the hood.
Deployment status: complete, progressing, failed.
To slow down deployment set `max unavailable 10 or less` and `max surge 2 or smth`.
For Blue/Green deployments - create blue deployment and blue service
then green deployment and green service and change docker image selector in green service.

**Ingress** manages external access to the services in a cluster, typically HTTP,
can provide load balancing, SSL termination and name-based virtual hosting.

**Volume**:
`emptyDir` - lifecycle tied to pod,
`hostPath` - path on worker node,
`persistentVolume` declares volume and `persistentVolumeClaim` uses that volume,
`nfs, configMap/secret, cloud`.

Persistent Volume - to store data permanently.

accessModes:
* ReadOnlyMany  — read-only by many nodes.
* ReadWriteOnce — volume mounted as read-write by a single node.
* ReadWriteMany — read-write by many nodes.

storageClassName: manual, gp2, aws-efs.

**ConfigMap**, `ls /etc/config`.

**DaemonSet** ensures that all (or some) Nodes run a copy of a Pod.

**Job** creates one or more Pods and ensures that a specified number of them successfully complete.
`restartPolicy: Never|OnFailure`

**CronJob**.

A federation is a cluster of clusters.

KOps stands for k8s operations (for AWS).

CRD - Custom Resource Definition.

````sh
kind: ConfigMap                  # separate configuration information from application definition
kind: HorizontalPodAutoscaler    #
kind: Ingress                    # specify how incoming network traffic should be routed to services and pods
kind: Job
kind: LimitRange                 # default CPU/Memory requests/limits for namespace
kind: Namespace
kind: NetworkPolicy              # defines the network access rules between pods inside the cluster
kind: PersistentVolume
kind: Pod
kind: ReplicaSet (RS)            # ReplicationController v2.
kind: ReplicationController (RC)
kind: ResourceQuota
kind: Secret
kind: Deployment
kind: Service
kind: ServiceAccount
kind: StatefulSet
kind: StorageClass
kind: ScaledObject               # KEDA - Kubernetes Event-Driven Autoscaling
````

## yaml

````bash
livenessProbe:
  httpGet:
    path: /health
    port: 80
readinessProbe:
  httpGet:
    path: /health
    port: 80
````

````bash
spec:
  containers:
  - name: x
    image: cn007b/alpine
    resources:
      limits:
        memory: "128Mi"
        cpu: "200m" # 200 millicores, 0.2 CPU cores, 20% of a single CPU core.
      limits:
        cpu: "2"         # no more than 2 CPU cores.
        memory: "2048Mi" # no more than 2GB memory.
      requests:
        cpu: "0.2"       # at least 0.2 CPU cores.
        memory: "512Mi"  # at least 512MB memory.
    livenessProbe:
      httpGet:
        path: /health
        port: 8080
      command:
        - cat
        - /tmp/health
    env:
      - name: POD_IP
        valueFrom:
          fieldRef:
            fieldPath: status.podIP
            # also available:
            # spec.nodeName
            # status.hostIP
            # status.podIP
            # metadata.name
            # metadata.namespace
            # spec.serviceAccountName
            # metadata.uid
            # metadata.labels['key']
            # metadata.annotations['key']
      - name: MEMORY_LIMIT
        valueFrom:
          resourceFieldRef:
            containerName: random-generator
            resource: limits.memory
            # also available:
            # requests.cpu
            # limits.cpu
            # requests.memory
            # limits.memory
            # requests.hugepages-<size>
            # limits.hugepages-<size>
            # requests.ephemeral-storage
            # limits.ephemeral-storage
````

## sh

````bash
minikube version
minikube delete; rm -rf ~/.minikube # deletes local k8s cluster
minikube status
minikube addons enable ingress
minikube ip
# xhyve|hyperkit
minikube start --vm-driver=hyperkit # won't work on ARM
minikube start --vm-driver=hyperkit --mount --mount-string="$HOME/web/kovpak/gh:/gh"
minikube start --driver=docker
minikube start --driver=docker --mount --mount-string="$HOME/web/kovpak/gh:/gh"
# minikube start --vm-driver=hyperkit --mount --mount-string="$HOME/web/kovpak/gh/.data/.k8s/mysql:/var/lib/mysql"
minikube stop
minikube dashboard

# grafana
open https://<your master ip>/api/v1/proxy/namespaces/kubesystem/services/monitoring-grafana

kubectl version
kubectl api-resources # all objects (kinds)
kubectl api-versions | sort
kubectl --kubeconfig=$cfg
kubectl cluster-info
kubectl cluster-info dump
kubectl config get-contexts
kubectl config current-context
kubectl config view
kubectl get namespaces
kubectl proxy # proxy to k8s api
kubectl explain pods

kubectl create -f ed/sh.kubernetes/examples/sh/pod.yaml
kubectl edit $resource $resource_name # edit yaml in vim

kubectl describe rc $rc # replication controller
kubectl describe svc $svc
kubectl describe pods

kubectl --kubeconfig=$f -n my_namespace get all
kubectl get all
kubectl get events
kubectl get nodes
kubectl describe node $n
kubectl get rc
kubectl get rc -l app=log # label
kubectl get rs
kubectl get svc # services
kubectl get ep # endpoints
kubectl get componentstatuses # cluster status
kubectl get configmaps

kubectl get pv           # persistent volumes
kubectl get pvc          # persistent volumes claims
kubectl get storageclass

kubectl get pods
kubectl get pods --watch
kubectl get pods --show-all
kubectl get pods -o wide # shows pods and nodes
kubectl get pods -o json
kubectl get pods -o jsonpath='{.items[*].metadata.name}'

kubectl -n=$ns describe pod $p

lbl='app='
p=`kubectl get pods -l $lbl -o jsonpath='{.items[*].metadata.name}'`
kubectl logs -f $p

kubectl logs $p
kubectl logs $p -c $containerName
kubectl logs -f $p # follow
kubectl logs -p $p # previous

kubectl port-forward $pod 8181:8080 # portOnHost:portInPod
kubectl port-forward $svc 9200

# run pod
kubectl run log-pod --image=cn007b/pi
kubectl port-forward $p 8181:8080
curl 'localhost:8181?yes'

kubectl patch --type=merge pod $p -p '{"metadata":{"labels":{"main":"cli"}}}'

kubectl delete pod $pod
kubectl delete pod $pod --grace-period=60

kubectl get deployments
# update image to initiate deploy
kubectl set image deployment $depl mycontainer=myimage:latest
kubectl set image deployment $depl mycontainer=myimage:latest --record
# simple deploy after update to new docker image version in yaml file
kubectl apply -f deployment.yaml --record
#
kubectl rollout status deployment $depl
kubectl rollout pause deployment
kubectl rollout resume deployment
kubectl rollout undo deployment
kubectl rollout history deployment log-depl # history
kubectl expose deployment $depl --port=8080 --target-port=80
#
kubectl scale deployment $depl --replicas=11

# ssh into pod
kubectl exec -it $pod /bin/bash
kubectl exec -it $pod -c $container -- /bin/bash

# copy files to/from containers running in the pod
# kubectl cp --container=kuard /etc/os-release kuard:/tmp/

kubectl run mysql-client --image=mysql:5.7.27 -it --rm --restart=Never -- \
  mysql -h$mysqlHost -P3306 -udbu -pdbp -Dtest -e 'select * from tmp'

kubectl get secrets
kubectl create secret generic dev-db-secret --from-literal=username=devuser --from-literal=password=dbpass
kubectl get secret dev-db-secret -o jsonpath='{.data.username}' | base64 --decode
kubectl get secret dev-db-secret -o jsonpath='{.data.password}' | base64 --decode
````
