Beanstalkd
-

````
# connect to bt
telnet 127.0.0.1 11300

# list of available tubes
list-tubes

# stats about certain tube
stats-tube tube_chat

# switch tube
use tube_chat
use tube_pictureCompression

# add job to tube
# where 95 - str length of data
put 1024 0 60 95
{"login":"user_46205","password":"4rfOoPdwzveveN/XWRFC","target_user_id":202636,"message":"x5"}

put 1024 0 60 95
{"folder":"000009345","thumbnail":true,"file":"public/photo_2016-07-05_17-07-34_thumbnail.jpg"}

# delete job from tube
use tube_documents
delete {jobId}
````

https://github.com/kr/beanstalkd/blob/master/doc/protocol.txt
