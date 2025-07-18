package lexer

import (
	"testing"

	"interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `foo = 4;
bar = foo + 3;
baz = fn() bar - 1;
blah = fn(x y) x * y;
x == y ? 1 : 0;
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "foo"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "bar"},
		{token.ASSIGN, "="},
		{token.IDENT, "foo"},
		{token.PLUS, "+"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "baz"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.IDENT, "bar"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "blah"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.IDENT, "x"},
		{token.ASTERISK, "*"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "x"},
		{token.EQ, "=="},
		{token.IDENT, "y"},
		{token.QUESTION, "?"},
		{token.INT, "1"},
		{token.COLON, ":"},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
