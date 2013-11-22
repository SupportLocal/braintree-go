package braintree

import (
	"encoding/xml"
	"testing"
)

func TestCreditCardErrors(t *testing.T) {
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

	var ces CreditCardErrors
	err := xml.Unmarshal(response, &ces)
	if err != nil {
		t.Fatal(err)
	}
	if len(ces) != 1 {
		t.Log(ces)
		t.Fatal(len(ces))
	}
	ce := ces[0]
	if ce.Message != "Customer ID is required." {
		t.Fatal(ce.Message)
	}
	if len(ce.Errors) != 1 {
		t.Log(ce)
		t.Fatal(len(ce.Errors))
	}
	e := ce.Errors[0]
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
