package lexer

import (
	"fmt"
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer/isValidNumber"

	"testing"
)

func TestIsValidNumber(t *testing.T) {
	count.Count("11,")
	input := count.BufferUtama
	bools := isValidNumber.IsValidNumber(input)
	fmt.Println(bools)

}

func TestLexsertoToken(t *testing.T) {

	m := "11+11986768787"
	count.Count(m)

	// count.BufferUtama = count.BufferUtama[:0]
	// count.Count("1,...	+afdg,3")
	// count.BufferUtama = count.BufferUtama[:0]
	// count.Count("-113,4+-231")
	// count.BufferUtama = count.BufferUtama[:0]
	// count.Count("0)-(-3)")

	input := count.BufferUtama
	sliceToken, err, errIndex := LexsertoToken(input)

	if err != nil {
		t.Error(err)
		fmt.Println("errIndexes:", errIndex)
	}
	for _, v := range sliceToken {

		fmt.Print(*v, ",", v.GetRawValue(m), ",", v.GetSliceIndex())
		fmt.Println()
	}

	fmt.Println("panjang input", len(input))

}
