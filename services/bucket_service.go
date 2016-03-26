package services

import (
	"github.com/jcarley/s3lite/domain"
)

type BucketService interface {
	AddBucket(bucket *domain.Bucket) error
}
