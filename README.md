# templar

Utility for rendering a set of templates across a collection of folders in a project.

## Examples

```yaml
---
# templar.yaml
variables:
    foo:
        type: string
    bar:
        type: number
    baz:
        default: something
commands:
    hello:
        command: 'echo hello'
```

```yaml
---
# .templar.yaml
values:
    bar: 2
```

```bash
templar init

# Prompt for foo
# bar = 2
# baz = something
```

```bash
templar run hello

# Output: hello
```

```bash
# Render file content from templates
templar render

# Drift check of the content only
templar render --check
```
