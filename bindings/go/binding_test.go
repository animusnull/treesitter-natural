package tree_sitter_treesitter_natural_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_treesitter_natural "github.com/animusnull/treesitter-natural/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_treesitter_natural.Language())
	if language == nil {
		t.Errorf("Error loading TreesitterNatural grammar")
	}
}
