package models

type BlobStorage interface {
	Create(bucket *Bucket) (uint, error)
	Exists(bucket *Bucket) bool
}
