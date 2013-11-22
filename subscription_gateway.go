package braintree

type SubscriptionGateway struct {
	*Braintree
}

func (g *SubscriptionGateway) Create(subscription *Subscription) error {
	err := g.requestXML("POST", "subscriptions", subscription, subscription)
	return err
}

func (g *SubscriptionGateway) Update(subscription *Subscription) error {
	err := g.requestXML("PUT", "subscriptions/"+subscription.Id, subscription, subscription)
	return err
}

func (g *SubscriptionGateway) Find(subId string) (Subscription, error) {
	var subscription Subscription
	err := g.requestXML("GET", "subscriptions/"+subId, nil, &subscription)
	return subscription, err
}

func (g *SubscriptionGateway) Cancel(subId string) (Subscription, error) {
	var subscription Subscription
	err := g.requestXML("PUT", "subscriptions/"+subId+"/cancel", nil, &subscription)
	return subscription, err
}
