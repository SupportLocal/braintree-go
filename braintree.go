package braintree

import (
	"bytes"
	"compress/gzip"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Environment string

const (
	sandbox    = "https://sandbox.braintreegateway.com"
	production = "https://www.braintreegateway.com"
)

func NewProduction(merchId, pubKey, privKey string) *Braintree {
	return newBraintree(production+"/merchants/"+merchId, merchId, pubKey, privKey)
}

func NewSandbox(merchId, pubKey, privKey string) *Braintree {
	return newBraintree(sandbox+"/merchants/"+merchId, merchId, pubKey, privKey)
}

func newBraintree(baseURL, merchId, pubKey, privKey string) *Braintree {
	return &Braintree{
		BaseURL:    baseURL,
		MerchantId: merchId,
		PublicKey:  pubKey,
		PrivateKey: privKey,
	}
}

type Braintree struct {
	BaseURL    string
	MerchantId string
	PublicKey  string
	PrivateKey string
	Logger     *log.Logger
}

func (g *Braintree) requestXML(method, path string, xmlIn, xmlOut interface{}) error {
	path = g.BaseURL + "/" + path

	var reqBuf bytes.Buffer

	if xmlIn != nil {
		xmlBody, err := xml.Marshal(xmlIn)
		if err != nil {
			return err
		}

		_, err = reqBuf.Write(xmlBody)
		if err != nil {
			return err
		}

	}

	req, err := http.NewRequest(method, path, &reqBuf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/xml")
	req.Header.Set("Accept", "application/xml")
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "Braintree Go 0.1.0")
	req.Header.Set("X-ApiVersion", "3")
	req.SetBasicAuth(g.PublicKey, g.PrivateKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	b, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	xmlBytes, err := ioutil.ReadAll(b)
	if err != nil {
		return err
	}

	if xmlOut != nil {
		if err := xml.Unmarshal(xmlBytes, xmlOut); err != nil {
			return err
		}
	}

	if resp.StatusCode > 299 {
		return fmt.Errorf("Status Code %d", resp.StatusCode)
	}

	return nil
}

func (g *Braintree) Transaction() *TransactionGateway {
	return &TransactionGateway{g}
}

func (g *Braintree) CreditCard() *CreditCardGateway {
	return &CreditCardGateway{g}
}

func (g *Braintree) Customer() *CustomerGateway {
	return &CustomerGateway{g}
}

func (g *Braintree) Discount() *Discount {
	return &Discount{g}
}

func (g *Braintree) Subscription() *SubscriptionGateway {
	return &SubscriptionGateway{g}
}

func (g *Braintree) Plan() *PlanGateway {
	return &PlanGateway{g}
}

func (g *Braintree) Address() *AddressGateway {
	return &AddressGateway{g}
}

func (g *Braintree) AddOn() *AddOnGateway {
	return &AddOnGateway{g}
}

func ParseDate(s string) (time.Time, error) {
	const fmt = "2006-01-02T15:04:05Z"
	return time.Parse(fmt, s)
}
