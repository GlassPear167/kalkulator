package evaluator

import (
	"fmt"
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/parser"
	"kalkulator/Calculate/printb"
	"testing"
)

func TestManatap(t *testing.T) {
	n := "1234567890,"

	// fmt.Println(n[0])

	b := fmt.Errorf("%v", n[0])
	c := fmt.Errorf("%.15f", (1 + 0.00000001))
	if b != nil {
		t.Errorf("%v", b)
		t.Errorf("%v", c)
	}
}

func TestEvalauator1(t *testing.T) {

	m := "-(-1)"
	m = "1+2(-(3+4)-3/5)+3*8-10*8*8*8*8*4/9*293877909834092380298879786782348*(-1)"
	// m = "1+2(-(3+4)-3)"
	// m = "2(7)"

	// m = "1+4"
	count.Count(m)
	input := count.BufferUtama
	tokens, err, indexErr := lexer.LexsertoToken(input)

	fmt.Print("hasil lexer  : ")
	for _, v := range tokens {

		fmt.Print(v.GetRawValue(m))
	}

	fmt.Println()
	printb.PrintDetail(tokens, m)

	if err != nil {
		t.Error(err, indexErr)
	}
	tokens, err, indexErr = parser.Parser(tokens, m)

	fmt.Print("hasil parser  :  ")
	for _, v := range tokens {

		fmt.Print(v.GetRawValue(m))
	}
	fmt.Println()
	printb.PrintDetail(tokens, m)

	if err != nil {

		t.Error(err, indexErr)
	}

	hasil, err, indexErr := Evaluate(tokens, m)

	if err != nil {
		t.Error(err, indexErr)
	}

	fmt.Println("Hasil: ", hasil)

}

func TestR(t *testing.T) {
	stack := []int{3, 4}

	A := stack[len(stack)-1]

	stack = stack[:len(stack)-1] //pop

	B := stack[len(stack)-1]

	stack = stack[:len(stack)-1] //pop

	result := A + B

	stack = append(stack, result)

	fmt.Println(A, B, stack)
}
func TestPOP(t *testing.T) {
	stak := []float64{1, 2, 3, 4}

	a, stak := pop(stak)
	fmt.Println(stak, a)
}

func TestAppend(t *testing.T) {

	n := []float64{1, 2, 3}
	n = append(n, 4)
	fmt.Println(n)
	a, n := pop(n)
	fmt.Println(a, n)

}
