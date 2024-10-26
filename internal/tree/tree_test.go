package tree

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecursePath(t *testing.T) {
	cases := []struct {
		name      string
		paths     []string
		root      string
		returnErr bool
		pathCount int
	}{
		{
			name: "EmptyRoot",
			paths: []string{
				"1",
			},
			root:      "1",
			returnErr: false,
			pathCount: 1,
		},
		{
			name: "FullTree",
			paths: []string{
				"1/2/3/4/5/6/7/8",
				"1/9/10",
			},
			root:      "1",
			returnErr: false,
			pathCount: 10,
		},
		{
			name:      "BadPath",
			paths:     []string{},
			root:      "1",
			returnErr: true,
			pathCount: 0,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			// Setup tmp dir
			dir, err := os.MkdirTemp("", tc.name)
			require.NoError(t, err)
			defer os.RemoveAll(dir)

			// Create all paths
			for _, path := range tc.paths {
				path := filepath.Join(dir, path)
				err := os.MkdirAll(path, os.ModePerm)
				require.NoError(t, err)
			}

			var paths []string
			recurseDir := filepath.Join(dir, tc.root)
			err = recursePath(recurseDir, &paths)

			// Validate errors
			if err != nil {
				require.True(t, tc.returnErr)
				return
			} else {
				require.False(t, tc.returnErr)
			}

			// Validate count
			require.Equal(t, tc.pathCount, len(paths))
		})
	}
}
