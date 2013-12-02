package braintree

type (
	DiscountList struct {
		XMLName   string    `xml:"discounts"`
		Discounts Discounts `xml:"discount"`
	}

	Discounts []Discount
	Discount  struct {
		Modification
	}
)

func (d Discount) ToSubscriptionDiscount() SubscriptionDiscount {
	return SubscriptionDiscount{
		Id:     d.Id,
		Amount: d.Amount,
	}
}
