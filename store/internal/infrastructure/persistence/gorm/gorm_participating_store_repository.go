package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
	"gorm.io/gorm"
)

type GormParticipatingStoreRepository struct {
	db          *gorm.DB
	storeMapper mapper.StoreMapperIntf
}

var _ repository.ParticipatingStoreRepository = (*GormParticipatingStoreRepository)(nil)

func NewGormParticipatingStoreRepository(db *gorm.DB) *GormParticipatingStoreRepository {
	return &GormParticipatingStoreRepository{
		db:          db,
		storeMapper: mapper.NewStoreMapper(),
	}
}

func (r *GormParticipatingStoreRepository) FindAll(ctx context.Context) ([]*aggregate.StoreAgg, error) {
	var stores []*po.Store
	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Select("id, name, location, participating").
		Where("participating = ?", true).
		Scan(&stores)

	if err := result.Error; err != nil {
		return nil, err
	}

	return r.storeMapper.ToDomainList(stores), nil
}
