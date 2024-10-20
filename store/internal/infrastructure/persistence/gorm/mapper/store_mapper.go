package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
)

type StoreMapperIntf interface {
	ToPersistent(store *aggregate.StoreAgg) *po.Store
	ToDomain(store *po.Store) *aggregate.StoreAgg
	ToDomainList(stores []*po.Store) []*aggregate.StoreAgg
}

type StoreMapper struct{}

var _ StoreMapperIntf = (*StoreMapper)(nil)

func NewStoreMapper() *StoreMapper {
	return &StoreMapper{}
}

func (m *StoreMapper) ToPersistent(store *aggregate.StoreAgg) *po.Store {
	return &po.Store{
		ID:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (m *StoreMapper) ToDomain(store *po.Store) *aggregate.StoreAgg {
	return &aggregate.StoreAgg{
		ID:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}
}

func (m *StoreMapper) ToDomainList(stores []*po.Store) []*aggregate.StoreAgg {
	var result []*aggregate.StoreAgg
	for _, store := range stores {
		result = append(result, m.ToDomain(store))
	}
	return result
}
