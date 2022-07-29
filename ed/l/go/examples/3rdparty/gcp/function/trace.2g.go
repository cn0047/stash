// This is 2 gen func.

package p

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"go.opencensus.io/trace"
)

var (
	traceExporter *stackdriver.Exporter
)

func init() {
	l("init1") // invoked by GCP

	functions.HTTP("mainEntryPoint", MainTraceHandler2Gen)

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		//ProjectID: "",
		DefaultTraceAttributes: map[string]interface{}{
			"service-name": "ftrace2g",
		},
		//NumberOfWorkers: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.ProbabilitySampler(1.0),
	})
	if err := exporter.StartMetricsExporter(); err != nil {
		log.Fatal(err)
	}
	traceExporter = exporter
	//defer exporter.Flush()
	//defer exporter.StopMetricsExporter()

	l("init2") // invoked by GCP
}

func main() {
	l("main") // not invoked by GCP

	MainFunc()
}

func MainFunc() {
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("failed to start, err: %v\n", err)
	}
}

func l(d interface{}) {
	j, _ := json.Marshal(map[string]interface{}{"d": d})
	_, _ = http.Post("https://realtimelog.herokuapp.com:443/rkc8q6llprn", "application/json", bytes.NewBuffer(j))
}

func doSomething2(cntx context.Context) {
	ctx, span := trace.StartSpan(cntx, "ftrace2g.doSomething2")
	defer span.End()

	req, err := http.NewRequest("GET", "https://www.youtube.com/watch?v=mZkSBnQUMiU", nil)
	if err != nil {
		log.Printf("failed to create new request err: %v", err)
		return
	}
	req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Printf("failed to send request err: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("got response status code: %v", res.StatusCode)
		return
	}

	doSomething1(ctx)
}

func doSomething1(cntx context.Context) {
	ctx, span := trace.StartSpan(cntx, "ftrace2g.doSomething1")
	defer span.End()

	req, err := http.NewRequest("GET", "https://www.youtube.com/watch?v=mZkSBnQUMiU", nil)
	if err != nil {
		log.Printf("failed to create new request err: %v", err)
		return
	}
	req.WithContext(ctx)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Printf("failed to send request err: %v", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Printf("got response status code: %v", res.StatusCode)
		return
	}
}

func MainTraceHandler2Gen(w http.ResponseWriter, r *http.Request) {
	defer traceExporter.Flush()
	ctx, span := trace.StartSpan(r.Context(), "ftrace2g.MainTraceHandler2Gen")
	defer span.End()

	doSomething2(ctx)

	msg := fmt.Sprintf("traced at: %s", time.Now().Format(time.Kitchen))
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
	log.Printf("%s", msg)
}
