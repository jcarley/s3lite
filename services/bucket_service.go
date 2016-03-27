package services

import (
	"errors"

	"github.com/jcarley/s3lite/domain"
)

type BucketServicer interface {
	AddBucket(bucket *domain.Bucket) error
	DeleteBucketById(id string) error
}

type BucketService struct {
	datastore domain.BucketDatastore
}

func NewBucketService(datastore domain.BucketDatastore) *BucketService {
	return &BucketService{
		datastore: datastore,
	}
}

func (this *BucketService) AddBucket(bucket *domain.Bucket) error {
	if bucket == nil {
		return errors.New("Must supply a bucket")
	}

	id, err := this.datastore.CreateBucket(bucket)
	if err != nil {
		return err
	}

	bucket.Id = id

	return nil
}
