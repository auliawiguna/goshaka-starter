package helpers

import (
	"bytes"
	"fmt"
	"goshaka/configs"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Upload file to AWS S3
//
//	param file  *multipart.FileHeader
//	return any, error
func UploadFileToS3(file *multipart.FileHeader) (any, error) {
	// Open the file for use
	uploadFile, err := file.Open()
	if err != nil {
		return false, fmt.Errorf("failed to open file: %w", err)
	}
	defer uploadFile.Close()

	fileBuffer := make([]byte, file.Size)
	_, err = uploadFile.Read(fileBuffer)

	if err != nil {
		return false, fmt.Errorf("failed to read file: %w", err)
	}

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(configs.GetEnv("AWS_DEFAULT_REGION")),
		Credentials: credentials.NewStaticCredentials(configs.GetEnv("AWS_ACCESS_KEY_ID"), configs.GetEnv("AWS_SECRET_ACCESS_KEY"), ""),
	})

	if err != nil {
		return false, fmt.Errorf("failed to set session: %w", err)
	}

	svc := s3.New(sess)
	// Upload the file to S3.
	s3PutObjectOutput, err := svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(configs.GetEnv("AWS_BUCKET")),
		Key:                  aws.String(file.Filename),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(file.Size),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})
	if err != nil {
		return false, fmt.Errorf("failed to upload file to S3: %w", err)
	}

	// Convert the interface to a map
	resultMap := make(map[string]interface{})
	resultMap["BucketKeyEnabled"] = s3PutObjectOutput.BucketKeyEnabled
	resultMap["ChecksumCRC32"] = s3PutObjectOutput.ChecksumCRC32
	resultMap["ChecksumCRC32C"] = s3PutObjectOutput.ChecksumCRC32C
	resultMap["ChecksumSHA1"] = s3PutObjectOutput.ChecksumSHA1
	resultMap["ChecksumSHA256"] = s3PutObjectOutput.ChecksumSHA256
	resultMap["ETag"] = s3PutObjectOutput.ETag
	resultMap["Expiration"] = s3PutObjectOutput.Expiration
	resultMap["RequestCharged"] = s3PutObjectOutput.RequestCharged
	resultMap["SSECustomerAlgorithm"] = s3PutObjectOutput.SSECustomerAlgorithm
	resultMap["SSECustomerKeyMD5"] = s3PutObjectOutput.SSECustomerKeyMD5
	resultMap["SSEKMSEncryptionContext"] = s3PutObjectOutput.SSEKMSEncryptionContext
	resultMap["SSEKMSKeyId"] = s3PutObjectOutput.SSEKMSKeyId
	resultMap["ServerSideEncryption"] = s3PutObjectOutput.ServerSideEncryption
	resultMap["VersionId"] = s3PutObjectOutput.VersionId

	// Generate a pre-signed URL for the uploaded file
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(configs.GetEnv("AWS_BUCKET")),
		Key:    aws.String(file.Filename),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	// assign presigned url
	// Add a new attribute to the map
	resultMap["AWSUrl"] = urlStr

	return resultMap, nil
}
