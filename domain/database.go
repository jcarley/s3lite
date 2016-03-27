package domain

type UploadDatastore interface {
	GetUploadById(uploadId string) *Upload
	CreateUpload(upload *Upload) (string, error)
}

type BucketDatastore interface {
	GetBucketById(bucketId string) *Bucket
	CreateBucket(bucket *Bucket) (string, error)
	DeleteBucketById(bucketId string) error
}
