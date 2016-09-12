RabbitMQ
-

````
sudo service rabbitmq-server status

sudo rabbitmqctl status

sudo rabbitmqctl list_queues
sudo rabbitmqctl list_queues name messages_ready messages_unacknowledged

sudo rabbitmqctl list_exchanges

sudo rabbitmqctl list_bindings
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
