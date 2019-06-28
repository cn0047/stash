MYSQL
-

#### MYSQL cluster

# init master node
docker run -it --rm -p 3307:3306 --name mysql-master --hostname mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-master.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -v $PWD/docker/.data/mysql:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu -e MYSQL_PASSWORD=dbp -e MYSQL_DATABASE=test mysql:latest

# replication user on master
docker exec mysql-master mysql -uroot -proot -e "CREATE USER 'repl'@'%' IDENTIFIED BY 'slavepass'"
docker exec mysql-master mysql -uroot -proot -e "GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%'"

# init slave 1 node
docker run -it --rm -p 3308:3306 --name mysql-slave-1 --link mysql-master \
    -v $PWD/docker/mysql/mysql-bin.log:/var/log/mysql/mysql-bin.log \
    -v $PWD/docker/mysql/config-slave-1.cnf:/etc/mysql/mysql.conf.d/mysqld.cnf \
    -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=dbu2 -e MYSQL_PASSWORD=dbp2 -e MYSQL_DATABASE=test mysql:latest

# start slave 1
docker exec mysql-slave-1 mysql -uroot -proot -e "CHANGE MASTER TO MASTER_HOST='mysql-master', MASTER_USER='repl', MASTER_PASSWORD='slavepass'"
docker exec mysql-slave-1 mysql -uroot -proot -e "START SLAVE"
docker exec mysql-slave-1 mysql -uroot -proot -e "SHOW SLAVE STATUS \G"

# test
docker exec -ti mysql-master mysql -P3307 -udbu -pdbp -Dtest
