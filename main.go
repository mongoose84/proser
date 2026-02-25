package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mongoose84/proser/config"
	"github.com/mongoose84/proser/filesystem"
	"github.com/mongoose84/proser/generator"
	"github.com/mongoose84/proser/input"
	"github.com/mongoose84/proser/project"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("PROSER - PROSE File Setup Tool")
	fmt.Println("===========================================")
	fmt.Println()

	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		printHelp()
		os.Exit(0)
	}

	// Parse target path argument
	targetPath := "."
	if len(os.Args) > 1 {
		targetPath = os.Args[1]
	}

	// Resolve to absolute path
	absTarget, err := filepath.Abs(targetPath)
	if err != nil {
		fmt.Printf("❌ Error resolving target path: %v\n", err)
		os.Exit(1)
	}

	// Verify target path exists and is a directory
	fileInfo, err := os.Stat(absTarget)
	if err != nil {
		fmt.Printf("❌ Error accessing target path: %v\n", err)
		os.Exit(1)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("❌ Target path is not a directory: %s\n", absTarget)
		os.Exit(1)
	}

	fmt.Printf("📁 Target directory: %s\n\n", absTarget)

	// Create dependencies
	fs := filesystem.NewOsFileSystem()
	collector := input.NewInteractiveCollector(os.Stdin)

	// Let user pick project type
	projectType := selectProjectType(collector)

	// Ask if user wants to generate all files
	fmt.Println("\n🤖 Quick setup or custom configuration?")
	fmt.Println()

	earlyAnswers, err := collector.Collect([]input.Question{
		{Key: "generate_all_files", Prompt: "Generate all recommended files for this project type? (yes/no)", DefaultValue: "yes"},
	})
	if err != nil {
		fmt.Printf("❌ Error collecting input: %v\n", err)
		os.Exit(1)
	}

	var answers map[string]string

	// Check if user wants all files or custom selection
	generateAll := earlyAnswers["generate_all_files"]
	if generateAll == "yes" || generateAll == "y" || generateAll == "Y" || generateAll == "YES" {
		// Quick setup: collect only essential questions
		answers = collectQuickSetup(collector, projectType)
	} else {
		// Custom setup: collect all detailed questions
		fmt.Println("\nPlease answer the following questions about your project:")
		fmt.Println()
		answers, err = collector.Collect(projectType.Questions())
		if err != nil {
			fmt.Printf("❌ Error collecting input: %v\n", err)
			os.Exit(1)
		}
	}

	// Build config from answers
	cfg := config.FromAnswers(answers)

	// Display configuration summary
	displaySummary(cfg)

	fmt.Println("\n📝 Generating files based on your configuration...")

	// Create generation context
	ctx := generator.GenerateContext{
		Config:     cfg,
		TargetPath: absTarget,
		FS:         fs,
	}

	// Create writer
	writer := generator.NewWriter(fs)

	// Run all generators for this project type
	for _, gen := range projectType.Generators() {
		if err := writer.RunGenerator(gen, ctx); err != nil {
			fmt.Printf("❌ Error running generator %s: %v\n", gen.Name(), err)
			os.Exit(1)
		}
		fmt.Printf("  ✓ Generated %s files\n", gen.Name())
	}

	fmt.Println("\n✅ Setup complete!")
	fmt.Println("📁 Files created in .github/")
	fmt.Println("📄 AGENTS.md created at project root")
	fmt.Println("\n🎉 Your project is now configured for PROSE Architectural Style for AI-Native Development!")
	fmt.Println("💡 Ask your AI agent to expand AGENTS.md with project-specific details.")
}

// selectProjectType prompts the user to select a project type
func selectProjectType(collector input.InputCollector) project.ProjectType {
	fmt.Println("Select project type:")
	types := project.GetAll()
	for i, pt := range types {
		fmt.Printf("  %d. %s - %s\n", i+1, pt.Name(), pt.Description())
	}
	fmt.Println()

	// For now, just ask for a selection
	answers, err := collector.Collect([]input.Question{
		{Key: "type", Prompt: "Enter project type (fullstack/frontend/backend)", DefaultValue: "fullstack"},
	})
	if err != nil || answers["type"] == "" {
		// Default to fullstack
		pt, _ := project.Get("fullstack")
		return pt
	}

	// Get the selected project type
	pt, exists := project.Get(answers["type"])
	if !exists {
		fmt.Printf("⚠️  Unknown project type '%s', using fullstack\n", answers["type"])
		pt, _ = project.Get("fullstack")
	}

	return pt
}

// collectQuickSetup collects only essential questions and auto-enables all appropriate files
func collectQuickSetup(collector input.InputCollector, projectType project.ProjectType) map[string]string {
	fmt.Println("\n⚡ Quick setup mode - collecting essential project information:")
	fmt.Println()

	// Determine which questions to ask based on project type
	var questions []input.Question

	// Always ask general questions
	questions = append(questions, input.Question{Key: "project_name", Prompt: "Project name", DefaultValue: "my-project"})
	questions = append(questions, input.Question{Key: "description", Prompt: "Project description", DefaultValue: "A software project"})
	questions = append(questions, input.Question{Key: "code_style", Prompt: "Code style guidelines", DefaultValue: "Follow standard formatting"})
	questions = append(questions, input.Question{Key: "security", Prompt: "Security requirements", DefaultValue: "Follow OWASP top 10"})

	// Ask tech stack questions based on project type
	switch projectType.Name() {
	case "fullstack":
		questions = append(questions, input.Question{Key: "frontend_language", Prompt: "Frontend language", DefaultValue: "JavaScript"})
		questions = append(questions, input.Question{Key: "frontend_framework", Prompt: "Frontend framework", DefaultValue: "React"})
		questions = append(questions, input.Question{Key: "backend_language", Prompt: "Backend language", DefaultValue: "Go"})
		questions = append(questions, input.Question{Key: "backend_framework", Prompt: "Backend framework", DefaultValue: "None"})
		questions = append(questions, input.Question{Key: "backend_database", Prompt: "Database", DefaultValue: "PostgreSQL"})
		questions = append(questions, input.Question{Key: "testing_framework", Prompt: "Testing framework", DefaultValue: "Jest"})

	case "frontend":
		questions = append(questions, input.Question{Key: "frontend_language", Prompt: "Frontend language", DefaultValue: "JavaScript"})
		questions = append(questions, input.Question{Key: "frontend_framework", Prompt: "Frontend framework", DefaultValue: "React"})
		questions = append(questions, input.Question{Key: "frontend_build_tool", Prompt: "Build tool", DefaultValue: "Vite"})
		questions = append(questions, input.Question{Key: "testing_framework", Prompt: "Testing framework", DefaultValue: "Jest"})

	case "backend":
		questions = append(questions, input.Question{Key: "backend_language", Prompt: "Backend language", DefaultValue: "Go"})
		questions = append(questions, input.Question{Key: "backend_framework", Prompt: "Backend framework", DefaultValue: "None"})
		questions = append(questions, input.Question{Key: "backend_database", Prompt: "Database", DefaultValue: "PostgreSQL"})
		questions = append(questions, input.Question{Key: "testing_framework", Prompt: "Testing framework", DefaultValue: "Go testing"})
	}

	// Collect answers
	answers, err := collector.Collect(questions)
	if err != nil {
		fmt.Printf("❌ Error collecting input: %v\n", err)
		os.Exit(1)
	}

	// Auto-populate remaining answers for all files
	allAnswers := config.DefaultAnswersForProjectType(projectType.Name(), answers)

	fmt.Println()
	fmt.Println("✨ Enabling all recommended files:")
	fmt.Println("  • Core instructions (copilot-instructions, domain-specific)")
	fmt.Println("  • Agent definitions (architect, engineers, code reviewer, etc.)")
	fmt.Println("  • Prompt templates (code review, feature spec, refactor, bug fix)")
	fmt.Println("  • Specification templates (feature, API, component)")
	fmt.Println("  • AGENTS.md discovery file")

	return allAnswers
}

// displaySummary shows the configuration summary
func displaySummary(cfg config.ProjectConfig) {
	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Project: %s\n", cfg.General.ProjectName)
	if cfg.General.Description != "" {
		fmt.Printf("  Description: %s\n", cfg.General.Description)
	}
	if cfg.HasFrontend() {
		fmt.Printf("  Frontend: %s", cfg.Frontend.Language)
		if cfg.Frontend.Framework != "" {
			fmt.Printf(" with %s", cfg.Frontend.Framework)
		}
		if cfg.Frontend.BuildTool != "" {
			fmt.Printf(" (%s)", cfg.Frontend.BuildTool)
		}
		fmt.Println()
	}
	if cfg.HasBackend() {
		fmt.Printf("  Backend: %s", cfg.Backend.Language)
		if cfg.Backend.Framework != "" && cfg.Backend.Framework != "None" {
			fmt.Printf(" with %s", cfg.Backend.Framework)
		}
		if cfg.Backend.Database != "" {
			fmt.Printf(" + %s", cfg.Backend.Database)
		}
		fmt.Println()
	}
	if cfg.Testing.Framework != "" {
		fmt.Printf("  Testing: %s (%s)\n", cfg.Testing.Framework, cfg.Testing.Strategy)
	}
	if cfg.General.CodeStyle != "" {
		fmt.Printf("  Code Style: %s\n", cfg.General.CodeStyle)
	}
	if cfg.HasBackend() && cfg.Backend.APIRules != "" {
		fmt.Printf("  API Rules: %s\n", cfg.Backend.APIRules)
	}
	if cfg.General.Security != "" {
		fmt.Printf("  Security: %s\n", cfg.General.Security)
	}
	if cfg.General.CustomRules != "" && cfg.General.CustomRules != "None" {
		fmt.Printf("  Custom Rules: %s\n", cfg.General.CustomRules)
	}
}

// printHelp displays usage information
func printHelp() {
	fmt.Println("Usage: proser [options] [target-path]")
	fmt.Println()
	fmt.Println("Arguments:")
	fmt.Println("  target-path    Path to the project to set up (default: current directory)")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println()
	fmt.Println("Description:")
	fmt.Println("  PROSER generates GitHub Copilot PROSE files for your project.")
	fmt.Println("  It creates .github/copilot-instructions.md, .instructions.md files,")
	fmt.Println("  and an AGENTS.md file at the project root based on your configuration.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  proser                    # Setup in current directory")
	fmt.Println("  proser /path/to/project   # Setup in specified directory")
}
