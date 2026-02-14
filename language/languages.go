package language

// registerDefaultLanguages populates the registry with built-in language definitions
func registerDefaultLanguages(r *Registry) {
	// Go
	r.RegisterLanguage(&LanguageInfo{
		Name:           "go",
		Aliases:        []string{"golang"},
		FileExtensions: []string{".go"},
		Guidelines: []string{
			"- Follow Go best practices and idioms",
			"- Use proper error handling with explicit error returns",
			"- Leverage Go's concurrency primitives (goroutines, channels)",
			"- Follow the effective Go guidelines",
		},
		BestPractices: []string{
			"- Use `gofmt` for code formatting",
			"- Handle errors explicitly, never ignore them",
			"- Use interfaces to define behavior",
			"- Prefer composition over inheritance",
			"- Package names should be lowercase, single words",
			"- Use meaningful variable and function names",
		},
		TestingPatterns: []string{
			"- [ ] Table-driven test patterns where appropriate",
			"- [ ] Benchmark tests for performance-critical code",
		},
		OutputChecklist: []string{
			"- [ ] Comprehensive error handling with wrapped errors",
			"- [ ] Unit tests with table-driven test patterns",
			"- [ ] Benchmark tests for performance-critical code",
		},
		ContextFiles: []string{"go.mod", "go.sum", "main.go"},
	})

	// Python
	r.RegisterLanguage(&LanguageInfo{
		Name:           "python",
		Aliases:        []string{"py"},
		FileExtensions: []string{".py"},
		Guidelines: []string{
			"- Follow Python best practices and idioms",
			"- Use proper exception handling",
			"- Follow PEP 8 style guide",
			"- Use type hints where appropriate",
		},
		BestPractices: []string{
			"- Follow PEP 8 style guidelines",
			"- Use virtual environments for dependencies",
			"- Write docstrings for all functions and classes",
			"- Use list comprehensions where appropriate",
			"- Follow naming conventions (snake_case)",
			"- Use context managers for resource management",
		},
		TestingPatterns: []string{
			"- [ ] Pytest fixtures for test setup",
			"- [ ] Parameterized tests for multiple scenarios",
		},
		OutputChecklist: []string{
			"- [ ] Proper exception handling with specific exception types",
			"- [ ] Unit tests with pytest or unittest",
			"- [ ] Type hints for better code maintainability",
		},
		ContextFiles: []string{"requirements.txt", "setup.py", "pyproject.toml"},
	})

	// Java
	r.RegisterLanguage(&LanguageInfo{
		Name:           "java",
		Aliases:        []string{},
		FileExtensions: []string{".java"},
		Guidelines: []string{
			"- Follow Java best practices and idioms",
			"- Use proper exception handling",
			"- Leverage Java's strong typing system",
			"- Follow SOLID principles",
		},
		BestPractices: []string{
			"- Follow Java naming conventions (camelCase, PascalCase)",
			"- Use proper exception handling with try-catch blocks",
			"- Leverage Java's object-oriented features",
			"- Use design patterns where appropriate",
			"- Follow dependency injection principles",
		},
		TestingPatterns: []string{
			"- [ ] JUnit test cases with proper assertions",
			"- [ ] Mock objects for unit testing",
		},
		OutputChecklist: []string{
			"- [ ] Proper exception handling with specific exception types",
			"- [ ] Unit tests with JUnit",
			"- [ ] JavaDoc documentation for all public methods",
		},
		ContextFiles: []string{"pom.xml", "build.gradle", "src/main/java"},
	})

	// JavaScript
	r.RegisterLanguage(&LanguageInfo{
		Name:           "javascript",
		Aliases:        []string{"js", "node", "node.js"},
		FileExtensions: []string{".js", ".mjs", ".cjs"},
		Guidelines: []string{
			"- Follow modern JavaScript best practices (ES6+)",
			"- Use proper module imports/exports",
		},
		BestPractices: []string{
			"- Use modern ES6+ syntax",
			"- Follow consistent code style (ESLint)",
			"- Use async/await for asynchronous code",
			"- Implement proper error handling",
			"- Use meaningful variable names",
		},
		TestingPatterns: []string{
			"- [ ] Jest test suites with describe/it blocks",
			"- [ ] Mock functions and modules where needed",
		},
		OutputChecklist: []string{
			"- [ ] Component documentation with usage examples",
			"- [ ] Proper TypeScript/JSDoc annotations",
			"- [ ] Unit tests with Jest/React Testing Library",
		},
		ContextFiles: []string{"package.json", "package-lock.json"},
	})

	// TypeScript
	r.RegisterLanguage(&LanguageInfo{
		Name:           "typescript",
		Aliases:        []string{"ts"},
		FileExtensions: []string{".ts", ".tsx"},
		Guidelines: []string{
			"- Follow TypeScript best practices with strict mode",
			"- Use proper type definitions and interfaces",
			"- Leverage TypeScript's type system for runtime safety",
		},
		BestPractices: []string{
			"- Use strict TypeScript configuration",
			"- Define interfaces for all data structures",
			"- Leverage type inference where possible",
			"- Use generics for reusable code",
			"- Avoid using `any` type",
		},
		TestingPatterns: []string{
			"- [ ] Jest test suites with proper type checking",
			"- [ ] Type-safe mock implementations",
		},
		OutputChecklist: []string{
			"- [ ] Component documentation with usage examples",
			"- [ ] Proper TypeScript type annotations",
			"- [ ] Unit tests with Jest/React Testing Library",
		},
		ContextFiles: []string{"package.json", "tsconfig.json"},
	})

	// Rust
	r.RegisterLanguage(&LanguageInfo{
		Name:           "rust",
		Aliases:        []string{"rs"},
		FileExtensions: []string{".rs"},
		Guidelines: []string{
			"- Follow Rust best practices and idioms",
			"- Leverage Rust's ownership system",
			"- Use proper error handling with Result types",
			"- Follow the Rust API guidelines",
		},
		BestPractices: []string{
			"- Use `rustfmt` for code formatting",
			"- Follow Rust naming conventions",
			"- Leverage the type system for safety",
			"- Use pattern matching extensively",
			"- Document all public APIs",
		},
		TestingPatterns: []string{
			"- [ ] Unit tests with #[test] attribute",
			"- [ ] Integration tests in tests/ directory",
		},
		OutputChecklist: []string{
			"- [ ] Proper error handling with Result and Option types",
			"- [ ] Unit tests and documentation tests",
			"- [ ] Benchmark tests for performance-critical code",
		},
		ContextFiles: []string{"Cargo.toml", "Cargo.lock"},
	})

	// C#
	r.RegisterLanguage(&LanguageInfo{
		Name:           "csharp",
		Aliases:        []string{"c#", "cs"},
		FileExtensions: []string{".cs"},
		Guidelines: []string{
			"- Follow C# best practices and idioms",
			"- Use proper exception handling",
			"- Leverage C#'s strong typing system",
			"- Follow SOLID principles",
		},
		BestPractices: []string{
			"- Follow C# naming conventions (PascalCase)",
			"- Use LINQ for data operations",
			"- Leverage async/await for asynchronous operations",
			"- Use dependency injection",
			"- Follow .NET coding guidelines",
		},
		TestingPatterns: []string{
			"- [ ] Unit tests with xUnit or NUnit",
			"- [ ] Mock objects with Moq or similar",
		},
		OutputChecklist: []string{
			"- [ ] Proper exception handling",
			"- [ ] Unit tests with xUnit/NUnit",
			"- [ ] XML documentation comments for public APIs",
		},
		ContextFiles: []string{"*.csproj", "*.sln"},
	})
}

// registerDefaultFrameworks populates the registry with built-in framework definitions
func registerDefaultFrameworks(r *Registry) {
	// React
	r.RegisterFramework(&FrameworkInfo{
		Name:     "react",
		Language: "javascript",
		Guidelines: []string{
			"- Follow React component lifecycle patterns",
			"- Use React Hooks for state management",
			"- Implement proper component composition",
			"- Follow React best practices and patterns",
		},
	})

	// Vue
	r.RegisterFramework(&FrameworkInfo{
		Name:     "vue",
		Language: "javascript",
		Guidelines: []string{
			"- Follow Vue component structure (template, script, style)",
			"- Use Vue composition API where appropriate",
			"- Implement proper prop validation",
			"- Follow Vue style guide and best practices",
		},
	})

	// Angular
	r.RegisterFramework(&FrameworkInfo{
		Name:     "angular",
		Language: "typescript",
		Guidelines: []string{
			"- Follow Angular style guide",
			"- Use dependency injection for services",
			"- Implement proper component communication",
			"- Use RxJS observables effectively",
		},
	})

	// Jest (testing framework)
	r.RegisterFramework(&FrameworkInfo{
		Name:     "jest",
		Language: "javascript",
		Guidelines: []string{
			"## Testing Framework (Jest)",
			"- Write descriptive test names that explain what is being tested",
			"- Use `describe` blocks to group related tests",
			"- Use `it` or `test` for individual test cases",
			"- Mock external dependencies appropriately",
		},
	})

	// pytest (testing framework)
	r.RegisterFramework(&FrameworkInfo{
		Name:     "pytest",
		Language: "python",
		Guidelines: []string{
			"## Testing Framework (pytest)",
			"- Write descriptive test function names (test_*)",
			"- Use fixtures for test setup and teardown",
			"- Leverage pytest's powerful assertion introspection",
			"- Use parametrize for testing multiple scenarios",
		},
	})

	// JUnit (testing framework)
	r.RegisterFramework(&FrameworkInfo{
		Name:     "junit",
		Language: "java",
		Guidelines: []string{
			"## Testing Framework (JUnit)",
			"- Use @Test annotations for test methods",
			"- Implement proper setup and teardown with @Before/@After",
			"- Use assertions effectively",
			"- Group related tests in test classes",
		},
	})

	// Go testing
	r.RegisterFramework(&FrameworkInfo{
		Name:     "go testing",
		Language: "go",
		Guidelines: []string{
			"## Testing Framework (Go testing)",
			"- Use table-driven tests for multiple scenarios",
			"- Name test functions with Test prefix",
			"- Use subtests with t.Run() for organization",
			"- Write benchmark tests for performance-critical code",
		},
	})
}
