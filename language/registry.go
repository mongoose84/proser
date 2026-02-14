package language

// LanguageInfo contains metadata and guidelines for a programming language
type LanguageInfo struct {
	Name            string
	Aliases         []string // Alternative names (e.g., "js" for "javascript")
	FileExtensions  []string // e.g., []string{".go"}
	Guidelines      []string // Language-specific guideline lines
	TestingPatterns []string // Language-specific testing pattern lines
	ContextFiles    []string // Important context files (e.g., "go.mod", "package.json")
	OutputChecklist []string // Structured output checklist items
	BestPractices   []string // Best practice lines
}

// FrameworkInfo contains metadata and guidelines for a framework
type FrameworkInfo struct {
	Name       string
	Language   string   // The language this framework is for
	Guidelines []string // Framework-specific guideline lines
}

// Registry manages language and framework information
type Registry struct {
	languages  map[string]*LanguageInfo
	frameworks map[string]*FrameworkInfo
}

// NewRegistry creates a new empty registry
func NewRegistry() *Registry {
	return &Registry{
		languages:  make(map[string]*LanguageInfo),
		frameworks: make(map[string]*FrameworkInfo),
	}
}

// RegisterLanguage adds a language to the registry
func (r *Registry) RegisterLanguage(lang *LanguageInfo) {
	// Register by primary name
	r.languages[lang.Name] = lang

	// Register aliases
	for _, alias := range lang.Aliases {
		r.languages[alias] = lang
	}
}

// RegisterFramework adds a framework to the registry
func (r *Registry) RegisterFramework(fw *FrameworkInfo) {
	r.frameworks[fw.Name] = fw
}

// LookupLanguage finds a language by name or alias (case-insensitive)
func (r *Registry) LookupLanguage(name string) (*LanguageInfo, bool) {
	// Normalize to lowercase for lookup
	lang, exists := r.languages[name]
	return lang, exists
}

// LookupFramework finds a framework by name (case-insensitive)
func (r *Registry) LookupFramework(name string) (*FrameworkInfo, bool) {
	fw, exists := r.frameworks[name]
	return fw, exists
}

// NewDefaultRegistry creates a registry pre-populated with common languages and frameworks
func NewDefaultRegistry() *Registry {
	r := NewRegistry()
	registerDefaultLanguages(r)
	registerDefaultFrameworks(r)
	return r
}
