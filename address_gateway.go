package braintree

type AddressGateway struct {
	*Braintree
}

// Create creates a new address for the specified customer id.
func (g *AddressGateway) Create(address *Address) error {
	// TODO: Find out what this really means and determine if this is needed or not
	// Copy address so that field sanitation won't affect original
	err := g.requestXML("POST", "customers/"+address.CustomerId+"/addresses", address, address)
	return err
}

// Delete deletes the address for the specified id and customer id.
func (g *AddressGateway) Delete(customerId, addrId string) error {
	err := g.requestXML("DELETE", "customers/"+customerId+"/addresses/"+addrId, nil, nil)
	return err
}
