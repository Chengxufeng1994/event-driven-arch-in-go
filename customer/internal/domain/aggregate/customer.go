package aggregate

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/internal/domain/event"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/stackus/errors"
)

const CustomerAggregate = "customers.CustomerAggregate"

var (
	ErrNameCannotBeBlank       = errors.Wrap(errors.ErrBadRequest, "the customer name cannot be blank")
	ErrCustomerIDCannotBeBlank = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
	ErrSmsNumberCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the SMS number cannot be blank")
	ErrCustomerAlreadyEnabled  = errors.Wrap(errors.ErrBadRequest, "the customer is already enabled")
	ErrCustomerAlreadyDisabled = errors.Wrap(errors.ErrBadRequest, "the customer is already disabled")
	ErrCustomerNotAuthorized   = errors.Wrap(errors.ErrUnauthorized, "customer is not authorized")
)

type Customer struct {
	ddd.AggregateBase
	Name      string
	SmsNumber string
	Enabled   bool
}

func NewCustomer(id string) *Customer {
	return &Customer{
		AggregateBase: ddd.NewAggregateBase(id, CustomerAggregate),
	}
}

func RegisterCustomer(id, name, smsNumber string) (*Customer, error) {
	if id == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	if name == "" {
		return nil, ErrNameCannotBeBlank
	}

	if smsNumber == "" {
		return nil, ErrSmsNumberCannotBeBlank
	}

	customer := NewCustomer(id)
	customer.Name = name
	customer.SmsNumber = smsNumber
	customer.Enabled = true

	customer.AddEvent(event.CustomerRegisteredEvent, event.NewCustomerRegistered(id, name, smsNumber, true))

	return customer, nil
}

func (Customer) Key() string { return CustomerAggregate }

func (c *Customer) Authorize( /* TODO authorize what? */ ) error {
	if !c.Enabled {
		return ErrCustomerNotAuthorized
	}

	c.AddEvent(event.CustomerAuthorizedEvent, event.NewCustomerAuthorized(c.ID()))

	return nil
}

func (c *Customer) Enable() error {
	if c.Enabled {
		return ErrCustomerAlreadyEnabled
	}

	c.Enabled = true

	c.AddEvent(event.CustomerEnabledEvent, event.NewCustomerEnabled(c.ID()))

	return nil
}

func (c *Customer) Disable() error {
	if !c.Enabled {
		return ErrCustomerAlreadyDisabled
	}

	c.Enabled = false

	c.AddEvent(event.CustomerEnabledEvent, event.NewCustomerEnabled(c.ID()))

	return nil
}
