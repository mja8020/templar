package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/mja8020/templar/internal/tree"
	"github.com/mja8020/templar/internal/utils"
	"gopkg.in/yaml.v2"
)

func (c *Stack) loadTree() (err error) {
	c.Tree, err = tree.NewTree(c.Target)
	if err != nil {
		return
	}
	c.Root = c.Tree.Root.Path

	return
}

func (c *Stack) loadConfig() (err error) {
	path := filepath.Join(c.Root, c.RootConfig)

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return nil
	}

	content, err := utils.FileRead(path)
	if err != nil {
		return nil
	}

	err = yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		return nil
	}

	return nil
}

func (c *Stack) buildFolders() (err error) {
	c.Folders = map[string]Folder{}

	// Start traversing the tree, load the config, merge etc

	return
}
