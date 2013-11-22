package braintree

import (
	"net/http"
	"testing"
)

func TestAddressCreate(t *testing.T) {
	var response = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<address>
  <id>gb</id>
  <customer-id>71086244</customer-id>
  <first-name>Jenna</first-name>
  <last-name>Smith</last-name>
  <company>Braintree</company>
  <street-address>1 E Main St</street-address>
  <extended-address>Suite 403</extended-address>
  <locality>Chicago</locality>
  <region>Illinois</region>
  <postal-code>60622</postal-code>
  <country-code-alpha2>US</country-code-alpha2>
  <country-code-alpha3>USA</country-code-alpha3>
  <country-code-numeric>840</country-code-numeric>
  <country-name>United States of America</country-name>
  <created-at type="datetime">2013-11-20T23:05:04Z</created-at>
  <updated-at type="datetime">2013-11-20T23:05:04Z</updated-at>
</address> `)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		writeZip(w, response)
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	addr := Address{
		CustomerId:         "71086244",
		FirstName:          "Jenna",
		LastName:           "Smith",
		Company:            "Braintree",
		StreetAddress:      "1 E Main St",
		ExtendedAddress:    "Suite 403",
		Locality:           "Chicago",
		Region:             "Illinois",
		PostalCode:         "60622",
		CountryCodeAlpha2:  "US",
		CountryCodeAlpha3:  "USA",
		CountryCodeNumeric: "840",
		CountryName:        "United States of America",
	}

	err := gw.Address().Create(&addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(addr)

	if addr.Id == "" {
		t.Fatal()
	}
	if addr.CustomerId != "71086244" {
		t.Fatal(addr.CustomerId, "Customer ID did not match")
	}
	if addr.FirstName != "Jenna" {
		t.Fatal(addr.FirstName, "FirstName failed to save")
	}
	if addr.LastName != "Smith" {
		t.Fatal()
	}
	if addr.Company != "Braintree" {
		t.Fatal()
	}
	if addr.StreetAddress != "1 E Main St" {
		t.Fatal()
	}
	if addr.ExtendedAddress != "Suite 403" {
		t.Fatal()
	}
	if addr.Locality != "Chicago" {
		t.Fatal()
	}
	if addr.Region != "Illinois" {
		t.Fatal()
	}
	if addr.PostalCode != "60622" {
		t.Fatal()
	}
	if addr.CountryCodeAlpha2 != "US" {
		t.Fatal()
	}
	if addr.CountryCodeAlpha3 != "USA" {
		t.Fatal()
	}
	if addr.CountryCodeNumeric != "840" {
		t.Fatal()
	}
	if addr.CountryName != "United States of America" {
		t.Fatal()
	}
	if addr.CreatedAt == "" {
		t.Fatal()
	}
	if addr.UpdatedAt == "" {
		t.Fatal()
	}

}

func TestAddressDelete(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		writeZip(w, []byte(``))
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	err := gw.Address().Delete("71086244", "gb")
	if err != nil {
		t.Fatal(err)
	}

}
