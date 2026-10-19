package sorting

import (
	"fmt"
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/printb"

	"testing"
)

func TestShuntingYardShorting(t *testing.T) {

	teks := "1-(1+2)*3/4" // harus : 1 1 2 + 3 * 4 / -
	teks = "1+2(-(3+4)-3/5)+3*8-10"
	fmt.Println(teks)
	count.Count(teks)
	input := count.BufferUtama

	tokens, err, indexErr := lexer.LexsertoToken(input)

	if err != nil {
		t.Error(err, indexErr)
	}

	if err != nil {
		t.Error(err, indexErr)
	}
	mantap := ShuntingYardShorting(tokens)
	for _, v := range mantap {
		fmt.Print(string(v.GetRawValue(teks)))
	}

	fmt.Println()
	printb.PrintDetail(mantap, teks)

}

func TestLoopBackWard(t *testing.T) {
	operatorStack := []int{1, 2, 3, 5, 6}

	returnQueue := []int{}

	for len(operatorStack) != 0 {
		returnQueue = append(returnQueue, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	fmt.Println(returnQueue)
}

func moveBackwardTest(from []int, jumlah int) ([]int, []int) {

	buffer := []int{}
	for i := jumlah; i > 0; i-- {

		if len(from) != 0 {

			buffer = append(buffer, from[len(from)-1])
			from = from[:len(from)-1]
		}
	}

	return buffer, from
}

func TestMoveBackwardTest(t *testing.T) {

	m := []int{1, 2, 3, 4, 5}

	count := 0

	for i := len(m) - 1; 0 < i; i-- {
		if m[i] == 2 {
			break
		}
		count++
		fmt.Println(count)
	}
	fmt.Println()
	fmt.Println(count)
	n, m := moveBackwardTest(m, count)

	fmt.Println(n, m)
}
