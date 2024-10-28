package config

import "github.com/mja8020/templar/internal/tree"

// Maybe workspace?
type Config struct {
	// Entrypoint when running
	Target string `yaml:"-"`
	// Configuration root path
	Root string `yaml:"-"`

	Layers    []Layer             `yaml:"layers,omitempty"`    // First is considered default
	Variables map[string]Variable `yaml:"variables,omitempty"` // Optionally define what values to expect

	// The tree structure of the folders
	Tree *tree.Tree `yaml:"-"`

	// Details for the folders, key is folder name relative to root
	Folders map[string]Folder `yaml:"folders,omitempty"`

	// Don't serialize, maybe we want to enable later?
	RootConfig     string `yaml:"-"`
	Folderconfig   string `yaml:"-"`
	TemplateFolder string `yaml:"-"`
}

type Variable struct {
	Type       string      `yaml:"type,omitempty"`
	Validation string      `yaml:"validation,omitempty"`
	Default    interface{} `yaml:"default,omitempty"`
}

type Layer struct {
	Levels []string `yaml:"levels,omitempty"`
	Match  string   `yaml:"match,omitempty"` // Determines which layer is selected?
}

func NewConfig(target string) (config Config, err error) {
	config = Config{}

	config.RootConfig = "templar.yaml"
	config.Folderconfig = ".templar.yaml"
	config.TemplateFolder = ".templar"

	config.Target = target

	config.Tree, err = tree.NewTree(target)
	if err != nil {
		return
	}
	config.Root = config.Tree.Root.Name

	config.Layers = []Layer{}
	config.Variables = map[string]Variable{}
	config.Folders = map[string]Folder{}

	return
}
