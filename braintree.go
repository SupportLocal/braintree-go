package braintree

import (
	"bytes"
	"encoding/xml"
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
	return newBraintree(production, merchId, pubKey, privKey)
}

func NewSandbox(merchId, pubKey, privKey string) *Braintree {
	return newBraintree(sandbox, merchId, pubKey, privKey)
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

func (g *Braintree) MerchantURL() string {
	return g.BaseURL + "/merchants/" + g.MerchantId
}

func (g *Braintree) execute(method, path string, xmlObj interface{}) (*Response, error) {
	var buf bytes.Buffer
	if xmlObj != nil {
		xmlBody, err := xml.Marshal(xmlObj)
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(xmlBody)
		if err != nil {
			return nil, err
		}
	}

	url := g.MerchantURL() + "/" + path

	if g.Logger != nil {
		g.Logger.Printf("> %s %s\n%s", method, url, buf.String())
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
	req.SetBasicAuth(g.PublicKey, g.PrivateKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	btr := &Response{
		Response: resp,
	}
	err = btr.unpackBody()
	if err != nil {
		return nil, err
	}

	if g.Logger != nil {
		g.Logger.Printf("<\n%s", string(btr.Body))
	}

	err = btr.apiError()
	if err != nil {
		return nil, err
	}
	return btr, nil
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
