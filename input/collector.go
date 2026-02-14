package input

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Question represents a single input question
type Question struct {
	Key          string
	Prompt       string
	DefaultValue string
}

// InputCollector abstracts user input collection for testability
type InputCollector interface {
	// Collect asks the user all questions and returns answers keyed by Question.Key
	Collect(questions []Question) (map[string]string, error)
}

// InteractiveCollector collects input interactively from a reader
type InteractiveCollector struct {
	reader *bufio.Reader
}

// NewInteractiveCollector creates a new interactive collector
func NewInteractiveCollector(r io.Reader) *InteractiveCollector {
	return &InteractiveCollector{
		reader: bufio.NewReader(r),
	}
}

// Collect asks all questions and collects answers
func (c *InteractiveCollector) Collect(questions []Question) (map[string]string, error) {
	answers := make(map[string]string)

	for _, q := range questions {
		answer, err := c.prompt(q.Prompt, q.DefaultValue)
		if err != nil {
			return nil, fmt.Errorf("failed to collect input for %s: %w", q.Key, err)
		}
		answers[q.Key] = answer
	}

	return answers, nil
}

// prompt asks a single question and returns the answer
func (c *InteractiveCollector) prompt(prompt, defaultValue string) (string, error) {
	fmt.Printf("%s [%s] (type 'skip' to omit): ", prompt, defaultValue)
	input, err := c.reader.ReadString('\n')
	if err != nil {
		// In case of read error, return default value
		return defaultValue, nil
	}
	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "skip" {
		return "", nil
	}
	if input == "" {
		return defaultValue, nil
	}
	return input, nil
}
