package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDisplay(t *testing.T) {
	// Root has empty name so we have an empty leading line
	expected := `

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
	assert.Equal(t, expected, display)
}
