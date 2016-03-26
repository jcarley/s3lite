package test

import "github.com/jcarley/s3lite/domain"

type MockUploadService struct {
	MockService
}

func NewMockUploadService() *MockUploadService {
	return &MockUploadService{
		MockService{
			callChain:     make(map[string]struct{ Count int }),
			methodWatches: make(map[string]*MethodWatch),
		}}
}

func (this *MockUploadService) AddPart(partNumber int, uploadId string, body []byte) (etag string, err error) {
	this.AddCallChainFunc("AddPart")

	if methodWatch, ok := this.MethodWatches()["AddPart"]; ok {

		this.GetReturnArg(methodWatch, 0, &etag, "")
		this.GetReturnArg(methodWatch, 1, &err, nil)

		return
	}

	return "", nil
}

func (this *MockUploadService) CreateUpload(filename string, bucket string, key string) (upload domain.Upload, err error) {
	this.AddCallChainFunc("CreateUpload")
	upload = domain.NewUpload()
	err = nil
	return
}
