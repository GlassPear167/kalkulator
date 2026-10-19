package lexeringNumber

import (
	"fmt"

	"kalkulator/Calculate/lexer/isValidNumber"
	"kalkulator/Calculate/model"
	"slices"
)

/*
	model loop biasa bukan range

output dari int pada fungsi lexeringNumber akan memberikan token juga angka sejauh mana index bergeser
*/
func LexeringNumber(index int, inputRune []rune) (*model.Token, []int, error) {

	buffer := []rune{}
	bufferIndx := []int{}

	for i := index; i < len(inputRune); i++ {
		v := inputRune[i]
		thisIsANumber := slices.Contains(model.Numbers, v)
		if thisIsANumber {
			buffer = append(buffer, v)
			bufferIndx = append(bufferIndx, i)
			continue
		}
		if !thisIsANumber {
			break
		}
	}

	// validate
	if !isValidNumber.IsValidNumber(buffer) {
		return model.Newtoken(model.Number).ToNotValid(), bufferIndx, fmt.Errorf("bukan nomor yang valid: %v", string(buffer))
	}
	token := model.Newtoken(model.Number).AddIndex(bufferIndx[0], bufferIndx[len(buffer)-1])
	return token, bufferIndx, nil
}
