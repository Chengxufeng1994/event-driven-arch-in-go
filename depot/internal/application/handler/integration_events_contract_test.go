//go:build contract

package handler

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/pact-foundation/pact-go/v2/matchers"
	v4 "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
)

type String = matchers.String
type Map = matchers.Map

var Like = matchers.Like

func TestStoresConsumer(t *testing.T) {
	type mocks struct {
		stores   *repository.MockStoreCacheRepository
		products *repository.MockProductCacheRepository
	}
	type rawEvent struct {
		Name    string
		Payload map[string]any
	}

	reg := registry.New()
	err := storev1.RegistrationsWithSerde(serdes.NewJSONSerde(reg))
	if err != nil {
		t.Fatal(err)
	}

	pact, err := v4.NewAsynchronousPact(v4.Config{
		Provider: "stores-pub",
		Consumer: "depot-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		given   []models.ProviderState
		content Map
		on      func(m mocks)
	}{
		"a StoreCreated message": {
			content: Map{
				"Name": String(storev1.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":       String("store-id"),
					"name":     String("NewStore"),
					"location": String("NewLocation"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore", "NewLocation").Return(nil)
			},
		},
		"a StoreRebranded message": {
			content: Map{
				"Name": String(storev1.StoreRebrandedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("RebrandedStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Rename", mock.Anything, "store-id", "RebrandedStore").Return(nil)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   repository.NewMockStoreCacheRepository(t),
				products: repository.NewMockProductCacheRepository(t),
			}
			if tt.on != nil {
				tt.on(m)
			}

			handlers := NewIntegrationEventHandlers(m.products, m.stores)
			msgHandlerFn := func(m v4.AsynchronousMessage) error {
				event := m.Body.(*rawEvent)

				data, err := json.Marshal(event.Payload)
				if err != nil {
					return err
				}

				payload := reg.MustDeserialize(event.Name, data)

				return handlers.HandleEvent(context.Background(), ddd.NewEvent(event.Name, payload))
			}

			message := pact.AddAsynchronousMessage()
			for _, given := range tt.given {
				message = message.GivenWithParameter(given)
			}

			err := message.
				ExpectsToReceive(name).
				WithJSONContent(tt.content).
				AsType(&rawEvent{}).
				ConsumedBy(msgHandlerFn).
				Verify(t)

			assert.NoError(t, err)
		})
	}
}
