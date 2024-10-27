package ddd

type Entity interface {
	GetID() string
}

type EntityBase struct {
	ID string
}

var _ Entity = (*EntityBase)(nil)

func NewEntityBase(id string) EntityBase {
	return EntityBase{ID: id}
}

func (e EntityBase) GetID() string {
	return e.ID
}
