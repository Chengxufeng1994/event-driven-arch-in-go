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

type GormMallRepository struct {
	db          *gorm.DB
	storeMapper mapper.StoreMapperIntf
}

var _ repository.MallRepository = (*GormMallRepository)(nil)

func NewGormMallRepository(db *gorm.DB) *GormMallRepository {
	return &GormMallRepository{
		db:          db,
		storeMapper: mapper.NewStoreMapper(),
	}
}

// AddStore implements repository.MallRepository.
func (r *GormMallRepository) AddStore(ctx context.Context, storeID string, name string, location string) error {

	result := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"name", "location", "participating",
			}),
		}).
		Create(&po.Store{
			ID:            storeID,
			Name:          name,
			Location:      location,
			Participating: false,
		})

	return errors.Wrap(result.Error, "inserting store")

}

// RenameStore implements repository.MallRepository.
func (r *GormMallRepository) RenameStore(ctx context.Context, storeID string, name string) error {
	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Where("id = ?", storeID).
		Update("name", name)

	return result.Error
}

// SetStoreParticipation implements repository.MallRepository.
func (r *GormMallRepository) SetStoreParticipation(ctx context.Context, storeID string, participating bool) error {
	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Where("id = ?", storeID).
		Update("participating", participating)

	return result.Error
}

// All implements repository.MallRepository.
func (r *GormMallRepository) All(ctx context.Context) ([]*aggregate.MallStore, error) {
	var stores []*po.Store

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Select("id, name, location, participating").
		Find(&stores)
	err := result.Error
	if err != nil {
		return nil, err
	}

	storesDomain := make([]*aggregate.MallStore, 0, len(stores))
	for _, store := range stores {
		storesDomain = append(storesDomain, &aggregate.MallStore{
			ID:            store.ID,
			Name:          store.Name,
			Location:      store.Location,
			Participating: store.Participating,
		})
	}

	return storesDomain, nil
}

// AllParticipating implements repository.MallRepository.
func (r *GormMallRepository) AllParticipating(ctx context.Context) ([]*aggregate.MallStore, error) {
	var stores []*po.Store

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Select("id, name, location, participating").
		Where("participating = ?", true).
		Find(&stores)
	err := result.Error
	if err != nil {
		return nil, err
	}

	storesDomain := make([]*aggregate.MallStore, 0, len(stores))
	for _, store := range stores {
		storesDomain = append(storesDomain, &aggregate.MallStore{
			ID:            store.ID,
			Name:          store.Name,
			Location:      store.Location,
			Participating: store.Participating,
		})
	}

	return storesDomain, nil
}

// Find implements repository.MallRepository.
func (r *GormMallRepository) Find(ctx context.Context, storeID string) (*aggregate.MallStore, error) {
	var store po.Store

	result := r.db.WithContext(ctx).
		Model(&po.Store{}).
		Where("id = ?", storeID).
		First(&store)
	err := result.Error
	if err != nil {
		return nil, err
	}

	return &aggregate.MallStore{
		ID:            store.ID,
		Name:          store.Name,
		Location:      store.Location,
		Participating: store.Participating,
	}, nil
}
