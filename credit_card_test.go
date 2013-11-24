package braintree

import (
	"fmt"
	"net/http"
	"testing"
)

var testCreditCard = CreditCard{
	CustomerId:     "test_customer",
	Number:         testCreditCards["visa"].Number,
	ExpirationDate: "05/14",
	CVV:            "100",
	Options: &CreditCardOptions{
		VerifyCard: true,
	},
}

func testCreditCardCreate(credit_card *CreditCard) error {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", fmt.Sprintf("create_test_credit_card_%s_%s", credit_card.CustomerId, credit_card.Number), http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	return gw.CreditCard().Create(credit_card)
}

func testCreditCardDelete(token string) error {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "delele_test_credit_card_"+token, http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	return gw.CreditCard().Delete(token)
}

func TestCreditCardCreate(t *testing.T) {
	customer := testCustomer
	if err := testCustomerCreate(&customer); err != nil {
		t.Fatal(err, "Unable to set up test customer")
	}

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "create", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()
	card := CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}
	err := g.Create(&card)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card)

	if card.Token == "" {
		t.Fatal("invalid token")
	}

	_ = testCustomerDelete(customer.Id)
}

func TestCreditCardUpdate(t *testing.T) {
	customer := testCustomer
	customer.Id = "cc_update"
	_ = testCustomerDelete(customer.Id)
	_ = testCustomerCreate(&customer)
	testCard := testCreditCard
	testCard.CustomerId = customer.Id
	if err := testCreditCardCreate(&testCard); err != nil {
		t.Fatal(err, "Unable to set up test credit card")
	}

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "update", http.StatusCreated); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	card := CreditCard{
		Token:          testCard.Token,
		Number:         testCreditCards["mastercard"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	err := g.Update(&card)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(card)

	if card.Token != testCard.Token {
		t.Fatal(card.Token, "Card token should be "+testCard.Token)
	}
	if card.CardType != "MasterCard" {
		t.Fatal(card.CardType, "Card type should be MasterCard")
	}
	if !card.Success() {
		t.Fatal(card.Success(), "card.Success() should be true (no errors)")
	}

	_ = testCreditCardDelete(card.Token)
	_ = testCustomerDelete(customer.Id)
}

func TestCreditCardDelete(t *testing.T) {
	customer := testCustomer
	customer.Id = "cc_credit_card_delete"
	_ = testCustomerDelete(customer.Id)
	_ = testCustomerCreate(&customer)
	testCard := CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["discover"].Number,
		ExpirationDate: "05/14",
		CVV:            "100",
		Options: &CreditCardOptions{
			VerifyCard: true,
		},
	}

	if err := testCreditCardCreate(&testCard); err != nil {
		t.Fatal(err, "Failed to set up test credit card to delete")
	}
	t.Log(testCard)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "delete", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	err := g.Delete(testCard.Token)
	if err != nil {
		t.Fatal(err)
	}

	_ = testCustomerDelete(customer.Id)
}

func TestCreateCreditCardWithExpirationMonthAndYear(t *testing.T) {
	customer := testCustomer
	customer.Id = "cc_with_exp_m_v"
	_ = testCustomerDelete(customer.Id)
	_ = testCustomerCreate(&customer)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "create_with_exp_m_v", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card := CreditCard{
		CustomerId:      customer.Id,
		Number:          testCreditCards["visa"].Number,
		ExpirationMonth: "05",
		ExpirationYear:  "2014",
		CVV:             "100",
	}
	err := g.Create(&card)

	if err != nil {
		t.Fatal(err)
	}
	if card.Token == "" {
		t.Fatal("invalid token")
	}

	_ = testCustomerDelete(customer.Id)
}

func TestCreateCreditCardInvalidInput(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "create_with_inv_exp_m_v", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card := CreditCard{
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
	}
	_ = g.Create(&card)

	t.Log(card)

	// This test should fail because customer id is required
	if card.Success() {
		t.Fatal(card.ErrorMessage)
	}
	if card.ErrorCount() != 1 {
		t.Fatal(card.ErrorCount(), "Error count should be 1")
	}
	expectedError := "Customer ID is required."
	if card.ErrorMessage != expectedError {
		t.Fatal(fmt.Sprintf("Error should be %q, got %q", expectedError, card.ErrorMessage))
	}
}

func TestFindCreditCard(t *testing.T) {
	customer := testCustomer
	customer.Id = "cc_find"
	_ = testCustomerCreate(&customer)
	testCard := testCreditCard
	testCard.CustomerId = customer.Id
	_ = testCreditCardCreate(&testCard)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "find", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card, err := g.Find(testCard.Token)

	t.Log(card)

	if err != nil {
		t.Fatal(err)
	}
	if card.Token != testCard.Token {
		t.Fatal("tokens do not match")
	}
	if card.CardType != "Visa" {
		t.Fatal(fmt.Sprintf("CardType should be %q, got %q", card.CardType, "Visa"))
	}

	_ = testCustomerDelete(customer.Id)
}

func TestFindCreditCard404(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "find_404", http.StatusNotFound); err != nil {
			panic(err)
		}
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
	customer := testCustomer
	customer.Id = "cc_w_venmo_sdk_code"
	_ = testCustomerCreate(&customer)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "create_with_vemo_sdk_code", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card := CreditCard{
		CustomerId:                customer.Id,
		VenmoSDKPaymentMethodCode: "stub-" + testCreditCards["visa"].Number,
	}
	err := g.Create(&card)
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}
	_ = testCustomerDelete(customer.Id)
}

func TestSaveCreditCardWithVenmoSDKSession(t *testing.T) {
	customer := testCustomer
	customer.Id = "cc_w_venmo_sdk_sesson"
	_ = testCustomerCreate(&customer)

	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "credit_card", "create_with_vemo_sdk_session", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	g := gw.CreditCard()

	card := CreditCard{
		CustomerId:     customer.Id,
		Number:         testCreditCards["visa"].Number,
		ExpirationDate: "05/14",
		Options: &CreditCardOptions{
			VenmoSDKSession: "stub-session",
		},
	}
	err := g.Create(&card)
	if err != nil {
		t.Fatal(err)
	}
	if !card.VenmoSDK {
		t.Fatal("venmo card not marked")
	}

	_ = testCustomerDelete(customer.Id)
}
