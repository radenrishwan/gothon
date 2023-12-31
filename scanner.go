package gothon

import "log"

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
		default:
			// TODO: change to error later
			log.Println("unexpected token: ", string(c))
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
