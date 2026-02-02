<div align="center">
  <img src="./docs/logo.png" alt="Mintmpl Logo" width="300">

  # Mintmpl

  **Transform your working code into intelligent templates**

  [![Go Version](https://img.shields.io/badge/Go-1.24.3+-00ADD8?style=flat&logo=go)](https://go.dev/)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

</div>

---

## What is Mintmpl?

Mintmpl is a powerful CLI tool that generates [Copier](https://copier.readthedocs.io/) templates from real, working source code. Unlike traditional templating tools, Mintmpl uses **AST-aware transformations** to intelligently convert your projects into reusable scaffolding templates.

Instead of manually creating templates from scratch, take your battle-tested code and mint it into a template that others can use to bootstrap new projects.

## Why Mintmpl?

🎯 **AST-Aware** - Understands code structure using tree-sitter parsers, transforming only what makes sense in context

🌍 **Multi-Language** - Supports 13+ languages including Python, Go, TypeScript, Java, C#, and more

⚡ **Smart Transformations** - Convert strings, identifiers, namespaces, and class names with precision

🎨 **Flexible** - Define transformations via simple YAML specifications

🔄 **Copier Integration** - Generates ready-to-use Copier templates with full variable support

## Quick Start

### Installation

```bash
go install github.com/tnaucoin/mintmpl/cmd/mintmpl@latest
```

Or download pre-built binaries from [releases](https://github.com/TNAucoin/mintmpl/releases).

### Basic Usage

1. Create a `.mintmpl.yml` specification in your project:

```yaml
name: my-awesome-template
version: 1.0.0

variables:
  project_name:
    type: str
    description: "Name of your project"
    default: "my-project"
    transforms:
      - match: "example-project"
        context: ["string", "identifier"]

  author_name:
    type: str
    description: "Project author"
    default: "Your Name"
    transforms:
      - match: "John Doe"
        context: ["string"]
```

2. Generate your template:

```bash
mintmpl generate --source . --output ./my-template
```

3. Use your new template:

```bash
copier copy ./my-template ./new-project
```

## Supported Languages

| Language | Extensions | AST Features |
|----------|-----------|--------------|
| Python | `.py` | ✅ Full support |
| Go | `.go` | ✅ Full support |
| TypeScript | `.ts`, `.tsx` | ✅ Full support |
| JavaScript | `.js`, `.jsx` | ✅ Full support |
| Java | `.java` | ✅ Full support |
| C# | `.cs` | ✅ Full support |
| YAML | `.yaml`, `.yml` | ✅ Full support |
| JSON | `.json` | ✅ Full support |
| TOML | `.toml` | ✅ Full support |
| XML | `.xml` | ✅ Full support |
| Markdown | `.md` | ✅ Full support |
| INI | `.ini` | ✅ Full support |
| Plaintext | `.txt` | Basic support |

## Features

### Context-Aware Transformations

Transform code based on AST node types:

- **`string`** - String literals only
- **`identifier`** - Variable and function names
- **`namespace`** - Package/module declarations
- **`class`** - Class definitions
- **`comment`** - Code comments
- **`any`** - Any occurrence

### Advanced Matching

```yaml
transforms:
  # Exact match, case-sensitive
  - match: "MyClass"
    context: ["class"]

  # Partial match, case-insensitive
  - match: "example"
    context: ["string"]
    partial: true
    case_sensitive: false

  # With Jinja2 filters
  - match: "example-package"
    context: ["identifier"]
    replacement: "{{ package_name | replace('-', '_') }}"
```

### Conditional Files

Include files based on user choices:

```yaml
variables:
  use_async:
    type: bool
    description: "Use async implementation?"
    default: false

conditional_paths:
  - path: "src/async_*.py"
    condition: "{{ use_async }}"
```

### Exclusion Patterns

```yaml
exclude:
  - "*.pyc"
  - "__pycache__/"
  - "node_modules/"
  - ".git/"

no_transform:
  - "data/*.json"
  - "assets/*"
```

## CLI Commands

```bash
# Generate a template
mintmpl generate --source ./my-project --output ./template

# Specify custom spec file
mintmpl generate --spec ./custom-spec.yml

# Check version
mintmpl version
```

## How It Works

1. **Parse Specification** - Reads `.mintmpl.yml` from your source directory
2. **Analyze Code** - Uses tree-sitter to build AST for each file
3. **Smart Replace** - Matches transformation rules against AST nodes
4. **Generate Template** - Creates Copier template with `.jinja` files and `copier.yaml`

```
Your Project          Mintmpl           Copier Template
    ↓                    ↓                     ↓
  main.py    ──────>  [Transform]  ──────>  main.py.jinja
class Example         ↓ AST-aware           class {{ class_name }}
                      ↓ Replace
  .mintmpl.yml        ↓ Variables          copier.yaml
variables:            ↓                    questions:
  class_name: ...     ↓                      class_name: ...
```

## Contributing

We welcome contributions! Whether it's:

- 🐛 Bug reports
- 💡 Feature requests
- 📝 Documentation improvements
- 🔧 Code contributions

Please open an issue or pull request on [GitHub](https://github.com/TNAucoin/mintmpl).

## Roadmap

- [ ] Template validation command
- [ ] Template inspection and preview
- [ ] Support for more languages (Rust, Ruby, etc.)
- [ ] Template marketplace/registry
- [ ] Interactive mode for creating specifications
- [ ] Diff preview before generation

## License

This project is licensed under the MIT License.

## Credits

Built with:

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [tree-sitter](https://tree-sitter.github.io/) - Parser generator
- [go-sitter-forest](https://github.com/alexaandru/go-sitter-forest) - Multi-language bindings

Powered by [Copier](https://copier.readthedocs.io/) for template rendering.

---

<div align="center">

  **[Issues](https://github.com/TNAucoin/mintmpl/issues)** • **[Discussions](https://github.com/TNAucoin/mintmpl/discussions)**

  Made with ⚡ by the Mintmpl community

</div>
