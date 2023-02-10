package iap

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"google.golang.org/api/idtoken"
)

const (
	clientID = ""
	host     = ""
)

func Run(projectID string) {
	ctx := context.Background()

	err := makeIAPRequest(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func makeIAPRequest(ctx context.Context) error {
	client, err := idtoken.NewClient(ctx, clientID)
	if err != nil {
		return fmt.Errorf("failed to create idtoken.NewClient: %+v", err)
	}

	req, err := http.NewRequest("GET", host+"/v1/ping", nil)
	req.Header.Set("X-Env", "test")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request, err: %+v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body, err: %+v", err)
	}

	fmt.Printf("Response: %+v \n", body)

	return nil
}
