package gorm

import (
	"context"
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/ddd"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/gorm"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry/serdes"
	"github.com/stretchr/testify/assert"
)

type testAggregate struct {
	es.AggregateBase
	Name string
}

var _ es.EventSourcedAggregate = (*testAggregate)(nil)

func NewTestAggregate(id, name string) *testAggregate {
	return &testAggregate{
		AggregateBase: es.NewAggregateBase(id, name),
	}
}

func (t *testAggregate) ApplyEvent(evt ddd.Event) error {
	switch e := evt.Payload().(type) {
	case *testChangeName:
		t.Name = e.Name
	}

	return nil
}

type testChangeName struct {
	Name string
}

func (t *testChangeName) Key() string {
	return "test.ChangeName"
}

func TestGormEventStore(t *testing.T) {
	ctx := context.Background()

	db, err := gorm.NewGormDB(&config.Infrastructure{
		GORM: &config.GORM{
			Debug:                   true,
			DBType:                  "postgres",
			DSN:                     "host=localhost user=postgres password=postgres dbname=mallbots port=5432 sslmode=disable TimeZone=Asia/Taipei",
			IgnoreErrRecordNotFound: false,
			ParameterizedQueries:    false,
			Colorful:                true,
			PrepareStmt:             true,
		},
	})
	assert.NoError(t, err)
	db.Exec("CREATE SCHEMA IF NOT EXISTS test")
	db.Exec(`CREATE TABLE IF NOT EXISTS test.events (
		stream_id text NOT NULL,
		stream_name text NOT NULL,
		stream_version int NOT NULL,
		event_id text NOT NULL,
		event_name text NOT NULL,
		event_data bytea NOT NULL,
		occurred_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)

	registry := registry.New()
	jsonSerde := serdes.NewJSONSerde(registry)
	err = jsonSerde.Register(&testChangeName{})
	assert.NoError(t, err)

	eventStore := NewGormEventStore("test.events", db, registry)

	agg := NewTestAggregate("1", "testAggregate")
	agg.AddEvent("test.ChangeName", &testChangeName{Name: "changeTest"})

	for _, event := range agg.Events() {
		_ = agg.ApplyEvent(event)
	}
	err = eventStore.Save(ctx, agg)
	assert.NoError(t, err)

	agg.CommitEvents()

	newAgg := NewTestAggregate("1", "testAggregate")
	err = eventStore.Load(ctx, newAgg)
	assert.NoError(t, err)

	assert.Equal(t, agg.Name, newAgg.Name)

	db.Exec("DROP SCHEMA IF EXISTS test CASCADE")
}
