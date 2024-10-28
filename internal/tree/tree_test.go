package tree

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {

	badPath, err := os.MkdirTemp("", "badPath")
	require.NoError(t, err)

	cases := []struct {
		name      string
		path      string
		returnErr bool
		size      int
	}{
		{
			name:      "badPath",
			path:      badPath,
			returnErr: true,
			size:      0,
		},
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
			size:      7,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree, err := NewTree(tc.path)
			if err != nil {
				require.True(t, tc.returnErr)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tc.size, tree.Size)
		})
	}
}
