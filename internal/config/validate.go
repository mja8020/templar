package config

import "fmt"

func (c *Stack) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("Stack name is required")
	}

	return nil
}
