package models

import (
	"bytes"
	"fmt"

	"github.com/jcarley/s3lite/app/lib/encoding"
)

type Upload struct {
	Id       uint
	Filename string
	Bucket   string
	Key      string
	UploadId string
}

func NewUpload() *Upload {
	return &Upload{}
}

func (u *Upload) GetNewUploadId() {
	var buffer bytes.Buffer
	token := "ABCDE123456"
	key := "1234567890"
	date := "20140215"
	rand := "58346548349467547569373"
	fmt.Fprintf(&buffer, "--{%s}--{%s}--{%s}--{%s}--", token, key, date, rand)
	hash := u.getHash(&buffer)
	u.UploadId = hash
}

func (u *Upload) AddPart(partNumber int, uploadId string, body []byte) string {
	// add part

	// generate etag
	etag := u.etag(partNumber, uploadId, "application/octet-stream", len(body))
	return etag
}

func (u *Upload) etag(partNumber int, uploadId string, mimeType string, fileFragmentSize int) string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "--{%i}--{%s}--{%s}--{%i}--", partNumber, uploadId, mimeType, fileFragmentSize)
	return u.getHash(&buffer)
}

func (u *Upload) getHash(buffer *bytes.Buffer) string {
	h := encoding.NewHasher("SHA1", buffer)
	hash := h.HashString()
	return hash
}
