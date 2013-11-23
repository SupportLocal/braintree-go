package braintree

type Customer struct {
	Id          string       `xml:"id,omitempty"`
	FirstName   string       `xml:"first-name,omitempty"`
	LastName    string       `xml:"last-name,omitempty"`
	Company     string       `xml:"company,omitempty"`
	Email       string       `xml:"email,omitempty"`
	Phone       string       `xml:"phone,omitempty"`
	Fax         string       `xml:"fax,omitempty"`
	Website     string       `xml:"website,omitempty"`
	CreditCard  *CreditCard  `xml:"credit-card,omitempty"`
	CreditCards *CreditCards `xml:"credit-cards,omitempty"`
	ApiErrors
}

// DefaultCreditCard returns the default credit card, or nil
func (c *Customer) DefaultCreditCard() *CreditCard {
	for _, card := range c.CreditCards.CreditCard {
		if card.Default {
			return card
		}
	}
	return nil
}
