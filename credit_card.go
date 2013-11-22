package braintree

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type (
	CreditCard struct {
		CustomerId                string             `xml:"customer-id,omitempty"`
		Token                     string             `xml:"token,omitempty"`
		Number                    string             `xml:"number,omitempty"`
		ExpirationDate            string             `xml:"expiration-date,omitempty"`
		ExpirationMonth           string             `xml:"expiration-month,omitempty"`
		ExpirationYear            string             `xml:"expiration-year,omitempty"`
		CVV                       string             `xml:"cvv,omitempty"`
		VenmoSDKPaymentMethodCode string             `xml:"venmo-sdk-payment-method-code,omitempty"`
		VenmoSDK                  bool               `xml:"venmo-sdk,omitempty"`
		Options                   *CreditCardOptions `xml:"options,omitempty"`
		CreatedAt                 string             `xml:"created-at,omitempty"`
		UpdatedAt                 string             `xml:"updated-at,omitempty"`
		Bin                       string             `xml:"bin,omitempty"`
		CardType                  string             `xml:"card-type,omitempty"`
		CardholderName            string             `xml:"cardholder-name,omitempty"`
		CustomerLocation          string             `xml:"customer-location,omitempty"`
		ImageURL                  string             `xml:"image-url,omitempty"`
		Default                   bool               `xml:"default,omitempty"`
		Expired                   bool               `xml:"expired,omitempty"`
		Last4                     string             `xml:"last-4,omitempty"`
		Commercial                string             `xml:"commercial,omitempty"`
		Debit                     string             `xml:"debit,omitempty"`
		DurbinRegulated           string             `xml:"durbin-regulated,omitempty"`
		Healthcare                string             `xml:"healthcare,omitempty"`
		Payroll                   string             `xml:"payroll,omitempty"`
		Prepaid                   string             `xml:"prepaid,omitempty"`
		CountryOfIssuance         string             `xml:"country-of-issuance,omitempty"`
		IssuingBank               string             `xml:"issuing-bank,omitempty"`
		UniqueNumberIdentifier    string             `xml:"unique-number-identifier,omitempty"`
		BillingAddress            *Address           `xml:"billing-address,omitempty"`
		Subscriptions             *Subscriptions     `xml:"subscriptions,omitempty"`
		Errors                    CreditCardErrors
	}

	CreditCards struct {
		CreditCard []*CreditCard `xml:"credit-card"`
	}

	CreditCardOptions struct {
		VerifyCard                    bool   `xml:"verify-card,omitempty"`
		VenmoSDKSession               string `xml:"venmo-sdk-session,omitempty"`
		MakeDefault                   bool   `xml:"make-default,omitempty"`
		FailOnDuplicatePaymentMethod  bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
		VerificationMerchantAccountId string `xml:"verification-merchant-account-id,omitempty"`
		UpdateExistingToken           string `xml:"update-existing-token,omitempty"`
	}
)

func (card *CreditCard) Create(gateway Braintree) error {
	url := gateway.MerchantURL() + "/payment_methods"
	req, err := xmlRequest(gateway, "POST", url, card)
	if err != nil {
		return err
	}
	resp, err := doRequest(req)
	if err != nil {
		return err
	}
	xmlBytes, err := unzipRequest(resp)
	if err != nil {
		return err
	}
	if err := xml.Unmarshal(xmlBytes, &card.Errors); err != nil {
		return err
	}
	if card.Errors.Count() == 0 && resp.StatusCode == http.StatusCreated {
		err := xml.Unmarshal(xmlBytes, card)
		return err
	} else {
		return fmt.Errorf("braintree returned invalid response (%d)", resp.StatusCode)
	}
}

// AllSubscriptions returns all subscriptions for this card, or nil if none present.
func (card *CreditCard) AllSubscriptions() []*Subscription {
	if card.Subscriptions != nil {
		subs := card.Subscriptions.Subscription
		if len(subs) > 0 {
			a := make([]*Subscription, 0, len(subs))
			for _, s := range subs {
				a = append(a, s)
			}
			return a
		}
	}
	return nil
}
