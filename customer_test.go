package braintree

import (
	"fmt"
	"net/http"
	"testing"
)

var testCustomer = Customer{
	Id:        "test_customer",
	FirstName: "Jim",
	LastName:  "Bean",
}

func testCustomerFindOrCreate(customer Customer) (Customer, error) {
	if c, err := testCustomerFind(customer.Id); err != nil {
		err = testCustomerCreate(&c)
		return c, err
	} else {
		return c, nil
	}
}

func testCustomerFind(customerId string) (Customer, error) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "find_test_customer_"+customerId, http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	return gw.Customer().Find(customerId)
}

func testCustomerCreate(customer *Customer) error {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "create_test_customer_"+customer.Id, http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	return gw.Customer().Create(customer)
}

func testCustomerDelete(customerId string) error {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "delele_test_customer_"+customerId, http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	return gw.Customer().Delete(customerId)
}

func TestCustomerCreateWithCVVError(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "create_with_cvv_error", 422); err != nil {
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
	if err == nil {
		t.Fatal(err, "Should have an error")
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
	// Delete it and fail if it isn't there.  Never know what state the sandbox could be in
	_ = testCustomerDelete("test_customer_update")
	testCustomer := testCustomer
	testCustomer.Id = "test_customer_update"
	if err := testCustomerCreate(&testCustomer); err != nil {
		t.Fatal(err, "Unable to set up test customer")
	}

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "update", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	customer := Customer{
		Id:        testCustomer.Id,
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

	_ = testCustomerDelete("test_customer_update")
}

func TestCustomerFind(t *testing.T) {
	// Delete it and fail if it isn't there.  Never know what state the sandbox could be in
	_ = testCustomerDelete("test_customer_find")
	testCustomer := testCustomer
	testCustomer.Id = "test_customer_find"

	if err := testCustomerCreate(&testCustomer); err != nil {
		t.Fatal(err, "Unable to set up test customer")
	}

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "find", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c, err := gw.Customer().Find("test_customer_find")

	t.Log(c)

	if err != nil {
		t.Fatal(err)
	}
	if c.Id != "test_customer_find" {
		t.Fatal("ids do not match")
	}

	// Clean up sandbox if needed
	_ = testCustomerDelete("test_customer_find")
}

func TestCustomerDelete(t *testing.T) {
	// Delete it and fail if it isn't there.  Never know what state the sandbox could be in
	_ = testCustomerDelete("test_customer_delete")
	testCustomer := testCustomer
	testCustomer.Id = "test_customer_delete"

	if err := testCustomerCreate(&testCustomer); err != nil {
		t.Fatalf("Unable to delete test customer %s\n%s", testCustomer.Id, err)
	}

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "delete", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	err := gw.Customer().Delete("test_customer_delete")
	if err != nil {
		t.Fatal(err)
	}
	_ = testCustomerDelete("test_customer_delete")
}

func TestCustomerFind404(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "customer", "find404", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	c4, err := gw.Customer().Find("test_customer_find_404")
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
