package cli

import (
	"strings"
	"testing"
)

func TestScanner_Success(t *testing.T) {
	var tests = []struct {
		s   string
		tok Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: EOF},
		{s: `#`, tok: ILLEGAL, lit: `#`},
		{s: ` `, tok: WS, lit: " "},
		{s: "\t", tok: WS, lit: "\t"},
		{s: "\n", tok: WS, lit: "\n"},

		// Misc characters
		{s: `*`, tok: ASTERISK, lit: "*"},
		{s: `,`, tok: COMMA, lit: ","},
		{s: `=`, tok: EQUAL, lit: "="},

		// Identifiers
		{s: `foo`, tok: IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: IDENT, lit: `Zx12_3U_`},
		{s: `1`, tok: IDENT, lit: `1`},

		// Keywords
		{s: `FROM`, tok: FROM, lit: "FROM"},
		{s: `SELECT`, tok: SELECT, lit: "SELECT"},
		{s: `WHERE`, tok: WHERE, lit: "WHERE"},
	}

	for i, tt := range tests {
		s := NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
