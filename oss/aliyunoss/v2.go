package aliyunoss

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// Config holds the configuration for the Aliyun OSS client.
type Config struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Region          string
}

// AliyunOSS is a client for Aliyun OSS operations.
type AliyunOSS struct {
	client *oss.Client
	config Config
}

// New creates a new AliyunOSS client.
func New(config Config) *AliyunOSS {
	provider := credentials.NewStaticCredentialsProvider(
		config.AccessKeyID,
		config.AccessKeySecret,
	)
	cfg := oss.LoadDefaultConfig().
		WithEndpoint(config.Endpoint).
		WithRegion(config.Region).
		WithCredentialsProvider(provider)

	client := oss.NewClient(cfg)

	return &AliyunOSS{
		client: client,
		config: config,
	}
}

// UploadFile uploads a file from a local path to the specified bucket.
func (a *AliyunOSS) UploadFile(ctx context.Context, bucketName, objectName, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = a.client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
		Body:   file,
	})
	return err
}

// Upload uploads data from an io.Reader to the specified bucket.
func (a *AliyunOSS) Upload(ctx context.Context, bucketName, objectName string, reader io.Reader) error {
	_, err := a.client.PutObject(ctx, &oss.PutObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
		Body:   reader,
	})
	return err
}

// DownloadFile downloads an object from the specified bucket to a local file.
func (a *AliyunOSS) DownloadFile(ctx context.Context, bucketName, objectName, filePath string) error {
	resp, err := a.client.GetObject(ctx, &oss.GetObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
	})
	if err != nil {
		return err
	}
	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	_, err = io.Copy(file, resp.Body)
	return err
}

// DeleteObject deletes an object from the specified bucket.
func (a *AliyunOSS) DeleteObject(ctx context.Context, bucketName, objectName string) error {
	_, err := a.client.DeleteObject(ctx, &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
	})
	return err
}

// Presign generates a presigned URL for an object.
func (a *AliyunOSS) Presign(ctx context.Context, bucketName, objectName string, expiredInSec int64) (string, error) {
	if expiredInSec <= 0 {
		expiredInSec = 60 // Default to 60 seconds
	}
	duration := time.Duration(expiredInSec) * time.Second
	getObjRequest := &oss.GetObjectRequest{
		Bucket: oss.Ptr(bucketName),
		Key:    oss.Ptr(objectName),
	}
	resp, err := a.client.Presign(ctx, getObjRequest, oss.PresignExpires(duration))
	if err != nil {
		return "", err
	}
	return resp.URL, nil
}

// CreateBucket creates a new bucket.
func (a *AliyunOSS) CreateBucket(ctx context.Context, bucketName string) error {
	_, err := a.client.PutBucket(ctx, &oss.PutBucketRequest{
		Bucket: oss.Ptr(bucketName),
		Acl:    oss.BucketACLPrivate,
		CreateBucketConfiguration: &oss.CreateBucketConfiguration{
			StorageClass:       oss.StorageClassStandard,
			DataRedundancyType: oss.DataRedundancyLRS,
		},
	})
	return err
}

// DeleteBucket deletes a bucket.
func (a *AliyunOSS) DeleteBucket(ctx context.Context, bucketName string) error {
	_, err := a.client.DeleteBucket(ctx, &oss.DeleteBucketRequest{
		Bucket: oss.Ptr(bucketName),
	})
	return err
}

// ListBuckets lists all buckets.
func (a *AliyunOSS) ListBuckets(ctx context.Context) ([]string, error) {
	resp, err := a.client.ListBuckets(ctx, &oss.ListBucketsRequest{})
	if err != nil {
		return nil, err
	}

	var bucketNames []string
	for _, bucket := range resp.Buckets {
		bucketNames = append(bucketNames, *bucket.Name)
	}
	return bucketNames, nil
}

// BucketExists checks if a bucket exists.
func (a *AliyunOSS) BucketExists(ctx context.Context, bucketName string) (bool, error) {
	_, err := a.client.IsBucketExist(ctx, bucketName)
	if err != nil {
		var serviceError *oss.ServiceError
		if errors.As(err, &serviceError) && serviceError.Code == "NoSuchBucket" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
