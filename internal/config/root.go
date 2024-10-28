package config

import "github.com/mja8020/templar/internal/tree"

// Maybe workspace?
type Root struct {
	Folder

	Path   string `yaml:"-"`
	Target string `yaml:"-"`

	// Settings only applicable in root
	Layers    []Layer             `yaml:"layers,omitempty"`    // Layers are only defined in root, first is considered default
	Variables map[string]Variable `yaml:"variables,omitempty"` // Optionally define what values to expect
	// The key is the folder path relative to the root? If we use a tree structure here with references
	// to parent/child we won't be able to serialize.  We can possibly store the Tree as just a tree
	// of node names with the folder values kept here.
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

func NewRoot(target string) (root Root, err error) {
	root = Root{}

	root.RootConfig = "templar.yaml"
	root.Folderconfig = ".templar.yaml"
	root.TemplateFolder = ".templar"

	root.Target = target

	tree, err := tree.NewTree(target)
	if err != nil {
		return
	}

	root.Path = tree.Root.Name

	return
}
