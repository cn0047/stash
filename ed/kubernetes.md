Kubernetes (K8s)
-

````bash
# osx

brew install kubernetes-cli
brew cask install minikube
brew install docker-machine-driver-xhyve
open https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver
````

````bash
kubectl version
kubectl cluster-info
kubectl config current-context
kubectl get nodes
kubectl get pods
kubectl explain pods
kubectl create -f pod.yml

minikube version
minikube delete; rm -rf ~/.minikube
minikube status
minikube dashboard
minikube start --vm-driver=xhyve
minikube start --vm-driver=hyperkit
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

````yaml
apiVersion: extensions/v1beta1
kind: Deployment
````
