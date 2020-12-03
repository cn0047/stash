# MySQL

tag=5.7.27
tag=latest
docker run -it --rm --net=xnet -p 3307:3306 --name xmysql --hostname xmysql \
    -v $PWD/.data/.docker/mysql_$tag:/var/lib/mysql -v /tmp:/tmp \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test \
    mysql:$tag

# general_log
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "set global general_log_file='/tmp/mysql.general.log';"
docker exec -ti xmysql mysql -P3307 -uroot -proot -e "set global general_log = 1;"
docker exec -ti xmysql tail -f /tmp/mysql.general.log

docker exec -ti xmysql mysql -P3307 -uroot -proot
docker exec -ti xmysql mysql -P3307 -udbu -pdbp -Dtest

#### MySQL cluster

# init master node
docker run -it --rm -p 3307:3306 --name mysql-master --hostname mysql-master \
    -v $PWD/ed/db.mysql/examples.replication/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/ed/db.mysql/examples.replication/config-master.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -v $PWD/.data/.docker/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test mysql:5.7.27

# replication user on master
docker exec mysql-master mysql -uroot -proot -e "CREATE USER 'repl'@'%' IDENTIFIED BY 'slavepass'"
docker exec mysql-master mysql -uroot -proot -e "GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%'"

# init slave 1 node
docker run -it --rm -p 3308:3306 --name mysql-slave-1 --link mysql-master \
    -v $PWD/ed/db.mysql/examples.replication/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/ed/db.mysql/examples.replication/config-slave-1.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu2 -e MYSQL_PASSWORD=dbp2 -e MYSQL_DATABASE=test mysql:5.7.27

# start slave 1
docker exec mysql-slave-1 mysql -uroot -proot -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='repl', MASTER_PASSWORD='slavepass'"
docker exec mysql-slave-1 mysql -uroot -proot -e "START SLAVE"
docker exec mysql-slave-1 mysql -uroot -proot -e "SHOW SLAVE STATUS \G"

# test
docker exec mysql-master mysql -uroot -proot -Dtest -e "CREATE TABLE tmp(code INT); INSERT INTO tmp VALUES (200);"
docker exec mysql-slave-1 mysql -uroot -proot -Dtest -e "SELECT * FROM tmp;"

docker exec -ti mysql-master mysql -P3307 -udbu -pdbp -Dtest
