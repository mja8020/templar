package config

// Maybe workspace?
type Root struct {
	Folder

	// Settings only applicable in root
	Layers    []Layer             // Layers are only defined in root, first is considered default
	Variables map[string]Variable // Optionally define what values to expect
	// The key is the folder path relative to the root? If we use a tree structure here with references
	// to parent/child we won't be able to serialize.  We can possibly store the Tree as just a tree
	// of node names with the folder values kept here.
	Folders map[string]Folder
}

type Variable struct {
	Type       string
	Validation string
	Default    interface{}
}

type Layer struct {
	Levels []string
	Match  string // Determines which layer is selected?
}
