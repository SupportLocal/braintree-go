package braintree

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCustomerCreateWithCVVError(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "create_with_cvv_error", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	oc := Customer{
		Id:        "81827736",
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
	err := gw.Customer().Create(&oc)
	t.Log(oc)
	if err != nil {
		// Oddly enough they don't err this out in a status code...
		t.Fatal("Should not have an error")
	}
	if oc.Success() {
		t.Fatal("Should receive an error when creating an invalid customer")
	}
	if oc.ErrorMessage != "Gateway Rejected: cvv" {
		t.Fatal(fmt.Sprintf("%q", oc.ErrorMessage), fmt.Sprintf("Should of received error %q", "Gateway Rejected: cvv"))
	}
}

func TestCustomerCreate(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "create", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	customer := Customer{
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

	err := gw.Customer().Create(&customer)

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
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "update", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	customer := Customer{
		Id:        "81827736",
		FirstName: "John",
	}
	err := gw.Customer().Update(&customer)

	t.Log(customer)

	if err != nil {
		t.Fatal(err)
	}
	if customer.FirstName != "John" {
		t.Fatal("first name not changed")
	}
}

func TestCustomerFind(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "find", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c3, err := gw.Customer().Find("81827736")

	t.Log(c3)

	if err != nil {
		t.Fatal(err)
	}
	if c3.Id != "81827736" {
		t.Fatal("ids do not match")
	}
}

func TestCustomerDelete(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "delete", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	err := gw.Customer().Delete("81827736")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCustomerFind404(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "find404", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c4, err := gw.Customer().Find("81827736")
	if err == nil {
		t.Fatal("should return EOF")
	}
	if err.Error() != "EOF" {
		t.Fatal(err)
	}
	if c4.Id != "" {
		t.Fatal(c4.Id)
	}
}
