package config

import "github.com/mja8020/templar/internal/templates"

// Maybe project?
type Folder struct {
	Values      map[string]interface{}
	Datasources map[string]templates.DataSource
	Commands    map[string]Command
}

type Command struct {
	Environment map[string]string // Environment variables (values are templates)
	Command     string            // The comand to run i.e. terraform plan (values are templates)
	Success     []int             // Exit codes that are considered successfull
}
