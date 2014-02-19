package domain

type BlobStorage interface {
  Create(bucket *Bucket) (bool, error)
  Exists(bucket *Bucket) bool
}



