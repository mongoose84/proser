package config

// DefaultAnswersForProjectType generates default answers for all files based on project type
func DefaultAnswersForProjectType(projectType string, generalAnswers map[string]string) map[string]string {
	answers := make(map[string]string)

	// Copy general answers (project name, description, etc.)
	for k, v := range generalAnswers {
		answers[k] = v
	}

	// Set defaults for fields not provided
	if answers["custom_rules"] == "" {
		answers["custom_rules"] = "None"
	}
	if answers["api_rules"] == "" {
		answers["api_rules"] = "RESTful API design"
	}
	if answers["testing_strategy"] == "" {
		answers["testing_strategy"] = "Unit and Integration tests"
	}
	if answers["frontend_build_tool"] == "" {
		answers["frontend_build_tool"] = "Vite"
	}

	// Enable agents, prompts, and specs by default
	answers["enable_agents"] = "yes"
	answers["enable_prompts"] = "yes"
	answers["enable_specs"] = "yes"

	// Enable appropriate agents based on project type
	answers["agent_architect"] = "yes"
	answers["agent_code_reviewer"] = "yes"
	answers["agent_technical_writer"] = "yes"
	answers["agent_tester"] = "yes"
	answers["agent_devops"] = "no" // typically opt-in

	// Enable all prompt templates
	answers["prompt_code_review"] = "yes"
	answers["prompt_feature_spec"] = "yes"
	answers["prompt_refactor"] = "yes"
	answers["prompt_bug_fix"] = "yes"
	answers["prompt_pr_description"] = "yes"

	switch projectType {
	case "fullstack":
		// Enable both frontend and backend agents
		answers["agent_frontend"] = "yes"
		answers["agent_backend"] = "yes"

		// Enable all spec templates for fullstack
		answers["spec_feature_template"] = "yes"
		answers["spec_api_endpoint"] = "yes"
		answers["spec_component"] = "yes"

	case "frontend":
		// Enable only frontend agent
		answers["agent_frontend"] = "yes"
		answers["agent_backend"] = "no"

		// Enable frontend-relevant specs
		answers["spec_feature_template"] = "yes"
		answers["spec_api_endpoint"] = "no"
		answers["spec_component"] = "yes"

	case "backend":
		// Enable only backend agent
		answers["agent_frontend"] = "no"
		answers["agent_backend"] = "yes"

		// Enable backend-relevant specs
		answers["spec_feature_template"] = "yes"
		answers["spec_api_endpoint"] = "yes"
		answers["spec_component"] = "no"
	}

	return answers
}
