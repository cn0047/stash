k8s examples
-
# build docker image
ctx=ed/sh/docker/examples.Dockerfile
# 1
GOOS=linux go build -o $ctx/xgoapp ed/l/go/examples/whatever/pingRealtimeLog.go
# 2 (mysql)
docker run -it --rm -v $PWD:/gh -w /gh -e ctx=$ctx -e GOPATH='/gh/ed/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && GOOS=linux go build -o /gh/$ctx/xgoapp src/mysql/simple.go'
# push to dockerhub
docker build -t cn007b/pi -f $ctx/go.x.Dockerfile $ctx
docker push cn007b/pi
mv $ctx/xgoapp /tmp
# test docker image
docker run -ti --rm -p 8080:8080 cn007b/pi
curl 'http://localhost:8080?x=1&y=2'

# sh example
# pod
open https://realtimelog.herokuapp.com:443/ping
kubectl apply --force=true -f ed/sh/kubernetes/examples/sh/pod.yaml
kubectl logs -f ksh-pod klog-pod
kubectl logs -f ksh-pod klog-pod-2
kubectl exec -it ksh-pod /bin/bash
kubectl delete pod ksh-pod
# rc
kubectl apply --force=true -f ed/sh/kubernetes/examples/sh/rc.yaml
kubectl describe rc ksh-rc
kubectl delete rc ksh-rc
# depl
kubectl apply --force=true -f ed/sh/kubernetes/examples/sh/depl.yaml
kubectl describe deployment ksh-depl
kubectl rollout status deployment ksh-depl
kubectl delete deployment ksh-depl

# log example
# pod
kubectl apply --force=true -f ed/sh/kubernetes/examples/log/pod.yaml
kubectl logs -f pod/log-pod
kubectl delete pod log-pod
# rc
kubectl apply --force=true -f ed/sh/kubernetes/examples/log/rc.yaml
kubectl delete rc log-rc
# svc
kubectl apply --force=true -f ed/sh/kubernetes/examples/log/svc.yaml
minikube service log-service --url
kubectl delete svc log-service
# or svc
kubectl expose rc log-rc --port=8080 --target-port=8080 --name=log-svc --type=LoadBalancer
minikube service log-svc --url
kubectl delete svc log-svc
# or svc.2
kubectl apply --force=true -f ed/sh/kubernetes/examples/log/svc.2.yaml
kubectl delete svc log-service-2
# depl
kubectl apply --force=true -f ed/sh/kubernetes/examples/log/depl.yaml
kubectl rollout status deployment log-depl
kubectl rollout history deployment log-depl
kubectl describe deployment log-depl
kubectl get deployment log-depl -o yaml
kubectl set image deployment log-depl log-container=cn007b/pi --record
kubectl set image deployment log-depl log-container=cn007b/pi:latest --record
kubectl scale deployment log-depl --replicas=5
kubectl delete deployment log-depl

# go.db example
# pv
kubectl apply -f ed/sh/kubernetes/examples/go.db/mysql.pv.yaml
kubectl get pv mysql-pv
kubectl describe pv mysql-pv
kubectl delete pv mysql-pv
# pvc
kubectl apply -f ed/sh/kubernetes/examples/go.db/mysql.pvc.yaml
kubectl get pvc mysql-pvc
kubectl describe pvc mysql-pvc
kubectl delete pvc mysql-pvc
# cm mysql
kubectl apply -f ed/sh/kubernetes/examples/go.db/mysql.cm.yaml
kubectl get cm/mysql-config
kubectl delete cm/mysql-config
# pod mysql
kubectl apply -f ed/sh/kubernetes/examples/go.db/mysql.v3.pod.yaml
kubectl exec -it mysql mysql -- -P3307 -uroot -proot -Dtest -e 'create table tmp(n int, m varchar(100))'
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e 'show tables'
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e "insert into tmp values (1, 'one')"
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e "insert into tmp values (2, 'two')"
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e 'select * from tmp'
mysqlHost=`kubectl get pod/mysql --template={{.status.podIP}}`; echo $mysqlHost
kubectl run mysql-client --image=mysql:5.7.27 -it --rm --restart=Never -- \
  mysql -h$mysqlHost -P3306 -udbu -pdbp -Dtest -e 'select * from tmp'
kubectl delete pod/mysql
# pod go
kubectl apply --force=true -f ed/sh/kubernetes/examples/go.db/go.pod.yaml
kubectl describe pod/go-db-pod
kubectl logs -f pod/go-db-pod
kubectl delete pod/go-db-pod
# svc go
kubectl apply --force=true -f ed/sh/kubernetes/examples/go.db/go.svc.yaml
minikube service go-db-service --url
kubectl delete service/go-db-service
