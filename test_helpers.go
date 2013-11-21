package braintree

import (
	"bytes"
	"compress/gzip"
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

func writeZip(w http.ResponseWriter, data []byte) {
	var b bytes.Buffer
	gw := gzip.NewWriter(&b)
	gw.Write(data)
	gw.Close()
	w.Write(b.Bytes())
}
