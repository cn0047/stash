RabbitMQ
-

````
sudo service rabbitmq-server status

sudo rabbitmqctl status
sudo rabbitmqctl list_queues
````

Declaring a queue is idempotent - it will only be created if it doesn't exist already.
Keep in mind that messages are sent asynchronously from the server to the clients.

https://www.rabbitmq.com/tutorials/tutorial-two-php.html
