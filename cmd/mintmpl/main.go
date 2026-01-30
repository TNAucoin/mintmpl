package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	sitter "github.com/alexaandru/go-tree-sitter-bare"
	"github.com/spf13/cobra"
	"github.com/wts-paradigm/mintmpl/internal/spec"
	"github.com/wts-paradigm/mintmpl/internal/transformer"
	"go.yaml.in/yaml/v3"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "mintmpl",
	Short: "Mintmpl Copier templates from source code",
	Long:  "Mintmpl generates Copier templates from working source code using AST-aware transformations",
}

func init() {
	rootCmd.AddCommand(generateCmd)
	validateCmd := &cobra.Command{}
	rootCmd.AddCommand(validateCmd)
	inspectCmd := &cobra.Command{}
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mintmpl %s (%s)\n", version, commit)
	},
}

var (
	genSource       string
	genOutput       string
	genSpec         string
	genGithubOutput string
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a copier template from source",
	Long:  "Generate a Copier template by transforming source using AST-aware replacements",
	RunE:  runGenerate,
}

func init() {
	generateCmd.Flags().StringVarP(&genSource, "source", "s", ".", "Source Directory")
	generateCmd.Flags().StringVarP(&genOutput, "output", "o", "template-output", "Output directory of generated copier template")
	generateCmd.Flags().StringVarP(&genSpec, "spec", "", "", "Path to spec file (Default: SOURCE/.mintmpl.yml)")
	generateCmd.Flags().StringVarP(&genGithubOutput, "github-output", "", "", "Path to GitHub output file")
}

func runGenerate(cmd *cobra.Command, args []string) error {
	source, err := filepath.Abs(genSource)
	if err != nil {
		return fmt.Errorf("resolving source path: %w", err)
	}
	output, err := filepath.Abs(genOutput)
	if err != nil {
		return fmt.Errorf("resolving output path %w", err)
	}
	specFile := genSpec
	if specFile == "" {
		specFile = filepath.Join(source, ".mintmpl.yml")
	}

	templateSpec, err := spec.Load(specFile)
	if err != nil {
		return fmt.Errorf("loading spec: %w", err)
	}

	fmt.Printf("Source: %s\n", source)
	fmt.Printf("Output: %s\n", output)
	fmt.Printf("Spec: %s\n", specFile)
	fmt.Println()

	if err := os.RemoveAll(output); err != nil {
		return fmt.Errorf("cleaning output directory: %w", err)
	}

	templateDir := filepath.Join(output, "template")
	if err := os.MkdirAll(templateDir, 0755); err != nil {
		return fmt.Errorf("creating template directory: %w", err)
	}

	transforms := templateSpec.BuildTransforms()
	trans := transformer.New(transforms)

	excludes := append(spec.GetDefaultExcludes(), templateSpec.Exclude...)

	var fileProcessed, filesTransformed int
	var warnings []string

	err = filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}

		if shouldExclude(relPath, excludes) {
			if d.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if d.IsDir() {
			return nil
		}

		fileProcessed++

		content, err := os.ReadFile(path)
		if err != nil {
			warnings = append(warnings, fmt.Sprintf("Could not read %s: %v", relPath, err))
			return nil
		}

		destPath := filepath.Join(templateDir, relPath)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("creating directory for %s: %w", relPath, err)
		}

		if shouldSkipTransform(relPath, templateSpec.NoTransform) {
			return os.WriteFile(destPath, content, 0644)
		}

		transformed, wasTransformed := trans.TransformFile(path, content)

		if wasTransformed {
			filesTransformed++
			destPath = destPath + ".jinja"
		}

		return os.WriteFile(destPath, transformed, 0644)
	})

	if err != nil {
		return fmt.Errorf("walking source directory: %w", err)
	}

	if err := generateCopierYAML(templateSpec, output); err != nil {
		return fmt.Errorf("generating copier yaml: %w", err)
	}

	fmt.Printf("\nTemplate generation finished.\n")
	fmt.Printf("	Files Processed: %d\n", fileProcessed)
	fmt.Printf("	Files Transformed: %d\n", filesTransformed)

	for _, w := range warnings {
		fmt.Printf("::warning::%s\n", w)
	}

	if genGithubOutput != "" {
		f, err := os.OpenFile(genGithubOutput, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(f, "files-processed=%d\n", fileProcessed)
			fmt.Fprintf(f, "files-transformed=%d\n", filesTransformed)
			fmt.Fprintf(f, "template-path=%s\n", output)
			f.Close()
		}
	}
	return nil
}

func shouldExclude(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if path == pattern {
			return true
		}

		if strings.HasSuffix(pattern, "/") || strings.HasSuffix(pattern, "/**") {
			prefix := strings.TrimSuffix(strings.TrimSuffix(pattern, "/**"), "/")
			if path == prefix || strings.HasPrefix(path, prefix+string(filepath.Separator)) {
				return true
			}
		}

		if matched, _ := filepath.Match(pattern, path); matched {
			return true
		}
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

func shouldSkipTransform(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if matched, _ := filepath.Match(pattern, path); matched {
			return true
		}
		if matched, _ := filepath.Match(pattern, filepath.Base(path)); matched {
			return true
		}
	}
	return false
}

func generateCopierYAML(s *spec.Spec, outputDir string) error {
	config := make(map[string]interface{})

	config["_min_copier_version"] = "9.0.0"
	config["_subdirectory"] = "template"
	config["_jinja_extensions"] = []string{"jinja2_time.TimeExtension"}

	for name, varConfig := range s.Variables {
		varDef := map[string]interface{}{
			"type":    varConfig.Type,
			"help":    varConfig.Description,
			"default": varConfig.Default,
		}
		if len(varConfig.Choices) > 0 {
			varDef["choices"] = varConfig.Choices
		}
		config[name] = varDef
	}

	if len(s.ConditionalPaths) > 0 {
		var excludes []string
		for pathPattern, condition := range s.ConditionalPaths {
			excludes = append(excludes, fmt.Sprintf("{%% if not %s %%}%s{%% endif %%}", condition, pathPattern))
		}
		config["_exclude"] = excludes
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("marshaling copier.yaml: %w", err)
	}

	copierPath := filepath.Join(outputDir, "copier.yaml")
	if err := os.WriteFile(copierPath, data, 0644); err != nil {
		return fmt.Errorf("writing copier.yaml: %w", err)
	}

	fmt.Printf("Generated: %s\n", copierPath)
	return nil
}

func printAST(node *sitter.Node, source []byte, depth, maxDepth int) {
	if depth > maxDepth {
		return
	}

	indent := strings.Repeat("  ", depth)
	nodeText := string(source[node.StartByte():node.EndByte()])
	if len(nodeText) > 50 {
		nodeText = nodeText[:50] + "..."
	}
	nodeText = strings.ReplaceAll(nodeText, "\n", "\\n")

	fmt.Printf("%s[%s] %q @ L%d:%d\n", indent, node.Type(), nodeText, node.StartPoint().Row+1, node.StartPoint().Column)

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(uint32(i))
		printAST(&child, source, depth+1, maxDepth)
	}
}

func printMatchingNodes(node *sitter.Node, source []byte, pattern string, depth int) {
	nodeText := string(source[node.StartByte():node.EndByte()])

	if strings.Contains(nodeText, pattern) {
		indent := strings.Repeat("  ", depth)
		display := nodeText
		if len(display) > 60 {
			display = display[:60] + "..."
		}
		display = strings.ReplaceAll(display, "\n", "\\n")
		fmt.Printf("%s[%s] %q @ L%d:%d\n", indent, node.Type(), display, node.StartPoint().Row+1, node.StartPoint().Column)
	}

	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(uint32(i))
		printMatchingNodes(&child, source, pattern, depth+1)
	}
}
