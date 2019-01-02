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
minikube start --vm-driver=hyperkit
minikube stop
minikube dashboard

kubectl version
kubectl cluster-info
kubectl cluster-info dump
kubectl config current-context
kubectl config view
kubectl get events
kubectl get nodes
kubectl get pods
kubectl get pods --show-all
kubectl get pods -o wide
kubectl get rc
kubectl get svc # services
kubectl get deployments
kubectl explain pods
kubectl create -f ed/kubernetes/examples/sh/pod.yaml
kubectl edit <resource> <resource_name>

# sh example
kubectl delete rc ksh-rc
kubectl apply --force=true -f ed/kubernetes/examples/sh/rc.yaml
kubectl describe rc ksh-rc
kubectl get pods -l app=ksh
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
kubectl describe rc log-rc
kubectl get rc -l app=log
kubectl get pods -l app=log

kubectl delete svc log-service
kubectl apply --force=true -f ed/kubernetes/examples/log/svc.yaml
kubectl describe svc log-service
kubectl get svc log-service
kubectl get ep log-service

curl 'http://localhost:8080?x=1&y=2'
````

minikube - tool to use K8s on dev.
kubelet - main k8s agent.
container engine - docker or rkt.
kube-proxy - k8s networking.

Kubernetes - orchestrator for microservice apps.
Kubernetes has three namespaces by default:
* default
* kube-system
* kube-public

To store data permanently, Kubernetes uses Persistent Volumes.

API server (`kube-apiserver`) - provides an HTTP/HTTPS0server,
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
A federation is a cluster of clusters.
KOps stands for Kubernetes operations (for AWS).

#### Master.

#### Node (minion).

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

````yaml
apiVersion: v1
kind: Pod
metadata:
    name: my-pod
spec:
    replicas: 10
````

#### Service - is LB for pods.

````yaml
apiVersion: v1
kind: Service
metadata:
    name: my-svc
    labesl:
        app: hw
spec:
    type: NodePort
    ports:
    - port: 8080
      protocol: TCP
````

#### Deployment

Deployment is the best primitive to manage and deploy our software
in Kubernetes. It supports gracefully deploying, rolling updating,
and rolling back pods and ReplicaSets.

````yaml
apiVersion: extensions/v1beta1
kind: Deployment
````
