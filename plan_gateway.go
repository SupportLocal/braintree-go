package braintree

type PlanGateway struct {
	*Braintree
}

// All returns all available plans
func (g *PlanGateway) All() (Plans, error) {
	var planList PlanList
	err := g.requestXML("GET", "plans", nil, &planList)
	return planList.Plans, err
}

// Find returns the plan with the specified id, or nil
func (g *PlanGateway) Find(id string) (Plan, error) {
	plans, err := g.All()
	if err != nil {
		return Plan{}, err
	}
	for _, p := range plans {
		if p.Id == id {
			return p, nil
		}
	}
	return Plan{}, nil
}
