package lexer

import (
	"testing"

	"github.com/Anand-Chaudhary/interpreter/token"
)

func TestNextToken(t *testing.T){
	input :+ `=+(){},;`

	tests := []struct {
		expectedType	token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"}
	}
}
