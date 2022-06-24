#### RabbitMQ

# init rabbit
docker run -it --rm --hostname localhost --name rabbit rabbitmq:latest

# test rabbitmq queues
docker exec rabbit rabbitmqctl list_queues name messages messages_ready messages_unacknowledged
