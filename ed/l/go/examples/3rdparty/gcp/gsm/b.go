package gsm

import (
	"context"
	"fmt"
	"log"

	sm "cloud.google.com/go/secretmanager/apiv1"
	"google.golang.org/api/option"
	smpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func Run(projectID string, SAFilePath string) {
	var err error
	ctx := context.Background()

	c, err := getClient(ctx, projectID, SAFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	secretName := "test"
	_, err = get(ctx, c, secretName)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient(ctx context.Context, projectID string, SAFilePath string) (*sm.Client, error) {
	opts := []option.ClientOption{
		option.WithCredentialsFile(SAFilePath),
	}
	c, err := sm.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create new secret manager client, err: %w", err)
	}

	return c, nil
}

func get(ctx context.Context, client *sm.Client, secretName string) (string, error) {
	r, err := client.GetSecret(ctx, &smpb.GetSecretRequest{Name: secretName})
	if err != nil {
		return "", fmt.Errorf("failed to get secret, err: %w", err)
	}

	fmt.Printf("r: %+v \n", r)

	return "", nil
}
