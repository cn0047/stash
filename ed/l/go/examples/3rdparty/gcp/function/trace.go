// This is 1 gen func.

package p

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

var (
	traceExporter *stackdriver.Exporter
)

func init() {
	l("init") // invoked by GCP

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		//ProjectID: "",
		DefaultTraceAttributes: map[string]interface{}{
			"service-name": "ftrace1g",
		},
		NumberOfWorkers: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.ProbabilitySampler(1.0), // % of requests to trace
	})
	if err := exporter.StartMetricsExporter(); err != nil {
		log.Fatal(err)
	}
	traceExporter = exporter
	//defer exporter.Flush()
	//defer exporter.StopMetricsExporter()
}

func main() {
	l("main") // not invoked by GCP
}

func l(d interface{}) {
	j, _ := json.Marshal(map[string]interface{}{"d": d})
	_, _ = http.Post("https://realtimelog.herokuapp.com:443/rkc8q6llprn", "application/json", bytes.NewBuffer(j))
}

func doSomething(ctx context.Context) {
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

func MainTraceHandler1Gen(w http.ResponseWriter, r *http.Request) {
	defer traceExporter.Flush()

	doSomething(context.Background())

	msg := fmt.Sprintf("traced at: %s", time.Now().Format(time.Kitchen))
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
	log.Printf("%s", msg)
}
