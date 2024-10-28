package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDisplay(t *testing.T) {
	expected := `
../../fixtures/1_basic
├── a
│   ├── b
│   └── c
│       └── d
└── e
    └── f
`

	tree, err := NewTree("../../fixtures/1_basic")
	require.NoError(t, err)

	display := tree.Display()
	t.Log(display)
	assert.Equal(t, expected, display)
}
