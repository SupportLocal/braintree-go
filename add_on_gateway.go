package braintree

type AddOnGateway struct {
	*Braintree
}

func (g *AddOnGateway) All() (AddOns, error) {
	var addOnList AddOnList
	err := g.requestXML("GET", "add_ons", nil, &addOnList)
	return addOnList.AddOns, err
}
