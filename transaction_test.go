package braintree

import (
	"net/http"
	"testing"
)

func TestTransactionCreate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>br66hg</id>
  <status>authorized</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>100.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id nil="true"/>
  <created-at type="datetime">2013-11-20T23:05:36Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:37Z</updated-at>
  <customer>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>I</avs-postal-code-response-code>
  <avs-street-address-response-code>I</avs-street-address-response-code>
  <cvv-response-code>I</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>62Q9QZ</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token nil="true"/>
    <bin>411111</bin>
    <last-4>1111</last-4>
    <card-type>Visa</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:37Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	tx := Transaction{
		Type:   "sale",
		Amount: 100.00,
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	err := gw.Transaction().Create(&tx)

	t.Log(tx)

	if err != nil {
		t.Fatal(err)
	}
	if tx.Id == "" {
		t.Fatal("Received invalid ID on new transaction")
	}
	if tx.Status != "authorized" {
		t.Fatal(tx.Status)
	}

}

func TestTransactionSettle(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>br66hg</id>
  <status>submitted_for_settlement</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>10.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id nil="true"/>
  <created-at type="datetime">2013-11-20T23:05:36Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:37Z</updated-at>
  <customer>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>I</avs-postal-code-response-code>
  <avs-street-address-response-code>I</avs-street-address-response-code>
  <cvv-response-code>I</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>62Q9QZ</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token nil="true"/>
    <bin>411111</bin>
    <last-4>1111</last-4>
    <card-type>Visa</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:37Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:37Z</timestamp>
      <status>submitted_for_settlement</status>
      <amount>10.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>

	`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	tx2, err := gw.Transaction().SubmitForSettlement("br66hg", 10)

	t.Log(tx2)

	if err != nil {
		t.Fatal(err)
	}
	if x := tx2.Status; x != "submitted_for_settlement" {
		t.Fatal(x)
	}
	if x := tx2.Amount; x != 10 {
		t.Fatal(x)
	}
}

func TestTransactionVoid(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>br66hg</id>
  <status>voided</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>10.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id nil="true"/>
  <created-at type="datetime">2013-11-20T23:05:36Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:38Z</updated-at>
  <customer>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>I</avs-postal-code-response-code>
  <avs-street-address-response-code>I</avs-street-address-response-code>
  <cvv-response-code>I</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>62Q9QZ</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token nil="true"/>
    <bin>411111</bin>
    <last-4>1111</last-4>
    <card-type>Visa</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:37Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:37Z</timestamp>
      <status>submitted_for_settlement</status>
      <amount>10.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:38Z</timestamp>
      <status>voided</status>
      <amount>10.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	tx3, err := gw.Transaction().Void("br66hg")

	t.Log(tx3)

	if err != nil {
		t.Fatal(err)
	}
	if x := tx3.Status; x != "voided" {
		t.Fatal(x)
	}
}

func TestTransactionSearch(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<credit-card-transactions type="collection">
  <current-page-number type="integer">1</current-page-number>
  <page-size type="integer">50</page-size>
  <total-items type="integer">1</total-items>
  <transaction>
    <id>4wdkfr</id>
    <status>authorized</status>
    <type>sale</type>
    <currency-iso-code>USD</currency-iso-code>
    <amount>196.00</amount>
    <merchant-account-id>foo</merchant-account-id>
    <order-id nil="true"/>
    <created-at type="datetime">2013-11-20T23:05:38Z</created-at>
    <updated-at type="datetime">2013-11-20T23:05:39Z</updated-at>
    <customer>
      <id nil="true"/>
      <first-name>Erik-1384988738</first-name>
      <last-name nil="true"/>
      <company nil="true"/>
      <email nil="true"/>
      <website nil="true"/>
      <phone nil="true"/>
      <fax nil="true"/>
    </customer>
    <billing>
      <id nil="true"/>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <street-address nil="true"/>
      <extended-address nil="true"/>
      <locality nil="true"/>
      <region nil="true"/>
      <postal-code nil="true"/>
      <country-name nil="true"/>
      <country-code-alpha2 nil="true"/>
      <country-code-alpha3 nil="true"/>
      <country-code-numeric nil="true"/>
    </billing>
    <refund-id nil="true"/>
    <refund-ids type="array"/>
    <refunded-transaction-id nil="true"/>
    <settlement-batch-id nil="true"/>
    <shipping>
      <id nil="true"/>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <street-address nil="true"/>
      <extended-address nil="true"/>
      <locality nil="true"/>
      <region nil="true"/>
      <postal-code nil="true"/>
      <country-name nil="true"/>
      <country-code-alpha2 nil="true"/>
      <country-code-alpha3 nil="true"/>
      <country-code-numeric nil="true"/>
    </shipping>
    <custom-fields>
    </custom-fields>
    <avs-error-response-code nil="true"/>
    <avs-postal-code-response-code>I</avs-postal-code-response-code>
    <avs-street-address-response-code>I</avs-street-address-response-code>
    <cvv-response-code>I</cvv-response-code>
    <gateway-rejection-reason nil="true"/>
    <processor-authorization-code>R6CW4Y</processor-authorization-code>
    <processor-response-code>1000</processor-response-code>
    <processor-response-text>Approved</processor-response-text>
    <voice-referral-number nil="true"/>
    <purchase-order-number nil="true"/>
    <tax-amount nil="true"/>
    <tax-exempt type="boolean">false</tax-exempt>
    <credit-card>
      <token nil="true"/>
      <bin>411111</bin>
      <last-4>1111</last-4>
      <card-type>Visa</card-type>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <customer-location>US</customer-location>
      <cardholder-name nil="true"/>
      <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
      <prepaid>Unknown</prepaid>
      <healthcare>Unknown</healthcare>
      <debit>Unknown</debit>
      <durbin-regulated>Unknown</durbin-regulated>
      <commercial>Unknown</commercial>
      <payroll>Unknown</payroll>
      <issuing-bank>Unknown</issuing-bank>
      <country-of-issuance>Unknown</country-of-issuance>
      <venmo-sdk type="boolean">false</venmo-sdk>
    </credit-card>
    <status-history type="array">
      <status-event>
        <timestamp type="datetime">2013-11-20T23:05:39Z</timestamp>
        <status>authorized</status>
        <amount>196.00</amount>
        <user>corylanou-sandbox</user>
        <transaction-source>API</transaction-source>
      </status-event>
    </status-history>
    <plan-id nil="true"/>
    <subscription-id nil="true"/>
    <subscription>
      <billing-period-end-date nil="true"/>
      <billing-period-start-date nil="true"/>
    </subscription>
    <add-ons type="array"/>
    <discounts type="array"/>
    <descriptor>
      <name nil="true"/>
      <phone nil="true"/>
    </descriptor>
    <recurring type="boolean">false</recurring>
    <channel nil="true"/>
    <service-fee-amount nil="true"/>
    <escrow-status nil="true"/>
    <disbursement-details>
      <disbursement-date nil="true"/>
      <settlement-amount nil="true"/>
      <settlement-currency-iso-code nil="true"/>
      <settlement-currency-exchange-rate nil="true"/>
      <funds-held nil="true"/>
    </disbursement-details>
  </transaction>
</credit-card-transactions>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	var query SearchQuery
	f := query.AddTextField("customer-first-name")
	f.Is = "Jimmy"

	result, err := txg.Search(query)
	if err != nil {
		t.Fatal(err)
	}

	tx := result.Transactions[0]
	if tx.Amount != 196.0 {
		t.Fatal(tx.Amount)
	}
}

func TestTransactionCreateWhenGatewayRejected(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<api-error-response>
  <errors>
    <errors type="array"/>
  </errors>
  <params>
    <transaction>
      <type>sale</type>
      <amount>2010</amount>
      <credit-card>
        <expiration-date>05/14</expiration-date>
      </credit-card>
    </transaction>
  </params>
  <message>Card Issuer Declined CVV</message>
  <transaction>
    <id>cqgfzr</id>
    <status>processor_declined</status>
    <type>sale</type>
    <currency-iso-code>USD</currency-iso-code>
    <amount>2010.00</amount>
    <merchant-account-id>foo</merchant-account-id>
    <order-id nil="true"/>
    <created-at type="datetime">2013-11-20T23:05:41Z</created-at>
    <updated-at type="datetime">2013-11-20T23:05:42Z</updated-at>
    <customer>
      <id nil="true"/>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <email nil="true"/>
      <website nil="true"/>
      <phone nil="true"/>
      <fax nil="true"/>
    </customer>
    <billing>
      <id nil="true"/>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <street-address nil="true"/>
      <extended-address nil="true"/>
      <locality nil="true"/>
      <region nil="true"/>
      <postal-code nil="true"/>
      <country-name nil="true"/>
      <country-code-alpha2 nil="true"/>
      <country-code-alpha3 nil="true"/>
      <country-code-numeric nil="true"/>
    </billing>
    <refund-id nil="true"/>
    <refund-ids type="array"/>
    <refunded-transaction-id nil="true"/>
    <settlement-batch-id nil="true"/>
    <shipping>
      <id nil="true"/>
      <first-name nil="true"/>
      <last-name nil="true"/>
      <company nil="true"/>
      <street-address nil="true"/>
      <extended-address nil="true"/>
      <locality nil="true"/>
      <region nil="true"/>
      <postal-code nil="true"/>
      <country-name nil="true"/>
      <country-code-alpha2 nil="true"/>
      <country-code-alpha3 nil="true"/>
      <country-code-numeric nil="true"/>
    </shipping>
    <custom-fields>
    </custom-fields>
    <avs-error-response-code nil="true"/>
    <avs-postal-code-response-code>I</avs-postal-code-response-code>
    <avs-street-address-response-code>I</avs-street-address-response-code>
    <cvv-response-code>I</cvv-response-code>
    <gateway-rejection-reason nil="true"/>
    <processor-authorization-code nil="true"/>
    <processor-response-code>2010</processor-response-code>
    <processor-response-text>Card Issuer Declined CVV</processor-response-text>
    <voice-referral-number nil="true"/>
    <purchase-order-number nil="true"/>
    <tax-amount nil="true"/>
    <tax-exempt type="boolean">false</tax-exempt>
    <credit-card>
      <token nil="true"/>
      <bin>411111</bin>
      <last-4>1111</last-4>
      <card-type>Visa</card-type>
      <expiration-month>05</expiration-month>
      <expiration-year>2014</expiration-year>
      <customer-location>US</customer-location>
      <cardholder-name nil="true"/>
      <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
      <prepaid>Unknown</prepaid>
      <healthcare>Unknown</healthcare>
      <debit>Unknown</debit>
      <durbin-regulated>Unknown</durbin-regulated>
      <commercial>Unknown</commercial>
      <payroll>Unknown</payroll>
      <issuing-bank>Unknown</issuing-bank>
      <country-of-issuance>Unknown</country-of-issuance>
      <venmo-sdk type="boolean">false</venmo-sdk>
    </credit-card>
    <status-history type="array">
      <status-event>
        <timestamp type="datetime">2013-11-20T23:05:42Z</timestamp>
        <status>processor_declined</status>
        <amount>2010.00</amount>
        <user>corylanou-sandbox</user>
        <transaction-source>API</transaction-source>
      </status-event>
    </status-history>
    <plan-id nil="true"/>
    <subscription-id nil="true"/>
    <subscription>
      <billing-period-end-date nil="true"/>
      <billing-period-start-date nil="true"/>
    </subscription>
    <add-ons type="array"/>
    <discounts type="array"/>
    <descriptor>
      <name nil="true"/>
      <phone nil="true"/>
    </descriptor>
    <recurring type="boolean">false</recurring>
    <channel nil="true"/>
    <service-fee-amount nil="true"/>
    <escrow-status nil="true"/>
    <disbursement-details>
      <disbursement-date nil="true"/>
      <settlement-amount nil="true"/>
      <settlement-currency-iso-code nil="true"/>
      <settlement-currency-exchange-rate nil="true"/>
      <funds-held nil="true"/>
    </disbursement-details>
  </transaction>
</api-error-response>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	tx := Transaction{
		Type:   "sale",
		Amount: 2010.00,
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
		},
	}
	err := txg.Create(&tx)
	if err != nil {
		t.Fatal("Did not receive error when creating invalid transaction")
	}
	if tx.Success() {
		t.Fatal(tx.ErrorMessage, "Did not receive error when creating invalid transaction")
	}
	if tx.ErrorMessage != "Card Issuer Declined CVV" {
		t.Fatal(err)
	}
}

func TestFindTransaction(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>kk833g</id>
  <status>authorized</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>100.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id nil="true"/>
  <created-at type="datetime">2013-11-20T23:05:44Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:44Z</updated-at>
  <customer>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>I</avs-postal-code-response-code>
  <avs-street-address-response-code>I</avs-street-address-response-code>
  <cvv-response-code>I</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>1S78FS</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token nil="true"/>
    <bin>555555</bin>
    <last-4>4444</last-4>
    <card-type>MasterCard</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/mastercard.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:44Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	foundTransaction, err := txg.Find("kk833g")
	if err != nil {
		t.Fatal(err)
	}

	if "kk833g" != foundTransaction.Id {
		t.Fatal("transaction ids do not match")
	}
}

func TestFindNonExistantTransaction(t *testing.T) {
	var response = []byte(``)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	_, err := txg.Find("bad_transaction_id")
	if err == nil {
		t.Fatal("Did not receive error when finding an invalid tx ID")
	}
	if err.Error() != "EOF" {
		t.Fatal(err)
	}
}

func TestAllTransactionFields(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>9vfk6w</id>
  <status>submitted_for_settlement</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>100.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id>my_custom_order</order-id>
  <created-at type="datetime">2013-11-20T23:05:46Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:46Z</updated-at>
  <customer>
    <id>13282481</id>
    <first-name>Lionel</first-name>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id>7s</id>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address>1 E Main St</street-address>
    <extended-address nil="true"/>
    <locality>Chicago</locality>
    <region>IL</region>
    <postal-code>60637</postal-code>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address>1 E Main St</street-address>
    <extended-address nil="true"/>
    <locality>Chicago</locality>
    <region>IL</region>
    <postal-code>60637</postal-code>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>M</avs-postal-code-response-code>
  <avs-street-address-response-code>M</avs-street-address-response-code>
  <cvv-response-code>M</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>DYVS9S</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token>54nzn6</token>
    <bin>411111</bin>
    <last-4>1111</last-4>
    <card-type>Visa</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/visa.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:46Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:46Z</timestamp>
      <status>submitted_for_settlement</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	tx := &Transaction{
		Type:    "sale",
		Amount:  100.00,
		OrderId: "my_custom_order",
		CreditCard: &CreditCard{
			Number:         testCreditCards["visa"].Number,
			ExpirationDate: "05/14",
			CVV:            "100",
		},
		Customer: &Customer{
			FirstName: "Lionel",
		},
		BillingAddress: &Address{
			StreetAddress: "1 E Main St",
			Locality:      "Chicago",
			Region:        "IL",
			PostalCode:    "60637",
		},
		ShippingAddress: &Address{
			StreetAddress: "1 E Main St",
			Locality:      "Chicago",
			Region:        "IL",
			PostalCode:    "60637",
		},
		Options: &TransactionOptions{
			SubmitForSettlement:              true,
			StoreInVault:                     true,
			AddBillingAddressToPaymentMethod: true,
			StoreShippingAddressInVault:      true,
		},
	}

	err := txg.Create(tx)
	if err != nil {
		t.Fatal(err)
	}

	if tx.Type != tx.Type {
		t.Fail()
	}
	if tx.Amount != tx.Amount {
		t.Fail()
	}
	if tx.OrderId != tx.OrderId {
		t.Fail()
	}
	if tx.Customer.FirstName != tx.Customer.FirstName {
		t.Fail()
	}
	if tx.BillingAddress.StreetAddress != tx.BillingAddress.StreetAddress {
		t.Fail()
	}
	if tx.BillingAddress.Locality != tx.BillingAddress.Locality {
		t.Fail()
	}
	if tx.BillingAddress.Region != tx.BillingAddress.Region {
		t.Fail()
	}
	if tx.BillingAddress.PostalCode != tx.BillingAddress.PostalCode {
		t.Fail()
	}
	if tx.ShippingAddress.StreetAddress != tx.ShippingAddress.StreetAddress {
		t.Fail()
	}
	if tx.ShippingAddress.Locality != tx.ShippingAddress.Locality {
		t.Fail()
	}
	if tx.ShippingAddress.Region != tx.ShippingAddress.Region {
		t.Fail()
	}
	if tx.ShippingAddress.PostalCode != tx.ShippingAddress.PostalCode {
		t.Fail()
	}
	if tx.CreditCard.Token == "" {
		t.Fail()
	}
	if tx.Customer.Id == "" {
		t.Fail()
	}
	if tx.Status != "submitted_for_settlement" {
		t.Fail()
	}
}

func TestTransactionCreateFromPaymentMethodCode(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<transaction>
  <id>ct247g</id>
  <status>authorized</status>
  <type>sale</type>
  <currency-iso-code>USD</currency-iso-code>
  <amount>100.00</amount>
  <merchant-account-id>foo</merchant-account-id>
  <order-id nil="true"/>
  <created-at type="datetime">2013-11-20T23:05:47Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:48Z</updated-at>
  <customer>
    <id>34430666</id>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <email nil="true"/>
    <website nil="true"/>
    <phone nil="true"/>
    <fax nil="true"/>
  </customer>
  <billing>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </billing>
  <refund-id nil="true"/>
  <refund-ids type="array"/>
  <refunded-transaction-id nil="true"/>
  <settlement-batch-id nil="true"/>
  <shipping>
    <id nil="true"/>
    <first-name nil="true"/>
    <last-name nil="true"/>
    <company nil="true"/>
    <street-address nil="true"/>
    <extended-address nil="true"/>
    <locality nil="true"/>
    <region nil="true"/>
    <postal-code nil="true"/>
    <country-name nil="true"/>
    <country-code-alpha2 nil="true"/>
    <country-code-alpha3 nil="true"/>
    <country-code-numeric nil="true"/>
  </shipping>
  <custom-fields>
  </custom-fields>
  <avs-error-response-code nil="true"/>
  <avs-postal-code-response-code>I</avs-postal-code-response-code>
  <avs-street-address-response-code>I</avs-street-address-response-code>
  <cvv-response-code>I</cvv-response-code>
  <gateway-rejection-reason nil="true"/>
  <processor-authorization-code>P4XG4R</processor-authorization-code>
  <processor-response-code>1000</processor-response-code>
  <processor-response-text>Approved</processor-response-text>
  <voice-referral-number nil="true"/>
  <purchase-order-number nil="true"/>
  <tax-amount nil="true"/>
  <tax-exempt type="boolean">false</tax-exempt>
  <credit-card>
    <token>8zv3pb</token>
    <bin>601111</bin>
    <last-4>1117</last-4>
    <card-type>Discover</card-type>
    <expiration-month>05</expiration-month>
    <expiration-year>2014</expiration-year>
    <customer-location>US</customer-location>
    <cardholder-name nil="true"/>
    <image-url>https://assets.braintreegateway.com/payment_method_logo/discover.png?environment=sandbox&amp;merchant_id=foo</image-url>
    <prepaid>Unknown</prepaid>
    <healthcare>Unknown</healthcare>
    <debit>Unknown</debit>
    <durbin-regulated>Unknown</durbin-regulated>
    <commercial>Unknown</commercial>
    <payroll>Unknown</payroll>
    <issuing-bank>Unknown</issuing-bank>
    <country-of-issuance>Unknown</country-of-issuance>
    <venmo-sdk type="boolean">false</venmo-sdk>
  </credit-card>
  <status-history type="array">
    <status-event>
      <timestamp type="datetime">2013-11-20T23:05:48Z</timestamp>
      <status>authorized</status>
      <amount>100.00</amount>
      <user>corylanou-sandbox</user>
      <transaction-source>API</transaction-source>
    </status-event>
  </status-history>
  <plan-id nil="true"/>
  <subscription-id nil="true"/>
  <subscription>
    <billing-period-end-date nil="true"/>
    <billing-period-start-date nil="true"/>
  </subscription>
  <add-ons type="array"/>
  <discounts type="array"/>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <recurring type="boolean">false</recurring>
  <channel nil="true"/>
  <service-fee-amount nil="true"/>
  <escrow-status nil="true"/>
  <disbursement-details>
    <disbursement-date nil="true"/>
    <settlement-amount nil="true"/>
    <settlement-currency-iso-code nil="true"/>
    <settlement-currency-exchange-rate nil="true"/>
    <funds-held nil="true"/>
  </disbursement-details>
</transaction>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	txg := gw.Transaction()

	tx := Transaction{
		Type:               "sale",
		CustomerID:         "34430666",
		Amount:             100,
		PaymentMethodToken: "8zv3pb",
	}
	err := txg.Create(&tx)

	if err != nil {
		t.Fatal(err)
	}
	if tx.Id == "" {
		t.Fatal("invalid tx id")
	}
}
