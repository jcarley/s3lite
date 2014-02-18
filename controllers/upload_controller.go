package controllers

import (
  "strconv"
  "strings"
  "regexp"
  "io/ioutil"
  "net/http"

  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/encoding"
  "github.com/jcarley/s3lite/webservice"
  "github.com/codegangsta/martini"
)

var rxFilename = regexp.MustCompile(`filename=(.*)$`)
var rxBucket = regexp.MustCompile(`^([\S\w\-]+)\.s3`)
var rxKey = regexp.MustCompile(`^.*\/([a-zA-z\/]*)$`)

type UploadController struct {}

func NewUploadController() *UploadController {
  return &UploadController{}
}

func (u *UploadController) InitiateMultipartUpload(req *http.Request, db domain.Database) (int, string) {

  upload := domain.NewUpload()
  upload.GetNewUploadId()
  upload.Filename = u.parseFilename(req)
  upload.Bucket = u.parseBucket(req)
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

func (u *UploadController) parseFilename(req *http.Request) string {

  matches := rxFilename.FindStringSubmatch(req.Header.Get("Content-Disposition"))

  filename := ""
  if len(matches) > 1 {
    filename = matches[1]
  }

  return filename
}

func (u *UploadController) parseBucket(req *http.Request) string {
  matches := rxBucket.FindStringSubmatch(req.Host)

  bucket := "default"
  if len(matches) > 1 {
    bucket = matches[1]
  }
  return bucket
}

func (u *UploadController) parseKey(req *http.Request) string {

  key := req.URL.Path

  if strings.Index(key, "/") == 0 {
    key = key[1:]
  }

  return key
}


