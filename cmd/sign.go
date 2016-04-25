package main

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))
	for i := 1; i <= 10; i++ {
		params := s3.UploadPartInput{
			Bucket:     aws.String("BucketName"), // Required
			Key:        aws.String("ObjectKey"),  // Required
			PartNumber: aws.Int64(int64(i)),      // Required
			UploadId:   aws.String("1234567890"), // Required
			Body:       nil,
			// Body:       bytes.NewReader([]byte("PAYLOAD")),
			// ContentLength:        aws.Int64(1),
			// RequestPayer:         aws.String("RequestPayer"),
			// SSECustomerAlgorithm: aws.String("SSECustomerAlgorithm"),
			// SSECustomerKey:       aws.String("SSECustomerKey"),
			// SSECustomerKeyMD5:    aws.String("SSECustomerKeyMD5"),
		}
		req, _ := svc.UploadPartRequest(&params)
		str, _ := req.Presign(15 * time.Minute)
		log.Println(str + "\n")
	}
}
