package domain

import (
  "sync"
  "errors"
)

var (
  ErrAlreadyExists = errors.New("upload already exists")
)

type InMemoryDatabase struct {
  sync.RWMutex
  uploads map[uint]*Upload
  seq uint
}

func NewInMemoryDatabase() *InMemoryDatabase {
  return &InMemoryDatabase{
    uploads: make(map[uint]*Upload),
  }
}

func (db *InMemoryDatabase) GetUploadByUploadId(uploadId string) *Upload {
  db.RLock()
  defer db.RUnlock()

  var res *Upload

  for _, v := range db.uploads {
    if v.UploadId == uploadId {
      res = v
      break
    }
  }

  return res
}

func (db *InMemoryDatabase) CreateUpload(upload *Upload) (uint, error) {
  db.Lock()
  defer db.Unlock()

  if !db.isUnique(upload) {
    return 0, ErrAlreadyExists
  }

  db.seq++
  upload.Id = db.seq
  db.uploads[upload.Id] = upload
  return upload.Id, nil
}

func (db *InMemoryDatabase) isUnique(upload *Upload) bool {
  for _, v := range db.uploads {
    if v.UploadId == upload.UploadId {
      return false
    }
  }
  return true
}
