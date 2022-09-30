Quiz Engine
-

## Tech stack:

* Golang.
* GCP Spanner.

## Data schema:

* Users - to store data about application user.
* Quizzes - to store data about quizzes, questions and answers stored as nested JSON document.
* Submissions - to store quiz submissions by users.

## How to scale:

* Application - via GCP Cloud Run autoscaling.
* Spanner - via GCP.

## Build and run application:

````sh
make lint
make test

docker build -t quizengine -f Dockerfile .
GOOS=darwin make build

make spanner-start
make run
make spanner-stop
````

## REST API usage examples:

````sh
# health check
curl "http://localhost:9000/health"

# create users
curl -i -X POST "http://localhost:9000/users" -d '{
	"email":"eric@gmail.com",
	"first_name":"Eric",
	"last_name":"Author",
	"password":"pa-ss-123"
}'
curl -i -X POST "http://localhost:9000/users" -d '{
	"email":"kevin@gmail.com",
	"first_name":"Kevin",
	"last_name":"Unknown",
	"password":"pa-ss-123"
}'
curl -i -X POST "http://localhost:9000/users" -d '{
	"email":"bob@gmail.com",
	"first_name":"Bob",
	"last_name":"Reader",
	"password":"pa-ss-123"
}'

# login
curl -i -X POST "http://localhost:9000/auth" -d '{"email":"eric-none@gmail.com", "password":"pa-ss-123======"}'
curl -i -X POST "http://localhost:9000/auth" -d '{"email":"eric@gmail.com", "password":"pa-ss-123"}'
curl -i -X POST "http://localhost:9000/auth" -d '{"email":"kevin@gmail.com", "password":"pa-ss-123"}'
curl -i -X POST "http://localhost:9000/auth" -d '{"email":"bob@gmail.com", "password":"pa-ss-123"}'

t="X-Token: $tkn"


# list quizzes
curl -i -X GET -H $t "http://localhost:9000/v1/quizzes"
curl -i -X GET -H $t "http://localhost:9000/v1/quizzes?author=$uid"

# put quiz 1
curl -i -X PUT -H $t "http://localhost:9000/v1/quizzes/" -d '{
	"published": false,
	"questions": [
		{
			"title":"Moon is a star? (q 2)",
			"answers":[
				{"title":"yes", "correct_answer": false},
				{"title":"maybe", "correct_answer": false},
				{"title":"probably", "correct_answer": false},
				{"title":"never", "correct_answer": true},
				{"title":"no", "correct_answer": true}
			]
		}
	]
}'

# publish quiz 1
curl -i -X PUT -H $t "http://localhost:9000/v1/quizzes/7d01d80c" -d '{
	"published": true,
	"questions": [
		{
			"title":"Is Earth flat?",
			"answers":[
				{"title":"yes", "correct_answer": false},
				{"title":"no", "correct_answer": true}
			]
		},
		{
			"title":"Is Sun bigger than Earth?",
			"answers":[
				{"title":"yes", "correct_answer": true},
				{"title":"no", "correct_answer": false}
			]
		},
		{
			"title":"Moon is a star?",
			"answers":[
				{"title":"yes", "correct_answer": false},
				{"title":"maybe", "correct_answer": false},
				{"title":"probably", "correct_answer": false},
				{"title":"never", "correct_answer": true},
				{"title":"no", "correct_answer": true}
			]
		}
	]
}'

# put quiz 2
curl -i -X PUT -H $t "http://localhost:9000/v1/quizzes/a001d333" -d '{
	"published": true,
	"questions": [
		{
			"title":"Temperature can be measured in?",
			"answers":[
				{"title":"Kelvin", "correct_answer": true},
				{"title":"Fahrenheit", "correct_answer": true},
				{"title":"Gram", "correct_answer": false},
				{"title":"Celsius", "correct_answer": true},
				{"title":"Liters", "correct_answer": false}
			]
		}
	]
}'

# get quiz by id
curl -i -X GET -H $t "http://localhost:9000/v1/quizzes/7d01d80c"

# delete quiz by id
curl -i -X DELETE -H $t "http://localhost:9000/v1/quizzes/7d01d80c"

# submit quiz 1
curl -X POST -H $t "http://localhost:9000/v1/submissions" -d '{
	"quiz_id": "7d01d80c",
	"answers": [
		[1],
		[0],
		[2]
	]
}'

# submit quiz 1 and skip question 3
curl -X POST -H $t "http://localhost:9000/v1/submissions" -d '{
	"quiz_id": "7d01d80c",
	"answers": [
		[1],
		[0],
		[]
	]
}'

# submit quiz 2
curl -X POST -H $t "http://localhost:9000/v1/submissions" -d '{
	"quiz_id": "a001d333",
	"answers": [
		[1, 2, 3]
	]
}'

# list submissions
curl -i -X GET -H $t "http://localhost:9000/v1/submissions?user_id=$uid"
curl -i -X GET -H $t "http://localhost:9000/v1/submissions?quiz_author=$uid"
````

## TODO:

* CI/CD.
* Introduce pagination for list EPs.
* Implement existing in code TODOs.
* Add more tests.
* Introduce cache.
* Finish swagger file.
