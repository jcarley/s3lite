package test

type MockBucketService struct {
	callChain map[string]struct {
		Count int
	}
	currentCallMethod string
	methodWatches     map[string]*MethodWatch
}

func NewMockBucketService() *MockBucketService {
	return &MockBucketService{
		callChain:     make(map[string]struct{ Count int }),
		methodWatches: make(map[string]*MethodWatch),
	}
}
