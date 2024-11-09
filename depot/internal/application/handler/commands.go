package handler

import (
	"context"

	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/application/usecase"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/valueobject"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/google/uuid"
)

type CommandHandlers[T ddd.Command] struct {
	app usecase.ShoppingListUseCase
}

var _ ddd.CommandHandler[ddd.Command] = (*CommandHandlers[ddd.Command])(nil)

func RegisterCommandHandlers(subscriber am.CommandSubscriber, handlers ddd.CommandHandler[ddd.Command]) error {
	cmdMsgHandler := am.CommandMessageHandlerFunc(func(ctx context.Context, cmdMsg am.IncomingCommandMessage) (ddd.Reply, error) {
		return handlers.HandleCommand(ctx, cmdMsg)
	})

	return subscriber.Subscribe(depotv1.CommandChannel, cmdMsgHandler, am.MessageFilter{
		depotv1.CreateShoppingListCommand,
		depotv1.CancelShoppingListCommand,
		depotv1.InitiateShoppingCommand,
	}, am.GroupName("depot-commands"))
}

func NewCommandHandlers(app usecase.ShoppingListUseCase) ddd.CommandHandler[ddd.Command] {
	return &CommandHandlers[ddd.Command]{
		app: app,
	}
}

func (h *CommandHandlers[T]) HandleCommand(ctx context.Context, cmd T) (ddd.Reply, error) {
	switch cmd.CommandName() {
	case depotv1.CreateShoppingListCommand:
		return h.doCreateShoppingList(ctx, cmd)
	case depotv1.CancelShoppingListCommand:
		return h.doCancelShoppingList(ctx, cmd)
	case depotv1.InitiateShoppingCommand:
		return h.doInitiateShopping(ctx, cmd)
	}

	return nil, nil
}

func (h CommandHandlers[T]) doCreateShoppingList(ctx context.Context, cmd T) (ddd.Reply, error) {
	payload := cmd.Payload().(*depotv1.CreateShoppingList)

	id := uuid.New().String()

	items := make([]valueobject.OrderItem, 0, len(payload.GetItems()))
	for _, item := range payload.GetItems() {
		items = append(items, valueobject.OrderItem{
			StoreID:   item.GetStoreId(),
			ProductID: item.GetProductId(),
			Quantity:  int(item.GetQuantity()),
		})
	}

	err := h.app.CreateShoppingList(ctx, command.CreateShoppingList{
		ID:      id,
		OrderID: payload.GetOrderId(),
		Items:   items,
	})

	return ddd.NewReply(depotv1.CreatedShoppingListReply, &depotv1.CreatedShoppingList{Id: id}), err
}

func (h CommandHandlers[T]) doCancelShoppingList(ctx context.Context, cmd T) (ddd.Reply, error) {
	payload := cmd.Payload().(*depotv1.CancelShoppingList)

	err := h.app.CancelShoppingList(ctx, command.CancelShoppingList{ID: payload.GetId()})
	// returning nil returns a simple Success or Failure reply; err being nil determines which
	return nil, err
}

func (h CommandHandlers[T]) doInitiateShopping(ctx context.Context, cmd T) (ddd.Reply, error) {
	payload := cmd.Payload().(*depotv1.InitiateShopping)

	err := h.app.InitiateShopping(ctx, command.InitiateShopping{ID: payload.GetId()})

	// returning nil returns a simple Success or Failure reply; err being nil determines which
	return nil, err
}
