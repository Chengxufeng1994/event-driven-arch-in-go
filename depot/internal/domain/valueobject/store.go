package valueobject

type Store struct {
	ID       string
	Name     string
	Location string
}

func NewStore(name, location string) Store {
	return Store{
		Name:     name,
		Location: location,
	}
}
