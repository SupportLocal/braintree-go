package braintree

import (
	"net/http"
	"testing"
)

func TestFindAllPlans(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<plans type="array">
  <plan>
    <id>community-sponsor-yearly</id>
    <merchant-id>foo</merchant-id>
    <billing-day-of-month nil="true"/>
    <billing-frequency type="integer">12</billing-frequency>
    <currency-iso-code>USD</currency-iso-code>
    <description></description>
    <name>Community Sponsor (Yearly Subscription)</name>
    <number-of-billing-cycles nil="true"/>
    <price>129.00</price>
    <trial-duration nil="true"/>
    <trial-duration-unit>day</trial-duration-unit>
    <trial-period type="boolean">false</trial-period>
    <created-at type="datetime">2013-11-12T03:36:26Z</created-at>
    <updated-at type="datetime">2013-11-12T03:36:26Z</updated-at>
    <add-ons type="array"/>
    <discounts type="array">
      <discount>
        <amount>64.50</amount>
        <created-at type="datetime">2013-11-12T03:35:33Z</created-at>
        <description></description>
        <id>smb-saturday</id>
        <kind>discount</kind>
        <merchant-id>foo</merchant-id>
        <name>Small Business Saturday</name>
        <never-expires type="boolean">false</never-expires>
        <number-of-billing-cycles type="integer">1</number-of-billing-cycles>
        <updated-at type="datetime">2013-11-12T03:35:33Z</updated-at>
      </discount>
    </discounts>
  </plan>
  <plan>
    <id>standard-membership</id>
    <merchant-id>foo</merchant-id>
    <billing-day-of-month nil="true"/>
    <billing-frequency type="integer">12</billing-frequency>
    <currency-iso-code>USD</currency-iso-code>
    <description>Blah blah blah.</description>
    <name>Standard Membership</name>
    <number-of-billing-cycles nil="true"/>
    <price>199.00</price>
    <trial-duration nil="true"/>
    <trial-duration-unit>day</trial-duration-unit>
    <trial-period type="boolean">false</trial-period>
    <created-at type="datetime">2013-07-15T16:12:48Z</created-at>
    <updated-at type="datetime">2013-07-16T23:11:15Z</updated-at>
    <add-ons type="array"/>
    <discounts type="array"/>
  </plan>
  <plan>
    <id>test_plan_2</id>
    <merchant-id>foo</merchant-id>
    <billing-day-of-month type="integer">31</billing-day-of-month>
    <billing-frequency type="integer">1</billing-frequency>
    <currency-iso-code>USD</currency-iso-code>
    <description>test_plan_2_description</description>
    <name>test_plan_2_name</name>
    <number-of-billing-cycles nil="true"/>
    <price>20.00</price>
    <trial-duration nil="true"/>
    <trial-duration-unit>day</trial-duration-unit>
    <trial-period type="boolean">false</trial-period>
    <created-at type="datetime">2013-11-20T19:34:14Z</created-at>
    <updated-at type="datetime">2013-11-20T19:34:14Z</updated-at>
    <add-ons type="array"/>
    <discounts type="array"/>
  </plan>
  <plan>
    <id>test_plan</id>
    <merchant-id>foo</merchant-id>
    <billing-day-of-month nil="true"/>
    <billing-frequency type="integer">1</billing-frequency>
    <currency-iso-code>USD</currency-iso-code>
    <description>test_plan_desc</description>
    <name>test_plan_name</name>
    <number-of-billing-cycles type="integer">2</number-of-billing-cycles>
    <price>10.00</price>
    <trial-duration type="integer">14</trial-duration>
    <trial-duration-unit>day</trial-duration-unit>
    <trial-period type="boolean">true</trial-period>
    <created-at type="datetime">2013-11-20T19:33:22Z</created-at>
    <updated-at type="datetime">2013-11-20T19:37:03Z</updated-at>
    <add-ons type="array"/>
    <discounts type="array"/>
  </plan>
</plans>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.Plan()

	plans, err := g.All()
	if err != nil {
		t.Fatal(err)
	}
	if len(plans) == 0 {
		t.Fatal(plans)
	}

	var plan Plan
	for _, p := range plans {
		if p.Id == "test_plan" {
			plan = p
			break
		}
	}

	t.Log(plan)

	if x := plan.Id; x != "test_plan" {
		t.Fatal(x)
	}
	if x := plan.MerchantId; x == "" {
		t.Fatal(x)
	}
	if x := plan.BillingFrequency; x != "1" {
		t.Fatal(x)
	}
	if x := plan.CurrencyISOCode; x != "USD" {
		t.Fatal(x)
	}
	if x := plan.Description; x != "test_plan_desc" {
		t.Fatal(x)
	}
	if x := plan.Name; x != "test_plan_name" {
		t.Fatal(x)
	}
	if x := plan.NumberOfBillingCycles; x != "2" {
		t.Fatal(x)
	}
	if x := plan.Price; x != 10.0 {
		t.Fatal(x)
	}
	if x := plan.TrialDuration; x != "14" {
		t.Fatal(x)
	}
	if x := plan.TrialDurationUnit; x != "day" {
		t.Fatal(x)
	}
	if x := plan.TrialPeriod; x != "true" {
		t.Fatal(x)
	}
	if x := plan.CreatedAt; x == "" {
		t.Fatal(x)
	}
	if x := plan.UpdatedAt; x == "" {
		t.Fatal(x)
	}

}

func TestFindPlan(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<plans type="array">
  <plan>
    <id>test_plan_2</id>
    <merchant-id>foo</merchant-id>
    <billing-day-of-month type="integer">31</billing-day-of-month>
    <billing-frequency type="integer">1</billing-frequency>
    <currency-iso-code>USD</currency-iso-code>
    <description>test_plan_2_description</description>
    <name>test_plan_2_name</name>
    <number-of-billing-cycles nil="true"/>
    <price>20.00</price>
    <trial-duration nil="true"/>
    <trial-duration-unit>day</trial-duration-unit>
    <trial-period type="boolean">false</trial-period>
    <created-at type="datetime">2013-11-20T19:34:14Z</created-at>
    <updated-at type="datetime">2013-11-20T19:34:14Z</updated-at>
    <add-ons type="array"/>
    <discounts type="array"/>
  </plan>
	</plans>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.Plan()

	plan2, err := g.Find("test_plan_2")
	if err != nil {
		t.Fatal(err)
	}
	if plan2.Id != "test_plan_2" {
		t.Fatal(plan2)
	}
	if x := plan2.BillingDayOfMonth; x != "31" {
		t.Fatal(x)
	}
}
