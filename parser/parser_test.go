package parser

import (
	"interpreter/ast"
	"interpreter/lexer"
	"testing"
)

func TestAssignStatements(t *testing.T) {
	input := `
foo = 4
bar = 57
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 2 {
		t.Fatalf("program does not contain 2 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"foo"},
		{"bar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testAssignStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testAssignStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "" {
		t.Errorf("s.TokenLiterla not an empty string, go =%q", s.TokenLiteral())
		return false
	}

	assignStmt, ok := s.(*ast.AssignmentStatement)
	if !ok {
		t.Errorf("s not *ast.AssignmentStatment, got=%T", s)
		return false
	}

	if assignStmt.Name.Value != name {
		t.Errorf("assignStmt name not '%s', got=%s", name, assignStmt.Name.Value)
		return false
	}

	if assignStmt.Name.TokenLiteral() != name {
		t.Errorf("assignStmt name token literal not '%s', got=%s", name, assignStmt.Name.TokenLiteral())
		return false
	}

	return true
}
