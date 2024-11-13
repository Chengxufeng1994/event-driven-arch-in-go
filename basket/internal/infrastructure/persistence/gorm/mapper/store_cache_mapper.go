package mapper

import (
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/domain/entity"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket/internal/infrastructure/persistence/gorm/po"
)

type StoreCacheMapperIntf interface {
	ToPersistence(store entity.Store) (po.StoreCache, error)
	ToDomain(store po.StoreCache) (entity.Store, error)
}

type StoreCacheMapper struct{}

var _ StoreCacheMapperIntf = (*StoreCacheMapper)(nil)

func NewStoreCacheMapper() *StoreCacheMapper {
	return &StoreCacheMapper{}
}

func (s *StoreCacheMapper) ToPersistence(store entity.Store) (po.StoreCache, error) {
	return po.StoreCache{
		ID:   store.ID,
		Name: store.Name,
	}, nil
}

func (s *StoreCacheMapper) ToDomain(store po.StoreCache) (entity.Store, error) {
	return entity.NewStore(store.ID, store.Name), nil
}
