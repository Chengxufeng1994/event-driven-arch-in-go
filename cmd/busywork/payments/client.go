package payments

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/paymentsclient"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/paymentsclient/models"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/paymentsclient/payment"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type Client interface {
	AuthorizePayment(ctx context.Context, customerID string, amount float64) (string, error)
}

type client struct {
	c *paymentsclient.APIPaymentV1MessagesProto
}

func NewClient(transport runtime.ClientTransport) Client {
	return &client{
		c: paymentsclient.New(transport, strfmt.Default),
	}
}

func (c *client) AuthorizePayment(ctx context.Context, customerID string, amount float64) (string, error) {
	resp, err := c.c.Payment.AuthorizePayment(&payment.AuthorizePaymentParams{
		Body: &models.V1AuthorizePaymentRequest{
			Amount:     amount,
			CustomerID: customerID,
		},
		Context: ctx,
	})
	if err != nil {
		return "", err
	}
	return resp.GetPayload().ID, nil
}
