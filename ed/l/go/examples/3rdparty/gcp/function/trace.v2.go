// This is 1 gen func.

package p

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	tracerProvider *sdktrace.TracerProvider
)

func init() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		log.Fatal(err)
	}
	tracerProvider = sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
	otel.SetTracerProvider(tracerProvider)
}

func main() {
}

func doSomething2(ctx context.Context) {
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

func doSomething1(ctx context.Context) {
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

func MainTraceHandlerV21Gen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer tracerProvider.ForceFlush(ctx)

	doSomething2(ctx)

	msg := fmt.Sprintf("traced with traceV21gen at: %s", time.Now().Format(time.Kitchen))
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
	log.Printf("%s", msg)
}
