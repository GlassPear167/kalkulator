package rules

import (
	"fmt"

	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/printb"
	"testing"
)

func TestPluss(t *testing.T) {

	m := "3+*2"
	count.Count(m)
	// count.Count("3+*-/2")
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("before")
	printb.PrintDetail(tokens, m)

	fmt.Println("after")

	err, indexErr = RulePlus(tokens, 1, m)
	if err != nil {

		t.Error(err, indexErr)
	}
}
func TestMinus(t *testing.T) {

	m := "3+*-/2"
	count.Count(m)
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("before")
	printb.PrintDetail(tokens, m)

	fmt.Println("after")

	err, indexErr = RuleMinus(tokens, 4, m)
	if err != nil {
		t.Error(err, indexErr)
	}
}
func TestTimes(t *testing.T) {
	m := "3+*-/2"
	count.Count(m)
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("before")
	printb.PrintDetail(tokens, m)

	fmt.Println("after")

	err, indexErr = RuleTimes(tokens, 2, m)
	if err != nil {
		t.Error(err, indexErr)
	}
}
func TestDevide(t *testing.T) {

	m := "3+*-/2"
	count.Count(m)
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("before")
	printb.PrintDetail(tokens, m)

	fmt.Println("after")

	err, indexErr = RuleDevide(tokens, 4, m)
	if err != nil {
		t.Error(err, indexErr)
	}
}
