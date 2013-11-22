package braintree

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"
)

func doRequest(r *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(r)
}

func unzipRequest(resp *http.Response) ([]byte, error) {
	b, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
