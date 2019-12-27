k8s examples
-

# sh example
# pod
open https://realtimelog.herokuapp.com:443/ping
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/sh/pod.yaml
kubectl logs -f ksh-pod klog-pod
kubectl logs -f ksh-pod klog-pod-2
kubectl exec -it ksh-pod /bin/bash
kubectl delete pod ksh-pod
# rc
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/sh/rc.yaml
kubectl describe rc ksh-rc
p=`kubectl get pods -l app=ksh -o jsonpath='{.items[*].metadata.name}'`
kubectl logs $p
kubectl delete rc ksh-rc

# log example
# build docker image
ctx=ed/sh/sh.docker/examples.Dockerfile
GOOS=linux go build -o $ctx/xgoapp ed/go/examples/whatever/pingRealtimeLog.go
docker build -t cn007b/pi -f $ctx/go.x.Dockerfile $ctx
docker push cn007b/pi
rm $ctx/xgoapp
# test docker image
docker run -ti --rm -p 8080:8080 cn007b/pi
curl 'http://localhost:8080?x=1&y=2'
# pod
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/log/pod.yaml
kubectl logs -f pod/log-pod
kubectl delete pod log-pod
# rc
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/log/rc.yaml
kubectl delete rc log-rc
# svc
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/log/svc.yaml
minikube service log-service --url
kubectl delete svc log-service
# or svc
kubectl expose rc log-rc --port=8080 --target-port=8080 --name=log-svc --type=LoadBalancer
minikube service log-svc --url
kubectl delete svc log-svc

# go.db example NOT FINISHED
# cm mysql
kubectl apply -f ed/sh/sh.kubernetes/examples/go.db/mysql.v1.cm.yaml
kubectl delete cm/mysql-config
# pod mysql
kubectl delete pod/mysql
kubectl apply -f ed/sh/sh.kubernetes/examples/go.db/mysql.v1.pod.yaml
kubectl exec -it mysql /bin/bash \
  # mkdir -p /gh/.data/.k8s/mysql
  # ln -d /gh/.data/.k8s/mysql /var/lib/mysql
  # ln -d /var/lib/mysql /gh/.data/.k8s/mysql
  # chown mysql.mysql -R /gh/.data/.k8s/mysql
  # cp /etc/mysql/conf.d/custom_mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf
  # service mysql reload
kubectl exec -it mysql mysql -- -P3307 -uroot -proot
# kubectl exec -it mysql mysql -- -P3307 -uroot -proot -Dtest -e "SET GLOBAL datadir='/gh/.data/.k8s/mysql'"
kubectl exec -it mysql mysql -- -P3307 -uroot -proot -Dtest -e 'create table tmp(n int, m varchar(100))'
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e "SHOW VARIABLES where Variable_name = 'datadir'"
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e 'show tables'
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e "insert into tmp values (1, 'one')"
kubectl exec -it mysql mysql -- -P3307 -udbu -pdbp -Dtest -e 'select * from tmp'
mysqlHost=`kubectl get pod/mysql --template={{.status.podIP}}`; echo $mysqlHost
kubectl run mysql-client --image=mysql:5.7.27 -it --rm --restart=Never -- \
  mysql -h$mysqlHost -P3306 -udbu -pdbp -Dtest -e 'select * from tmp'
kubectl delete pod/mysql
# build docker image
ctx=ed/sh/sh.docker/examples.Dockerfile
docker run -it --rm -v $PWD:/gh -w /gh -e ctx=$ctx -e GOPATH='/gh/ed/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && GOOS=linux go build -o /gh/$ctx/xgoapp src/mysql/simple.go'
docker build -t cn007b/pi -f $ctx/go.x.Dockerfile $ctx
docker push cn007b/pi
rm $ctx/xgoapp
# pod
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/go.db/go.pod.yaml
kubectl logs -f pod/go-db-pod
kubectl delete pod/go-db-pod
# svc
kubectl apply --force=true -f ed/sh/sh.kubernetes/examples/go.db/go.svc.yaml
minikube service go-db-service --url
kubectl delete service/go-db-service


