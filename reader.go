package crisp

import "errors"
import "unicode"
import "unicode/utf8"
import "log"

type tokenType int

const (
	StringToken tokenType = iota
	IntToken
	SymbolToken
	ParenOpenToken
	ParenCloseToken
	BracketOpenToken
	BracketCloseToken
)

type Token struct {
	kind  tokenType
	start int
	end   int
	value string
}

func isSymbolStart(r rune) bool {
	return !unicode.IsDigit(r) && !unicode.IsSpace(r) && !unicode.IsControl(r) && r != '(' && r != ')'
}

func isSymbolPart(r rune) bool {
	return isSymbolStart(r) || unicode.IsDigit(r)
}

func lex(input string) <-chan Token {
	tokens := make(chan Token, 1)
	position := 0
	length := len(input)
	go func() {
		defer close(tokens)
		for position < length {
			for {
				char, _ := utf8.DecodeRuneInString(input[position:])
				if unicode.IsSpace(char) {
					position++
				} else {
					break
				}
			}
			startPosition := position
			char, _ := utf8.DecodeRuneInString(input[position:])
			if char == '"' {
				position++
				var stringValue string
				for position < length && input[position] != '"' {
					stringValue = stringValue + string(input[position])
					position++
				}
				position++
				tokens <- Token{kind: StringToken, value: stringValue, start: startPosition, end: position}
				continue
			} else if unicode.IsDigit(char) {
				if char == '0' {
					//TODO
					continue
				}
				var intValue string
				for unicode.IsDigit(char) {
					intValue += string(char)
					position++
					if position >= length {
						break
					}
					char, _ = utf8.DecodeRuneInString(input[position:])
				}
				tokens <- Token{kind: IntToken, value: intValue, start: startPosition, end: position}
			} else if isSymbolStart(char) {
				var symbolValue string
				for isSymbolPart(char) {
					symbolValue += string(char)
					position++
					if position >= length {
						break
					}
					char, _ = utf8.DecodeRuneInString(input[position:])
				}
				tokens <- Token{kind: SymbolToken, value: symbolValue, start: startPosition, end: position}
			} else if char == '(' {
				tokens <- Token{kind: ParenOpenToken, value: "(", start: startPosition, end: position}
				position++
			} else if char == ')' {
				tokens <- Token{kind: ParenCloseToken, value: ")", start: startPosition, end: position}
				position++
			}
		}
	}()
	return tokens
}

func parse(token Token, tokens <-chan Token) (Form, error) {
	switch token.kind {
	case StringToken:
		return &String{token}, nil
	case IntToken:
		return &Int{token}, nil
	case SymbolToken:
		return &Symbol{token}, nil
	case ParenOpenToken:
		//		var start Token = token
		items := make([]Form, 0) // TODO Better guess of size
		for next := <-tokens; next.kind != ParenCloseToken; next = <-tokens {
			item, err := parse(next, tokens)
			if err != nil {
				return nil, err
			}
			items = append(items, item)
		}
		return &List{items}, nil
	}

	log.Println()

	return nil, errors.New("Could not parse") //TODO
}

type readResult struct {
	form Form
	err  error
}

func Read(input string) <-chan readResult {
	readForms := make(chan readResult, 1)
	go func() {
		defer close(readForms)
		tokens := lex(input)
		for token := range tokens {
			form, err := parse(token, tokens)
			if err != nil {
				readForms <- readResult{nil, err}
			} else {
				readForms <- readResult{form, nil}
			}
		}
	}()
	return readForms
}
