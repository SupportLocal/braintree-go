package braintree

import (
	"net/http"
	"testing"
)

func TestCustomerCreateWithCVVError(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<api-error-response>
  <errors>
    <errors type="array"/>
    <credit-card>
      <errors type="array"/>
    </credit-card>
  </errors>
  <params>
    <customer>
      <first-name>Lionel</first-name>
      <last-name>Barrow</last-name>
      <company>Braintree</company>
      <email>lionel.barrow@example.com</email>
      <phone>312.555.1234</phone>
      <fax>614.555.5678</fax>
      <website>http://www.example.com</website>
      <credit-card>
        <expiration-date>05/14</expiration-date>
        <options>
          <verify-card>true</verify-card>
        </options>
        <device-session-id nil="true"/>
        <fraud-merchant-id nil="true"/>
      </credit-card>
    </customer>
  </params>
  <message>Gateway Rejected: cvv</message>
  <verification>
    <status>gateway_rejected</status>
    <cvv-response-code>N</cvv-response-code>
    <avs-error-response-code nil="true"/>
    <avs-postal-code-response-code>I</avs-postal-code-response-code>
    <avs-street-address-response-code>I</avs-street-address-response-code>
    <gateway-rejection-reason>cvv</gateway-rejection-reason>
    <merchant-account-id>foo</merchant-account-id>
    <processor-response-code>1000</processor-response-code>
    <processor-response-text>Approved</processor-response-text>
    <id>8jrkmb</id>
    <billing>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <street-address nil="true"/>
      <extended-address nil="true"/>
      <locality nil="true"/>
      <region nil="true"/>
      <postal-code nil="true"/>
      <country-name nil="true"/>
    </billing>
    <credit-card>
      <token nil="true"/>
      <bin>411111</bin>
      <last-4>1111</last-4>
      <card-type>Visa</card-type>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <customer-location>US</customer-location>
      <cardholder-name nil="true"/>
      <prepaid>Unknown</prepaid>
      <healthcare>Unknown</healthcare>
      <debit>Unknown</debit>
      <durbin-regulated>Unknown</durbin-regulated>
      <commercial>Unknown</commercial>
      <payroll>Unknown</payroll>
      <issuing-bank>Unknown</issuing-bank>
      <country-of-issuance>Unknown</country-of-issuance>
    </credit-card>
    <created-at type="datetime">2013-11-20T23:05:23Z</created-at>
    <updated-at type="datetime">2013-11-20T23:05:23Z</updated-at>
  </verification>
</api-error-response>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	oc := &Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "200",
			Options: &CreditCardOptions{
				VerifyCard: true,
			},
		},
	}

	// Create with errors
	_, err := gw.Customer().Create(oc)
	if err == nil {
		t.Fatal("Did not receive error when creating invalid customer")
	}
}

func TestCustomerCreate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<customer>
  <id>35182871</id>
  <merchant-id>foo</merchant-id>
  <first-name>Lionel</first-name>
  <last-name>Barrow</last-name>
  <company>Braintree</company>
  <email>lionel.barrow@example.com</email>
  <phone>312.555.1234</phone>
  <fax>614.555.5678</fax>
  <website>http://www.example.com</website>
  <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
  <custom-fields>
  </custom-fields>
  <credit-cards type="array">
    <credit-card>
      <bin>411111</bin>
      <card-type>Visa</card-type>
      <cardholder-name nil="true"/>
      <commercial>Unknown</commercial>
      <country-of-issuance>Unknown</country-of-issuance>
      <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
      <customer-id>35182871</customer-id>
      <customer-location>US</customer-location>
      <debit>Unknown</debit>
      <default type="boolean">true</default>
      <durbin-regulated>Unknown</durbin-regulated>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <expired type="boolean">false</expired>
      <healthcare>Unknown</healthcare>
      <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
      <issuing-bank>Unknown</issuing-bank>
      <last-4>1111</last-4>
      <payroll>Unknown</payroll>
      <prepaid>Unknown</prepaid>
      <subscriptions type="array"/>
      <token>gz27pb</token>
      <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
      <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
      <venmo-sdk type="boolean">false</venmo-sdk>
      <verifications type="array"/>
    </credit-card>
  </credit-cards>
  <addresses type="array"/>
</customer>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	oc := &Customer{
		FirstName: "Lionel",
		LastName:  "Barrow",
		Company:   "Braintree",
		Email:     "lionel.barrow@example.com",
		Phone:     "312.555.1234",
		Fax:       "614.555.5678",
		Website:   "http://www.example.com",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "",
		},
	}

	customer, err := gw.Customer().Create(oc)

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}
	if customer.Id == "" {
		t.Fatal("invalid customer id")
	}
	if card := customer.DefaultCreditCard(); card == nil {
		t.Fatal("invalid credit card")
	}
	if card := customer.DefaultCreditCard(); card.Token == "" {
		t.Fatal("invalid token")
	}

}

func TestCustomerUpdate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<customer>
  <id>35182871</id>
  <merchant-id>foo</merchant-id>
  <first-name>John</first-name>
  <last-name>Barrow</last-name>
  <company>Braintree</company>
  <email>lionel.barrow@example.com</email>
  <phone>312.555.1234</phone>
  <fax>614.555.5678</fax>
  <website>http://www.example.com</website>
  <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
  <custom-fields>
  </custom-fields>
  <credit-cards type="array">
    <credit-card>
      <bin>411111</bin>
      <card-type>Visa</card-type>
      <cardholder-name nil="true"/>
      <commercial>Unknown</commercial>
      <country-of-issuance>Unknown</country-of-issuance>
      <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
      <customer-id>35182871</customer-id>
      <customer-location>US</customer-location>
      <debit>Unknown</debit>
      <default type="boolean">true</default>
      <durbin-regulated>Unknown</durbin-regulated>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <expired type="boolean">false</expired>
      <healthcare>Unknown</healthcare>
      <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
      <issuing-bank>Unknown</issuing-bank>
      <last-4>1111</last-4>
      <payroll>Unknown</payroll>
      <prepaid>Unknown</prepaid>
      <subscriptions type="array"/>
      <token>gz27pb</token>
      <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
      <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
      <venmo-sdk type="boolean">false</venmo-sdk>
      <verifications type="array"/>
    </credit-card>
  </credit-cards>
  <addresses type="array"/>
</customer>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c2, err := gw.Customer().Update(&Customer{
		Id:        "35182871",
		FirstName: "John",
	})

	t.Log(c2)

	if err != nil {
		t.Fatal(err)
	}
	if c2.FirstName != "John" {
		t.Fatal("first name not changed")
	}
}

func TestCustomerFind(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<customer>
  <id>35182871</id>
  <merchant-id>foo</merchant-id>
  <first-name>John</first-name>
  <last-name>Barrow</last-name>
  <company>Braintree</company>
  <email>lionel.barrow@example.com</email>
  <phone>312.555.1234</phone>
  <fax>614.555.5678</fax>
  <website>http://www.example.com</website>
  <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
  <custom-fields>
  </custom-fields>
  <credit-cards type="array">
    <credit-card>
      <bin>411111</bin>
      <card-type>Visa</card-type>
      <cardholder-name nil="true"/>
      <commercial>Unknown</commercial>
      <country-of-issuance>Unknown</country-of-issuance>
      <created-at type="datetime">2013-11-20T23:05:24Z</created-at>
      <customer-id>35182871</customer-id>
      <customer-location>US</customer-location>
      <debit>Unknown</debit>
      <default type="boolean">true</default>
      <durbin-regulated>Unknown</durbin-regulated>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <expired type="boolean">false</expired>
      <healthcare>Unknown</healthcare>
      <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
      <issuing-bank>Unknown</issuing-bank>
      <last-4>1111</last-4>
      <payroll>Unknown</payroll>
      <prepaid>Unknown</prepaid>
      <subscriptions type="array"/>
      <token>gz27pb</token>
      <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
      <updated-at type="datetime">2013-11-20T23:05:24Z</updated-at>
      <venmo-sdk type="boolean">false</venmo-sdk>
      <verifications type="array"/>
    </credit-card>
  </credit-cards>
  <addresses type="array"/>
</customer>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c3, err := gw.Customer().Find("35182871")

	t.Log(c3)

	if err != nil {
		t.Fatal(err)
	}
	if c3.Id != "35182871" {
		t.Fatal("ids do not match")
	}
}

func TestCustomerDelete(t *testing.T) {
	var response = []byte(``)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	err := gw.Customer().Delete("35182871")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCustomerFind404(t *testing.T) {
	var response = []byte(``)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c4, err := gw.Customer().Find("35182871")
	if err == nil {
		t.Fatal("should return 404")
	}
	if err.Error() != "Not Found (404)" {
		t.Fatal(err)
	}
	if c4 != nil {
		t.Fatal(c4)
	}
}
