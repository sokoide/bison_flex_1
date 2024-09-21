package main

// Scope represents a single scope level
type Scope struct {
	variables map[string]int
	parent    *Scope
}

// ScopeStack manages the stack of scopes
type ScopeStack struct {
	current *Scope
}

// NewScopeStack creates a new ScopeStack with a global scope
func NewScopeStack() *ScopeStack {
	return &ScopeStack{
		current: &Scope{
			variables: make(map[string]int),
			parent:    nil,
		},
	}
}

// PushScope creates a new scope
func (s *ScopeStack) PushScope() {
	newScope := &Scope{
		variables: make(map[string]int),
		parent:    s.current,
	}
	s.current = newScope
}

// PopScope removes the current scope
func (s *ScopeStack) PopScope() {
	if s.current.parent != nil {
		s.current = s.current.parent
	}
}

// Set sets a variable in the current scope
func (s *ScopeStack) Set(name string, value int) {
	s.current.variables[name] = value
}

// Get retrieves a variable from the current scope or its parents
func (s *ScopeStack) Get(name string) (int, *Scope) {
	scope := s.current
	for scope != nil {
		if value, ok := scope.variables[name]; ok {
			return value, scope
		}
		scope = scope.parent
	}
	return 0, nil
}

// Set sets a variable in the 's' scope
func (s *Scope) Set(name string, value int) {
	s.variables[name] = value
}
