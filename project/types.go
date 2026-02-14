package project

import (
	"github.com/mongoose84/proser/generator"
	"github.com/mongoose84/proser/input"
)

// FullstackProject represents a project with both frontend and backend
type FullstackProject struct{}

func (p *FullstackProject) Name() string {
	return "fullstack"
}

func (p *FullstackProject) Description() string {
	return "Full-stack application with frontend and backend"
}

func (p *FullstackProject) Generators() []generator.Generator {
	return []generator.Generator{
		&generator.CopilotInstructionsGenerator{},
		&generator.FrontendInstructionsGenerator{},
		&generator.BackendInstructionsGenerator{},
		&generator.TestingInstructionsGenerator{},
		&generator.AgentMdGenerator{},
	}
}

func (p *FullstackProject) Questions() []input.Question {
	return append(append(
		generalQuestions(),
		frontendQuestions()...),
		append(backendQuestions(), testingQuestions()...)...,
	)
}

// FrontendProject represents a frontend-only project
type FrontendProject struct{}

func (p *FrontendProject) Name() string {
	return "frontend"
}

func (p *FrontendProject) Description() string {
	return "Frontend application only"
}

func (p *FrontendProject) Generators() []generator.Generator {
	return []generator.Generator{
		&generator.CopilotInstructionsGenerator{},
		&generator.FrontendInstructionsGenerator{},
		&generator.TestingInstructionsGenerator{},
		&generator.AgentMdGenerator{},
	}
}

func (p *FrontendProject) Questions() []input.Question {
	return append(append(
		generalQuestions(),
		frontendQuestions()...),
		testingQuestions()...,
	)
}

// BackendProject represents a backend-only project
type BackendProject struct{}

func (p *BackendProject) Name() string {
	return "backend"
}

func (p *BackendProject) Description() string {
	return "Backend/API service only"
}

func (p *BackendProject) Generators() []generator.Generator {
	return []generator.Generator{
		&generator.CopilotInstructionsGenerator{},
		&generator.BackendInstructionsGenerator{},
		&generator.TestingInstructionsGenerator{},
		&generator.AgentMdGenerator{},
	}
}

func (p *BackendProject) Questions() []input.Question {
	return append(append(
		generalQuestions(),
		backendQuestions()...),
		testingQuestions()...,
	)
}
