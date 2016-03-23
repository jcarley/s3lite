package testing

import (
	"errors"
	"reflect"

	"github.com/jcarley/s3lite/domain"
)

type MockUploadService struct {
	callChain map[string]struct {
		Count int
	}
	currentCallMethod string
	methodWatches     map[string]*MethodWatch
}

type MethodWatch struct {
	MethodName string
	ReturnArgs []interface{}
}

func NewMockUploadService() *MockUploadService {
	return &MockUploadService{
		callChain:     make(map[string]struct{ Count int }),
		methodWatches: make(map[string]*MethodWatch),
	}
}

func (this *MockUploadService) AddPart(partNumber int, uploadId string, body []byte) (etag string, err error) {
	this.addCallChainFunc("AddPart")

	if methodWatch, ok := this.methodWatches["AddPart"]; ok {

		getReturnArg(methodWatch, 0, &etag, "")
		getReturnArg(methodWatch, 1, &err, nil)

		return
	}

	return "", nil
}

func getReturnArg(methodWatch *MethodWatch, idx int, value interface{}, defaultValue interface{}) error {

	dataValue := reflect.ValueOf(value)
	argsValue := reflect.ValueOf(methodWatch.ReturnArgs[idx])

	if dataValue.Kind() != reflect.Ptr {
		return errors.New("result must be a pointer")
	}

	dataElem := dataValue.Elem()
	if !dataElem.CanAddr() {
		return errors.New("result must be addressable (a pointer)")
	}

	if !argsValue.IsValid() {
		dataValue.Elem().Set(reflect.Zero(dataValue.Elem().Type()))
		return nil
	}

	dataValue.Elem().Set(argsValue)

	return nil
}

func (this *MockUploadService) CreateUpload(filename string, bucket string, key string) (upload domain.Upload, err error) {
	this.addCallChainFunc("CreateUpload")
	upload = domain.NewUpload()
	err = nil
	return
}

func (this *MockUploadService) Called(methodName string) *MockUploadService {
	this.currentCallMethod = methodName
	return this
}

func (this *MockUploadService) Times(count int) bool {
	if value, ok := this.callChain[this.currentCallMethod]; !ok {
		return false
	} else {
		return value.Count == count
	}

	return false
}

func (this *MethodWatch) Return(args ...interface{}) {
	this.ReturnArgs = args
}

func (this *MockUploadService) On(methodName string) *MethodWatch {
	methodWatch := &MethodWatch{
		MethodName: methodName,
	}
	this.methodWatches[methodName] = methodWatch
	return methodWatch
}

func (this *MockUploadService) addCallChainFunc(methodName string) {
	if value, ok := this.callChain[methodName]; !ok {
		value = struct{ Count int }{1}
		this.callChain[methodName] = value
	} else {
		value.Count++
	}
}
