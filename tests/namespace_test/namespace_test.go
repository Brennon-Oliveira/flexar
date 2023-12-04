package namespace_test

import (
	"github.com/Brennon-Oliveira/flexar/compilation"
	"github.com/Brennon-Oliveira/flexar/program_data"
	"github.com/Brennon-Oliveira/flexar/utils"
	"testing"
)

func TestNamespaceSemantic(t *testing.T) {
	var files []string
	compilation.FindFiles(".", &files)

	for _, file := range files {
		utils.SetCurrentFile(file)
		compilation.Compile(file)
	}

	namespaces := program_data.GetNamespaces()
	for _, namespace := range namespaces {
		t.Logf("Namespace %s: \n", namespace.Name)
		for _, file := range namespace.Files {
			t.Logf("\t%s\n", file)
		}
	}

	if len(namespaces) != 2 {
		t.Errorf("Expected 2 namespaces, got %d", len(namespaces))
	}

	if namespaces["aux.test"].Name != "aux.test" {
		t.Errorf("Expected namespace_test, got %s", namespaces["aux.test"].Name)
	}

	if namespaces["aux.test"].Files[0] != "aux.test.fl" {
		t.Errorf("Expected aux.test.fl, got %s", namespaces["aux.test"].Files[0])
	}

}
