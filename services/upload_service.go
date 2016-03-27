package services

import (
	"github.com/jcarley/s3lite/domain"
)

type UploadServicer interface {
	AddPart(partNumber int, uploadId string, body []byte) (etag string, err error)
	CreateUpload(filename string, bucket string, key string) (upload *domain.Upload, err error)
}

type UploadService struct {
	datastore domain.UploadDatastore
}

func NewUploadService(datastore domain.UploadDatastore) UploadServicer {
	return &UploadService{
		datastore: datastore,
	}
}

func (this *UploadService) AddPart(partNumber int, uploadId string, body []byte) (etag string, err error) {
	return "", nil
}

func (this *UploadService) CreateUpload(filename string, bucket string, key string) (upload *domain.Upload, err error) {
	return nil, nil
}

// func (this *UploadService) AddPart(partNumber int, uploadId string, body []byte) (etag string, err error) {
// return "", nil
// }

// func (this *UploadService) CreateUpload(filename string, bucket string, key string) (upload domain.Upload, err error) {
// upload = domain.NewUpload()
// err = nil
// return
// }
