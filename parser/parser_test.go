package parser_test

import (
	"github.com/firesquid6/do-shell/parser"
	"github.com/firesquid6/do-shell/tree"
  "github.com/firesquid6/do-shell/token"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		input       []token.Token
		expectation *tree.Program
    expectedErr []error
	}{}

  for _, tt := range tests {
    p := parser.New(tt.input)
    p.ParseProgram()

    if len(p.Errors) != len(tt.expectedErr) {
      t.Errorf("Expected %d errors, got %d", len(tt.expectedErr), len(p.Errors))
    }

    for i, err := range p.Errors {
      if err != tt.expectedErr[i] {
        t.Errorf("Expected error %s, got %s", tt.expectedErr[i], err)
      }
    }

    if !reflect.DeepEqual(p.Tree, tt.expectation) {
      t.Log("Expected:")
      printProgram(tt.expectation, t)
      t.Log("Got:")
      printProgram(p.Tree, t)

      t.Errorf("Expected %s, got %s", tt.expectation, p.Tree)
    }
  }

}


func printProgram(program *tree.Program, t *testing.T) {
  t.Log(program.String())
}
