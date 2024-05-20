package lexer

import (
  "testing"
)


func TestPeekAhead(t *testing.T) {
  source := "abcdefg"
  i := 2  // c

  if !peekAhead(source, 'd', &i) {
    t.Errorf("Expected to peek ahead and find 'd'")
  }
}
