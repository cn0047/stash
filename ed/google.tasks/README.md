Tasks
-

[docs](https://cloud.google.com/tasks/docs/)

````bash
gcloud beta tasks queues list

gueueName=default
gcloud beta tasks queues describe $gueueName

curl 'https://cloudtasks.googleapis.com/$discovery/rest?version=v2beta3' | jq

# queues list
curl 'https://cloudtasks.googleapis.com/v2beta3/projects/prj/locations/us-central1/queues'
````
