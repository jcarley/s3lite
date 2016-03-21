package testing

import (
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

		// getReturnArg(methodWatch, 0, &etag, "")

		x := ""
		v := reflect.ValueOf(x)
		v.SetString(methodWatch.ReturnArgs[0].(string))

		// fmt.Printf("%#v\n", etag)

		err = nil

		// getReturnArg(methodWatch, 1, &err, nil)
		// switch methodWatch.ReturnArgs[0].(type) {
		// case string:
		// etag = methodWatch.ReturnArgs[0].(string)
		// default:
		// etag = ""
		// }

		// switch methodWatch.ReturnArgs[1].(type) {
		// case error:
		// err = methodWatch.ReturnArgs[1].(error)
		// default:
		// err = nil
		// }

		return
	}

	return "", nil
}

func getReturnArg(methodWatch *MethodWatch, idx int, value interface{}, defaultValue interface{}) {

	v := reflect.ValueOf(value)
	x := reflect.ValueOf(methodWatch.ReturnArgs[idx])
	v.Set(x)
	return

	// switch methodWatch.ReturnArgs[idx].(type) {
	// case string:
	// fmt.Println("************************** HERE")
	// newValue := methodWatch.ReturnArgs[idx].(string)
	// value = &newValue
	// case error:
	// value = methodWatch.ReturnArgs[idx].(error)
	// default:
	// value = defaultValue
	// }
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
