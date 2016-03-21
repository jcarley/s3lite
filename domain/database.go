package domain

type Database interface {
	GetUploadByUploadId(uploadId string) *Upload
	CreateUpload(upload *Upload) (uint, error)
}
