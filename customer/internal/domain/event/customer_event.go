package event

const (
	CustomerRegisteredEvent = "customers.CustomerRegistered"
	CustomerSmsChangedEvent = "customers.CustomerSmsChanged"
	CustomerAuthorizedEvent = "customers.CustomerAuthorized"
	CustomerEnabledEvent    = "customers.CustomerEnabled"
	CustomerDisabledEvent   = "customers.CustomerDisabled"
)

type CustomerRegistered struct {
	Name      string
	SmsNumber string
	Enabled   bool
}

func NewCustomerRegistered(name, smsNumber string, enabled bool) *CustomerRegistered {
	return &CustomerRegistered{
		Name:      name,
		SmsNumber: smsNumber,
		Enabled:   enabled,
	}
}

func (CustomerRegistered) Key() string { return CustomerRegisteredEvent }

type CustomerSmsChanged struct {
	SmsNumber string
}

func NewCustomerSmsChanged(smsNumber string) *CustomerSmsChanged {
	return &CustomerSmsChanged{
		SmsNumber: smsNumber,
	}
}

func (CustomerSmsChanged) Key() string { return CustomerSmsChangedEvent }

type CustomerAuthorized struct{}

func NewCustomerAuthorized() *CustomerAuthorized {
	return &CustomerAuthorized{}
}

func (CustomerAuthorized) Key() string { return CustomerAuthorizedEvent }

type CustomerEnabled struct{}

func NewCustomerEnabled() *CustomerEnabled {
	return &CustomerEnabled{}
}

func (CustomerEnabled) Key() string { return CustomerEnabledEvent }

type CustomerDisabled struct{}

func NewCustomerDisabled() *CustomerDisabled {
	return &CustomerDisabled{}
}

func (CustomerDisabled) Key() string { return CustomerDisabledEvent }
