package config

// Maybe workspace?
type Root struct {
	Folder
	Layers    map[string]Layer    // Layers are only defined in root
	Variables map[string]Variable // Optionally define what values to expect
}

type Variable struct {
	Type       string
	Validation string
	Default    interface{}
}

type Layer struct{}
