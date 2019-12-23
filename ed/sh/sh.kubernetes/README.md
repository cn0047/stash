Kubernetes (K8s)
-
<br>1.10

[docs](https://kubernetes.io/docs/)

````bash
# osx
brew install kubernetes-cli
brew cask install minikube
brew install docker-machine-driver-xhyve
open https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver
````

Kubernetes - orchestrator for microservice apps (desired state management tool).
<br>Kubernetes has three namespaces by default: default, kube-system, kube-public.
<br>Minikube - tool to use K8s on dev.

#### Master.

`etcd` (cluster data store) is an open source distributed key-value store (https://coreos.com/etcd).
Kubernetes stores all the RESTful API objects here.

`kube-controller-manager` - controls lots of different things
in the cluster. Replication Controller Manager ensures all the Replication
Controllers run on the desired container amount.

`kube-scheduler` - decides which node is suitable for pods to run on,
according to the resource capacity or the balance of the resource utilization on the node.

`kube-apiserver` (API server) - provides an HTTP/HTTPS server,
which provides a RESTful API for all the components in the Kubernetes master.
All k8s commands come here.

#### Node (minion).

A node is a worker machine in k8s, previously known as a minion.

`kubelet` is a major process (agent) in the nodes, which reports node activities back
to kubeapiserver periodically, such as pod health, node health, and liveness probe.

`container engine` - docker or rkt.

`kube-proxy` - k8s networking, handles the routing between pod load balancer (a.k.a. service)
and pods, it also provides the routing from outside to service.

`/spec /healthz /pods` - contains the services
necessary to run pods and is managed by the master components.

#### Manifest file.

Manifest file - describe desired state.

#### Pod.

Pod - (like container in docker)
description of a set of container(s) that need to run together.
Pod is the smallest deployable unit in Kubernetes.
It can contain one or more containers.
Most of the time, we just need one container per pod.
Pod is also designed as mortal.
Containers in pod shares same IP address and mounted volumes.

#### Service - is LB for pods.

Service - object that describes a set of pods that provide a useful service.

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

#### Deployment.

Deployment is the best primitive to manage and deploy our software in k8s.
It supports gracefully deploying, rolling updating,
and rolling back pods and ReplicaSets.

## Kinds.

Persistent Volume - to store data permanently.
DaemonSet where we need an agent to run on every single node in the cluster.
A federation is a cluster of clusters.
KOps stands for Kubernetes operations (for AWS).

````sh
kind: ConfigMap                  # separate configuration information from application definition
kind: Deployment                 # controls the creation and destruction of pods
kind: HorizontalPodAutoscaler    #
kind: Ingress                    # specify how incoming network traffic should be routed to services and pods
kind: Job
kind: LimitRange
kind: NetworkPolicy              # defines the network access rules between pods inside the cluster
kind: PersistentVolume
kind: Pod
kind: ReplicationController (RC)
kind: ReplicaSet (RS)            # ReplicationController v2.
kind: ResourceQuota
kind: Secret
kind: Service
kind: ServiceAccount
kind: StatefulSet
kind: StorageClass
````

## SH.

````bash
minikube version
minikube delete; rm -rf ~/.minikube
minikube status
# xhyve|hyperkit
minikube start --vm-driver=hyperkit --mount --mount-string="$HOME/web/kovpak/gh/ed:/ed"
minikube stop
minikube dashboard

# grafana
open https://<your master ip>/api/v1/proxy/namespaces/kubesystem/services/monitoring-grafana

kubectl version
kubectl api-resources # all objects (kinds)
kubectl cluster-info
kubectl cluster-info dump
kubectl config current-context
kubectl config view
kubectl explain pods
kubectl create -f ed/sh.kubernetes/examples/sh/pod.yaml
kubectl edit <resource> <resource_name>

kubectl describe rc $rc # replication controller
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
kubectl apply --force=true -f ed/sh.kubernetes/examples/sh/pod.yaml
# &
kubectl delete rc ksh-rc
kubectl apply --force=true -f ed/sh.kubernetes/examples/sh/rc.yaml
kubectl describe rc ksh-rc
kubectl get pods -l app=ksh

# copy files to/from containers running in the pod
kubectl cp --container=kuard /etc/os-release kuard:/tmp/
````

````bash
# log example

ctx=ed/sh.docker/examples.Dockerfile
GOOS=linux go build -o $ctx/xgoapp ed/go/examples/whatever/pingRealtimeLog.go
docker build -t cn007b/pi -f $ctx/go.x.Dockerfile $ctx
docker push cn007b/pi
rm $ctx/xgoapp

docker run -ti --rm -p 8080:8080 cn007b/pi
curl 'http://localhost:8080?x=1&y=2'

# pod
kubectl delete pod log-pod
kubectl apply --force=true -f ed/sh.kubernetes/examples/log/pod.yaml
kubectl logs -f pod/log-pod

kubectl delete rc log-rc
kubectl apply --force=true -f ed/sh.kubernetes/examples/log/rc.yaml

kubectl delete svc log-service
kubectl apply --force=true -f ed/sh.kubernetes/examples/log/svc.yaml
minikube service log-service --url
# or
kubectl delete svc log-svc
kubectl expose rc log-rc --port=8080 --target-port=8080 --name=log-svc --type=LoadBalancer
minikube service log-svc --url
````

````sh
# go.db example

kubectl delete cm/mysql-config
kubectl apply -f ed/sh.kubernetes/examples/go.db/mysql.v1.cm.yaml
kubectl delete pod/mysql
kubectl apply -f ed/sh.kubernetes/examples/go.db/mysql.v1.pod.yaml
kubectl exec -it mysql /bin/bash
kubectl exec -it mysql mysql -- -P3307 -uroot -proot
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest
mysqlHost=`kubectl get pod/mysql --template={{.status.podIP}}`; echo $mysqlHost
kubectl run mysql-client --image=mysql:5.7.27 -it --rm --restart=Never -- \
  mysql -h$mysqlHost -P3306 -uroot -proot -Dtest -e 'select * from test_mysql'

ctx=ed/sh.docker/examples.Dockerfile
docker run -it --rm -v $PWD:/gh -w /gh -e ctx=$ctx -e GOPATH='/gh/ed/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && GOOS=linux go build -o /gh/$ctx/xgoapp src/mysql/simple.go'
docker build -t cn007b/pi -f $ctx/go.x.Dockerfile $ctx
docker push cn007b/pi
rm $ctx/xgoapp

# run xmysql container before next command
docker run -ti --rm -p 8080:8080 --net=xnet -v $PWD:/gh -w /gh -e ctx=$ctx \
  xubuntu sh -c '/gh/$ctx/xgoapp'
open http://localhost:8080

# pod
kubectl delete pod/go-db-pod
kubectl apply --force=true -f ed/sh.kubernetes/examples/go.db/go.pod.yaml
kubectl logs -f pod/go-db-pod

# svc
kubectl delete service/go-db-service
kubectl apply --force=true -f ed/sh.kubernetes/examples/go.db/go.svc.yaml
minikube service go-db-service --url
````
