package nekoutils

import (
	"github.com/sagernet/sing-box/option"
)

// Function type for proxy selection callback
type ProxySelectedCallback func(groupTag, proxyDisplayName, proxyType, proxyAddress string)

// Function type for geo rules
type GeoRulesFunc func(name string) ([]option.HeadlessRule, error)

// Global variables that are assigned from libcore
var (
	Selector_OnProxySelected  ProxySelectedCallback
	GetGeoSiteHeadlessRules   GeoRulesFunc
	GetGeoIPHeadlessRules     GeoRulesFunc
)

// Utility functions used in http.go (commented out in the original code)

// Map transforms a slice using the provided function
func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, item := range slice {
		result[i] = fn(item)
	}
	return result
}

// Filter returns a new slice containing only elements that satisfy the predicate
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// Contains checks if a slice contains a specific element
func Contains[T comparable](slice []T, element T) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}