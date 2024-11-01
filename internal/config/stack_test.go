package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStackLoad(t *testing.T) {
	fixture := "../../fixtures/1_basic"

	stack, err := NewStack(fixture)
	require.NoError(t, err)

	assert.Contains(t, stack.Root, "1_basic", "Missing stack root")
	assert.Equal(t, stack.Name, "basic", "Invalid name")
	assert.Len(t, stack.Layers, 0, "layers should not be configured")
	assert.Len(t, stack.Variables, 0, "variables should not be configured")
}
