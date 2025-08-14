package process

// Searcher represents a process searcher
type Searcher struct {
	// Implementation details
}

// NewSearcher creates a new process searcher
func NewSearcher() *Searcher {
	return &Searcher{}
}

// FindProcessInfo finds process information
func (s *Searcher) FindProcessInfo(pid int) (*ProcessInfo, error) {
	return &ProcessInfo{
		PID: pid,
	}, nil
}

// ProcessInfo represents process information
type ProcessInfo struct {
	PID int
	// Additional fields
}

// Info represents process information (alias for ProcessInfo)
type Info = ProcessInfo