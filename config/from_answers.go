package config

import "strings"

// FromAnswers creates a ProjectConfig from user input answers
func FromAnswers(answers map[string]string) ProjectConfig {
	cfg := ProjectConfig{
		General: GeneralConfig{
			ProjectName: answers["project_name"],
			Description: answers["description"],
			CodeStyle:   answers["code_style"],
			Security:    answers["security"],
			CustomRules: answers["custom_rules"],
		},
		Testing: TestingConfig{
			Framework: answers["testing_framework"],
			Strategy:  answers["testing_strategy"],
		},
	}

	// Frontend config (only if not skipped)
	frontendLang := answers["frontend_language"]
	if frontendLang != "" && strings.ToLower(frontendLang) != "skip" {
		cfg.Frontend = &FrontendConfig{
			Language:  frontendLang,
			Framework: answers["frontend_framework"],
			BuildTool: answers["frontend_build_tool"],
		}
	}

	// Backend config (only if not skipped)
	backendLang := answers["backend_language"]
	if backendLang != "" && strings.ToLower(backendLang) != "skip" {
		cfg.Backend = &BackendConfig{
			Language:  backendLang,
			Framework: answers["backend_framework"],
			Database:  answers["backend_database"],
			APIRules:  answers["api_rules"],
		}
	}

	return cfg
}
