package domain

type UploadDatastore interface {
	GetUploadById(uploadId string) *Upload
	CreateUpload(upload *Upload) (uint, error)
}

type BucketDatastore interface {
	GetBucketById(bucketId string) *Bucket
	CreateBucket(bucket *Bucket) (string, error)
}
