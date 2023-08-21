package storage

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// ConnectMinio establishes a connection to Minio
func ConnectMinio(useSSL bool, endpoint, bucketName, bucketLocation string) (*minio.Client, error) {
	// Initialize minio client object.
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewEnvMinio(),
		Secure: useSSL,
	})
	if errInit != nil {
		return nil, errInit
	}

	// Make a new bucket
	err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: bucketLocation})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return nil, err
		}
	}
	return minioClient, nil
}
