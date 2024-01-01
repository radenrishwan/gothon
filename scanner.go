package gothon

type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source:  source,
		tokens:  []Token{},
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current

		// scan token
		c := s.advance()
		switch c {
		// check line
		case '\n':
			s.line++
		case ' ':
		case '\r':
		case '\t':
		case '(':
			s.addTokenWithoutLiteral(LEFT_PAREN)
		case ')':
			s.addTokenWithoutLiteral(RIGHT_PAREN)
		case '{':
			s.addTokenWithoutLiteral(LEFT_BRACE)
		case '}':
			s.addTokenWithoutLiteral(RIGHT_BRACE)
		case ',':
			s.addTokenWithoutLiteral(COMMA)
		case '.':
			// TODO: check if before is not a number & after is not a number, if so, error
			s.addTokenWithoutLiteral(DOT)
		case '-':
			s.addTokenWithoutLiteral(MINUS)
		case '+':
			s.addTokenWithoutLiteral(PLUS)
		case ';':
			s.addTokenWithoutLiteral(SEMICOLON)
		case '*':
			s.addTokenWithoutLiteral(STAR)
		// slash or comment
		case '/':
			if s.match('/') {
				for s.peek() != '\n' && !s.isAtEnd() {
					s.advance()
				}
			} else {
				s.addTokenWithoutLiteral(SLASH)
			}
		case '!':
			if s.match('=') {
				s.addTokenWithoutLiteral(BANG_EQUAL)
			} else {
				s.addTokenWithoutLiteral(BANG)

			}
		case '=':
			if s.match('=') {
				s.addTokenWithoutLiteral(EQUAL_EQUAL)
			} else {
				s.addTokenWithoutLiteral(EQUAL)
			}
		case '<':
			if s.match('=') {
				s.addTokenWithoutLiteral(LESS_EQUAL)
			} else {
				s.addTokenWithoutLiteral(LESS)
			}
		case '>':
			if s.match('=') {
				s.addTokenWithoutLiteral(GREATER_THAN)
			} else {
				s.addTokenWithoutLiteral(GREATER)
			}
		case '"':
			for s.peek() != '"' && !s.isAtEnd() {
				if s.peek() == '\n' {
					s.line++
				}
				s.advance()
			}

			if s.isAtEnd() {
				Error(s.line, "unterminated string")
			}

			s.advance()

			value := s.source[s.start+1 : s.current-1]
			s.addToken(STRING, value)
		case 'o':
			if s.match('r') {
				s.addTokenWithoutLiteral(OR)
			}
		default:
			// check if it is a digit
			if isDigit(c) {
				s.number()

				continue
			}

			if isAlpha(c) {
				for isAlphaNumeric(s.peek()) {
					s.advance()
				}

				text := s.source[s.start:s.current]
				tokenType := keywords[text]

				if tokenType == "" {
					tokenType = IDENTIFIER
				}

				s.addTokenWithoutLiteral(tokenType)

				continue
			}

			Error(s.line, "unexpected character")
		}
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", "eof", s.line))

	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	s.current++

	return s.source[s.current-1]
}

func (s *Scanner) addToken(t TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(t, text, literal, s.line))
}

func (s *Scanner) addTokenWithoutLiteral(t TokenType) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(t, text, t, s.line))
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if (s.current + 1) >= len(s.source) {
		return '\000'
	}

	return s.source[s.current+1]
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}

	// check if there is a fraction
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	s.addToken(NUMBER, s.source[s.start:s.current])
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= 'A' && c <= 'Z' {
		return true
	}

	if c == '_' {
		return true
	}

	return false
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
