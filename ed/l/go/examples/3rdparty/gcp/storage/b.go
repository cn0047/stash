package storage

import (
	"context"
	"errors"
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
	//uploadFile(ctx, c, "/tmp/debug.txt", bucketName)
	//delete(ctx, c, bucketName, "/tmp/debug.txt")
	//list(ctx, c, bucketName, "")
	//copy(ctx, c, bucketName, "debug.txt", "debug2.txt")
	//move(ctx, c, bucketName, "debug.txt", "debug3.txt")
	update(ctx, c, bucketName, "test.3.processed.txt")
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

func list(ctx context.Context, c *storage.Client, bucketName string, path string) {
	query := &storage.Query{Prefix: path, Delimiter: "/"}

	it := c.Bucket(bucketName).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatalf("failed to list objects, err: %v", err)
		}
		fmt.Printf(" - %v %v\n", attrs.Prefix, attrs.Name)
	}
}

func find(ctx context.Context, c *storage.Client, bucketName string, prefix string) {
	query := &storage.Query{
		Prefix:    prefix,
		Delimiter: "/",
	}

	it := c.Bucket(bucketName).Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			log.Fatalf("failed to perform GCP Storage get next, error: %v", err)
		}

		if attrs.Name != "" {
			// File.
			fmt.Printf("  File: gs://%s/%s (Size: %d bytes)\n", bucketName, attrs.Name, attrs.Size)
		} else if attrs.Prefix != "" {
			// Directory.
			fmt.Printf("  Dir:  gs://%s/%s\n", bucketName, attrs.Prefix)
		}
	}
}

func delete(ctx context.Context, c *storage.Client, bucketName string, pathToObject string) {
	b := c.Bucket(bucketName)
	o := b.Object(pathToObject)
	err := o.Delete(ctx)
	if err != nil {
		log.Fatalf("failed to delete object %q from bucket %q, error: %v", pathToObject, bucketName, err)
	}

	fmt.Printf("deleted: gs://%s/%s\n", bucketName, pathToObject)
}

func copy(ctx context.Context, c *storage.Client, bucketName, srcPath, dstPath string) {
	b := c.Bucket(bucketName)
	src := b.Object(srcPath)
	dst := b.Object(dstPath)
	_, err := dst.CopierFrom(src).Run(ctx)
	if err != nil {
		log.Fatalf("failed to copy object %q to %q, err: %v", srcPath, dstPath, err)
	}
}

func move(ctx context.Context, c *storage.Client, bucketName, srcPath, dstPath string) {
	b := c.Bucket(bucketName)
	src := b.Object(srcPath)
	dst := b.Object(dstPath)
	_, err := dst.CopierFrom(src).Run(ctx)
	if err != nil {
		log.Fatalf("failed to copy object %q to %q, err: %v", srcPath, dstPath, err)
	}
	if err := src.Delete(ctx); err != nil {
		log.Fatalf("failed to delete source object %q after copy, err: %v", srcPath, err)
	}
}

func update(ctx context.Context, c *storage.Client, bucketName string, pathToObject string) {
	b := c.Bucket(bucketName)
	v := storage.ObjectAttrsToUpdate{Metadata: map[string]string{
		"x-processed": "true",
	}}
	res, err := b.Object(pathToObject).Update(ctx, v)
	if err != nil {
		log.Fatalf("failed to update object %q, err: %v", pathToObject, err)
	}
	fmt.Printf("updated: %s, result: %v", pathToObject, res)
}
