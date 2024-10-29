package config

// .templar.yaml
type Folder struct {
	Common

	Path   string                 `yaml:"-"` // Full path
	Values map[string]interface{} `yaml:"values,omitempty"`
}

func NewFolder(path string) (folder Folder, err error) {
	folder = Folder{}

	// TODO

	return
}
