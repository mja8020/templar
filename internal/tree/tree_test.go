package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecursePath(t *testing.T) {
	cases := []struct {
		name      string
		path      string
		returnErr bool
		size      int
	}{
		{
			name:      "0_onlyroot",
			path:      "../../fixtures/0_onlyroot",
			returnErr: false,
			size:      1,
		},
		{
			name:      "1_basic",
			path:      "../../fixtures/1_basic",
			returnErr: false,
			size:      5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := NewTree(tc.path)
			require.NoError(t, err)

			assert.Equal(t, tc.size, tree.Size)
		})
	}
}
