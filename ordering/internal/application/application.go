package application

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/in/query"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/port/out/client"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/internal/domain/repository"
)

type (
	OrderApplication struct {
		appCommands
		appQueries
	}

	appCommands struct {
		command.CreateOrderHandler
		command.CancelOrderHandler
		command.ReadyOrderHandler
		command.CompleteOrderHandler
	}

	appQueries struct {
		query.GetOrderHandler
	}
)

var _ usecase.OrderUseCase = (*OrderApplication)(nil)

func NewOrderApplication(
	orderRepository repository.OrderRepository,
	customerClient client.CustomerClient,
	paymentClient client.PaymentClient,
	shoppingClient client.ShoppingClient,
	domainEventPublisher ddd.EventPublisher,
) *OrderApplication {
	return &OrderApplication{
		appCommands: appCommands{
			CreateOrderHandler:   command.NewCreateOrderHandler(orderRepository, customerClient, paymentClient, shoppingClient, domainEventPublisher),
			CancelOrderHandler:   command.NewCancelOrderHandler(orderRepository, shoppingClient, domainEventPublisher),
			ReadyOrderHandler:    command.NewReadyOrderHandler(orderRepository, domainEventPublisher),
			CompleteOrderHandler: command.NewCompleteOrderHandler(orderRepository),
		},
		appQueries: appQueries{
			GetOrderHandler: query.NewGetOrderHandler(
				orderRepository,
			),
		},
	}
}
