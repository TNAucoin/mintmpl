# Mintmpl Extras

This directory contains additional tools and resources for working with Mintmpl.

## Claude Code Agent: Mintmpl Spec Generator

The `mintmpl-spec-generator.yaml` is a Claude Code agent that helps developers create `.mintmpl.yml` specification files for their projects.

### What It Does

The agent provides expert guidance for:
- Analyzing your codebase to identify what should be templated
- Creating properly structured `.mintmpl.yml` files
- Understanding node types and transformation strategies
- Language-specific patterns for Go, Java, C#, TypeScript, Python, Rust, and more
- Best practices for exclusions and conditional paths

### How to Use

#### Option 1: Install in Your Claude Code Configuration

1. Copy the agent to your Claude Code agents directory:
   ```bash
   mkdir -p ~/.claude/agents
   cp mintmpl-spec-generator.yaml ~/.claude/agents/
   ```

2. Restart Claude Code or reload your configuration

3. In any project, you can now invoke the agent:
   ```bash
   # In Claude Code
   @mintmpl-spec-generator help me create a .mintmpl.yml for this Go project
   ```

#### Option 2: Use Directly in Your Project

1. Copy the agent file to your project:
   ```bash
   mkdir -p .claude/agents
   cp /path/to/mintmpl/extras/mintmpl-spec-generator.yaml .claude/agents/
   ```

2. The agent will be available in Claude Code when working in this project

3. Use it by mentioning the agent:
   ```bash
   @mintmpl-spec-generator I need help creating a specification file
   ```

### Example Usage

**Create a spec for a Go project:**
```
@mintmpl-spec-generator I have a Go microservice called "user-service"
that I want to template. Can you help me create a .mintmpl.yml file?
```

**Create a spec for a TypeScript library:**
```
@mintmpl-spec-generator Help me create a .mintmpl.yml for my TypeScript
React component library
```

**Add a new variable to existing spec:**
```
@mintmpl-spec-generator I need to add a database_host variable to my
existing .mintmpl.yml file
```

**Troubleshoot transforms:**
```
@mintmpl-spec-generator My class name transform isn't working. The class
is called "UserService" but it's not being replaced in all places.
```

### What the Agent Knows

The agent has comprehensive knowledge about:

- **Mintmpl Architecture**: How AST-based transformations work
- **Node Types**: string, identifier, class, namespace, comment, any
- **Transform Strategies**: exact match, partial match, case sensitivity
- **Language Patterns**: Naming conventions and common patterns for multiple languages
- **Best Practices**: Exclusions, conditional paths, variable naming
- **Jinja2 Filters**: How to use filters for case conversion and formatting

### Agent Capabilities

The agent will:
1. ✅ Analyze your codebase using Read, Grep, and Glob tools
2. ✅ Ask clarifying questions about what should be templated
3. ✅ Identify language-specific patterns and conventions
4. ✅ Generate well-documented `.mintmpl.yml` files
5. ✅ Explain node types and transformation strategies
6. ✅ Provide language-specific exclusion patterns
7. ✅ Help troubleshoot transform issues

### Supported Languages

The agent provides patterns and examples for:
- **Go**: Module names, package names, struct/interface names
- **Java**: Group IDs, artifact IDs, class names, package names
- **C#**: Namespaces, assembly names, class names
- **TypeScript/JavaScript**: Package names, class names, module names
- **Python**: Package names, class names, module names
- **Rust**: Crate names, module names, struct/trait names
- **Multi-language**: Projects using multiple languages

### Tips for Best Results

1. **Be specific**: Tell the agent what language(s) your project uses
2. **Provide context**: Explain what type of project it is (library, service, app, etc.)
3. **Ask questions**: The agent can explain concepts if you're unsure
4. **Iterate**: Start with a basic spec and refine it
5. **Test**: Always test the generated spec with Mintmpl

### Example Workflow

```bash
# 1. Navigate to your project
cd ~/my-go-project

# 2. Invoke the agent in Claude Code
@mintmpl-spec-generator analyze this Go project and help me create a .mintmpl.yml

# 3. Answer the agent's questions about what should be customizable

# 4. Review and save the generated .mintmpl.yml file

# 5. Test it with Mintmpl
mintmpl generate --source . --output ./test-template

# 6. Verify the template with Copier
copier copy ./test-template ./test-output
```

### Troubleshooting

**Agent not showing up:**
- Ensure the file is in `~/.claude/agents/` or `.claude/agents/` in your project
- Restart Claude Code
- Check that the YAML is valid

**Agent doesn't understand my language:**
- The agent has patterns for common languages, but can adapt to others
- Provide specific examples from your codebase
- The agent can learn from your project structure

**Transforms not working as expected:**
- Ask the agent to explain node types
- Provide examples of where the value appears in your code
- The agent can help you choose the right node types and settings

## Contributing

If you have improvements or additional language patterns, contributions are welcome!

## License

Same as the Mintmpl project.
