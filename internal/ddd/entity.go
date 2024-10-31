package ddd

type IDer interface {
	ID() string
}

type EntityNamer interface {
	EntityName() string
}

type Entity interface {
	IDer
	IDSetter
	EntityNamer
	NameSetter
}

type EntityBase struct {
	id   string
	name string
}

var _ Entity = (*EntityBase)(nil)

func NewEntityBase(id, name string) EntityBase {
	return EntityBase{
		id:   id,
		name: name,
	}
}

func (e EntityBase) ID() string             { return e.id }
func (e EntityBase) EntityName() string     { return e.name }
func (e EntityBase) Equals(other IDer) bool { return e.id == other.ID() }

func (e *EntityBase) setID(id string)     { e.id = id }
func (e *EntityBase) setName(name string) { e.name = name }
