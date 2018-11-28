Kubernetes (K8s)
-

minikube - tool to use K8s on dev.
kubelet - main k8s agent.
container engine - docker or rkt.
kube-proxy - k8s networking.

Kubernetes - orchestrator for microservice apps.

* Master.
* Node (minion) (`/spec /healthz /pods`).
  Contains the services necessary to run pods and is managed by the master components.
* Manifest file - describe desired state.
* Pod - like container in docker.
* Service - is LB for pods.

````bash
# osx

# brew install kubectl
brew install kubernetes-cli
brew cask install minikube
# brew install docker-machine-driver-xhyve
open https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver
````
````bash
kubectl version
kubectl cluster-info
kubectl config current-context
kubectl get nodes

minikube version
minikube status
# minikube start --vm-driver=xhyve
minikube start --vm-driver=hyperkit
````
