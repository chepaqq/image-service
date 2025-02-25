package storage

import (
	"context"
	"io"
	"net/url"
)

type Storage interface {
	Upload(ctx context.Context, bucketName, objectName string, reader io.Reader, contentType string) (*url.URL, error)
	BucketExists(ctx context.Context, bucketName string) (bool, error)
	CreateBucket(ctx context.Context, bucketName, region string) error
}
