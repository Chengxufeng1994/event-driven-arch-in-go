package serdes

import (
	"testing"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
	"github.com/stretchr/testify/assert"
)

type Person struct {
	Name string
}

func (p *Person) Key() string {
	return p.Name
}

func TestJsonSerde(t *testing.T) {
	registry := registry.New()

	jsonSerde := NewJSONSerde(registry)
	_ = jsonSerde.Register(&Person{})

	person := &Person{Name: "Chengxufeng"}
	jsonData, _ := jsonSerde.serialize(person)

	newPerson := &Person{}
	_ = jsonSerde.deserialize(jsonData, newPerson)

	assert.Equal(t, person.Name, newPerson.Name)
}
