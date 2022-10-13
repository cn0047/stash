Beanstalkd
-

[protocol](https://github.com/kr/beanstalkd/blob/master/doc/protocol.txt)

````sh
sudo service beanstalkd status
ps aux | grep beanstalkd | grep grep -v
````

#### Commmands:

````
put         - create job
reserve     - reserve for the worker
delete      - delete job
release     - back job to ready status
bury        - change job status to buried
kick        - change job status to ready
touch       - allows a worker to request more time to work on a job (useful for jobs that potentially take a long time)
watch       - adds the named tube to the watch list
ignore      - removes the named tube from the watch list
peek        - let the client inspect a job in the system
peek-buried - return the next job in the list of buried jobs
quit
````

Examples:

````bash
# connect to bt
telnet 127.0.0.1 11300

# list of available tubes
list-tubes

# statistical information about the system as a whole
stats

# stats about certain tube
stats-tube tube_pictureCompression

# switch tube
use tube_chat

# add job to tube
# where 95 - str length of data
# use tube_chat
put 1024 0 60 91
{"login":"user_46205","password":"4rfOoPdwz====___","target_user_id":282040,"message":"x5"}
# use tube_pictureCompression
put 1024 0 60 95
{"folder":"000009345","thumbnail":true,"file":"public/photo_2016-07-05_17-07-34_thumbnail.jpg"}
# use tube_emails
put 1024 0 60 110
{"reason":"Marketing or partnership","email":"x@dkint.com","name":"xxx","message":"TEST"}

# delete job from tube
use tube_documents
delete {jobId}

# shows the next job to be processed
peek-ready

# kick jobs
# 5 - the number of jobs to kick
kick 5

# kick job by id
kick-job jobId

stats-job jobId
````

Job statuses:

* READY - put job.
* DELAYED - put with delay.
* RESERVED - reserve job for the worker.
* BURIED.
