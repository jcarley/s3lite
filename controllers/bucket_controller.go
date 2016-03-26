package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/services"
	"github.com/jcarley/s3lite/web"
)

type BucketController struct {
	service services.BucketService
}

func NewBucketController(service services.BucketService) *BucketController {
	return &BucketController{
		service: service,
	}
}

func (this *BucketController) Register(router *mux.Router) {
	router.Handle("/buckets/{id}", web.NewContextAdapter(web.ContextHandlerFunc(this.DeleteBucket))).Methods("DELETE")
}

// Buckets
func (this *BucketController) CreateBucket(rw http.ResponseWriter, req *http.Request) {

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
	fmt.Printf("%+v\n", req)

	vars := web.VarsFromContext(ctx)
	id := vars["id"]

	fmt.Println("==========================")
	fmt.Println(id)
	fmt.Println("==========================")
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
