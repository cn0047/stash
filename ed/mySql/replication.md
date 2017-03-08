Replication
-

*MySQL 5.5*

Depending on the configuration, you can replicate all databases, selected databases, or even selected tables within a database.

* Asynchronous replication - one server acts as the master, while one or more other servers act as slaves.
The master writes events to its binary log and slaves request them when they are ready.

* Synchronous replication which is a characteristic of MySQL Cluster.
When a master commits a transaction, all slaves also will have committed the transaction
before the master returns to the session that performed the transaction.

* Semisynchronous replication (is implemented using plugins).
The master waits only until at least one slave has received and logged the events.

There are two core types of replication format:

* Statement Based Replication (SBR) - which replicates entire SQL statements.
* Row Based Replication (RBR) - which replicates only the changed rows.

Replication between servers in MySQL is based on the binary logging mechanism.
The information in the binary log is stored in different logging formats according to the database changes being recorded.

#### Master Configuration

You will need to shut down your MySQL server and edit /etc/mysql/my.cnf (/etc/mysql/mysql.conf.d/mysqld.cnf):

````sql
[mysqld]
server-id=1
log_bin=/var/log/mysql/mysql-bin.log

# For the greatest possible durability and consistency
# in a replication setup using InnoDB with transactions,
# you should use:
innodb_flush_log_at_trx_commit=1
# and
sync_binlog=1
````
````
sudo service mysql restart
````

````sql
CREATE USER 'repl'@'192.168.56.%' IDENTIFIED BY 'slavepass';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'192.168.56.%';
````

````
USE test;
CREATE TABLE tt (msg VARCHAR(25) KEY);
INSERT INTO tt VALUES ('this'), ('is'), ('test');
````

````
mysqldump -uroot -proot --all-databases --master-data > dump.db.sql
scp dump.db.sql vagrant@192.168.56.103:/tmp
````

#### Setting the Replication Slave Configuration

 You should shut down your slave server and edit the configuration to specify a unique server ID. For example:
````sql
[mysqld]
server-id=2
````
````
sudo service mysql restart
````
````sql
CHANGE MASTER TO
    MASTER_HOST='192.168.56.102',
    MASTER_USER='repl',
    MASTER_PASSWORD='slavepass'
;
````
````
mysql -uroot -proot < /tmp/dump.db.sql
````
````
START SLAVE;
SHOW SLAVE STATUS \G
````

#### Unsafe Statements in Binary Logging

* Functions: RAND, UUID ...
* Trigger or stored program.
* INSERT ... ON DUPLICATE KEY UPDATE.
* Updates using LIMIT.

binlog_format=ROW

https://dev.mysql.com/doc/refman/5.7/en/replication-problems.html