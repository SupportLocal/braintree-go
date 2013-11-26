package braintree

type DiscountGateway struct {
	*Braintree
}

func (g *DiscountGateway) All() (Discounts, error) {
	var discountList DiscountList
	err := g.requestXML("GET", "discounts", nil, &discountList)
	return discountList.Discounts, err
}

func (g *DiscountGateway) Find(id string) (Discount, error) {
	discounts, err := g.All()
	if err != nil {
		return Discount{}, err
	}
	for _, d := range discounts {
		if d.Id == id {
			return d, nil
		}
	}
	return Discount{}, nil
}
