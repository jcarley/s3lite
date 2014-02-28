package controllers

import (
  "io/ioutil"
  "net/http"
  "regexp"
  "strconv"
  "strings"

  "github.com/codegangsta/martini"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/encoding"
  "github.com/jcarley/s3lite/webservice"
)

type UploadController struct{}

func (u *UploadController) InitiateMultipartUpload(req *http.Request, db domain.Database, blobStorage domain.BlobStorage) (int, string) {

  bucketName := u.parseBucket(req)
  bucket := domain.NewBucket(bucketName)

  // if !blobStorage.Exists(bucket) {
  // blobStorage.Create(bucket)
  // }

  upload := domain.NewUpload()
  upload.GetNewUploadId()
  upload.Filename = u.parseFilename(req)
  upload.Bucket = bucket.Name
  upload.Key = u.parseKey(req)

  db.CreateUpload(upload)

  result := webservice.InitiateMultipartUploadResult{UploadId: upload.UploadId, Bucket: upload.Bucket, Key: upload.Key}

  xmlEncoder := encoding.XmlEncoder{}
  xml, err := xmlEncoder.Encode(&result)
  if err != nil {
    panic(err)
  }
  return 201, xml
}

func (u *UploadController) UploadPart(params martini.Params, res http.ResponseWriter, req *http.Request) (int, string) {
  partNumber, err := strconv.Atoi(params["partNumber"])
  if err != nil {
    panic(err)
  }
  uploadId := params["uploadId"]
  body, err := ioutil.ReadAll(req.Body)
  if err != nil {
    panic(err)
  }

  upload := domain.NewUpload()
  etag := upload.AddPart(partNumber, uploadId, body)

  return 200, etag
}
