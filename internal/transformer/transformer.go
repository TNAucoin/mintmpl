package transformer

import (
	"context"
	"sort"
	"strings"

	sitter "github.com/alexaandru/go-tree-sitter-bare"
	"github.com/tnaucoin/mintmpl/internal/languages"
	"github.com/tnaucoin/mintmpl/internal/spec"
)

type Replacement struct {
	StartByte uint32
	EndByte   uint32
	OldText   string
	NewText   string
}

type Transformer struct {
	transforms []spec.Transform
	parsers    map[string]*sitter.Parser
}

func New(transforms []spec.Transform) *Transformer {
	return &Transformer{
		transforms: transforms,
		parsers:    make(map[string]*sitter.Parser),
	}
}

func (t *Transformer) getParser(lang *languages.LanguageConfig) *sitter.Parser {
	if parser, ok := t.parsers[lang.Name]; ok {
		return parser
	}
	parser := sitter.NewParser()
	parser.SetLanguage(lang.Language)
	t.parsers[lang.Name] = parser
	return parser
}

// Transform transforms the source using AST replacements
func (t *Transformer) Transform(source []byte, langConfig *languages.LanguageConfig) ([]byte, bool) {
	parser := t.getParser(langConfig)
	tree, err := parser.ParseString(context.Background(), nil, source)
	if err != nil {
		return source, false
	}
	rootNode := tree.RootNode()
	replacements := t.collectReplacements(&rootNode, source, langConfig, false)
	if len(replacements) == 0 {
		return source, false
	}

	// sort by position reverse order for safe replace
	sort.Slice(replacements, func(i, j int) bool {
		return replacements[i].StartByte > replacements[j].StartByte
	})

	result := source
	for _, r := range replacements {
		result = append(result[:r.StartByte], append([]byte(r.NewText), result[r.EndByte:]...)...)
	}

	return result, true
}

func (t *Transformer) collectReplacements(node *sitter.Node, source []byte, langConfig *languages.LanguageConfig, parentReplaced bool) []Replacement {
	var replacements []Replacement
	thisNodeReplaced := false

	if !parentReplaced {
		nodeType := node.Type()

		for _, transform := range t.transforms {
			if !langConfig.MatchesCategory(nodeType, transform.NodeTypes) {
				continue
			}
			nodeText := string(source[node.StartByte():node.EndByte()])

			if t.matches(nodeText, transform) {
				newText := t.apply(nodeText, transform)
				if newText != nodeText {
					replacements = append(replacements, Replacement{
						StartByte: uint32(node.StartByte()),
						EndByte:   uint32(node.EndByte()),
						OldText:   nodeText,
						NewText:   newText,
					})
					thisNodeReplaced = true
					break // only apply the first matching change ?
				}
			}
		}
	}

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(uint32(i))
		childReplacements := t.collectReplacements(&child, source, langConfig, thisNodeReplaced)
		replacements = append(replacements, childReplacements...)
	}

	return replacements
}

func (t *Transformer) matches(value string, transform spec.Transform) bool {
	if transform.ExactMatch {
		if transform.CaseSensitive {
			return value == transform.Match
		}
		return strings.EqualFold(value, transform.Match)
	}

	if transform.CaseSensitive {
		return strings.Contains(value, transform.Match)
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(transform.Match))
}

func (t *Transformer) apply(value string, transform spec.Transform) string {
	if transform.ExactMatch {
		return transform.Replace
	}
	if transform.CaseSensitive {
		return strings.ReplaceAll(value, transform.Match, transform.Replace)
	}
	// case-insensitive replacement
	return replaceAllCaseInsensitive(value, transform.Match, transform.Replace)
}

func replaceAllCaseInsensitive(s, old, new string) string {
	if old == "" {
		return s
	}

	var result strings.Builder
	lower := strings.ToLower(s)
	oldLower := strings.ToLower(old)

	start := 0
	for {
		idx := strings.Index(lower[start:], oldLower)
		if idx == -1 {
			result.WriteString(s[start:])
			break
		}

		result.WriteString(s[start : start+idx])
		result.WriteString(new)
		start = start + idx + len(old)
	}

	return result.String()
}

func (t *Transformer) TransformFile(path string, content []byte) ([]byte, bool) {
	langConfig := languages.GetLanguageForFile(path)

	if langConfig == nil {
		return content, false
	}

	if langConfig.Language == nil {
		return t.TransformPlaintext(content)
	}

	return t.Transform(content, langConfig)
}

func (t *Transformer) TransformPlaintext(content []byte) ([]byte, bool) {
	result := string(content)
	changed := false

	for _, transform := range t.transforms {
		var newResult string
		if transform.CaseSensitive {
			if strings.Contains(result, transform.Match) {
				newResult = strings.ReplaceAll(result, transform.Match, transform.Replace)
			} else {
				newResult = result
			}
		} else {
			newResult = replaceAllCaseInsensitive(result, transform.Match, transform.Replace)
		}
		if newResult != result {
			changed = true
			result = newResult
		}
	}
	return []byte(result), changed
}
