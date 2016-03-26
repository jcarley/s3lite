package web

import (
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type key int

const (
	RequestVarsKey key = 0
)

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (h ContextHandlerFunc) ServeHTTPContext(ctx context.Context, rw http.ResponseWriter, req *http.Request) {
	h(ctx, rw, req)
}

type ContextAdapter struct {
	ctx     context.Context
	handler ContextHandler
}

func NewContextAdapter(contextHandlerFunc ContextHandlerFunc) *ContextAdapter {
	return &ContextAdapter{context.Background(), contextHandlerFunc}
}

func (ca *ContextAdapter) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	context := NewContextWithRequestVars(ca.ctx, req)
	ca.handler.ServeHTTPContext(context, rw, req)
}

func NewContextWithRequestVars(ctx context.Context, req *http.Request) context.Context {
	vars := mux.Vars(req)
	return context.WithValue(ctx, RequestVarsKey, vars)
}

func VarsFromContext(ctx context.Context) map[string]string {
	return ctx.Value(RequestVarsKey).(map[string]string)
}
