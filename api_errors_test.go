package braintree

import (
	"encoding/xml"
	"testing"
)

func TestApiErrors(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<api-error-response>
  <errors>
    <errors type="array"/>
    <credit-card>
      <errors type="array">
        <error>
          <code>91704</code>
          <attribute type="symbol">customer_id</attribute>
          <message>Customer ID is required.</message>
        </error>
      </errors>
    </credit-card>
  </errors>
  <params>
    <credit-card>
      <expiration-date>05/14</expiration-date>
    </credit-card>
    <action>create</action>
    <controller>payment_methods</controller>
    <merchant-id>foo</merchant-id>
  </params>
  <message>Customer ID is required.</message>
</api-error-response>`)

	var aes struct {
		ApiErrors
	}

	err := xml.Unmarshal(response, &aes)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(aes)
	if aes.ErrorMessage != "Customer ID is required." {
		t.Fatal(aes.ErrorMessage)
	}
	if aes.ErrorCount() != 1 {
		t.Fatal(aes.ErrorCount())
	}
	if aes.CreditCardErrors == nil {
		t.Fatal(aes.CreditCardErrors)
	}

	if aes.CreditCardErrors.Count() != 1 {
		t.Fatal(aes.CreditCardErrors.Count())
	}

	cce := *aes.CreditCardErrors
	e := cce[0]
	if e.Message != "Customer ID is required." {
		t.Fatal(e.Message)
	}
	if e.Code != "91704" {
		t.Fatal(e.Code)
	}
	if e.Attribute != "customer_id" {
		t.Fatal(e.Attribute)
	}
	if e.Message != "Customer ID is required." {
		t.Fatal(e.Message)
	}

}
