package languages

import (
	"path/filepath"
	"slices"
	"strings"

	forest "github.com/alexaandru/go-sitter-forest"
	sitter "github.com/alexaandru/go-tree-sitter-bare"
)

// NodeCategory Represents categories of AST nodes we allow to be targeted
type NodeCategory string

const (
	CategoryString     NodeCategory = "string"
	CategoryIdentifier NodeCategory = "identifier"
	CategoryNamespace  NodeCategory = "namespace"
	CategoryClass      NodeCategory = "class"
	CategoryComment    NodeCategory = "comment"
	CategoryAny        NodeCategory = "any"
)

type LanguageConfig struct {
	Name            string
	Extensions      []string
	Filenames       []string
	Language        *sitter.Language
	StringTypes     []string
	IdentifierTypes []string
	NamespaceTypes  []string
	ClassTypes      []string
	CommentTypes    []string
}

var Languages = map[string]*LanguageConfig{
	"python": {
		Name:            "python",
		Extensions:      []string{".py"},
		Language:        forest.GetLanguage("python"),
		StringTypes:     []string{"string", "string_content"},
		IdentifierTypes: []string{"identifier"},
		ClassTypes:      []string{"class_definition"},
		CommentTypes:    []string{"comment"},
	},
	"java": {
		Name:            "java",
		Extensions:      []string{".java"},
		Language:        forest.GetLanguage("java"),
		StringTypes:     []string{"string_literal", "string_fragment"},
		IdentifierTypes: []string{"identifier"},
		NamespaceTypes:  []string{"package_declaration"},
		ClassTypes:      []string{"class_declaration"},
		CommentTypes:    []string{"line_comment", "block_comment"},
	},
	"csharp": {
		Name:            "csharp",
		Extensions:      []string{".cs"},
		Language:        forest.GetLanguage("c_sharp"),
		StringTypes:     []string{"string_literal", "string_literal_content", "verbatim_string_literal"},
		IdentifierTypes: []string{"identifier"},
		NamespaceTypes:  []string{"namespace_declaration", "file_scoped_namespace_declaration"},
		ClassTypes:      []string{"class_declaration"},
		CommentTypes:    []string{"comment", "multiline_comment"},
	},
	"typescript": {
		Name:            "typescript",
		Extensions:      []string{".ts", ".tsx"},
		Language:        forest.GetLanguage("typescript"),
		StringTypes:     []string{"string", "string_fragment", "template_string"},
		IdentifierTypes: []string{"identifier", "property_identifier"},
		ClassTypes:      []string{"class_declaration"},
		CommentTypes:    []string{"comment"},
	},
	"javascript": {
		Name:            "javascript",
		Extensions:      []string{".js", ".jsx"},
		Language:        forest.GetLanguage("javascript"),
		StringTypes:     []string{"string", "string_fragment", "template_string"},
		IdentifierTypes: []string{"identifier", "property_identifier"},
		ClassTypes:      []string{"class_declaration"},
		CommentTypes:    []string{"comment"},
	},
	"go": {
		Name:            "go",
		Extensions:      []string{".go"},
		Language:        forest.GetLanguage("go"),
		StringTypes:     []string{"raw_string_literal", "interpreted_string_literal"},
		IdentifierTypes: []string{"identifier", "type_identifier", "field_identifier", "package_identifier"},
		NamespaceTypes:  []string{"package_clause"},
		ClassTypes:      []string{"type_declaration"},
		CommentTypes:    []string{"comment"},
	},
	"yaml": {
		Name:            "yaml",
		Extensions:      []string{".yaml", ".yml"},
		Language:        forest.GetLanguage("yaml"),
		StringTypes:     []string{"string_scalar", "double_quote_scalar", "single_quote_scalar", "block_scalar"},
		IdentifierTypes: []string{"flow_node"},
		CommentTypes:    []string{"comment"},
	},
	"toml": {
		Name:            "toml",
		Extensions:      []string{".toml"},
		Language:        forest.GetLanguage("toml"),
		StringTypes:     []string{"string", "multi_line_string"},
		IdentifierTypes: []string{"bare_key"},
		CommentTypes:    []string{"comment"},
	},
	"json": {
		Name:            "json",
		Extensions:      []string{".json"},
		Language:        forest.GetLanguage("json"),
		StringTypes:     []string{"string", "string_content"},
		IdentifierTypes: []string{},
		CommentTypes:    []string{},
	},
	"xml": {
		Name:            "xml",
		Extensions:      []string{".xml", ".csproj", ".props", ".targets", ".nuspec", ".config"},
		Filenames:       []string{"Directory.Build.props", "Directory.Build.targets", "Directory.Packages.props"},
		Language:        forest.GetLanguage("xml"),
		StringTypes:     []string{"AttValue", "CharData", "CData"},
		IdentifierTypes: []string{"Name"},
		CommentTypes:    []string{"Comment"},
	},
	"markdown": {
		Name:            "markdown",
		Extensions:      []string{".md", ".markdown", ".mdx"},
		Filenames:       []string{"README", "CHANGELOG", "CONTRIBUTING"},
		Language:        forest.GetLanguage("markdown"),
		StringTypes:     []string{"inline", "text", "code_span", "link_text"},
		IdentifierTypes: []string{"link_destination"},
		CommentTypes:    []string{"html_comment"},
	},
	"ini": {
		Name:            "ini",
		Extensions:      []string{".ini", ".editorconfig", ".gitconfig"},
		Filenames:       []string{".editorconfig", ".gitconfig"},
		Language:        forest.GetLanguage("ini"),
		StringTypes:     []string{"setting_value"},
		IdentifierTypes: []string{"setting_name", "section_name"},
		CommentTypes:    []string{"comment"},
	},
	"plaintext": {
		Name:            "plaintext",
		Extensions:      []string{".txt", ".sln", ".env.example"},
		Filenames:       []string{"LICENSE", "LICENCE", "NOTICE", "AUTHORS", "CODEOWNERS", ".env.example"},
		Language:        nil, // No AST parsing, use simple string replacement
		StringTypes:     []string{},
		IdentifierTypes: []string{},
		CommentTypes:    []string{},
	},
}

// maps file extensions to language names
var extensionToLanguage = make(map[string]string)

var filenameToLanguage = make(map[string]string)

func init() {
	for name, config := range Languages {
		for _, ext := range config.Extensions {
			extensionToLanguage[ext] = name
		}
		for _, filename := range config.Filenames {
			filenameToLanguage[filename] = name
		}
	}
}

func GetLanguageForExtension(ext string) *LanguageConfig {
	if name, ok := extensionToLanguage[ext]; ok {
		return Languages[name]
	}
	return nil
}

func GetLanguageForFile(path string) *LanguageConfig {
	filename := filepath.Base(path)
	if name, ok := filenameToLanguage[filename]; ok {
		return Languages[name]
	}
	ext := strings.ToLower(filepath.Ext(path))
	return GetLanguageForExtension(ext)
}

func (lc *LanguageConfig) GetNodeCategory(nodeType string) NodeCategory {
	if slices.Contains(lc.StringTypes, nodeType) {
		return CategoryString
	}
	if slices.Contains(lc.IdentifierTypes, nodeType) {
		return CategoryIdentifier
	}

	if slices.Contains(lc.NamespaceTypes, nodeType) {
		return CategoryNamespace
	}
	if slices.Contains(lc.ClassTypes, nodeType) {
		return CategoryClass
	}
	if slices.Contains(lc.CommentTypes, nodeType) {
		return CategoryComment
	}
	return ""
}

// MatchesCategory checks if a node type matches given categories
func (lc *LanguageConfig) MatchesCategory(nodeType string, categories []NodeCategory) bool {
	nodeCategory := lc.GetNodeCategory(nodeType)
	if nodeCategory == "" {
		return false
	}

	for _, cat := range categories {
		if cat == CategoryAny || cat == nodeCategory {
			return true
		}
	}
	return false
}
