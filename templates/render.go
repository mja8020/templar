package templates

import (
	"bytes"
	"context"

	"github.com/hairyhenderson/gomplate/v4"
)

type DataSource gomplate.DataSource

func Render(template string, values map[string]interface{}, sources map[string]DataSource) (output string, err error) {
	if sources == nil {
		sources = map[string]DataSource{}
	}

	ctx := context.Background()

	options := gomplate.RenderOptions{}
	options.Datasources = map[string]gomplate.DataSource{}

	for name, datasource := range sources {
		options.Datasources[name] = gomplate.DataSource(datasource)
	}

	tr := gomplate.NewRenderer(options)

	buffer := new(bytes.Buffer)
	err = tr.Render(ctx, "template", template, buffer)
	if err != nil {
		return
	}

	output = buffer.String()

	return
}
