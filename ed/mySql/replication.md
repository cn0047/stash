Replication
-

*MySQL 5.5*

Depending on the configuration, you can replicate all databases, selected databases, or even selected tables within a database.
<br>Asynchronous replication - one server acts as the master, while one or more other servers act as slaves.
<br>Synchronous replication which is a characteristic of MySQL Cluster.
<br>There are two core types of replication format:
* Statement Based Replication (SBR) - which replicates entire SQL statements.
* Row Based Replication (RBR) - which replicates only the changed rows.

Replication between servers in MySQL is based on the binary logging mechanism. The information in the binary log is stored in different logging formats according to the database changes being recorded.

####Setting the Replication Master Configuration
You will need to shut down your MySQL server and edit the my.cnf or my.ini file:
````sql
[mysqld]
log-bin=mysql-bin
server-id=1
````
After making the changes, restart the server.
<br>For the greatest possible durability and consistency in a replication setup using InnoDB with transactions, you should use innodb_flush_log_at_trx_commit=1 and sync_binlog=1 in the master my.cnf file.

####Setting the Replication Slave Configuration
 You should shut down your slave server and edit the configuration to specify a unique server ID. For example:
````sql
[mysqld]
server-id=2
````
After making the changes, restart the server.
<br>You do not have to enable binary logging on the slave for replication to be enabled.

####Creating a User for Replication
Each slave must connect to the master using a MySQL user name and password.
````sql
CREATE USER 'repl'@'%.mydomain.com' IDENTIFIED BY 'slavepass';
GRANT REPLICATION SLAVE ON *.* TO 'repl'@'%.mydomain.com';
````

####Obtaining the Replication Master Binary Log Coordinates
Start a session on the master by connecting to it with the command-line client, and:
````sql
FLUSH TABLES WITH READ LOCK;
````
For InnoDB tables, note that FLUSH TABLES WITH READ LOCK also blocks COMMIT operations.
<br>If you exit the client, the lock is released.
<br>On the master:
````sql
SHOW MASTER STATUS;
+------------------+----------+--------------+------------------+
| File             | Position | Binlog_Do_DB | Binlog_Ignore_DB |
+------------------+----------+--------------+------------------+
| mysql-bin.000003 | 73       | test         | manual,mysql     |
+------------------+----------+--------------+------------------+
````

####Creating a Data Snapshot Using mysqldump



























[>>>](http://dev.mysql.com/doc/refman/5.5/en/replication-howto-mysqldump.html)