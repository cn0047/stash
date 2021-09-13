package storage

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Run(projectID, SAFilePath, bucketName string) {
	ctx := context.Background()
	c := getClientWithCredentialsFile(ctx, SAFilePath)
	//c := getClientWithCredentialsJSON(ctx, []byte(``))

	//printBuckets(ctx, c, projectID)
	//uploadTestFile(ctx, c, bucketName)
	uploadFile(ctx, c, "/tmp/debug.txt", bucketName)
}

func getClientWithAPIKey(ctx context.Context, key string) *storage.Client {
	client, err := storage.NewClient(ctx, option.WithAPIKey(key))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getClientWithCredentialsFile(ctx context.Context, SAFilePath string) *storage.Client {
	// INFO: option.WithServiceAccountFile is deprecated in favor of option.WithCredentialsFile.
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(SAFilePath))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func getClientWithCredentialsJSON(ctx context.Context, credentialsJSON []byte) *storage.Client {
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(credentialsJSON))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func uploadTestFile(ctx context.Context, c *storage.Client, bucketName string) {
	b := c.Bucket(bucketName)

	wc := b.Object("test.txt").NewWriter(ctx)
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{"x-test": "true"}

	if _, err := wc.Write([]byte("it works\n")); err != nil {
		log.Fatalf("unable to write data to bucket %v, error: %+v", bucketName, err)
	}

	if err := wc.Close(); err != nil {
		log.Fatalf("unable to close bucket %v file, error: %+v", bucketName, err)
	}
}

func uploadFile(ctx context.Context, c *storage.Client, srcFile string, bucketName string) {
	content, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Fatalf("failed to read file, error: %v", err)
	}

	b := c.Bucket(bucketName)

	wc := b.Object("test2.txt").NewWriter(ctx)
	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{"x-test": "true"}

	if _, err := io.Copy(wc, strings.NewReader(string(content))); err != nil {
		log.Fatalf("failed to perform GCP Storage write, error: %v", err)
	}

	if err := wc.Close(); err != nil {
		log.Fatalf("failed to perform GCP Storage close, error: %v", err)
	}
}

func printBuckets(ctx context.Context, c *storage.Client, projectID string) {
	it := c.Buckets(ctx, projectID)
	fmt.Println("Buckets:")
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(attrs.Name)
	}
}
