package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jcarley/s3lite/domain"
	// "github.com/jcarley/s3lite/domain"
	// "github.com/jcarley/s3lite/encoding"
	// "github.com/jcarley/s3lite/webservice"
)

var MissingUploadIdError = errors.New("Request is missing an upload id")
var MissingPartNumberError = errors.New("Request is missing a part number")
var MissingAuthorizationKeyError = errors.New("Request is missing an authorization key")
var MissingContentDispostionError = errors.New("Request is missing a content-dispostion")

var rxFilename = regexp.MustCompile(`filename=(.*)$`)
var rxBucket = regexp.MustCompile(`^([\S\w\-]+)\.s3`)
var rxKey = regexp.MustCompile(`^.*\/([a-zA-z\/]*)$`)

type ErrMessage struct {
	Msg string `json:"error"`
}

type InitiateMultipartUploadResult struct {
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	UploadId string `json:"upload_id"`
}

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

func (u *UploadController) InitiateMultipartUpload(rw http.ResponseWriter, req *http.Request) {

	upload := domain.NewUpload()
	upload.GetNewUploadId()
	upload.Filename = u.parseFilename(req)
	upload.Bucket = u.parseBucket(req)
	upload.Key = u.parseKey(req)

	// TODO: Need to save upload record to database
	// u.db.CreateUpload(upload)

	result := InitiateMultipartUploadResult{UploadId: upload.UploadId, Bucket: upload.Bucket, Key: upload.Key}
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(result)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (u *UploadController) UploadPart(rw http.ResponseWriter, req *http.Request) {

	uploadId := getHeaderValue("Upload-Id", req)
	if uploadId == "" {
		httpError(MissingUploadIdError, http.StatusBadRequest, rw)
		return
	}

	partNumber := getHeaderValue("Part-Number", req)
	if partNumber == "" {
		httpError(MissingPartNumberError, http.StatusBadRequest, rw)
		return
	}

	key := getHeaderValue("Authorization", req)
	if key == "" {
		httpError(MissingAuthorizationKeyError, http.StatusBadRequest, rw)
		return
	}

	filename := u.parseFilename(req)
	if filename == "" {
		httpError(MissingContentDispostionError, http.StatusBadRequest, rw)
		return
	}

	// partNumber, err := strconv.Atoi(params["partNumber"])
	// if err != nil {
	// panic(err)
	// }
	// uploadId := params["uploadId"]
	// body, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// panic(err)
	// }

	// upload := domain.NewUpload()
	// etag := upload.AddPart(partNumber, uploadId, body)

}

func (u *UploadController) parseFilename(req *http.Request) string {

	headerValue := getHeaderValue("Content-Disposition", req)
	matches := rxFilename.FindStringSubmatch(headerValue)

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

func getHeaderValue(name string, req *http.Request) string {
	headerValue := req.Header[name]
	if len(headerValue) > 0 {
		return headerValue[0]
	}
	return ""
}

func httpError(err error, status int, rw http.ResponseWriter) {
	errorMessage := ErrMessage{err.Error()}
	result, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	} else {
		http.Error(rw, string(result), status)
	}
}
