Kubernetes (K8s)
-

[docs](https://kubernetes.io/docs/)

````bash
# osx

brew install kubernetes-cli
brew cask install minikube
brew install docker-machine-driver-xhyve
open https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver
````

````bash
minikube version
minikube delete; rm -rf ~/.minikube
minikube status
minikube start --vm-driver=xhyve
minikube start --vm-driver=hyperkit \
  --mount --mount-string="$HOME/web/kovpak/gh/ed:/ed"
minikube stop
minikube dashboard

# grafana
open https://<your master ip>/api/v1/proxy/namespaces/kubesystem/services/monitoring-grafana

kubectl version
kubectl cluster-info
kubectl cluster-info dump
kubectl config current-context
kubectl config view
kubectl explain pods
kubectl create -f ed/kubernetes/examples/sh/pod.yaml
kubectl edit <resource> <resource_name>

kubectl describe rc $rc
kubectl describe svc $svc

kubectl get events
kubectl get nodes
kubectl get rc
kubectl get rc -l app=log # label
kubectl get rs
kubectl get svc # services
kubectl get deployments
kubectl get ep # endpoints
kubectl get componentstatuses # cluster status

kubectl get pods
kubectl get pods --show-all
kubectl get pods -o wide

# ssh into pod
kubectl exec -it $pod /bin/bash

# sh example
kubectl delete pod ksh-pod
kubectl apply --force=true -f ed/kubernetes/examples/sh/pod.yaml
# &
kubectl delete rc ksh-rc
kubectl apply --force=true -f ed/kubernetes/examples/sh/rc.yaml
kubectl describe rc ksh-rc
kubectl get pods -l app=ksh

# copy files to/from containers running in the pod
kubectl cp --container=kuard /etc/os-release kuard:/tmp/
````

````bash
# log example

cd ed/kubernetes/examples/log/ \
  && GOOS=linux go build ../../../go/examples/whatever/pingRealtimeLog.go \
  && docker build -t cn007b/pi . \
  && docker push cn007b/pi \
  && rm pingRealtimeLog \
  && cd -

docker run -ti --rm -p 8080:8080 cn007b/pi
curl 'http://localhost:8080?x=1&y=2'

kubectl delete rc log-rc
kubectl apply --force=true -f ed/kubernetes/examples/log/rc.yaml
kubectl get pods -l app=log

kubectl delete svc log-service
kubectl apply --force=true -f ed/kubernetes/examples/log/svc.yaml
minikube service log-service --url
# or
kubectl delete svc log-svc
kubectl expose rc log-rc --port=8080 --target-port=8080 \
  --name=log-svc --type=LoadBalancer
minikube service log-svc --url

# only pod
kubectl delete pod log-pod
kubectl apply --force=true -f ed/kubernetes/examples/log/pod.yaml
````
<br>minikube - tool to use K8s on dev.
<br>kubelet - main k8s agent.
<br>container engine - docker or rkt.
<br>kube-proxy - k8s networking.

Kubernetes - orchestrator for microservice apps.
Kubernetes has three namespaces by default:
* default
* kube-system
* kube-public

Kinds:

````sh
kind: ConfigMap
kind: DaemonSet
kind: Deployment
kind: HorizontalPodAutoscaler
kind: Ingress
kind: Job
kind: LimitRange
kind: NetworkPolicy
kind: PersistentVolume
kind: Pod
kind: ReplicationController
kind: ResourceQuota
kind: Secret
kind: Service
kind: ServiceAccount
kind: StatefulSet
kind: StorageClass
````

To store data permanently, Kubernetes uses Persistent Volumes.

API server (`kube-apiserver`) - provides an HTTP/HTTPS server,
which provides a RESTful API for all the components in the Kubernetes master.

Controller Manager (`kube-controller-manager`) - controls lots of different things
in the cluster. Replication Controller Manager ensures all the Replication
Controllers run on the desired container amount.

`etcd` is an open source distributed key-value store (https://coreos.com/etcd).
Kubernetes stores all the RESTful API objects here.

Scheduler (kube-scheduler) - decides which node is suitable for pods to run on,
according to the resource capacity or the balance of the resource utilization on the node.

`Kubelet` is a major process in the nodes, which reports node activities back
to kubeapiserver periodically, such as pod health, node health, and liveness probe.

Proxy (`kube-proxy`) - handles the routing between pod load balancer (a.k.a. service)
and pods, it also provides the routing from outside to service.

ReplicaSet (RS).
ReplicationController (RC).
DaemonSet where we need an agent to run on every single node in the cluster.
A federation is a cluster of clusters.
KOps stands for Kubernetes operations (for AWS).

#### Master.

#### Node (minion).

A node is a worker machine in k8s, previously known as a minion.

`/spec /healthz /pods`
CONTAINS the services necessary to run pods and is managed by the master components.

#### Manifest file.

Manifest file - describe desired state.

#### Pod

Pod - like container in docker.
Pod is the smallest deployable unit in Kubernetes.
It can contain one or more containers.
Most of the time, we just need one container per pod.
Pod is also designed as mortal.
Containers in pod shares same localhost.

#### Service - is LB for pods.

Types:

* ClusterIP (default)
  Gives service inside cluster that other apps inside cluster can access.
  There is no external access.

* NodePort
  Primitive way to get external traffic directly to service.
  Opens a specific port on all the nodes.

* LoadBalancer

* Ingress
  Something like smart router.

#### Deployment

Deployment is the best primitive to manage and deploy our software
in Kubernetes. It supports gracefully deploying, rolling updating,
and rolling back pods and ReplicaSets.
