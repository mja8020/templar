package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	type fixture struct {
		name     string
		template string
		values   map[string]interface{}
		expected string
		sources  map[string]DataSource
		err      string
	}

	fixtures := []fixture{}

	fixtures = append(fixtures, fixture{
		name:     "upper",
		template: `{{ "hello" | toUpper }}`,
		values:   map[string]interface{}{},
		expected: "HELLO",
		sources:  map[string]DataSource{},
		err:      "",
	})

	for _, fixture := range fixtures {
		output, err := Render(fixture.template, fixture.values, fixture.sources)

		if fixture.err == "" {
			assert.Empty(t, err, "expected no error but got %s", err)
		} else {
			assert.ErrorContainsf(t, err, fixture.err, "unexpected error %s", err)
		}

		assert.Equal(t, fixture.expected, output, "template %s did not match, expected `%s` got `%s", fixture.name, fixture.expected, output)
	}
}
