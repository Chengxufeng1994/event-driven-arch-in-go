package event

const (
	CustomerRegisteredEvent = "customers.CustomerRegistered"
	CustomerAuthorizedEvent = "customers.CustomerAuthorized"
	CustomerEnabledEvent    = "customers.CustomerEnabled"
	CustomerDisabledEvent   = "customers.CustomerDisabled"
)

type CustomerRegistered struct {
	CustomerID string
}

func NewCustomerRegistered(customerID string) *CustomerRegistered {
	return &CustomerRegistered{CustomerID: customerID}
}

func (CustomerRegistered) EventName() string { return CustomerRegisteredEvent }

type CustomerAuthorized struct {
	CustomerID string
}

func NewCustomerAuthorized(customerID string) *CustomerAuthorized {
	return &CustomerAuthorized{CustomerID: customerID}
}

func (CustomerAuthorized) EventName() string { return CustomerAuthorizedEvent }

type CustomerEnabled struct {
	CustomerID string
}

func NewCustomerEnabled(customerID string) *CustomerEnabled {
	return &CustomerEnabled{CustomerID: customerID}
}

func (CustomerEnabled) EventName() string { return CustomerEnabledEvent }

type CustomerDisabled struct {
	CustomerID string
}

func NewCustomerDisabled(customerID string) *CustomerDisabled {
	return &CustomerDisabled{CustomerID: customerID}
}

func (CustomerDisabled) EventName() string { return CustomerDisabledEvent }
