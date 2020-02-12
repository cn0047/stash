helm
-

[docs](https://helm.sh/docs/)

Helm - package manager for k8s.
Chart - package.

````sh
export HELM_HOST=localhost:44134

helm version
helm init  # init helm
helm reset # uninstalls tiller

helm --kubeconfig=$f list

helm create $chart

helm install $chart -f $chart/$env/values.yaml
helm install $chart --set varfoo=bar --dry-run --debug
helm upgrade $release
helm rollback $release
helm get $release
helm status $release
helm list # list releases
helm history $release
helm delete $release

cd ed/sh/helm/examples
chart=sh
env=dev
# ⬆
````
