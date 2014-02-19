package infrastructure

import (
  "sync"
  "errors"

  "github.com/jcarley/s3lite/domain"
)

var (
  ErrBucketAlreadyExists = errors.New("bucket already exists")
)

type InMemoryBlobStorage struct {
  sync.RWMutex
  buckets map[uint]*domain.Bucket
  seq uint
}

func NewInMemoryBlobStorage() *InMemoryBlobStorage {
  return &InMemoryBlobStorage{
    buckets: make(map[uint]*domain.Bucket),
  }
}

func (bs *InMemoryBlobStorage) Create(bucket *domain.Bucket) (uint, error) {
  bs.Lock()
  defer bs.Unlock()

  if !bs.Exists(bucket) {
    return 0, ErrBucketAlreadyExists
  }

  bs.seq++
  bucket.Id = bs.seq
  bs.buckets[bucket.Id] = bucket
  return bucket.Id, nil
}

func (bs *InMemoryBlobStorage) Exists(bucket *domain.Bucket) bool {
  bs.RLock()
  defer bs.RUnlock()

  var res *domain.Bucket

  for _, v := range bs.buckets {
    if v.Name == bucket.Name {
      res = v
      break
    }
  }

  return res == nil
}


