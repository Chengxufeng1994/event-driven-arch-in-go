package serdes

import (
	"encoding/json"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/registry"
)

type JSONSerde struct {
	r registry.Registry
}

var _ registry.Serde = (*JSONSerde)(nil)

func NewJSONSerde(r registry.Registry) *JSONSerde {
	return &JSONSerde{r: r}
}

func (c JSONSerde) Register(v registry.Registrable, options ...registry.BuildOption) error {
	return registry.Register(c.r, v, c.serialize, c.deserialize, options)
}

func (c JSONSerde) RegisterKey(key string, v interface{}, options ...registry.BuildOption) error {
	return registry.RegisterKey(c.r, key, v, c.serialize, c.deserialize, options)
}

func (c JSONSerde) RegisterFactory(key string, fn func() interface{}, options ...registry.BuildOption) error {
	return registry.RegisterFactory(c.r, key, fn, c.serialize, c.deserialize, options)
}

func (JSONSerde) serialize(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (JSONSerde) deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
