// Package storage implements the functionality of file storage ()
package storage

import (
	"bytes"
	"fmt"
	io "io"
	"io/ioutil"
	"time"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/Konstantsiy/image-converter/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// URLTimeout determines the validity period of the downloaded image URL.
const URLTimeout = 10 * time.Minute

// S3Storage implements the functionality of file storage (Amazon S3).
type S3Storage struct {
	svc    *s3.S3
	s3conf *config.AWSConfig
}

// NewStorage creates new file storage with the given S3 configs and bucket name.
func NewStorage(s3conf *config.AWSConfig) (*S3Storage, error) {
	err := validateAWSConfig(s3conf)
	if err != nil {
		return nil, err
	}

	svc, err := initS3ServiceClient(*s3conf)
	if err != nil {
		return nil, err
	}

	return &S3Storage{svc: svc, s3conf: s3conf}, nil
}

// validateAWSConfig validates AWS configurations.
func validateAWSConfig(s3conf *config.AWSConfig) error {
	if s3conf.BucketName == "" || s3conf.SecretAccessKey == "" || s3conf.AccessKeyID == "" || s3conf.Region == "" {
		return fmt.Errorf("AWS configurations should not be empty")
	}
	return nil
}

// initS3ServiceClient initializes SDK's service client.
func initS3ServiceClient(s3conf config.AWSConfig) (*s3.S3, error) {
	s3session, err := createSession(&s3conf)
	if err != nil {
		return nil, err
	}

	return s3.New(s3session), nil
}

// createSession creates and returns a new session.
func createSession(s3conf *config.AWSConfig) (*session.Session, error) {
	s3session, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3conf.Region),
		Credentials: credentials.NewStaticCredentials(s3conf.AccessKeyID, s3conf.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("can't create session, %w", err)
	}
	return s3session, nil
}

// UploadFile uploads the given file to the bucket.
func (s *S3Storage) UploadFile(file io.ReadSeeker, fileID string) error {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Body:   file,
		Bucket: aws.String(s.s3conf.BucketName),
		Key:    aws.String(fileID),
		ACL:    aws.String(s3.BucketCannedACLPublicRead),
	})
	if err != nil {
		return fmt.Errorf("can't upload file: %w", err)
	}

	return nil
}

// DownloadFile downloads a file from the storage by the given id.
func (s *S3Storage) DownloadFile(fileID string) (io.ReadSeeker, error) {
	resp, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.s3conf.BucketName),
		Key:    aws.String(fileID),
	})
	if err != nil {
		return nil, fmt.Errorf("can't download file with id %s: %w", fileID, err)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("can't serialize response body: %w", err)
	}

	return bytes.NewReader(buf), nil
}

// GetDownloadURL returns URL to download а file from the bucket by the given file id.
func (s *S3Storage) GetDownloadURL(fileID string) (string, error) {
	req, _ := s.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.s3conf.BucketName),
		Key:    aws.String(fileID),
	})

	url, err := req.Presign(URLTimeout)
	if err != nil {
		return "", fmt.Errorf("can't create requets's presigned URL, %w", err)
	}

	return url, err
}
