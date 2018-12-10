Task Queue
-

[appengine task deadline](https://cloud.google.com/appengine/docs/standard/go/taskqueue/push/#the_task_deadline)

In push queues tasks are delivered to a worker service based on the queue's configuration.
In pull queues the worker service must ask the queue for tasks.

Tasks handled by automatic scaling services must finish in ten minutes.
Tasks handled by basic and manual scaling services can run for up to 24 hours.

The task queue uses `token buckets` to control the rate of task execution.
Each named queue has a token bucket that holds tokens.
Default `bucket_size` is 5.
Each time your application executes a task, a token is removed from the bucket.
Your app continues processing tasks in the queue until the queue's bucket runs out of tokens.
App Engine refills the bucket with new tokens continuously
based on the rate that you specified for the queue.

Default `total_storage_limit` is 500M.

You can enqueue a task as part of a Cloud Datastore transaction!

````go
params := map[string][]string{
  "project": {prj.ID},
  "url":     {prj.URL},
  "method":  {prj.Method},
  "json":    {prj.JSON},
}
t := taskqueue.NewPOSTTask(workerPath, params)
_, err := taskqueue.Add(ctx, t, queueName)

# or
t := &taskqueue.Task{
  Path:    workerPath,
  Payload: jsonData,
  Header:  http.Header{"Content-Type": []string{"application/json"}},
  Method:  http.MethodPost,
}
taskqueue.Add(cctx.GAECtx, t, queueName)
````
