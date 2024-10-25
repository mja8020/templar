package templates

import (
	"bytes"
	"context"

	"github.com/hairyhenderson/gomplate/v4"
)

func Render(template string, values map[string]interface{}, sources []string) (output string, err error) {
	ctx := context.Background()

	options := gomplate.RenderOptions{}

	tr := gomplate.NewRenderer(options)

	buffer := new(bytes.Buffer)
	err = tr.Render(ctx, "template", template, buffer)
	if err != nil {
		return
	}

	output = buffer.String()

	return
}
