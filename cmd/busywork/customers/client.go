package customers

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/customersclient"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/customersclient/customer"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer/customersclient/models"
)

type Client interface {
	RegisterCustomer(ctx context.Context, name, smsNumber string) (string, error)
}

type client struct {
	c *customersclient.Customers
}

func NewClient(transport runtime.ClientTransport) Client {
	return &client{
		c: customersclient.New(transport, strfmt.Default),
	}
}

func (c *client) RegisterCustomer(ctx context.Context, name, smsNumber string) (string, error) {
	resp, err := c.c.Customer.RegisterCustomer(&customer.RegisterCustomerParams{
		Body: &models.V1RegisterCustomerRequest{
			Name:      name,
			SmsNumber: smsNumber,
		},
		Context: ctx,
	})
	if err != nil {
		return "", err
	}

	return resp.GetPayload().ID, nil
}
