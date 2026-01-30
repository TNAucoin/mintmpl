package spec

import (
	"fmt"
	"os"

	"github.com/tnaucoin/mintmpl/internal/languages"
	"go.yaml.in/yaml/v3"
)

type Spec struct {
	Name             string                     `yaml:"name"`
	Version          string                     `yaml:"version"`
	Variables        map[string]*VariableConfig `yaml:"variables"`
	ConditionalPaths map[string]string          `yaml:"conditional_paths"`
	Exclude          []string                   `yaml:"exclude"`
	NoTransform      []string                   `yaml:"no_transform"`
}

type VariableConfig struct {
	Type        string            `yaml:"type"`
	Description string            `yaml:"description"`
	Default     any               `yaml:"default"`
	Choices     []string          `yaml:"choices"`
	Transforms  []TransformConfig `yaml:"transforms"`
}

type TransformConfig struct {
	Match         string   `yaml:"match"`
	NodeTypes     []string `yaml:"node_types"`
	Filter        string   `yaml:"filter"`
	CaseSensitive *bool    `yaml:"case_sensitive"`
	ExactMatch    bool     `yaml:"exact_match"`
}

type Transform struct {
	Match         string
	Replace       string
	NodeTypes     []languages.NodeCategory
	CaseSensitive bool
	ExactMatch    bool
}

func Load(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading spec file: %w", err)
	}

	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("parsing spec file %w", err)
	}

	if spec.Name == "" {
		spec.Name = "template"
	}
	if spec.Version == "" {
		spec.Version = "1.0.0"
	}

	return &spec, nil
}

func (s *Spec) BuildTransforms() []Transform {
	var transforms []Transform

	for varName, varConfig := range s.Variables {
		for _, t := range varConfig.Transforms {
			nodeTypes := make([]languages.NodeCategory, 0, len(t.NodeTypes))
			for _, nt := range t.NodeTypes {
				if nt == "" {
					continue
				}
				nodeTypes = append(nodeTypes, languages.NodeCategory(nt))
			}

			if len(nodeTypes) == 0 {
				nodeTypes = []languages.NodeCategory{languages.CategoryString}
			}

			var replacement string
			if t.Filter != "" {
				replacement = fmt.Sprintf("{{ %s | %s }}", varName, t.Filter)
			} else {
				replacement = fmt.Sprintf("{{ %s }}", varName)
			}

			caseSensitive := true
			if t.CaseSensitive != nil {
				caseSensitive = *t.CaseSensitive
			}

			transforms = append(transforms, Transform{
				Match:         t.Match,
				Replace:       replacement,
				NodeTypes:     nodeTypes,
				CaseSensitive: caseSensitive,
				ExactMatch:    t.ExactMatch,
			})
		}
	}
	return transforms
}

func GetDefaultExcludes() []string {
	return []string{
		".git",
		".git/**",
		"__pycache__",
		"__pycache__/**",
		"*.pyc",
		".mintmpl.yml",
		".template-spec.yaml",
		"node_modules",
		"node_modules/**",
		"bin",
		"obj",
		".vs",
		".vs/**",
		".idea",
		"idea/**",
	}
}
