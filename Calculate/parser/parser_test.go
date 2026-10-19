package parser

import (
	"fmt"
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/parser/getpatern"
	"kalkulator/Calculate/printb"
	"testing"
)

func a(b []int, i int) []int {
	return b[:i-1]
}

func TestSringSlice(t *testing.T) {
	buffer := []int{0, 1, 2, 3}
	fmt.Println(len(buffer))
	fmt.Println(buffer[2:])
	buffer = a(buffer, 4)
	buffer2 := make([]int, len(buffer))[:0]
	fmt.Println(len(buffer), len(buffer2), cap(buffer2))

}
func TestGetPaternTest(t *testing.T) {
	fmt.Println(getpatern.GetPatern([]int{0, 1, 2, 3, 4}, []int{0, 1, 2}))

}

func TestParser(t *testing.T) {
	//"1+2+3+5+6+(9/(9-9)-(-6)(-8)*5))/8"
	// utils.Count("1+2+3+5+6(-1+3)/8")
	// utils.Count("1+2+3+5+6+(9/(9-9)-(-6)((-8)*5))/8")

	// utils.Count("1+2+3+5+6+(9/(9-9)-(-6q)((-8)*5))/8+*/()")
	// utils.Count("1+2+3+5+6+(9/(9-9)-(-6)((-8)*5))/8+*/()")
	// m := "1+2+3+5+6+(9/(9-9)-(-6)(-(-8)*5))/8/()"
	m := "-(-1)"
	m = "1+2(-(3+4)-3/5)+3*8-10"
	count.Count(m)
	// count.Count("3+*(0/0)()")
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("before")
	printb.PrintDetail(tokens, m)

	fmt.Println("after")
	final, err, indexErr := Parser(tokens, m)
	if err != nil {
		t.Error(err, indexErr)
	}

	printb.PrintDetail(final, m)

}

func TestM(t *testing.T) {

	m := []int{1, 2, 3, 4, 5}
	m = append(m, m[0])
	fmt.Println(m)

}
