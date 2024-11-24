//go:build contract

package handler

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	storev1 "github.com/Chengxufeng1994/event-driven-arch-in-go/store/api/store/v1"
	"github.com/pact-foundation/pact-go/v2/matchers"
	v4 "github.com/pact-foundation/pact-go/v2/message/v4"
	"github.com/pact-foundation/pact-go/v2/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
		Consumer: "baskets-sub",
		PactDir:  "./pacts",
	})
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]struct {
		given    []models.ProviderState
		metadata map[string]string
		content  Map
		on       func(m mocks)
	}{
		"a StoreCreated message": {
			metadata: map[string]string{
				"subject": storev1.StoreAggregateChannel,
			},
			content: Map{
				"Name": String(storev1.StoreCreatedEvent),
				"Payload": Like(Map{
					"id":   String("store-id"),
					"name": String("NewStore"),
				}),
			},
			on: func(m mocks) {
				m.stores.On("Add", mock.Anything, "store-id", "NewStore").Return(nil)
			},
		},
		"a StoreRebranded message": {
			metadata: map[string]string{
				"subject": storev1.StoreAggregateChannel,
			},
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

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			m := mocks{
				stores:   repository.NewMockStoreCacheRepository(t),
				products: repository.NewMockProductCacheRepository(t),
			}
			handlers := NewIntegrationEventHandlers(m.stores, m.products)
			if tc.on != nil {
				tc.on(m)
			}

			fn := func(m v4.AsynchronousMessage) error {
				event := m.Body.(*rawEvent)

				data, err := json.Marshal(event.Payload)
				if err != nil {
					panic(err)
				}
				payload := reg.MustDeserialize(event.Name, data)

				return handlers.HandleEvent(
					context.Background(),
					ddd.NewEvent(event.Name, payload),
				)
			}

			message := pact.AddAsynchronousMessage()
			for _, given := range tc.given {
				message = message.GivenWithParameter(given)
			}
			err := message.
				ExpectsToReceive(name).
				WithMetadata(tc.metadata).
				WithJSONContent(tc.content).
				AsType(&rawEvent{}).
				ConsumedBy(fn).
				Verify(t)

			assert.NoError(t, err)
		})
	}
}
