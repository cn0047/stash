package queue

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/taskqueue"

	"go-app/config"
)

func AddPingJob(ctx context.Context, msg string) error {
	return addJob(ctx, "ping", config.WorkerPathPing, msg)
}

func AddPingingJob(ctx context.Context, msg string) error {
	return addJob(ctx, "pinging", config.WorkerPathPinging, msg)
}

func addJob(ctx context.Context, queueName string, path string, msg string) error {
	params := map[string][]string{"msg": {msg}}
	t := taskqueue.NewPOSTTask(path, params)

	_, err := taskqueue.Add(ctx, t, queueName)
	if err != nil {
		return fmt.Errorf("[20180703-005] failded add task into %s queue, error: %v", queueName, err)
	}

	return nil
}
