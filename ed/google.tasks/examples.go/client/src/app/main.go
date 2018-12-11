/*

GOPATH=$PWD/ed/google.tasks/examples/client/
go get $GOPATH/...
go run ed/google.tasks/examples/client/src/app/main.go

*/

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/cloudtasks/apiv2beta3"
	"github.com/labstack/echo"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2beta3"
)

func main() {
	echoServer := echo.New()
	echoServer.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"c": 204})
	})
	echoServer.GET("/x", x)
	echoServer.POST("/y", y)
	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}

func x(ctx echo.Context) error {
	c := context.Background()
	client, err := cloudtasks.NewClient(c)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"err": err})
	}

	id := time.Now().Unix()
	n := strconv.Itoa(int(id))
	req := &tasks.CreateTaskRequest{
		Parent: "projects/prj/locations/us-central1/queues/default",
		Task: &tasks.Task{
			Name: "projects/prj/locations/us-central1/queues/default/tasks/" + n,
			PayloadType: &tasks.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &tasks.AppEngineHttpRequest{
					HttpMethod:       tasks.HttpMethod_POST,
					AppEngineRouting: &tasks.AppEngineRouting{Service: "x-test-service"},
					RelativeUri:      "/y",
					Body:             []byte(`{"c":"task-` + n + `"}`),
					Headers:          map[string]string{"Content-Type": "application/json"},
				},
			},
		},
	}
	respTask, err := client.CreateTask(c, req)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"err": err})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"c": respTask})
}

type J struct {
	C string `json:"c"`
}

func y(ctx echo.Context) error {
	j := J{}
	err := ctx.Bind(&j)
	if err != nil {
		return ctx.JSON(http.StatusOK, map[string]interface{}{"err": err})
	}

	p, _ := json.Marshal(map[string]interface{}{"j": j})
	http.Post("https://realtimelog.herokuapp.com/t", "application/json", bytes.NewBuffer(p))

	log.Printf("[x-test] %+v", j)
	fmt.Printf("[x-test] %+v", j)
	return ctx.JSON(http.StatusOK, map[string]interface{}{"j": j})
}
