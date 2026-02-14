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
	fmt.Println("PROSER - PROSE Framework Setup Tool")
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

	// Collect input for that project type
	fmt.Println("\nPlease answer the following questions about your project:")
	fmt.Println()
	answers, err := collector.Collect(projectType.Questions())
	if err != nil {
		fmt.Printf("❌ Error collecting input: %v\n", err)
		os.Exit(1)
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
	fmt.Println("📄 AGENT.md files created in subdirectories")
	fmt.Println("\n🎉 Your project is now configured for PROSE framework!")
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
	fmt.Println("  PROSER generates GitHub Copilot PROSE framework files for your project.")
	fmt.Println("  It creates .github/copilot-instructions.md, .instructions.md files,")
	fmt.Println("  and AGENT.md files in subdirectories based on your project configuration.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  proser                    # Setup in current directory")
	fmt.Println("  proser /path/to/project   # Setup in specified directory")
}
