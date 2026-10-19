package getpatern

import (
	"fmt"
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/printb"

	"testing"
)

// done
// func TestGetminusPatern(t *testing.T) {

// 	utils.Count("0)-(-3)")
// 	utils.BufferUtama = utils.BufferUtama[:0]
// 	utils.Count("-3+(-2)-3(-7)+4")

// 	input := utils.BufferUtama
// 	for _, v := range input {
// 		fmt.Print(string([]rune{v}))
// 	}

// 	tokens, err, indexerr := lexer.LexsertoToken(input)
// 	if err == nil {

// 		t.Error(err, "index:", indexerr)

// 	}

// 	fmt.Println()
// 	fmt.Println("token")
// 	fmt.Println()
// 	for _, v := range tokens {
// 		fmt.Println("value:", string(v.Value), "rawIndex:", v.RawIndex, "tipe token:", v.Types)
// 	}

// 	fmt.Println()
// 	fmt.Println("result:")
// 	result := GetminusPatern(tokens)
// 	for _, v := range result {
// 		fmt.Print(*v)
// 	}
// 	fmt.Println()
// 	for _, v := range result {
// 		fmt.Println("value:", string(v.Value), "rawIndex:", v.RawIndex, "tipe token:", v.Types)
// 	}
// }

func TestGetunaryMinusPatern(t *testing.T) {
	m := "-3+(-2)-3(-7)+4"
	count.Count(m)

	input := count.BufferUtama

	tokens, err, indexerr := lexer.LexsertoToken(input)
	if err == nil {

		t.Error(err, "index:", indexerr)

	}

	fmt.Println("token Before")
	printb.PrintDetail(tokens, m)

	tokenss := GetUnaryMinusPatern(tokens)

	fmt.Println("token after")
	printb.PrintDetail(tokenss, m)

}

func TestA(t *testing.T) {
	fmt.Println(len([]int{0, 2}))
}

func TestGetImplisitTimesPatern(t *testing.T) {

	m := "()()()()"
	// utils.Count("-3+(-2)(-3)(-7)((+4))")
	m = "1+2(-(3+4)-3/5)+3*8-10"

	count.Count(m)

	input := count.BufferUtama

	tokens, err, indexerr := lexer.LexsertoToken(input)
	if err == nil {

		t.Error(err, "index:", indexerr)

	}

	tokenss := GetImplisitTimesPatern(tokens)

	printb.PrintDetail(tokenss, m)

}
