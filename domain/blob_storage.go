package domain

type BlobStorage interface {
	Create(bucket *Bucket) (int64, error)
	Exists(bucket *Bucket) bool
}
