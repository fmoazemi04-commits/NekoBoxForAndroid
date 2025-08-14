package adapter

import (
	"context"
	"net"
)

// Inbound represents an inbound adapter
type Inbound interface {
	Start() error
	Close() error
}

// Outbound represents an outbound adapter
type Outbound interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
	ListenPacket(ctx context.Context, destination string) (net.PacketConn, error)
}

// Adapter represents a general network adapter
type Adapter interface {
	Type() string
	Tag() string
}

// WIFIState represents WiFi connection state
type WIFIState struct {
	SSID string
	BSSID string
}

// NetworkManager manages network interfaces
type NetworkManager interface {
	GetWIFIState() *WIFIState
}

// NetworkInterface represents a network interface
type NetworkInterface struct {
	Name string
	Index int
}