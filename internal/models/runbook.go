package models

// Runbook defines a troubleshooting workflow composed of symptoms and steps.
type Runbook struct {
	ID          string    `yaml:"id"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Symptoms    []Symptom `yaml:"symptoms"`
	Steps       []Step    `yaml:"steps"`
}

// Symptom describes a condition that identifies a troubleshooting scenario.
type Symptom struct {
	ID       string   `yaml:"id"`
	Name     string   `yaml:"name"`
	Keywords []string `yaml:"keywords"`
}

// Step defines a troubleshooting action.
type Step struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Command     string `yaml:"command"`
}
