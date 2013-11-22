package braintree

type CreditCardGateway struct {
	*Braintree
}

func (g *CreditCardGateway) Create(card *CreditCard) error {
	err := g.requestXML("POST", "payment_methods", card, card)
	return err
}

func (g *CreditCardGateway) Update(card *CreditCard) error {
	err := g.requestXML("PUT", "payment_methods/"+card.Token, card, card)
	return err
}

func (g *CreditCardGateway) Find(token string) (CreditCard, error) {
	var card CreditCard
	err := g.requestXML("GET", "payment_methods/"+token, nil, &card)
	return card, err
}

func (g *CreditCardGateway) Delete(card *CreditCard) error {
	err := g.requestXML("DELETE", "payment_methods/"+card.Token, nil, nil)
	return err
}
