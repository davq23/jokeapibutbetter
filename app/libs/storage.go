package libs

import (
	"time"

	"github.com/aws/aws-sdk-go/service/s3"
)

type StorageManager interface {
	GetSignedDownloadUrl(filePath string, bucket string, duration time.Duration) (string, error)
	GetSignedUploadUrl(
		filePath string,
		bucket string,
		fileTypes []string,
		duration time.Duration,
	) (string, error)
}

type S3StorageManager struct {
	s3Client *s3.S3
}

func NewS3StorageManager(s3Client *s3.S3) *S3StorageManager {
	return &S3StorageManager{
		s3Client: s3Client,
	}
}

func (sm S3StorageManager) GetSignedDownloadUrl(filePath string, bucket string, duration time.Duration) (string, error) {
	req, _ := sm.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &filePath,
	})

	url, _ := req.Presign(duration)

	return url, nil
}

func (sm S3StorageManager) GetSignedUploadUrl(filePath string, bucket string, fileTypes []string, duration time.Duration) (string, error) {
	req, _ := sm.s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &filePath,
	})

	url, _ := req.Presign(duration)

	return url, nil
}
