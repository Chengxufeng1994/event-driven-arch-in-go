package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/es"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
)

type StoreMapperIntf interface {
	ToPersistent(store *aggregate.Store) *po.Store
	ToDomain(store *po.Store) *aggregate.Store
	ToDomainList(stores []*po.Store) []*aggregate.Store
}

type StoreMapper struct{}

var _ StoreMapperIntf = (*StoreMapper)(nil)

func NewStoreMapper() *StoreMapper {
	return &StoreMapper{}
}

func (m *StoreMapper) ToPersistent(store *aggregate.Store) *po.Store {
	return &po.Store{
		ID:            store.ID(),
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (m *StoreMapper) ToDomain(store *po.Store) *aggregate.Store {
	return &aggregate.Store{
		AggregateBase: es.NewAggregateBase(store.ID, aggregate.StoreAggregate),
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (m *StoreMapper) ToDomainList(stores []*po.Store) []*aggregate.Store {
	var result []*aggregate.Store
	for _, store := range stores {
		result = append(result, m.ToDomain(store))
	}
	return result
}
