package dialer

import (
	"context"
	"net"
)

// Variables for dialer configuration
var (
	DoNotSelectInterface bool = false
)

// Dialer represents a network dialer
type Dialer struct {
	// Implementation details
}

// NewDialer creates a new dialer
func NewDialer() *Dialer {
	return &Dialer{}
}

// DialContext dials a network connection
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return net.Dial(network, address)
}