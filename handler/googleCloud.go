package handler

import (
	"context"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/fredliang44/multicloud-upload-service/utils"
)

// GoogleBucket var loacte google cloud buckey
var GoogleBucket = initBucket()

func initBucket() storage.BucketHandle {
	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := utils.Config.GoogleCloud.BucketName

	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	return *bucket
}

// FileWriter is a func to write file to bucket
func FileWriter(filename string, file io.Reader) (err error) {
	ctx := context.Background()
	wc := GoogleBucket.Object(filename).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}
