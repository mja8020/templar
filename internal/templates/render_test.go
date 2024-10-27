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

	// TODO: So many more, maybe build this fizture set somewhere else?
	fixtures := []fixture{}

	fixtures = append(fixtures, fixture{
		name:     "upper",
		template: `{{ "hello" | toUpper }}`,
		expected: "HELLO",
	})

	fixtures = append(fixtures, fixture{
		name:     "values",
		template: `{{ .FOO }}`,
		values: map[string]interface{}{
			"FOO": "bar",
		},
		expected: "bar",
	})

	fixtures = append(fixtures, fixture{
		name:     "values_combine",
		template: `{{ .FOO | toUpper }}`,
		values: map[string]interface{}{
			"FOO": "bar",
		},
		expected: "BAR",
	})

	fixtures = append(fixtures, fixture{
		name:     "error",
		template: `{{ !!!!! }}`,
		expected: "",
		err:      "unexpected",
	})

	// Confirm sprig is loaded
	fixtures = append(fixtures, fixture{
		name:     "sprig-add",
		template: `{{ add 1 2 3 }}`,
		expected: "6",
	})

	fixtures = append(fixtures, fixture{
		name:     "sprig-trim",
		template: `{{ trim "   hello    " }}`,
		expected: "hello",
	})

	for _, fixture := range fixtures {
		output, err := Render(fixture.template, fixture.values, fixture.sources)

		if fixture.err == "" {
			assert.Empty(t, err, "expected no error but got `%s`", err)
		} else {
			assert.ErrorContainsf(t, err, fixture.err, "unexpected error `%s`", err)
		}

		assert.Equal(t, fixture.expected, output, "template %s did not match, expected `%s` got `%s`", fixture.name, fixture.expected, output)
	}
}
