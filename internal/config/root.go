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
	config.Layers = []Layer{}
	config.Variables = map[string]Variable{}

	err = config.buildTree()
	if err != nil {
		return
	}

	err = config.buildFolders()
	if err != nil {
		return
	}

	return
}

func (c *Config) buildTree() (err error) {
	c.Tree, err = tree.NewTree(c.Target)
	if err != nil {
		return
	}
	c.Root = c.Tree.Root.Name

	return
}

func (c *Config) buildFolders() (err error) {
	c.Folders = map[string]Folder{}

	// Start traversing the tree, load the config, merge etc

	return
}
