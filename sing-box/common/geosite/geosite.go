package geosite

import (
	"fmt"
	"os"
	
	"github.com/sagernet/sing-box/option"
)

// Reader represents a geosite database reader
type Reader struct {
	// Implementation details would depend on the actual geosite format
	data map[string]*SourceSet
}

// SourceSet represents a set of domain rules
type SourceSet struct {
	Domain        []string
	DomainSuffix  []string
	DomainKeyword []string
	DomainRegex   []string
}

// Open opens a geosite database file
func Open(path string) (*Reader, interface{}, error) {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("geosite file not found: %s", path)
	}
	
	// Create a basic reader with empty data
	reader := &Reader{
		data: make(map[string]*SourceSet),
	}
	
	return reader, nil, nil
}

// Read reads a geosite code from the database
func (r *Reader) Read(code string) (*SourceSet, error) {
	if sourceSet, exists := r.data[code]; exists {
		return sourceSet, nil
	}
	
	// Return empty set for missing codes
	return &SourceSet{}, nil
}

// Compile converts a SourceSet to a DefaultHeadlessRule
func Compile(sourceSet *SourceSet) option.DefaultHeadlessRule {
	return option.DefaultHeadlessRule{
		Domain:        sourceSet.Domain,
		DomainSuffix:  sourceSet.DomainSuffix,
		DomainKeyword: sourceSet.DomainKeyword,
		DomainRegex:   sourceSet.DomainRegex,
	}
}