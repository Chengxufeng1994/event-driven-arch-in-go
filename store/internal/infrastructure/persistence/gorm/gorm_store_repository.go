package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/mapper"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/infrastructure/persistence/gorm/po"
	"github.com/stackus/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormStoreRepository struct {
	db          *gorm.DB
	storeMapper mapper.StoreMapperIntf
}

var _ repository.StoreRepository = (*GormStoreRepository)(nil)

func NewGormStoreRepository(db *gorm.DB) *GormStoreRepository {
	return &GormStoreRepository{db: db, storeMapper: mapper.NewStoreMapper()}
}

// Save implements repository.StoreRepository.
func (r *GormStoreRepository) Save(ctx context.Context, store *aggregate.StoreAgg) error {
	storePO := r.storeMapper.ToPersistent(store)

	result := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "location", "participating",
			}),
		}).
		Create(&storePO)

	return errors.Wrap(result.Error, "inserting store")
}

// Update implements repository.StoreRepository.
func (r *GormStoreRepository) Update(ctx context.Context, store *aggregate.StoreAgg) error {
	storePO := r.storeMapper.ToPersistent(store)

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Where("id = ?", storePO.ID).
		Select("name", "location", "participating").
		Updates(&po.Store{
			Name:          storePO.Name,
			Location:      storePO.Location,
			Participating: storePO.Participating,
		})

	return errors.Wrap(result.Error, "updating store")
}

// Delete implements repository.StoreRepository.
func (r *GormStoreRepository) Delete(ctx context.Context, storeID string) error {
	result := r.db.WithContext(ctx).
		Unscoped().
		Where("id = ?", storeID).
		Delete(&po.Store{})

	return errors.Wrap(result.Error, "deleting store")
}

// Find implements repository.StoreRepository.
func (r *GormStoreRepository) Find(ctx context.Context, storeID string) (*aggregate.StoreAgg, error) {
	var store po.Store

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Where("id = ?", storeID).
		First(&store)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return r.storeMapper.ToDomain(&store), nil
}

// FindAll implements repository.StoreRepository.
func (r *GormStoreRepository) FindAll(ctx context.Context) ([]*aggregate.StoreAgg, error) {
	var stores []*po.Store

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Find(&stores)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return r.storeMapper.ToDomainList(stores), nil
}
