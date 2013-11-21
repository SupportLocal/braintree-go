package braintree

import (
	"net/http"
	"testing"
)

func TestCreditCardCreate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <bin>411111</bin>
  <card-type>Visa</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:05Z</created-at>
  <customer-id>19976549</customer-id>
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
  <token>32nxpb</token>
  <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:05Z</updated-at>
  <venmo-sdk type="boolean">false</venmo-sdk>
  <verifications type="array">
    <verification>
      <status>verified</status>
      <cvv-response-code>M</cvv-response-code>
      <avs-error-response-code nil="true"/>
      <avs-postal-code-response-code>I</avs-postal-code-response-code>
      <avs-street-address-response-code>I</avs-street-address-response-code>
      <gateway-rejection-reason nil="true"/>
      <merchant-account-id>qrtm5hqfdgg7xkjs</merchant-account-id>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <id>jgj8mw</id>
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
        <token>32nxpb</token>
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
      <created-at type="datetime">2013-11-20T23:05:05Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:05Z</updated-at>
    </verification>
  </verifications>
</credit-card> `)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()
	card, err := g.Create(&CreditCard{
		CustomerId:     "19976549",
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card)

	if card.Token == "" {
		t.Fatal("invalid token")
	}

}

func TestCreditCardUpdate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <bin>555555</bin>
  <card-type>MasterCard</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:05Z</created-at>
  <customer-id>19976549</customer-id>
  <customer-location>US</customer-location>
  <debit>Unknown</debit>
  <default type="boolean">true</default>
  <durbin-regulated>Unknown</durbin-regulated>
  <expiration-month>05</expiration-month>
  <expiration-year>2014</expiration-year>
  <expired type="boolean">false</expired>
  <healthcare>Unknown</healthcare>
  <image-url>https://assets.braintreegateway.com/payment_method_logo/mastercard.png?environment=sandbox&amp;merchant_id=foo</image-url>
  <issuing-bank>Unknown</issuing-bank>
  <last-4>4444</last-4>
  <payroll>Unknown</payroll>
  <prepaid>Unknown</prepaid>
  <subscriptions type="array"/>
  <token>32nxpb</token>
  <unique-number-identifier>1bae4a11d2481f7150fc545a264f45b8</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:07Z</updated-at>
  <venmo-sdk type="boolean">false</venmo-sdk>
  <verifications type="array">
    <verification>
      <status>verified</status>
      <cvv-response-code>M</cvv-response-code>
      <avs-error-response-code nil="true"/>
      <avs-postal-code-response-code>I</avs-postal-code-response-code>
      <avs-street-address-response-code>I</avs-street-address-response-code>
      <gateway-rejection-reason nil="true"/>
      <merchant-account-id>qrtm5hqfdgg7xkjs</merchant-account-id>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <id>hh7z4m</id>
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
        <token>32nxpb</token>
        <bin>555555</bin>
        <last-4>4444</last-4>
        <card-type>MasterCard</card-type>
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
      <created-at type="datetime">2013-11-20T23:05:06Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:07Z</updated-at>
    </verification>
    <verification>
      <status>verified</status>
      <cvv-response-code>M</cvv-response-code>
      <avs-error-response-code nil="true"/>
      <avs-postal-code-response-code>I</avs-postal-code-response-code>
      <avs-street-address-response-code>I</avs-street-address-response-code>
      <gateway-rejection-reason nil="true"/>
      <merchant-account-id>qrtm5hqfdgg7xkjs</merchant-account-id>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <id>jgj8mw</id>
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
        <token>32nxpb</token>
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
      <created-at type="datetime">2013-11-20T23:05:05Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:05Z</updated-at>
    </verification>
  </verifications>
</credit-card>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card2, err := g.Update(&CreditCard{
		Token:          "32nxpb",
		Number:         testCreditCards["mastercard"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card2)

	if card2.Token != "32nxpb" {
		t.Fatal()
	}
	if card2.CardType != "MasterCard" {
		t.Fatal()
	}

}

func TestCreditCardDelete(t *testing.T) {
	var response = []byte(``)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()
	card2 := CreditCard{
		Token:          "32nxpb",
		Number:         testCreditCards["mastercard"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}

	err := g.Delete(&card2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateCreditCardWithExpirationMonthAndYear(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <bin>411111</bin>
  <card-type>Visa</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:08Z</created-at>
  <customer-id>57131083</customer-id>
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
  <token>gr75sr</token>
  <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:08Z</updated-at>
  <venmo-sdk type="boolean">false</venmo-sdk>
  <verifications type="array"/>
</credit-card>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Create(&CreditCard{
		CustomerId:      "57131083",
		Number:          testCreditCards["visa"].Number,
		ExpirationMonth: "05",
		ExpirationYear:  "2014",
		CVV:             "100",
	})

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}
}

func TestCreateCreditCardInvalidInput(t *testing.T) {
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

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		// This seems like it shouldn't be a status created repsonse if we get an error...  Is this an API bug?
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Create(&CreditCard{
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
	})

	t.Log(card)

	// This test should fail because customer id is required
	if err == nil {
		t.Fail()
	}

	// TODO: validate fields
}

func TestFindCreditCard(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <bin>411111</bin>
  <card-type>Visa</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:15Z</created-at>
  <customer-id>74987569</customer-id>
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
  <token>ghd3x6</token>
  <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:14Z</updated-at>
  <venmo-sdk type="boolean">false</venmo-sdk>
  <verifications type="array">
    <verification>
      <status>verified</status>
      <cvv-response-code>M</cvv-response-code>
      <avs-error-response-code nil="true"/>
      <avs-postal-code-response-code>I</avs-postal-code-response-code>
      <avs-street-address-response-code>I</avs-street-address-response-code>
      <gateway-rejection-reason nil="true"/>
      <merchant-account-id>foo</merchant-account-id>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <id>557nbb</id>
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
        <token>ghd3x6</token>
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
      <created-at type="datetime">2013-11-20T23:05:14Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:15Z</updated-at>
    </verification>
  </verifications>
</credit-card>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Find("ghd3x6")

	t.Log(card)

	if err != nil {
		t.Fatal(err)
	}
	if card.Token != "ghd3x6" {
		t.Fatal("tokens do not match")
	}
}

func TestFindCreditCardBadData(t *testing.T) {
	var response = []byte(``)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Find("invalid_token")

	t.Log(card)

	if err == nil {
		t.Fail()
	}
}

func TestSaveCreditCardWithVenmoSDKPaymentMethodCode(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <billing-address>
    <id>np</id>
    <customer-id>20414939</customer-id>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code>60614</postal-code>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
    <country-name nil="true"/>
    <created-at type="datetime">2013-11-20T23:05:21Z</created-at>
    <updated-at type="datetime">2013-11-20T23:05:21Z</updated-at>
  </billing-address>
  <bin>411111</bin>
  <card-type>Visa</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:21Z</created-at>
  <customer-id>20414939</customer-id>
  <customer-location>US</customer-location>
  <debit>Unknown</debit>
  <default type="boolean">true</default>
  <durbin-regulated>Unknown</durbin-regulated>
  <expiration-month>01</expiration-month>
  <expiration-year>2015</expiration-year>
  <expired type="boolean">false</expired>
  <healthcare>Unknown</healthcare>
  <image-url>https://assets.braintreegateway.com/payment_method_logo/visa_via_venmo.png?environment=sandbox&amp;merchant_id=foo</image-url>
  <issuing-bank>Unknown</issuing-bank>
  <last-4>1111</last-4>
  <payroll>Unknown</payroll>
  <prepaid>Unknown</prepaid>
  <subscriptions type="array"/>
  <token>5qqpbm</token>
  <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:21Z</updated-at>
  <venmo-sdk type="boolean">true</venmo-sdk>
  <verifications type="array"/>
</credit-card>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Create(&CreditCard{
		CustomerId:                "20414939",
		VenmoSDKPaymentMethodCode: "stub-" + testCreditCards["visa"].Number,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}
}

func TestSaveCreditCardWithVenmoSDKSession(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card>
  <bin>411111</bin>
  <card-type>Visa</card-type>
  <cardholder-name nil="true"/>
  <commercial>Unknown</commercial>
  <country-of-issuance>Unknown</country-of-issuance>
  <created-at type="datetime">2013-11-20T23:05:22Z</created-at>
  <customer-id>81995216</customer-id>
  <customer-location>US</customer-location>
  <debit>Unknown</debit>
  <default type="boolean">true</default>
  <durbin-regulated>Unknown</durbin-regulated>
  <expiration-month>05</expiration-month>
  <expiration-year>2014</expiration-year>
  <expired type="boolean">false</expired>
  <healthcare>Unknown</healthcare>
  <image-url>https://assets.braintreegateway.com/payment_method_logo/visa_via_venmo.png?environment=sandbox&amp;merchant_id=foo</image-url>
  <issuing-bank>Unknown</issuing-bank>
  <last-4>1111</last-4>
  <payroll>Unknown</payroll>
  <prepaid>Unknown</prepaid>
  <subscriptions type="array"/>
  <token>fj3tdw</token>
  <unique-number-identifier>0d4987f6852ffa8612463aff627d84b5</unique-number-identifier>
  <updated-at type="datetime">2013-11-20T23:05:22Z</updated-at>
  <venmo-sdk type="boolean">true</venmo-sdk>
  <verifications type="array"/>
</credit-card>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Create(&CreditCard{
		CustomerId:     "81995216",
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		Options: &CreditCardOptions{
			VenmoSDKSession: "stub-session",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}
}
