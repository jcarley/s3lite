package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/jcarley/s3lite/services"
	// "github.com/jcarley/s3lite/domain"
	// "github.com/jcarley/s3lite/encoding"
	// "github.com/jcarley/s3lite/webservice"
)

var MissingUploadIdError = errors.New("Request is missing an upload id")
var InvalidPartNumberError = errors.New("The part number is missing or invalid")
var MissingAuthorizationKeyError = errors.New("Request is missing an authorization key")
var MissingContentDispostionError = errors.New("Request is missing a content-dispostion")
var MissingContentBodyError = errors.New("Request does not have a body")

var rxFilename = regexp.MustCompile(`filename=(.*)$`)
var rxBucket = regexp.MustCompile(`^([\S\w\-]+)\.s3`)
var rxKey = regexp.MustCompile(`^.*\/([a-zA-z\/]*)$`)

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type InitiateMultipartUploadResult struct {
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	UploadId string `json:"upload_id"`
}

type UploadController struct {
	service services.UploadService
}

func NewUploadController(service services.UploadService) *UploadController {
	return &UploadController{
		service,
	}
}

func (this *UploadController) InitiateMultipartUpload(rw http.ResponseWriter, req *http.Request) {

	filename := this.parseFilename(req)
	bucket := this.parseBucket(req)
	key := this.parseKey(req)

	upload, err := this.service.CreateUpload(filename, bucket, key)
	if err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	}

	result := InitiateMultipartUploadResult{UploadId: upload.UploadId, Bucket: upload.Bucket, Key: upload.Key}
	encoder := json.NewEncoder(rw)
	err = encoder.Encode(result)
	if err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	}
}

func (this *UploadController) UploadPart(rw http.ResponseWriter, req *http.Request) {

	uploadId := getHeaderValue("Upload-Id", req)
	if uploadId == "" {
		httpError(MissingUploadIdError, http.StatusBadRequest, rw)
		return
	}

	partNumber, err := strconv.Atoi(getHeaderValue("Part-Number", req))
	if err != nil {
		httpError(InvalidPartNumberError, http.StatusBadRequest, rw)
		return
	}

	key := getHeaderValue("Authorization", req)
	if key == "" {
		httpError(MissingAuthorizationKeyError, http.StatusBadRequest, rw)
		return
	}

	filename := this.parseFilename(req)
	if filename == "" {
		httpError(MissingContentDispostionError, http.StatusBadRequest, rw)
		return
	}

	if req.Body == nil {
		httpError(MissingContentBodyError, http.StatusBadRequest, rw)
		return
	}
	var body []byte
	body, err = ioutil.ReadAll(req.Body)
	if len(body) == 0 {
		httpError(MissingContentBodyError, http.StatusBadRequest, rw)
		return
	}
	if err != nil {
		httpError(err, http.StatusBadRequest, rw)
		return
	}

	if etag, err := this.service.AddPart(partNumber, uploadId, body); err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	} else {
		encoder := json.NewEncoder(rw)
		successMessage := Message{"success", etag}
		err = encoder.Encode(successMessage)
		if err != nil {
			httpError(err, http.StatusInternalServerError, rw)
			return
		}
	}

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
	errorMessage := Message{"error", err.Error()}

	result, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
	} else {
		http.Error(rw, string(result), status)
	}
}
