package templates

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"runtime"

	"github.com/Masterminds/sprig/v3"
	"github.com/hairyhenderson/gomplate/v4"
	"github.com/mja8020/templar/internal/utils"
)

// Datasource - External URL to pull data from, see https://docs.gomplate.ca/datasources/
type DataSource gomplate.DataSource

// Render - Rendersa the template including any top level values and datasources
func Render(template string, values map[string]interface{}, sources map[string]DataSource) (output string, err error) {
	if values == nil {
		values = map[string]interface{}{}
	}
	if sources == nil {
		sources = map[string]DataSource{}
	}

	valuesPath, err := serializeValues(values)
	if runtime.GOOS == "windows" {
		valuesPath = "/" + filepath.ToSlash(valuesPath)
	}
	defer os.Remove(valuesPath)

	contextValues, err := url.Parse(fmt.Sprintf("file://%s", valuesPath))
	if err != nil {
		return
	}

	ctx := context.Background()

	options := gomplate.RenderOptions{}
	options.Funcs = sprig.FuncMap()
	options.MissingKey = "default" // Missing key behavior, see https://docs.gomplate.ca/usage/#--missing-key
	// Idealy we fork or enhance gomplate so we are not writing a temp json file for an internal process
	options.Context = map[string]gomplate.DataSource{
		".": {
			URL: contextValues,
		},
	}
	options.Datasources = map[string]gomplate.DataSource{} // Datasources only called when needed

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

func serializeValues(values map[string]interface{}) (path string, err error) {
	file, err := os.CreateTemp("", "templar_*.json")
	if err != nil {
		return
	}
	defer file.Close()

	json, err := utils.JSONMarshal(values, true)
	if err != nil {
		return
	}

	_, err = file.WriteString(json)
	if err != nil {
		return
	}

	path = file.Name()

	return
}
