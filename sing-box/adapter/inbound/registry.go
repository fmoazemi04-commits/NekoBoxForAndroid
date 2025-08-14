package inbound

import (
	"github.com/sagernet/sing-box/adapter"
)

// Registry manages inbound adapters
type Registry struct {
	adapters map[string]func() adapter.Inbound
}

// NewRegistry creates a new inbound registry
func NewRegistry() *Registry {
	return &Registry{
		adapters: make(map[string]func() adapter.Inbound),
	}
}

// Register registers an inbound adapter constructor
func (r *Registry) Register(name string, constructor func() adapter.Inbound) {
	r.adapters[name] = constructor
}