package braintree

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

func xmlRequest(gateway Braintree, method, url string, xmlObj interface{}) (*http.Request, error) {
	xmlBody, err := xml.Marshal(xmlObj)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = buf.Write(xmlBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "Braintree Go 0.1.0")
	req.Header.Set("X-ApiVersion", "3")
	req.SetBasicAuth(gateway.PublicKey, gateway.PrivateKey)

	return req, nil
}
