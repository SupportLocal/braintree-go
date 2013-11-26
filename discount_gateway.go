package braintree

type DiscountGateway struct {
	*Braintree
}

func (g *DiscountGateway) All() (Discounts, error) {
	var discountList DiscountList
	err := g.requestXML("GET", "discounts", nil, &discountList)
	return discountList.Discounts, err
}
