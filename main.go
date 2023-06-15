package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func connectToAWS(accessKey, secretKey, region, bucketName string) (*s3.S3, error) {
	// Create a new AWS session with the provided access key, secret key, and region
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	// Create an S3 client using the session
	s3Client := s3.New(sess)

	// Verify the connection by checking if the specified bucket exists
	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to AWS successfully")

	return s3Client, nil
}

func uploadFileToS3(s3Client *s3.S3, bucketName, objectKey, filePath string) error {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create an S3 uploader using the provided S3 client
	uploader := s3manager.NewUploaderWithClient(s3Client)

	// Upload the file to S3
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Println("File uploaded successfully")
	return nil
}
func main() {

	accessKey := ""
	secretKey := ""
	region := ""
	bucketName := ""
	objectKey := ""
	filePath := ""

	s3Client, err := connectToAWS(accessKey, secretKey, region, bucketName)
	if err != nil {
		panic(err)
	}

	// Use the s3Client for further interactions with AWS S3
	_ = s3Client

	err = uploadFileToS3(s3Client, bucketName, objectKey, filePath)
	if err != nil {
		panic(err)
	}
}
