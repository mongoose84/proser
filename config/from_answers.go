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

	// Chat modes config (only if enabled)
	if shouldEnable(answers["enable_chatmodes"]) {
		cfg.ChatModes = &ChatModesConfig{
			EnableArchitect:       shouldEnable(answers["chatmode_architect"]),
			EnableFrontend:        shouldEnable(answers["chatmode_frontend"]),
			EnableBackend:         shouldEnable(answers["chatmode_backend"]),
			EnableCodeReviewer:    shouldEnable(answers["chatmode_code_reviewer"]),
			EnableTechnicalWriter: shouldEnable(answers["chatmode_technical_writer"]),
			EnableDevOps:          shouldEnable(answers["chatmode_devops"]),
		}
	}

	// Prompts config (only if enabled)
	if shouldEnable(answers["enable_prompts"]) {
		cfg.Prompts = &PromptsConfig{
			EnableCodeReview:    shouldEnable(answers["prompt_code_review"]),
			EnableFeatureSpec:   shouldEnable(answers["prompt_feature_spec"]),
			EnableRefactor:      shouldEnable(answers["prompt_refactor"]),
			EnableBugFix:        shouldEnable(answers["prompt_bug_fix"]),
			EnablePRDescription: shouldEnable(answers["prompt_pr_description"]),
		}
	}

	// Specs config (only if enabled)
	if shouldEnable(answers["enable_specs"]) {
		cfg.Specs = &SpecsConfig{
			EnableFeatureTemplate: shouldEnable(answers["spec_feature_template"]),
			EnableAPIEndpoint:     shouldEnable(answers["spec_api_endpoint"]),
			EnableComponent:       shouldEnable(answers["spec_component"]),
		}
	}

	return cfg
}

// shouldEnable returns true if the answer is affirmative
func shouldEnable(answer string) bool {
	lower := strings.ToLower(strings.TrimSpace(answer))
	return lower == "yes" || lower == "y" || lower == "true" || lower == "1"
}
