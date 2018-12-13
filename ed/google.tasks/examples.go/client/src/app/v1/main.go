/*

GOPATH=$PWD/ed/google.tasks/examples.go/client/
cd $GOPATH && go get ./... && cd -
go run $GOPATH/src/app/main.go
curl 'localhost:8080/x?n=1'

cd $GOPATH
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app --name=xct --memory="100m" \
	-e GOPATH='/app' \
	-e GOOGLE_APPLICATION_CREDENTIALS='/app/src/app/sa.json' \
	cn007b/go:1.10-gae sh -c 'go run src/app/main.go'
docker exec -it xct sh -c 'watch -n 1 free'
docker exec -it xct sh -c 'vmstat 1'
curl 'localhost:8080/x?n=1500'

gcloud app deploy -q $GOPATH/src/app/app.yaml
curl 'https://x-test-dot-clique-dev.appspot.com/x?n=5'

*/

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"strconv"
	"time"

	"cloud.google.com/go/cloudtasks/apiv2beta3"
	"github.com/labstack/echo"
	"github.com/sevenNt/echo-pprof"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

const (
	ServiceName  = "x-test"
	PathToWorker = "/y"
	Parent       = "projects/clique-dev/locations/us-central1/queues/default"
	LogURL       = "https://realtimelog.herokuapp.com/t"
)

type Job struct {
	ID string `json:"id"`
	I  int    `json:"i"`
}

func main() {
	echoServer := echo.New()
	echoServer.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "active", "version": 2})
	})
	echoServer.GET("/x", x)
	echoServer.POST(PathToWorker, y)
	echopprof.Wrap(echoServer)
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}

func x(ctx echo.Context) error {
	n := ctx.QueryParam("n")
	if n == "" {
		n = "1"
	}

	val, err := strconv.Atoi(n)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"err": "ERR-0"})
	}

	er := addJobs(val)
	if er != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"err": er.Error()})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func y(ctx echo.Context) error {
	job := Job{}
	err := ctx.Bind(&job)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
	}

	data, err := handleJob(job)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"err": err.Error()})
	}

	return ctx.JSON(http.StatusOK, data)
}

func addJobs(n int) error {
	for i := 0; i < n; i++ {
		_, err := addJob(i)
		if err != nil {
			return fmt.Errorf("ERR-5: %s", err)
		}
		runtime.GC()
	}

	return nil
}

func addJob(i int) (interface{}, error) {
	c := context.Background()
	client, err := cloudtasks.NewClient(c)
	if err != nil {
		return nil, fmt.Errorf("ERR-1: %s", err)
	}

	id := time.Now().UnixNano()
	job := Job{ID: "task-" + strconv.Itoa(int(id)), I: i}

	j, err := json.Marshal(&job)
	if err != nil {
		return nil, fmt.Errorf("ERR-2: %s", err)
	}

	req := &tasks.CreateTaskRequest{
		Parent: Parent,
		Task: &tasks.Task{
			Name: Parent + "/tasks/" + job.ID,
			PayloadType: &tasks.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &tasks.AppEngineHttpRequest{
					HttpMethod:       tasks.HttpMethod_POST,
					AppEngineRouting: &tasks.AppEngineRouting{Service: ServiceName},
					RelativeUri:      PathToWorker,
					Body:             j,
					Headers:          map[string]string{"Content-Type": "application/json"},
				},
			},
		},
	}
	respTask, err := client.CreateTask(c, req)
	if err != nil {
		return nil, fmt.Errorf("ERR-3: %s", err)
	}

	er := l("1", map[string]interface{}{"i": i, "CreateTime": respTask.CreateTime})
	if er != nil {
		return nil, fmt.Errorf("ERR-7: %s", respTask)
	}

	//log.Printf("🎾 %+v, %+v", i, respTask.CreateTime)

	return map[string]interface{}{"res": respTask}, nil
}

func handleJob(job Job) (interface{}, error) {
	err := l("2", map[string]interface{}{"job": job})
	if err != nil {
		return nil, fmt.Errorf("ERR-6: %s", err)
	}

	return map[string]interface{}{"status": "done", "job": job}, nil
}

func l(key string, data interface{}) error {
	p, _ := json.Marshal(data)
	_, err := http.Post(LogURL+key, "application/json", bytes.NewBuffer(p))
	if err != nil {
		return fmt.Errorf("ERR-4: %s", err)
	}

	return nil
}
