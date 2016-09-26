RabbitMQ
-

[sourcecode](https://github.com/rabbitinaction/sourcecode)

````
erl -sname test
````
````
sudo service rabbitmq-server status

sudo rabbitmqctl status
sudo rabbitmqctl cluster_status
sudo rabbitmqctl stop_app

sudo rabbitmqctl list_queues
sudo rabbitmqctl list_queues name durable auto_delete messages consumers memory messages_ready messages_unacknowledged

sudo rabbitmqctl list_exchanges

sudo rabbitmqctl list_bindings

sudo rabbitmqctl list_users

sudo rabbitmqctl list_permissions

sudo rabbitmqctl change_password cashing-tier compl3xPassword
````

Declaring a queue is idempotent - it will only be created if it doesn't exist already.
Keep in mind that messages are sent asynchronously from the server to the clients.
The `durable` queue won't be lost even if RabbitMQ restarts.

Marking messages as persistent doesn't fully guarantee that a message won't be lost...
It tells RabbitMQ to save the message to disk.

Exchange - receives messages from producers and pushes them to queues.
There are a few exchange types available: direct, topic, headers, fanout.

That relationship between exchange and a queue is called a binding.

Avoid black hole messages:

* Have its delivery mode option set to 2 (persistent)
* Be published into a durable exchange
* Arrive in a durable queue

## Cluster

RabbitMQ only requires that one node in a cluster be a disk node.
Every other node can be a RAM node.

If you only have one disk node and that node happens to be down,
your cluster can continue to route messages but you can’t do any of the following:

* Create queues
* Create exchanges
* Create bindings
* Add users
* Change permissions
* Add or remove cluster nodes

When RAM nodes restart, they connect to the disk nodes they’re preconfigured with
to download the current copy of the cluster’s metadata.

## Rabbitmq management pugin

````
sudo /usr/lib/rabbitmq/bin/rabbitmq-plugins list
sudo /usr/lib/rabbitmq/bin/rabbitmq-plugins enable rabbitmq_management
sudo /usr/sbin/rabbitmqctl stop
sudo /usr/sbin/rabbitmq-server -detached

# go to: http://localhost:15672/
# use guest:guest

wget http://localhost:15672/cli/rabbitmqadmin
chmod +x rabbitmqadmin
./rabbitmqadmin -V "/" list exchanges
./rabbitmqadmin purge queue name=test
./rabbitmqadmin -u guest -p guest declare exchange name=cli_test type=direct
./rabbitmqadmin list connections name
````
