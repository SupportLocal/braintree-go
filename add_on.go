package braintree

type (
	AddOnList struct {
		XMLName string `xml:"add-ons"`
		AddOns  AddOns `xml:"add-on"`
	}

	AddOns []AddOn
	AddOn  struct {
		XMLName string `xml:"add-on"`
		Modification
	}
)
