package isValidNumber

import (
	"fmt"
	"kalkulator/Calculate/count"
	"testing"
)

func TestIsValidNumber(t *testing.T) {
	count.Count("11,")
	input := count.BufferUtama
	bools := IsValidNumber(input)
	fmt.Println(bools)

}

func TestTek(t *testing.T) {
	fmt.Println("mantap", string([]rune{9}), string([]rune{32}))
	fmt.Println("mantap", "spasi:", []rune(" "), "tab:", []rune("	"))
}
