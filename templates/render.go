package templates

import (
	"context"
	"fmt"
	"os"
)

func Render(template string, values map[string]interface{}, sources []string) (output string, err error) {
	ctx := context.Background()

	// create a new template renderer
	tr := NewRenderer(RenderOptions{})

	// render a template to stdout
	err := tr.Render(ctx, "mytemplate",
		`{{ "hello, world!" | toUpper }}`,
		os.Stdout)
	if err != nil {
		fmt.Println("gomplate error:", err)
	}

	return
}
