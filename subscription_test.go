package braintree

import (
	"net/http"
	"testing"
)

// Everything you wanted to know about subscriptions:
// https://www.braintreepayments.com/docs/ruby/subscriptions/create

func TestSubscriptionCreateWithDiscount(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "subscription", "create_with_discount", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	var customer = Customer{
		Id:        "tc_subscrip",
		FirstName: "Test",
		LastName:  "Subscription",
		Email:     "cory+bt@supportlocal.com",
	}

	if err := testCustomerCreate(&customer); err != nil {
		t.Fatal(err, "Unable to set up test customer")
	}

	creditCard := testCreditCard
	creditCard.Number = "6011111111111117"
	creditCard.CustomerId = customer.Id
	if err := testCreditCardCreate(&creditCard); err != nil {
		t.Fatal(err, "Unable to set up test credit card")
	}

	d := Discount{}
	d.Id = "test_discount"
	d.Amount = 5
	sub := Subscription{
		PaymentMethodToken: creditCard.Token,
		PlanId:             "test_plan",
	}
	sub.AddDiscount(d, 1)

	gw := Braintree{BaseURL: server.URL}

	err := gw.Subscription().Create(&sub)

	t.Log(sub)

	if err != nil {
		t.Log(sub.ErrorMessage)
		t.Fatal(err)
	}

	if !sub.Success() {
		t.Fatalf("Recieved an error of %q", sub.ErrorMessage)
	}

	if err := testCustomerDelete(customer.Id); err != nil {
		t.Fatal(err, "Unable to delete test customer")
	}

}

func TestSubscriptionCreate(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "subscription", "create", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	customer := testCustomer
	if c, err := testCustomerFindOrCreate(customer); err != nil {
		t.Fatal(err, "Unable to set up test customer")
	} else {
		customer = c
	}

	creditCard := testCreditCard
	creditCard.CustomerId = customer.Id
	if err := testCreditCardCreate(&creditCard); err != nil {
		t.Fatal(err, "Unable to set up test credit card")
	}

	gw := Braintree{BaseURL: server.URL}

	g := gw.Subscription()

	// Create
	subscription := Subscription{
		PaymentMethodToken: creditCard.Token,
		PlanId:             "test_plan",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	}

	err := g.Create(&subscription)

	t.Log("sub1", subscription)

	if err != nil {
		t.Fatal(err)
	}
	if subscription.Id == "" {
		t.Fatal("invalid subscription id")
	}

}

func TestSubscriptionUpdate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<subscription>
  <add-ons type="array"/>
  <balance>0.00</balance>
  <billing-day-of-month type="integer">20</billing-day-of-month>
  <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
  <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
  <current-billing-cycle type="integer">1</current-billing-cycle>
  <days-past-due nil="true"/>
  <discounts type="array"/>
  <failure-count type="integer">0</failure-count>
  <first-billing-date type="date">2013-11-20</first-billing-date>
  <id>4j2ntb</id>
  <merchant-account-id>foo</merchant-account-id>
  <never-expires type="boolean">false</never-expires>
  <next-bill-amount>10.00</next-bill-amount>
  <next-billing-period-amount>10.00</next-billing-period-amount>
  <next-billing-date type="date">2013-12-20</next-billing-date>
  <number-of-billing-cycles type="integer">2</number-of-billing-cycles>
  <paid-through-date type="date">2013-12-19</paid-through-date>
  <payment-method-token>fzqy62</payment-method-token>
  <plan-id>test_plan_2</plan-id>
  <price>10.00</price>
  <status>Active</status>
  <trial-duration nil="true"/>
  <trial-duration-unit nil="true"/>
  <trial-period type="boolean">false</trial-period>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <transactions type="array">
    <transaction>
      <id>3fg66b</id>
      <status>submitted_for_settlement</status>
      <type>sale</type>
      <currency-iso-code>USD</currency-iso-code>
      <amount>10.00</amount>
      <merchant-account-id>foo</merchant-account-id>
      <order-id nil="true"/>
      <created-at type="datetime">2013-11-20T23:05:32Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:33Z</updated-at>
      <customer>
        <id>17235299</id>
        <first-name>Lionel</first-name>
        <last-name>Barrow</last-name>
        <company>Braintree</company>
        <email>lionel.barrow@example.com</email>
        <website>http://www.example.com</website>
        <phone>312.555.1234</phone>
        <fax>614.555.5678</fax>
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
      <processor-authorization-code>BF91G9</processor-authorization-code>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <voice-referral-number nil="true"/>
      <purchase-order-number nil="true"/>
      <tax-amount nil="true"/>
      <tax-exempt type="boolean">false</tax-exempt>
      <credit-card>
        <token>fzqy62</token>
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
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>authorized</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
        <status-event>
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>submitted_for_settlement</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
      </status-history>
      <plan-id>test_plan</plan-id>
      <subscription-id>4j2ntb</subscription-id>
      <subscription>
        <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
        <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
      </subscription>
      <add-ons type="array"/>
      <discounts type="array"/>
      <descriptor>
        <name nil="true"/>
        <phone nil="true"/>
      </descriptor>
      <recurring type="boolean">true</recurring>
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
  </transactions>
</subscription>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.Subscription()

	subscription := Subscription{
		Id:     "4j2ntb",
		PlanId: "test_plan_2",
		Options: &SubscriptionOptions{
			ProrateCharges:                       true,
			RevertSubscriptionOnProrationFailure: true,
			StartImmediately:                     true,
		},
	}
	err := g.Update(&subscription)

	t.Log("sub2", subscription)

	if err != nil {
		t.Fatal(err)
	}
	if subscription.Id != "4j2ntb" {
		t.Fatal(subscription.Id)
	}
	if x := subscription.PlanId; x != "test_plan_2" {
		t.Fatal(x)
	}
}

func TestSubscriptionFind(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<subscription>
  <add-ons type="array"/>
  <balance>0.00</balance>
  <billing-day-of-month type="integer">20</billing-day-of-month>
  <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
  <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
  <current-billing-cycle type="integer">1</current-billing-cycle>
  <days-past-due nil="true"/>
  <discounts type="array"/>
  <failure-count type="integer">0</failure-count>
  <first-billing-date type="date">2013-11-20</first-billing-date>
  <id>4j2ntb</id>
  <merchant-account-id>foo</merchant-account-id>
  <never-expires type="boolean">false</never-expires>
  <next-bill-amount>10.00</next-bill-amount>
  <next-billing-period-amount>10.00</next-billing-period-amount>
  <next-billing-date type="date">2013-12-20</next-billing-date>
  <number-of-billing-cycles type="integer">2</number-of-billing-cycles>
  <paid-through-date type="date">2013-12-19</paid-through-date>
  <payment-method-token>fzqy62</payment-method-token>
  <plan-id>test_plan_2</plan-id>
  <price>10.00</price>
  <status>Active</status>
  <trial-duration nil="true"/>
  <trial-duration-unit nil="true"/>
  <trial-period type="boolean">false</trial-period>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <transactions type="array">
    <transaction>
      <id>3fg66b</id>
      <status>submitted_for_settlement</status>
      <type>sale</type>
      <currency-iso-code>USD</currency-iso-code>
      <amount>10.00</amount>
      <merchant-account-id>foo</merchant-account-id>
      <order-id nil="true"/>
      <created-at type="datetime">2013-11-20T23:05:32Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:33Z</updated-at>
      <customer>
        <id>17235299</id>
        <first-name>Lionel</first-name>
        <last-name>Barrow</last-name>
        <company>Braintree</company>
        <email>lionel.barrow@example.com</email>
        <website>http://www.example.com</website>
        <phone>312.555.1234</phone>
        <fax>614.555.5678</fax>
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
      <processor-authorization-code>BF91G9</processor-authorization-code>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <voice-referral-number nil="true"/>
      <purchase-order-number nil="true"/>
      <tax-amount nil="true"/>
      <tax-exempt type="boolean">false</tax-exempt>
      <credit-card>
        <token>fzqy62</token>
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
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>authorized</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
        <status-event>
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>submitted_for_settlement</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
      </status-history>
      <plan-id>test_plan</plan-id>
      <subscription-id>4j2ntb</subscription-id>
      <subscription>
        <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
        <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
      </subscription>
      <add-ons type="array"/>
      <discounts type="array"/>
      <descriptor>
        <name nil="true"/>
        <phone nil="true"/>
      </descriptor>
      <recurring type="boolean">true</recurring>
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
  </transactions>
</subscription>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.Subscription()

	sub3, err := g.Find("4j2ntb")
	if err != nil {
		t.Fatal(err)
	}
	if sub3.Id != "4j2ntb" {
		t.Fatal(sub3.Id)
	}
}

func TestSubscriptionCancel(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<subscription>
  <add-ons type="array"/>
  <balance>0.00</balance>
  <billing-day-of-month type="integer">20</billing-day-of-month>
  <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
  <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
  <current-billing-cycle type="integer">1</current-billing-cycle>
  <days-past-due nil="true"/>
  <discounts type="array"/>
  <failure-count type="integer">0</failure-count>
  <first-billing-date type="date">2013-11-20</first-billing-date>
  <id>4j2ntb</id>
  <merchant-account-id>foo</merchant-account-id>
  <never-expires type="boolean">false</never-expires>
  <next-bill-amount>10.00</next-bill-amount>
  <next-billing-period-amount>10.00</next-billing-period-amount>
  <next-billing-date type="date">2013-12-20</next-billing-date>
  <number-of-billing-cycles type="integer">2</number-of-billing-cycles>
  <paid-through-date type="date">2013-12-19</paid-through-date>
  <payment-method-token>fzqy62</payment-method-token>
  <plan-id>test_plan_2</plan-id>
  <price>10.00</price>
  <status>Canceled</status>
  <trial-duration nil="true"/>
  <trial-duration-unit nil="true"/>
  <trial-period type="boolean">false</trial-period>
  <descriptor>
    <name nil="true"/>
    <phone nil="true"/>
  </descriptor>
  <transactions type="array">
    <transaction>
      <id>3fg66b</id>
      <status>submitted_for_settlement</status>
      <type>sale</type>
      <currency-iso-code>USD</currency-iso-code>
      <amount>10.00</amount>
      <merchant-account-id>foo</merchant-account-id>
      <order-id nil="true"/>
      <created-at type="datetime">2013-11-20T23:05:32Z</created-at>
      <updated-at type="datetime">2013-11-20T23:05:33Z</updated-at>
      <customer>
        <id>17235299</id>
        <first-name>Lionel</first-name>
        <last-name>Barrow</last-name>
        <company>Braintree</company>
        <email>lionel.barrow@example.com</email>
        <website>http://www.example.com</website>
        <phone>312.555.1234</phone>
        <fax>614.555.5678</fax>
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
      <processor-authorization-code>BF91G9</processor-authorization-code>
      <processor-response-code>1000</processor-response-code>
      <processor-response-text>Approved</processor-response-text>
      <voice-referral-number nil="true"/>
      <purchase-order-number nil="true"/>
      <tax-amount nil="true"/>
      <tax-exempt type="boolean">false</tax-exempt>
      <credit-card>
        <token>fzqy62</token>
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
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>authorized</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
        <status-event>
          <timestamp type="datetime">2013-11-20T23:05:33Z</timestamp>
          <status>submitted_for_settlement</status>
          <amount>10.00</amount>
          <user>corylanou-sandbox</user>
          <transaction-source>Recurring</transaction-source>
        </status-event>
      </status-history>
      <plan-id>test_plan</plan-id>
      <subscription-id>4j2ntb</subscription-id>
      <subscription>
        <billing-period-end-date type="date">2013-12-19</billing-period-end-date>
        <billing-period-start-date type="date">2013-11-20</billing-period-start-date>
      </subscription>
      <add-ons type="array"/>
      <discounts type="array"/>
      <descriptor>
        <name nil="true"/>
        <phone nil="true"/>
      </descriptor>
      <recurring type="boolean">true</recurring>
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
  </transactions>
</subscription>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.Subscription()

	_, err := g.Cancel("4j2ntb")
	if err != nil {
		t.Fatal(err)
	}
}
