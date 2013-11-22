package braintree

type CustomerGateway struct {
	*Braintree
}

// Create creates a new customer from the passed in customer object.
// If no Id is set, Braintree will assign one.
func (g *CustomerGateway) Create(customer *Customer) error {
	err := g.requestXML("POST", "customers", customer, customer)
	return err
}

// Update updates any field that is set in the passed customer object.
// The Id field is mandatory.
func (g *CustomerGateway) Update(customer *Customer) error {
	err := g.requestXML("PUT", "customers/"+customer.Id, customer, customer)
	return err
}

// Find finds the customer with the given id.
func (g *CustomerGateway) Find(id string) (Customer, error) {
	var customer Customer
	err := g.requestXML("GET", "customers/"+id, nil, &customer)
	return customer, err
}

// Delete deletes the customer with the given id.
func (g *CustomerGateway) Delete(id string) error {
	err := g.requestXML("DELETE", "customers/"+id, nil, nil)
	return err
}
