package braintree

import "fmt"

const (
	SubscriptionStatusActive       = "Active"
	SubscriptionStatusCanceled     = "Canceled"
	SubscriptionStatusExpired      = "Expired"
	SubscriptionStatusPastDue      = "Past Due"
	SubscriptionStatusPending      = "Pending"
	SubscriptionStatusUnrecognized = "Unrecognized"
)

type (
	Subscriptions []Subscription
	Subscription  struct {
		Id                      string                      `xml:"id,omitempty"`
		Balance                 float64                     `xml:"balance,omitempty"`
		BillingDayOfMonth       string                      `xml:"billing-day-of-month,omitempty"`
		BillingPeriodEndDate    string                      `xml:"billing-period-end-date,omitempty"`
		BillingPeriodStartDate  string                      `xml:"billing-period-start-date,omitempty"`
		CurrentBillingCycle     string                      `xml:"current-billing-cycle,omitempty"`
		DaysPastDue             string                      `xml:"days-past-due,omitempty"`
		SubscriptionDiscounts   *SubscriptionDiscountObject `xml:"discounts,omitempty"`
		FailureCount            string                      `xml:"failure-count,omitempty"`
		FirstBillingDate        string                      `xml:"first-billing-date,omitempty"`
		MerchantAccountId       string                      `xml:"merchant-account-id,omitempty"`
		NeverExpires            string                      `xml:"never-expires,omitempty"` // bool
		NextBillAmount          float64                     `xml:"next-bill-amount,omitempty"`
		NextBillingPeriodAmount float64                     `xml:"next-billing-period-amount,omitempty"`
		NextBillingDate         string                      `xml:"next-billing-date,omitempty"`
		NumberOfBillingCycles   string                      `xml:"number-of-billing-cycles,omitempty"` // int
		PaidThroughDate         string                      `xml:"paid-through-date,omitempty"`
		PaymentMethodToken      string                      `xml:"payment-method-token,omitempty"`
		PlanId                  string                      `xml:"plan-id,omitempty"`
		Price                   float64                     `xml:"price,omitempty"`
		Status                  string                      `xml:"status,omitempty"`
		TrialDuration           string                      `xml:"trial-duration,omitempty"`
		TrialDurationUnit       string                      `xml:"trial-duration-unit,omitempty"`
		TrialPeriod             string                      `xml:"trial-period,omitempty"` // bool
		Transactions            *Transactions               `xml:"transactions,omitempty"`
		Options                 *SubscriptionOptions        `xml:"options,omitempty"`
		// AddOns                  []interface{} `xml:"add-ons,omitempty"`
		// Descriptor              interface{}   `xml:"descriptor,omitempty"`   // struct with name, phone
		ApiErrors
	}

	SubscriptionDiscountObject struct {
		Add *SubscriptionDiscountItems `xml:"add,omitempty"`
	}

	SubscriptionDiscountItems []SubscriptionDiscountItem
	SubscriptionDiscountItem  struct {
		Item *SubscriptionDiscounts `xml:"item,omitempty"`
		Type string                 `xml:"type,attr,omitempty"`
	}

	SubscriptionDiscounts []SubscriptionDiscount
	SubscriptionDiscount  struct {
		Id            string  `xml:"inherited-from-id,omitempty"`
		Amount        float64 `xml:"amount,omitempty"`
		BillingCycles string  `xml:"number-of-billing-cycles,omitempty"`
	}
	SubscriptionList struct {
		XMLName       string        `xml:"subscriptions"`
		Subscriptions Subscriptions `xml:"subscription"`
	}

	// TODO(eaigner): same considerations apply as with plan type marshalling

	SubscriptionOptions struct {
		DoNotInheritAddOnsOrDiscounts        bool `xml:"do-not-inherit-add-ons-or-discounts,omitempty"`
		ProrateCharges                       bool `xml:"prorate-charges,omitempty"`
		ReplaceAllAddOnsAndDiscounts         bool `xml:"replace-all-add-ons-and-discounts,omitempty"`
		RevertSubscriptionOnProrationFailure bool `xml:"revert-subscription-on-proration-failure,omitempty"`
		StartImmediately                     bool `xml:"start-immediately,omitempty"`
	}
)

func (s *Subscription) AddDiscount(discount Discount, billingCycles int) {
	if s.SubscriptionDiscounts == nil {
		s.SubscriptionDiscounts = &SubscriptionDiscountObject{}
	}
	if s.SubscriptionDiscounts.Add == nil {
		s.SubscriptionDiscounts.Add = &SubscriptionDiscountItems{}
	}
	item := SubscriptionDiscountItem{}
	item.Type = "array"
	sds := SubscriptionDiscounts{}
	sd := discount.ToSubscriptionDiscount()
	sd.BillingCycles = fmt.Sprintf("%d", billingCycles)

	sds = append(sds, sd)
	item.Item = &sds
	*s.SubscriptionDiscounts.Add = append(*s.SubscriptionDiscounts.Add, item)
}
