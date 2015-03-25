package crisp

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadString(t *testing.T) {
	var input = "\"hello world\""
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, StringToken, read.form.Tokens()[0].kind)
	require.Equal(t, "hello world", read.form.Tokens()[0].value)
}

func TestReadInteger(t *testing.T) {
	var input = "2"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, IntToken, read.form.Tokens()[0].kind)
	require.Equal(t, "2", read.form.Tokens()[0].value)
}

func TestReadIntegerTwoDigits(t *testing.T) {
	var input = "65"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, IntToken, read.form.Tokens()[0].kind)
	require.Equal(t, "65", read.form.Tokens()[0].value)
}

func TestReadIntegerMoreDigits(t *testing.T) {
	var input = "329158921"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, IntToken, read.form.Tokens()[0].kind)
	require.Equal(t, "329158921", read.form.Tokens()[0].value)
}

func TestReadSymbol(t *testing.T) {
	var input = "ok"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, SymbolToken, read.form.Tokens()[0].kind)
	require.Equal(t, "ok", read.form.Tokens()[0].value)
}

func TestReadSymbolDashes(t *testing.T) {
	var input = "do-that-thing??"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, SymbolToken, read.form.Tokens()[0].kind)
	require.Equal(t, "do-that-thing??", read.form.Tokens()[0].value)
}

func TestReadSkipSpaces(t *testing.T) {
	var input = "   \n  \n\t\nspace"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Tokens()))
	require.Equal(t, SymbolToken, read.form.Tokens()[0].kind)
	require.Equal(t, "space", read.form.Tokens()[0].value)
}

func TestReadList(t *testing.T) {
	var input = "(a list of things)"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 4, len(read.form.Children()))
}

func TestReadNestedList(t *testing.T) {
	var input = "(a list (of a list))"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 3, len(read.form.Children()))
}

func TestReadVeryNestedList(t *testing.T) {
	var input = "(((((((a \"list\" 20)))))))"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 1, len(read.form.Children()))
	require.Equal(t, input, fmt.Sprintf("%v", read.form))
}

func TestReadMacro(t *testing.T) {
	var expander MacroExpander
	var input = "(defn my-name () 2)"
	read := <-Read(input)
	require.Nil(t, read.err)
	require.Equal(t, 4, len(read.form.Children()))
	expanded, err := expander.Visit(read.form)
	require.Nil(t, err)
	require.Equal(t, input, fmt.Sprintf("%v", read.form))
	require.Equal(t, "(define my-name (fn () 2))", fmt.Sprintf("%v", expanded))
}
