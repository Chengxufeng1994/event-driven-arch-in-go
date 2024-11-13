package entity

type Store struct {
	ID   string
	Name string
}

func NewStore(id, name string) Store {
	return Store{
		ID:   id,
		Name: name,
	}
}
