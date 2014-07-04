package infrastructure

import (
  "sync"
  "errors"

  "github.com/jcarley/s3lite/domain"
)

var (
  ErrUploadAlreadyExists = errors.New("upload already exists")
)

type InMemoryDatabase struct {
  sync.RWMutex
  uploads map[uint]*domain.Upload
  seq uint
}

func NewInMemoryDatabase() *InMemoryDatabase {
  return &InMemoryDatabase{
    uploads: make(map[uint]*domain.Upload),
  }
}

func (db *InMemoryDatabase) GetUploadByUploadId(uploadId string) *domain.Upload {
  db.RLock()
  defer db.RUnlock()

  var res *domain.Upload

  for _, v := range db.uploads {
    if v.UploadId == uploadId {
      res = v
      break
    }
  }

  return res
}

func (db *InMemoryDatabase) CreateUpload(upload *domain.Upload) (uint, error) {
  db.Lock()
  defer db.Unlock()

  if !db.isUnique(upload) {
    return 0, ErrUploadAlreadyExists
  }

  db.seq++
  upload.Id = db.seq
  db.uploads[upload.Id] = upload
  return upload.Id, nil
}

func (db *InMemoryDatabase) isUnique(upload *domain.Upload) bool {
  for _, v := range db.uploads {
    if v.UploadId == upload.UploadId {
      return false
    }
  }
  return true
}
