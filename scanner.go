package main

import (
	"strconv"
)

const EndOfFile = 255

type scanner func() (Token)

type value struct {
	i   [2]int
	i64 int64
	f   float64
	p   interface{}
}

var scanners = make([]scanner, EndOfFile+1)

var tokenValue value

func setupScanner() {
	for i := uint8(0); i < EndOfFile; i++ {
		if isLetter(i) {
			scanners[i] = scanIdentifier
		} else if isDigital(i) {
			scanners[i] = scanNumericLiteral
		} else {
			scanners[i] = scanBadChar
		}
	}

	scanners[EndOfFile] = scanEOF
	scanners['\''] = scanCharLiteral
	scanners['"'] = scanStringLiteral
	scanners['+'] = scanPlus
	scanners['-'] = scanMinus
	scanners['*'] = scanStar
	scanners['/'] = scanSlash
	scanners['%'] = scanPercent
	scanners['<'] = scanLess
	scanners['>'] = scanGreat
	scanners['!'] = scanExclamation
	scanners['='] = scanEqual
	scanners['|'] = scanBar
	scanners['&'] = scanAmpersand
	scanners['^'] = scanCaret
	scanners['.'] = scanDot
	scanners['{'] = singleCharScanner
	scanners['}'] = singleCharScanner
	scanners['['] = singleCharScanner
	scanners[']'] = singleCharScanner
	scanners['('] = singleCharScanner
	scanners[')'] = singleCharScanner
	scanners[','] = singleCharScanner
	scanners[';'] = singleCharScanner
	scanners['~'] = singleCharScanner
	scanners['?'] = singleCharScanner
	scanners[':'] = singleCharScanner
	//todo: add keyword?
}

func isLetter(i uint8) (bool) {
	return (i >= 'a' && i <= 'z') || (i == '_') || (i >= 'A' && i <= 'Z')
}

func isDigital(i uint8) bool {
	return i >= '0' && i <= '9'
}

func isHexDigital(i uint8) bool {
	return isDigital(i) || (i >= 'A' && i < 'F')
}

func getNextToken() Token {
	var tok Token
	prevCoord = tokenCoord
	skipSpace()
	tokenCoord.line = input.line
	tokenCoord.col = input.Cursor() - input.lineHead + 1
	tok = scanners[input.Char()]()
	return tok
}

func scanIdentifier() (Token) {
	start := input.Cursor()

	if input.Char() == 'L' {
		if input.Peek() == '\'' {
			return scanCharLiteral()
		}
		if input.Peek() == '"' {
			return scanStringLiteral()
		}
	}
	input.Move()

	for ; isLetter(input.Char()) ||
		isDigital(input.Char()); {
		input.Move()
	}

	v := string(input.TokenTo(start, input.Cursor()))
	tok, ok := keywords[v]
	if !ok {
		tok = TK_ID
		tokenValue.p = v
	}
	return tok
}

func scanNumericLiteral() (Token) {
	start := input.Cursor()
	base := 10

	if input.Char() == '0' && input.Peek() == 'x' {
		input.Move().Move()
		start = input.Cursor()
		base = 16

		if !isHexDigital(input.Char()) {
			Error(&tokenCoord, "expect hex digit")
			tokenValue.i[0] = 0
			return TK_INT_CONST
		}

		for ; isDigital(input.Char()); {
			input.Move()
		}
	} else if input.Char() == '0' {
		base = 8
		input.Move()
		for ; isDigital(input.Char()); {
			input.Move()
		}
	} else {
		input.Move()
		for ; isDigital(input.Char()); {
			input.Move()
		}
	}

	if input.Char() == '.' {
		return scanFloatLiteral(start)
	} else {
		return scanIntLiteral(start, input.Cursor()-start, base)
	}
}

// .01 , 0.1 , 1e2 , -0.1 , 1e-1 , 1E1 , 1E-1
func scanFloatLiteral(start int) Token {
	// only parse as example 0.1
	input.Move()
	for ; isDigital(input.Char()); {
		input.Move()
	}

	v, err := strconv.ParseFloat(string(input.TokenTo(start, input.Cursor())), 64)
	if err != nil {
		Warning(&tokenCoord, "float literal error:%v", err)
	}

	tokenValue.f = v
	return TK_FLOAT_CONST
}

func scanIntLiteral(start int, len int, base int) Token {
	s := string(input.Token(start, len))
	v, err := strconv.ParseInt(s, base, 64)
	if err != nil {
		Error(&tokenCoord, "integer literal error:%v", err)
	}
	tokenValue.i64 = v
	return TK_INT_CONST
}

func scanBadChar() (Token) {
	Error(&tokenCoord, "illegal character:%c", input.Char())
	input.Move()
	return getNextToken()
}

func scanEOF() Token {
	return TK_END
}

// \n??
func scanCharLiteral() Token {

	ch := uint8(0)
	count := 0

	input.Move()
	for ; input.Char() != '\''; {
		if input.Char() == '\n' || input.Char() == EOF {
			break
		}

		if input.Char() == '\\' {
			ch = scanEscapeChar()
		} else {
			ch = input.Char()
			input.Move()
		}
		count++
	}

	if input.Char() != '\'' {
		Error(&tokenCoord, "expect '")
		goto end
	}
	input.Move()
	if count > 1 {
		Warning(&tokenCoord, "too many characters")
	}

end:
	tokenValue.i[0] = int(ch)
	tokenValue.i[1] = 0
	return TK_INT_CONST
}

func scanEscapeChar() uint8 {
	input.Move()
	ch := input.Char()
	input.Move()
	switch ch {
	case 'a':
		return '\a'
	case 'n':
		return 'n'
	case 'r':
		return '\r'
	case 't':
		return '\t'
	case '\'', '"', '?', '\\':
		return ch
		//case 'x':
		// ignore parse hex literal
		// ignore parse oct literal
	default:
		Warning(&tokenCoord, "unrecognized escape sequence:\\%c", input.Char())
		return input.Char()
	}
}

func scanStringLiteral() Token {

	ch := uint8(0)
	len := 0

	var cp []uint8

	input.Move()
	for ; input.Char() != '"'; {
		if input.Char() == '\n' || input.Char() == EOF {
			break
		}

		if input.Char() == '\\' {
			ch = scanEscapeChar()
		} else {
			ch = input.Char()
			input.Move()
		}
		len++
		cp = append(cp, ch)
	}

	if input.Char() != '"' {
		Error(&tokenCoord, "expect \"")
		goto end
	}
	input.Move()
end:
	tokenValue.p = cp
	return TK_STRING
}

// += ,++ ,+
func scanPlus() Token {
	if input.Char() != '+' {
		Fatal("unexpected character")
	}
	input.Move()
	if input.Char() == '+' {
		input.Move()
		return TK_INC
	} else if input.Char() == '=' {
		input.Move()
		return TK_ADD_ASSIGN
	} else {
		return TK_ADD
	}
}

// - , -= , -> , --
func scanMinus() Token {
	input.Move()
	if input.Char() == '-' {
		input.Move()
		return TK_DEC
	} else if input.Char() == '=' {
		input.Move()
		return TK_SUB_ASSIGN
	} else if input.Char() == '>' {
		input.Move()
		return TK_POINTER
	} else {
		return TK_SUB
	}
}

// * , *=
func scanStar() Token {
	input.Move()
	if input.Char() == '=' {
		input.Move()
		return TK_MUL_ASSIGN
	} else {
		return TK_MUL
	}
}

// / , /=
func scanSlash() Token {
	input.Move()
	if input.Char() == '=' {
		input.Move()
		return TK_DIV_ASSIGN
	} else {
		return TK_DIV
	}
}

// % ,%=
func scanPercent() Token {
	input.Move()
	if input.Char() == '=' {
		input.Move()
		return TK_MOD_ASSIGN
	} else {
		return TK_MOD
	}
}

// < , <<, <<=, <=,
func scanLess() Token {
	input.Move()
	if input.Char() == '<' {
		input.Move()
		if input.Char() == '=' {
			input.Move()
			return TK_LSHIFT_ASSIGN
		} else {
			return TK_LSHIFT
		}
	} else if input.Char() == '=' {
		input.Move()
		return TK_LESS_EQ
	} else {
		return TK_LESS
	}
}

// > , >> , >>= , >=
func scanGreat() Token {
	input.Move()
	if input.Char() == '>' {
		input.Move()
		if input.Char() == '=' {
			input.Move()
			return TK_RSHIFT_ASSIGN
		} else {
			return TK_RSHIFT
		}
	} else if input.Char() == '=' {
		input.Move()
		return TK_GREAT_EQ
	} else {
		return TK_GREAT
	}
}

// ! , !=
func scanExclamation() Token {
	input.Move()
	if input.Char() == '=' {
		input.Move()
		return TK_UNEQUAL
	} else {
		return TK_NOT
	}
}

// = , ==
func scanEqual() Token {
	input.Move()
	if input.Char() == '=' {
		input.Move()
		return TK_EQUAL
	} else {
		return TK_ASSIGN
	}
}

// | , |= , ||
func scanBar() Token {
	input.Move()
	if input.Char() == '|' {
		input.Move()
		return TK_OR
	} else if input.Char() == '=' {
		input.Move()
		return TK_BITOR_ASSIGN
	} else {
		return TK_BITOR
	}
}

// & , &= , &&
func scanAmpersand() Token {
	input.Move()
	if input.Char() == '&' {
		input.Move()
		return TK_AND
	} else if input.Char() == '=' {
		input.Move()
		return TK_BITAND_ASSIGN
	} else {
		return TK_BITAND
	}
}

// ^ , ^=
func scanCaret() Token {
	//todo
	input.Move()
	return TK_UNKNOWN
}

// . , 0.1 , ..
func scanDot() Token {
	if isDigital(input.Peek()) {
		return scanFloatLiteral(input.Cursor())
	}

	if input.Peek() == '.' && input.Peek2() == '.' {
		input.Move().Move()
		return TK_ELLIPSE
	} else {
		input.Move()
		return TK_DOT
	}
}

func singleCharScanner() Token {
	ch := input.Char()
	input.Move()
	switch ch {
	case '{':
		return TK_LEFT_BRACE
	case '}':
		return TK_RIGHT_BRACE
	case '[':
		return TK_LEFT_BRACKET
	case ']':
		return TK_RIGHT_BRACKET
	case '(':
		return TK_LEFT_PAREN
	case ')':
		return TK_RIGHT_PAREN
	case ',':
		return TK_COMMA
	case ';':
		return TK_SEMICOLON
	case '~':
		return TK_COMP
	case '?':
		return TK_QUESTION
	case ':':
		return TK_COLON
	default:
		return TK_BEGIN
	}
}

func scanPreProcessLine() {
	//line := 0
	input.Move()
	for input.Char() == ' ' || input.Char() == '\t' {
		input.Move()
		goto readline
	}
readline:
	for ; input.Char() != '\n' && input.Char() != EndOfFile; {
		input.Move()
	}
}

func scanComment() bool {
	if input.Peek() != '/' && input.Peek() != '*' {
		return false
	}
	input.Move()

	if input.Char() == '/' {
		// // comment
		input.Move()
		for ; input.Char() != '\n' && input.Char() != EndOfFile; {
			input.Move()
		}
	} else {
		// /* comment
		input.Move().Move()
		for ; input.Char() != '*' || input.Peek() != '/'; {
			if input.Char() == '\n' {
				tokenCoord.ppline++
				input.line++
			} else if input.Char() == EndOfFile || input.Peek() == EndOfFile {
				Error(&tokenCoord, "comment is not closed")
				return false
			}
			input.Move()
		}
		input.Move().Move()
	}
	return true
}
func finKeyword(str []byte) (Token) {
	return keywords[string(str)]
}

func skipSpace() {
	ch := input.Char()
	for ; ch == '\t' ||
		ch == '\v' ||
		ch == ' ' ||
		ch == '\r' ||
		ch == '\n' ||
		ch == '/' ||
		ch == '#'; {

		switch ch {
		case '\n':
			// move to next line
			tokenCoord.ppline++
			input.line++
			input.Move()
			input.lineHead = input.Cursor()
		case '#':
			scanPreProcessLine()
		case '/':
			if !scanComment() {
				return
			}
		default:
			input.Move()
		}
		ch = input.Char()
	}
}
