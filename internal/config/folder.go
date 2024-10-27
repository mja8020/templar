package config

import "github.com/mja8020/templar/internal/templates"

// Maybe project?
type Folder struct {
	Values      map[string]interface{}
	Datasources map[string]templates.DataSource
	Commands    map[string]Command
	Templates   map[string]Template
}

type Command struct {
	Environment map[string]string // Environment variables (values are templates)
	Command     string            // The comand to run i.e. terraform plan (values are templates)
	Success     []int             // Exit codes that are considered successfull
}

// Templates are considered configuration and are loaded with the rest of the configs, for now
// only content is kept but we may decide later to add options for the templates i.e. IgnoreChanges
type Template struct {
	Content string
}
