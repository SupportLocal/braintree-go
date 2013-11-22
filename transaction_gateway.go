package braintree

type TransactionGateway struct {
	*Braintree
}

// Create initiates a transaction.
func (g *TransactionGateway) Create(transaction *Transaction) error {
	err := g.requestXML("POST", "transactions", transaction, transaction)
	if err != nil {
		return err
	}
	return nil
}

// SubmitForSettlement submits the transaction with the specified id for settlement.
// If the amount is omitted, the full amount is settled.
func (g *TransactionGateway) SubmitForSettlement(id string, amount ...float64) (Transaction, error) {
	var transaction Transaction
	if len(amount) > 0 {
		transaction = Transaction{
			Amount: amount[0],
		}
	}
	err := g.requestXML("PUT", "transactions/"+id+"/submit_for_settlement", nil, &transaction)
	return transaction, err
}

// Void voids the transaction with the specified id if it has a status of authorized or
// submitted_for_settlement. When the transaction is voided Braintree will do an authorization
// reversal if possible so that the customer won’t have a pending charge on their card
func (g *TransactionGateway) Void(id string) (Transaction, error) {
	var transaction Transaction
	err := g.requestXML("PUT", "transactions/"+id+"/void", nil, &transaction)
	return transaction, err
}

// Find finds the transaction with the specified id.
func (g *TransactionGateway) Find(id string) (Transaction, error) {
	var transaction Transaction
	err := g.requestXML("GET", "transactions/"+id, nil, &transaction)
	return transaction, err
}

// Search finds all transactions matching the search query.
func (g *TransactionGateway) Search(query SearchQuery) (TransactionSearchResult, error) {
	var transactionSearchResult TransactionSearchResult
	err := g.requestXML("POST", "transactions/advanced_search", query, &transactionSearchResult)
	return transactionSearchResult, err
}
