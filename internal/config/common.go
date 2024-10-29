package config

import "github.com/mja8020/templar/internal/templates"

// Fields shared in both the stack and folder configs
type Common struct {
	DataSources map[string]templates.DataSource `yaml:"datasources,omitempty"`
	Commands    map[string]Command              `yaml:"commands,omitempty"`
	Templates   map[string]Template             `yaml:"templates,omitempty"`
}

type Command struct {
	Environment map[string]string `yaml:"environment,omitempty"` // Environment variables (values are templates)
	Command     string            `yaml:"command,omitempty"`     // The comand to run i.e. terraform (values are templates)
	Arguments   []string          `yaml:"arguments,omitempty"`   // The comand arguments (values are templates)
	Success     []int             `yaml:"success,omitempty"`     // Exit codes that are considered successfull
}

// Templates are considered configuration and are loaded with the rest of the configs, for now
// only content is kept but we may decide later to add options for the templates i.e. IgnoreChanges
type Template struct {
	Content string `yaml:"content,omitempty"`
}
