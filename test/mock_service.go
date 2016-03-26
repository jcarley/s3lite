package test

import (
	"errors"
	"reflect"
)

type MockService struct {
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

func (this *MockService) AddCallChainFunc(methodName string) {
	if value, ok := this.callChain[methodName]; !ok {
		value = struct{ Count int }{1}
		this.callChain[methodName] = value
	} else {
		value.Count++
	}
}

func (this *MockService) CurrentCallMethod() string {
	return this.currentCallMethod
}

func (this *MockService) SetCurrentCallMethod(methodName string) {
	this.currentCallMethod = methodName
}

func (this *MockService) MethodWatches() map[string]*MethodWatch {
	return this.methodWatches
}

func (this *MockService) GetReturnArg(methodWatch *MethodWatch, idx int, data interface{}, defaultValue interface{}) {

	dataValue := reflect.ValueOf(data)
	argsValue := reflect.ValueOf(methodWatch.ReturnArgs[idx])

	if dataValue.Kind() != reflect.Ptr {
		panic(errors.New("result must be a pointer"))
	}

	dataElem := dataValue.Elem()
	if !dataElem.CanAddr() {
		panic(errors.New("result must be addressable (a pointer)"))
	}

	if !argsValue.IsValid() {
		dataElem.Set(reflect.Zero(dataElem.Type()))
		return
	}

	dataElem.Set(argsValue)
}

func (this *MockService) Called(methodName string) *MockService {
	this.SetCurrentCallMethod(methodName)
	return this
}

func (this *MockService) Times(count int) bool {
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

func (this *MockService) On(methodName string) *MethodWatch {
	methodWatch := &MethodWatch{
		MethodName: methodName,
	}
	this.MethodWatches()[methodName] = methodWatch
	return methodWatch
}
