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
````
[mysqld]
log-bin=mysql-bin
server-id=1
````
After making the changes, restart the server.
<br>For the greatest possible durability and consistency in a replication setup using InnoDB with transactions, you should use innodb_flush_log_at_trx_commit=1 and sync_binlog=1 in the master my.cnf file.

####Setting the Replication Slave Configuration
 You should shut down your slave server and edit the configuration to specify a unique server ID. For example:
````
[mysqld]
server-id=2
````
After making the changes, restart the server.
<br>You do not have to enable binary logging on the slave for replication to be enabled.











[>>>](http://dev.mysql.com/doc/refman/5.5/en/replication-howto-repuser.html)