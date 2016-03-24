package controllers

import (
	"net/http"
)

type BucketController struct {
}

// Buckets
func (this *BucketController) CreateBucket(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) DeleteBucket(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) HeadBucket(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) ListBuckets(rw http.ResponseWriter, req *http.Request) {
}

// Objects
func (this *BucketController) CopyObject(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) DeleteObject(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) DeleteObjects(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) GetObject(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) HeadObject(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) ListObjects(rw http.ResponseWriter, req *http.Request) {
}

func (this *BucketController) PutObject(rw http.ResponseWriter, req *http.Request) {
}
