package valueobject

type Store struct {
	ID       string
	Name     string
	Location string
}

func NewStore(id, name, location string) Store {
	return Store{
		ID:       id,
		Name:     name,
		Location: location,
	}
}
