package braintree

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"
)

type testProxy struct {
	handlerFunc http.HandlerFunc
}

var testGateway = newBraintree(
	sandbox,
	os.Getenv("BRAINTREE_MERCH_ID"),
	os.Getenv("BRAINTREE_PUB_KEY"),
	os.Getenv("BRAINTREE_PRIV_KEY"),
)

func (p testProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.handlerFunc(w, r)
}

func newServer(handlerFunc http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(testProxy{handlerFunc: handlerFunc})
}

func serveRecording(w http.ResponseWriter, r *http.Request, directory, filename string, httpStatusCode int) error {
	// Check to see if we already have this recorded
	d, err := os.Getwd()
	if err != nil {
		return err
	}
	workingDir := filepath.Join(d, "test_recordings", directory)
	err = os.MkdirAll(workingDir, 0777)
	if err != nil {
		return err
	}
	var file []byte
	file, _ = ioutil.ReadFile(filepath.Join(workingDir, filename+"_response.xml"))

	if len(file) == 0 { // no file found, so time to record it
		// Keep a log of what we are doing and write it when we are done.
		log.Printf("No recorded file for %s. Retrieiving response from sandbox.", filepath.Join(directory, filename+"_response.xml"))

		var fileLog bytes.Buffer

		sandboxUrl := sandbox + "/merchants/" + os.Getenv("BRAINTREE_MERCH_ID") + r.URL.Path

		reqBytes, err := ioutil.ReadAll(r.Body)
		var reqBuf bytes.Buffer
		reqBuf.Write(reqBytes)

		fileLog.WriteString(fmt.Sprintf("Recorded at:...........%s\n", time.Now()))
		fileLog.WriteString(fmt.Sprintf("Request url:...........%s\n", sandboxUrl))
		fileLog.WriteString(fmt.Sprintf("Request Method:........%s\n", r.Method))
		fileLog.WriteString(fmt.Sprintf("Braintree Merchant Id:.%s\n", os.Getenv("BRAINTREE_MERCH_ID")))
		fileLog.WriteString(fmt.Sprintf("Braintree Private Key:.%s\n", os.Getenv("BRAINTREE_PUB_KEY")))
		fileLog.WriteString(fmt.Sprintf("Braintree Public Key:..%s\n", os.Getenv("BRAINTREE_PRIV_KEY")))
		fileLog.WriteString(fmt.Sprintf("Request Body:  \n\n%s\n\n", reqBuf.Bytes()))

		if err != nil {
			return err
		}

		req, err := http.NewRequest(r.Method, sandboxUrl, &reqBuf)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/xml")
		req.Header.Set("Accept", "application/xml")
		req.Header.Set("Accept-Encoding", "gzip")
		req.Header.Set("User-Agent", "Braintree Go 0.1.0")
		req.Header.Set("X-ApiVersion", "3")
		req.SetBasicAuth(os.Getenv("BRAINTREE_PUB_KEY"), os.Getenv("BRAINTREE_PRIV_KEY"))

		ioutil.WriteFile(filepath.Join(workingDir, filename+"_request.txt"), fileLog.Bytes(), 0666)
		resp, err := http.DefaultClient.Do(req)
		httpStatusCode = resp.StatusCode
		if err != nil {
			return err
		}
		b, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		respBytes, err := ioutil.ReadAll(b)
		file = respBytes
		ioutil.WriteFile(filepath.Join(workingDir, filename+"_response.xml"), file, 0666)
	}
	w.WriteHeader(httpStatusCode)
	writeZip(w, file)

	return nil
}

func writeZip(w http.ResponseWriter, data []byte) {
	var b bytes.Buffer
	gw := gzip.NewWriter(&b)
	gw.Write(data)
	gw.Close()
	w.Write(b.Bytes())
}
