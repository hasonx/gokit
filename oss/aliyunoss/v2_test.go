package aliyunoss

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// readConfigFromEnv reads OSS configuration from environment variables.
// Note: For a real test run, these environment variables must be set.
func readConfigFromEnv(t *testing.T) (Config, string) {
	endpoint := os.Getenv("ALI_OSS_ENDPOINT")
	accessKeyID := os.Getenv("ALI_OSS_ACCESS_KEY_ID")
	accessKeySecret := os.Getenv("ALI_OSS_ACCESS_KEY_SECRET")
	bucketName := os.Getenv("ALI_OSS_BUCKET_NAME")
	region := os.Getenv("ALI_OSS_REGION")

	if endpoint == "" || accessKeyID == "" || accessKeySecret == "" || bucketName == "" || region == "" {
		t.Skip("Skipping OSS integration tests: environment variables not set")
	}

	return Config{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		Region:          region,
	}, bucketName
}

func TestAliyunOSS(t *testing.T) {
	config, bucketName := readConfigFromEnv(t)
	ctx := context.Background()

	client := New(config)
	assert.NotNil(t, client)

	// Prepare a temporary file for testing
	content := []byte("hello, world!")
	tmpFile, err := os.CreateTemp("", "oss-test-*.txt")
	assert.NoError(t, err)
	_, err = tmpFile.Write(content)
	assert.NoError(t, err)
	_ = tmpFile.Close()
	defer func() {
		_ = os.Remove(tmpFile.Name()) // Clean up the temp file
	}()

	objectName := "test/test-object-" + time.Now().Format("20060102150405") + ".txt"
	downloadPath := tmpFile.Name() + ".download"
	defer func() {
		_ = os.Remove(downloadPath) // Clean up the downloaded file
	}()

	t.Run("BucketExists", func(t *testing.T) {
		exist, err := client.BucketExists(ctx, bucketName)
		assert.NoError(t, err)
		assert.True(t, exist)
	})

	t.Run("UploadFile", func(t *testing.T) {
		err := client.UploadFile(ctx, bucketName, objectName, tmpFile.Name())
		assert.NoError(t, err)
	})

	t.Run("Presign", func(t *testing.T) {
		signedURL, err := client.Presign(ctx, bucketName, objectName, 3600)
		assert.NoError(t, err)
		assert.NotEmpty(t, signedURL)
		t.Logf("Generated Signed URL: %s", signedURL)
	})

	t.Run("DownloadFile", func(t *testing.T) {
		err := client.DownloadFile(ctx, bucketName, objectName, downloadPath)
		assert.NoError(t, err)

		downloadedContent, err := os.ReadFile(downloadPath)
		assert.NoError(t, err)
		assert.Equal(t, content, downloadedContent)
	})

	t.Run("DeleteObject", func(t *testing.T) {
		// err := client.DeleteObject(ctx, bucketName, objectName)
		// assert.NoError(t, err)

		// Verify deletion by trying to download again, which should fail
		err = client.DownloadFile(ctx, bucketName, objectName, downloadPath)
		assert.Error(t, err) // Expect an error because the object should not exist
	})
}
