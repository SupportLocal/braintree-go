package braintree

import (
	"net/http"
	"testing"
)

func TestAddOn(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<add-ons type="array">
  <add-on>
    <amount>10.00</amount>
    <created-at type="datetime">2013-11-20T19:38:49Z</created-at>
    <description>A test add on</description>
    <id>test_add_on_id</id>
    <kind>add_on</kind>
    <merchant-id>foo</merchant-id>
    <name>test_add_on_name</name>
    <never-expires type="boolean">true</never-expires>
    <number-of-billing-cycles nil="true"/>
    <updated-at type="datetime">2013-11-20T19:38:49Z</updated-at>
  </add-on>
</add-ons>`)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	addOns, err := gw.AddOn().All()

	if err != nil {
		t.Error(err)
	} else if len(addOns) != 1 {
		t.Fail()
	}

	addOn := addOns[0]

	t.Log(addOn)

	if addOn.Id != "test_add_on_id" {
		t.Fail()
	} else if addOn.Amount != 10 {
		t.Fail()
	} else if addOn.Kind != ModificationKindAddOn {
		t.Fail()
	} else if addOn.Name != "test_add_on_name" {
		t.Fail()
	} else if addOn.NeverExpires != true {
		t.Fail()
	} else if addOn.Description != "A test add on" {
		t.Fail()
	}
}
