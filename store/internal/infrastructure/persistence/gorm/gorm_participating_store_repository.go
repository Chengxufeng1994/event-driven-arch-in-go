package gorm

import (
	"context"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/aggregate"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store/internal/domain/repository"
	"gorm.io/gorm"
)

type GormParticipatingStoreRepository struct {
	db *gorm.DB
}

// FindAll implements repository.ParticipatingStoreRepository.
func (g *GormParticipatingStoreRepository) FindAll(ctx context.Context) ([]aggregate.StoreAgg, error) {
	panic("unimplemented")
}

var _ repository.ParticipatingStoreRepository = (*GormParticipatingStoreRepository)(nil)

func NewGormParticipatingStoreRepository(db *gorm.DB) *GormParticipatingStoreRepository {
	return &GormParticipatingStoreRepository{db: db}
}
