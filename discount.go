package braintree

type (
	DiscountList struct {
		XMLName   string    `xml:"discounts"`
		Discounts Discounts `xml:"discount"`
	}

	Discounts []Discount
	Discount  struct {
		XMLName string `xml:"discount"`
		Modification
	}
)
