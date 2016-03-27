package services

import "github.com/jcarley/s3lite/domain"

type BucketServicer interface {
	AddBucket(bucket *domain.Bucket) error
	DeleteBucketById(id string) error
}

type BucketService struct {
}
