helm
-

[docs](https://helm.sh/docs/)
[hub](https://hub.helm.sh/)

Helm - package manager for k8s.
Chart - package.

````sh
# v2.16
download *.tar.gz from https://github.com/helm/helm/releases/tag/v2.16.12
mv ./Downloads/darwin-amd64/helm /usr/local/bin/helm
mv ./Downloads/darwin-amd64/tiller /usr/local/bin/tiller

mkdir -p /Users/k/.helm/repository/
touch    /Users/k/.helm/repository/repositories.yaml
mkdir -p /Users/k/.helm/repository/cache
touch    /Users/k/.helm/repository/cache/stable-index.yaml
````

````sh
export HELM_HOME=/Users/k/.helm
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
helm list              # list releases
helm history $release
helm delete $release

cd ed/sh/helm/examples
chart=sh
env=dev
# ⬆
````
