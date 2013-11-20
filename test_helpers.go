package braintree

import (
	"net/http"
	"net/http/httptest"
)

type testProxy struct {
	handlerFunc http.HandlerFunc
}

func (p testProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.handlerFunc(w, r)
}

func newServer(handlerFunc http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(testProxy{handlerFunc: handlerFunc})
}
