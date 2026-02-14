package generator

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
