package crisp

import "fmt"
import "strings"

type formType int

const (
	ListType formType = iota
	IntType
	SymbolType
	StringType
)

type Visitor interface {
	VisitList(*List) (interface{}, error)
	VisitInt(*Int) (interface{}, error)
	VisitString(*String) (interface{}, error)
	VisitSymbol(*Symbol) (interface{}, error)
}

type Form interface {
	Tokens() []Token
	Children() []Form
	Accept(Visitor) (interface{}, error)
	Kind() formType
}

type List interface {
}

type List interface {
}

type Int interface {
}

type Symbol interface {
}

type List struct {
	items []Form
}

type String struct {
	token Token
}

type Int struct {
	token Token
}

type Symbol struct {
	token Token
}

func (str *String) Tokens() []Token {
	return []Token{str.token}
}

func (i *Int) Tokens() []Token {
	return []Token{i.token}
}

func (s *Symbol) Tokens() []Token {
	return []Token{s.token}
}

func (list *List) Tokens() []Token {
	return []Token{} //TODO
}

func (list *List) Children() []Form {
	return list.items //TODO
}
func (str *String) Children() []Form {
	return nil //TODO
}
func (i *Int) Children() []Form {
	return nil //TODO
}
func (s *Symbol) Children() []Form {
	return nil //TODO
}

func (list *List) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitList(list)
}
func (str *String) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitString(str)
}
func (i *Int) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitInt(i)
}
func (s *Symbol) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitSymbol(s)
}

func (s *Symbol) String() string {
	return s.token.value
}
func (str *String) String() string {
	return fmt.Sprintf("\"%s\"", str.token.value)
}
func (i *Int) String() string {
	return i.token.value
}

func (list *List) String() string {
	var itemsStr = make([]string, len(list.items))
	for i, item := range list.items {
		itemsStr[i] = fmt.Sprintf("%v", item)
	}
	return fmt.Sprintf("(%s)", strings.Join(itemsStr, " "))
}

func (str *String) Kind() formType {
	return StringType
}

func (i *Int) Kind() formType {
	return IntType
}

func (s *Symbol) Kind() formType {
	return SymbolType
}

func (list *List) Kind() formType {
	return ListType
}
