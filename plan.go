package braintree

type (
	Plans []Plan
	Plan  struct {
		XMLName               string  `xml:"plan"`
		Id                    string  `xml:"id"`
		MerchantId            string  `xml:"merchant-id"`
		BillingDayOfMonth     string  `xml:"billing-day-of-month"` // int
		BillingFrequency      string  `xml:"billing-frequency"`    // int
		CurrencyISOCode       string  `xml:"currency-iso-code"`
		Description           string  `xml:"description"`
		Name                  string  `xml:"name"`
		NumberOfBillingCycles string  `xml:"number-of-billing-cycles"` // int
		Price                 float64 `xml:"price"`
		TrialDuration         string  `xml:"trial-duration"` // int
		TrialDurationUnit     string  `xml:"trial-duration-unit"`
		TrialPeriod           string  `xml:"trial-period"` // bool
		CreatedAt             string  `xml:"created-at"`
		UpdatedAt             string  `xml:"updated-at"`
		// AddOns                []interface{} `xml:"add-ons"`
		// Discounts             []interface{} `xml:"discounts"`
	}
	PlanList struct {
		XMLName string `xml:"plans"`
		Plans   []Plan `xml:"plan"`
	}
)
