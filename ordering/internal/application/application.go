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
		command.RejectOrderHandler
		command.ApproveOrderHandler
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
	shoppingClient client.ShoppingClient,
	publisher ddd.EventPublisher[ddd.Event],
) *OrderApplication {
	return &OrderApplication{
		appCommands: appCommands{
			CreateOrderHandler:   command.NewCreateOrderHandler(orderRepository, publisher),
			RejectOrderHandler:   command.NewRejectOrderHandler(orderRepository, publisher),
			ApproveOrderHandler:  command.NewApproveOrderHandler(orderRepository, publisher),
			CancelOrderHandler:   command.NewCancelOrderHandler(orderRepository, shoppingClient, publisher),
			ReadyOrderHandler:    command.NewReadyOrderHandler(orderRepository, publisher),
			CompleteOrderHandler: command.NewCompleteOrderHandler(orderRepository, publisher),
		},
		appQueries: appQueries{
			GetOrderHandler: query.NewGetOrderHandler(orderRepository),
		},
	}
}
