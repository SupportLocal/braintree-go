package braintree

import (
	"net/http"
	"testing"
)

func TestDiscountFind(t *testing.T) {
	server := newServer(func(w http.ResponseWriter, r *http.Request) {
		if err := serveRecording(w, r, "discount", "find", http.StatusOK); err != nil {
			panic(err)
		}
	})
	defer server.Close()

	gw := Braintree{BaseURL: server.URL}

	discount, err := gw.Discount().Find("test_discount")

	t.Log(discount)

	if err != nil {
		t.Fatal(err)
	}

	if discount.Id != "test_discount" {
		t.Fatalf("Expected id of %q, got %q", "test_discount", discount.Id)
	}

	if discount.Amount != 10 {
		t.Fatalf("Expected id of %d, got %d", 10, discount.Amount)
	}

}
