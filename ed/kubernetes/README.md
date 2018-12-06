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
minikube dashboard

kubectl version
kubectl cluster-info
kubectl config current-context
kubectl config view
kubectl get events
kubectl get nodes
kubectl get pods
kubectl get pods --show-all
kubectl get services
kubectl get deployments
kubectl explain pods
kubectl create -f ed/kubernetes/examples/sh/pod.yaml
````

````bash
kubectl apply --force=true -f ed/kubernetes/examples/sh/pod.yaml
kubectl logs myapp-sh-pod
kubectl delete pod myapp-sh-pod
````

````bash
cd ed/kubernetes/examples/pingRealtimeLog/ \
  && GOOS=linux go build ../../../go/examples/whatever/pingRealtimeLog.go \
  && docker build -t kube_ping_realtime_log . \
  && cd -

docker run -ti --rm  kube_ping_realtime_log

kubectl apply --force=true -f ed/kubernetes/examples/pingRealtimeLog/pod.yaml
kubectl logs myapp-kube-ping-realtime-log-pod
````

minikube - tool to use K8s on dev.
kubelet - main k8s agent.
container engine - docker or rkt.
kube-proxy - k8s networking.

Kubernetes - orchestrator for microservice apps.

To store data permanently, Kubernetes uses Persistent Volumes.

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
