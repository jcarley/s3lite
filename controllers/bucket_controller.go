package controllers

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/services"
	"github.com/jcarley/s3lite/web"
)

type BucketController struct {
	service services.BucketServicer
}

func NewBucketController(service services.BucketServicer) *BucketController {
	return &BucketController{
		service: service,
	}
}

func (this *BucketController) Register(router *mux.Router) {
	web.Post("/buckets", router, this.CreateBucket)
	web.Delete("/buckets/{id}", router, this.DeleteBucket)
}

// Buckets
func (this *BucketController) CreateBucket(ctx context.Context, rw http.ResponseWriter, req *http.Request) {

	var bucket domain.Bucket
	err := decode(req, &bucket)
	if err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	}

	err = this.service.AddBucket(&bucket)
	if err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	}

	err = encode(rw, &bucket)
	if err != nil {
		httpError(err, http.StatusInternalServerError, rw)
		return
	}
}

func (this *BucketController) DeleteBucket(ctx context.Context, rw http.ResponseWriter, req *http.Request) {

	vars := web.VarsFromContext(ctx)
	id := vars["id"]

	err := this.service.DeleteBucketById(id)
	if err != nil {
		if err == services.RecordNotFoundError {
			httpError(err, http.StatusNotFound, rw)
		} else {
			httpError(err, http.StatusInternalServerError, rw)
		}
	}

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
