package generator

// PromptsGenerator generates .github/prompts/*.prompt.md files
// TODO: Implement prompt generation
type PromptsGenerator struct{}

func (g *PromptsGenerator) Name() string {
	return "prompts"
}

func (g *PromptsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	// TODO: Generate .github/prompts/*.prompt.md files
	return make(map[string]string), nil
}

// ChatModesGenerator generates .github/chatmodes/*.chatmode.md files
// TODO: Implement chatmode generation
type ChatModesGenerator struct{}

func (g *ChatModesGenerator) Name() string {
	return "chatmodes"
}

func (g *ChatModesGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	// TODO: Generate .github/chatmodes/*.chatmode.md files
	return make(map[string]string), nil
}

// SpecsGenerator generates .github/specs/*.spec.md files
// TODO: Implement spec generation
type SpecsGenerator struct{}

func (g *SpecsGenerator) Name() string {
	return "specs"
}

func (g *SpecsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	// TODO: Generate .github/specs/*.spec.md files
	return make(map[string]string), nil
}

// SkillsGenerator generates .github/skills/*.skill.md files
// TODO: Implement skill generation
type SkillsGenerator struct{}

func (g *SkillsGenerator) Name() string {
	return "skills"
}

func (g *SkillsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	// TODO: Generate .github/skills/*.skill.md files
	return make(map[string]string), nil
}
