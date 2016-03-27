package domain

import (
	r "github.com/dancannon/gorethink"
)

type RethinkDatastore struct {
	session *r.Session
}

func NewRethinkDatastore(session *r.Session) *RethinkDatastore {
	return &RethinkDatastore{
		session: session,
	}
}

func (this *RethinkDatastore) GetUploadById(uploadId string) *Upload {
	return nil
}

func (this *RethinkDatastore) CreateUpload(upload *Upload) (string, error) {
	return "", nil
}

func (this *RethinkDatastore) GetBucketById(bucketId string) *Bucket {
	return nil
}

func (this *RethinkDatastore) CreateBucket(bucket *Bucket) (id string, err error) {
	return "", nil
}

func (this *RethinkDatastore) DeleteBucketById(bucketId string) error {
	return nil
}
