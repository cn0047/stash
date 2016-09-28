Beanstalkd
-

````
sudo service beanstalkd status

# connect to bt
telnet 127.0.0.1 11300

# list of available tubes
list-tubes

# stats about certain tube
stats-tube tube_pictureCompression

# switch tube
use tube_chat

# add job to tube
# where 95 - str length of data
put 1024 0 60 95
{"login":"user_46205","password":"4rfOoPdwzveveN/XWRFC","target_user_id":202636,"message":"x5"}

# delete job from tube
use tube_documents
delete {jobId}

#  shows the next job to be processed
use tube_documents
peek-ready
````

https://github.com/kr/beanstalkd/blob/master/doc/protocol.txt
