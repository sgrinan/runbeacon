package models

type Runbook struct {
	ID          string    `yaml:"id"`
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Symptoms    []Symptom `yaml:"symptoms"`
	Steps       []Step    `yaml:"steps"`
}

type Symptom struct {
	ID       string   `yaml:"id"`
	Name     string   `yaml:"name"`
	Keywords []string `yaml:"keywords"`
}

type Step struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Command     string `yaml:"command"`
}
