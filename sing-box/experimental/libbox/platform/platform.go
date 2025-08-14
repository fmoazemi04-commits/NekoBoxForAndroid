package platform

import (
	"github.com/sagernet/sing-box/adapter"
)

// Interface represents platform-specific functionality
type Interface interface {
	// Platform-specific methods would be defined here
	Initialize(networkManager adapter.NetworkManager) error
}

// DefaultInterface provides a default implementation
type DefaultInterface struct {
	// Implementation details
}

// Initialize initializes the platform interface
func (d *DefaultInterface) Initialize(networkManager adapter.NetworkManager) error {
	return nil
}

// NewInterface creates a new platform interface
func NewInterface() Interface {
	return &DefaultInterface{}
}

// Notification represents a platform notification
type Notification struct {
	Title   string
	Content string
	Icon    string
}

// NewNotification creates a new notification
func NewNotification(title, content string) *Notification {
	return &Notification{
		Title:   title,
		Content: content,
	}
}