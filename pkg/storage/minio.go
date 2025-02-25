package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinioStorage struct implements the Storage interface
type MinioStorage struct {
	client *minio.Client
}

// NewMinioStorage initializes Minio storage and ensures the bucket exists
func NewMinioStorage(endpoint, bucketName, region string, useSSL bool) (*MinioStorage, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewEnvMinio(),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	storage := &MinioStorage{client: client}
	// Ensure the bucket exists
	err = storage.ensureBucket(context.Background(), bucketName, region)
	if err != nil {
		return nil, err
	}

	return storage, nil
}

// ensureBucket checks if the bucket exists, and creates it if necessary
func (m *MinioStorage) ensureBucket(ctx context.Context, bucketName, region string) error {
	exists, err := m.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if !exists {
		err = m.CreateBucket(ctx, bucketName, region)
		if err != nil {
			return err
		}
	}

	return nil
}

// Upload uploads an object to MinIO and returns its public URL
func (m *MinioStorage) Upload(ctx context.Context, bucketName, objectName string, reader io.Reader, contentType string) (*url.URL, error) {
	_, err := m.client.PutObject(ctx, bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return nil, err
	}

	// Construct the URL
	objectURL := fmt.Sprintf("%s/%s/%s", m.client.EndpointURL(), bucketName, objectName)
	parsedURL, err := url.Parse(objectURL)
	if err != nil {
		return nil, err
	}

	return parsedURL, nil
}

// BucketExists checks if a bucket exists
func (m *MinioStorage) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	return m.client.BucketExists(ctx, bucketName)
}

// CreateBucket creates a new bucket if it does not exist
func (m *MinioStorage) CreateBucket(ctx context.Context, bucketName, region string) error {
	err := m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err != nil {
		exists, errCheck := m.BucketExists(ctx, bucketName)
		if errCheck == nil && exists {
			return nil
		}
		return err
	}
	return nil
}
