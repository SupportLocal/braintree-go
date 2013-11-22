package braintree

type (
	CreditCardErrors []CreditCardError
	CreditCardError  struct {
		Message string `xml:"message"`
		Errors  []struct {
			Code      string `xml:"code"`
			Attribute string `xml:"attribute"`
			Message   string `xml:"message"`
		} `xml:"errors>credit-card>errors>error"`
	}
)

func (cce CreditCardErrors) Count() int {
	var total int
	for _, e := range cce {
		total += len(e.Errors)
	}
	return total
}
