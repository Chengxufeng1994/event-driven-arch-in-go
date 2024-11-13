package handler

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/am"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/application/port/in/command"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/pact-foundation/pact-go/v2/message"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/pact-foundation/pact-go/v2/provider"
)

func TestStoresProducer(t *testing.T) {
	var err error

	stores := repository.NewFakeStoreRepository()
	products := repository.NewFakeProductRepository()
	mall := repository.NewFakeMallRepository()
	catalog := repository.NewFakeCatalogRepository()

	type rawEvent struct {
		Name    string
		Payload json.RawMessage
	}

	reg := registry.New()
	err = storev1.RegistrationsWithSerde(serdes.NewJSONSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	verifier := provider.NewVerifier()
	err = verifier.VerifyProvider(t, provider.VerifyRequest{
		Provider:                   "stores-pub",
		ProviderVersion:            "1.0.0",
		BrokerURL:                  "http://127.0.0.1:9292",
		BrokerUsername:             "pact_workshop",
		BrokerPassword:             "pact_workshop",
		PublishVerificationResults: true,
		AfterEach: func() error {
			stores.Reset()
			products.Reset()
			return nil
		},
		MessageHandlers: map[string]message.Handler{
			"a StoreCreated message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				// Assign
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.NewStoreApplication(stores, products, mall, catalog, dispatcher)
				publisher := am.NewFakeMessagePublisher[ddd.Event]()
				handler := NewDomainEventHandler(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				// Act
				err := app.CreateStore(context.Background(), command.CreateStore{
					ID:       "store-id",
					Name:     "NewStore",
					Location: "NewLocation",
				})
				if err != nil {
					return nil, nil, err
				}

				// Assert
				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
			"a StoreRebranded message": func(states []models.ProviderState) (message.Body, message.Metadata, error) {
				dispatcher := ddd.NewEventDispatcher[ddd.Event]()
				app := application.NewStoreApplication(stores, products, mall, catalog, dispatcher)
				publisher := am.NewFakeMessagePublisher[ddd.Event]()
				handler := NewDomainEventHandler(publisher)
				RegisterDomainEventHandlers(dispatcher, handler)

				store := aggregate.NewStore("store-id")
				store.Name = "NewStore"
				store.Location = "NewLocation"
				stores.Reset(store)

				err := app.RebrandStore(context.Background(), command.RebrandStore{
					ID:   "store-id",
					Name: "RebrandedStore",
				})
				if err != nil {
					return nil, nil, err
				}

				subject, event, err := publisher.Last()
				if err != nil {
					return nil, nil, err
				}

				return rawEvent{
						Name:    event.EventName(),
						Payload: reg.MustSerialize(event.EventName(), event.Payload()),
					}, map[string]any{
						"subject": subject,
					}, nil
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}