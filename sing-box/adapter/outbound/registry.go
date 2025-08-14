package outbound

import (
	"github.com/sagernet/sing-box/adapter"
)

// Registry manages outbound adapters
type Registry struct {
	adapters map[string]func() adapter.Outbound
}

// NewRegistry creates a new outbound registry
func NewRegistry() *Registry {
	return &Registry{
		adapters: make(map[string]func() adapter.Outbound),
	}
}

// Register registers an outbound adapter constructor
func (r *Registry) Register(name string, constructor func() adapter.Outbound) {
	r.adapters[name] = constructor
}