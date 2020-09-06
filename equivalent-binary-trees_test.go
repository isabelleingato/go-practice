package main

import(
	"testing"
	"golang.org/x/tour/tree"
);

func TestEquivalentBinaryTrees(t *testing.T) {
	if (!Same(tree.New(1), tree.New(1))) {
		t.Errorf("Expected true but was false")
	}
	if (Same(tree.New(1), tree.New(2))) {
		t.Errorf("Expected false but was true")
	}
}