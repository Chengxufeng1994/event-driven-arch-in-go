package internal

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec/internal/models"
	customerv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/customer/api/customer/v1"
	depotv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/depot/api/depot/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/sec"
	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	paymentv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/payment/api/payment/v1"
)

const CreateOrderSagaName = "cosec.CreateOrder"
const CreateOrderReplyChannel = "mallbots.cosec.replies.CreateOrder"

type createOrderSaga struct {
	sec.Saga[*models.CreateOrderData]
}

func NewCreateOrderSaga() sec.Saga[*models.CreateOrderData] {
	saga := createOrderSaga{
		Saga: sec.NewSaga[*models.CreateOrderData](CreateOrderSagaName, CreateOrderReplyChannel),
	}

	// 0. -RejectOrder
	saga.AddStep().
		Compensation(saga.rejectOrder)

	// 1. AuthorizeCustomer
	saga.AddStep().
		Action(saga.authorizeCustomer)

	// 2. CreateShoppingList, -CancelShoppingList
	saga.AddStep().
		Action(saga.createShoppingList).
		OnActionReply(depotv1.CreatedShoppingListReply, saga.onCreatedShoppingListReply).
		Compensation(saga.cancelShoppingList)

	// 3. ConfirmPayment
	saga.AddStep().
		Action(saga.confirmPayment)

	// 4. InitiateShopping
	saga.AddStep().
		Action(saga.initiateShopping)

	// 5. ApproveOrder
	saga.AddStep().
		Action(saga.approveOrder)

	return saga
}

func (s createOrderSaga) rejectOrder(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(orderv1.RejectOrderCommand, orderv1.CommandChannel, &orderv1.RejectOrder{Id: data.OrderID})
}

func (s createOrderSaga) authorizeCustomer(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(customerv1.AuthorizeCustomerCommand, customerv1.CommandChannel, &customerv1.AuthorizeCustomer{Id: data.CustomerID})
}

func (s createOrderSaga) createShoppingList(ctx context.Context, data *models.CreateOrderData) am.Command {
	items := make([]*depotv1.CreateShoppingList_Item, len(data.Items))
	for i, item := range data.Items {
		items[i] = &depotv1.CreateShoppingList_Item{
			ProductId: item.ProductID,
			StoreId:   item.StoreID,
			Quantity:  int32(item.Quantity),
		}
	}

	return am.NewCommand(depotv1.CreateShoppingListCommand, depotv1.CommandChannel, &depotv1.CreateShoppingList{
		OrderId: data.OrderID,
		Items:   items,
	})
}

func (s createOrderSaga) onCreatedShoppingListReply(ctx context.Context, data *models.CreateOrderData, reply ddd.Reply) error {
	payload := reply.Payload().(*depotv1.CreatedShoppingList)

	data.ShoppingID = payload.GetId()

	return nil
}

func (s createOrderSaga) cancelShoppingList(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(depotv1.CancelShoppingListCommand, depotv1.CommandChannel, &depotv1.CancelShoppingList{Id: data.ShoppingID})
}

func (s createOrderSaga) confirmPayment(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(paymentv1.ConfirmPaymentCommand, paymentv1.CommandChannel, &paymentv1.ConfirmPayment{
		Id:     data.PaymentID,
		Amount: data.Total,
	})
}

func (s createOrderSaga) initiateShopping(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(depotv1.InitiateShoppingCommand, depotv1.CommandChannel, &depotv1.InitiateShopping{Id: data.ShoppingID})
}

func (s createOrderSaga) approveOrder(ctx context.Context, data *models.CreateOrderData) am.Command {
	return am.NewCommand(orderv1.ApproveOrderCommand, orderv1.CommandChannel, &orderv1.ApproveOrder{
		Id:         data.OrderID,
		ShoppingId: data.ShoppingID,
	})
}