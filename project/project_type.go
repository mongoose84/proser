package project

import (
	"github.com/mongoose84/proser/generator"
	"github.com/mongoose84/proser/input"
)

// ProjectType defines a project type and its associated generators and questions
type ProjectType interface {
	// Name returns the project type identifier
	Name() string

	// Description returns a user-facing description
	Description() string

	// Generators returns the set of generators that should run for this project type
	Generators() []generator.Generator

	// Questions returns the input questions specific to this project type
	Questions() []input.Question
}

// Registry of available project types
var registry = make(map[string]ProjectType)

// Register adds a project type to the registry
func Register(pt ProjectType) {
	registry[pt.Name()] = pt
}

// GetAll returns all registered project types
func GetAll() []ProjectType {
	types := make([]ProjectType, 0, len(registry))
	for _, pt := range registry {
		types = append(types, pt)
	}
	return types
}

// Get returns a project type by name
func Get(name string) (ProjectType, bool) {
	pt, exists := registry[name]
	return pt, exists
}

// Initialize registers all built-in project types
func init() {
	Register(&FullstackProject{})
	Register(&FrontendProject{})
	Register(&BackendProject{})
}
