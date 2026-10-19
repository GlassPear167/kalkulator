package getpatern

import (
	"kalkulator/Calculate/model"
	"slices"
)

// mendapatkan matching patern pertama dari sliceSourcePatern
func GetPatern(patern []int, sliceSourcePatern []int) bool {

	trunes := []bool{}

	if len(sliceSourcePatern) < len(patern) {
		return false
	}
	for i, _ := range patern {
		if patern[i] == sliceSourcePatern[i] {
			trunes = append(trunes, true)

		} else {
			trunes = append(trunes, false)
		}
	}

	if slices.Contains(trunes, false) {
		return false
	}
	return true
}

// memeriksa apakah ada karakter yang cocok di kiri dan kanan
func LeftRightCheck(source []*model.Token, index int, leftSide []int, rightside []int) bool {

	if slices.Contains(leftSide, source[index-1].Types) && slices.Contains(rightside, source[index+1].Types) {
		return true
	}
	return false
}

func GetUnaryMinusPatern(SliceOfToken []*model.Token) []*model.Token {

	buffer := []*model.Token{}

	token := model.Newtoken(model.UnaryMinus).ToBlank("-")

	//filter
	typeToken := []int{}
	for _, v := range SliceOfToken {
		typeToken = append(typeToken, v.Types)
	}
	// patern := []int{model.LeftParenttesis, model.Minus, model.Number, model.RightParentesis}
	// panjanglompatan := len(patern) - 1

	for i, v := range SliceOfToken {
		if v.Types == model.Minus {

			switch {
			case i == 1:
				// token := model.Newtoken(model.UnaryMinus, v.Value).ReplaceAllIndexWith(v.RawIndex)
				buffer = append(buffer, token)
				continue
			case i > 1:

				left := []int{model.Pluss, model.Minus, model.Times, model.Devide, model.LeftParenttesis}
				if slices.Contains(left, SliceOfToken[i-1].Types) {
					buffer = append(buffer, token)
					continue
				}

			}
		}
		buffer = append(buffer, v)

	}

	return buffer
}

func GetImplisitTimesPatern(sliceOfToken []*model.Token) []*model.Token {

	buffer := []*model.Token{}
	filteredTokenBytype := []int{}

	for _, v := range sliceOfToken {
		filteredTokenBytype = append(filteredTokenBytype, v.Types)
	}

	patern := []int{model.RightParentesis, model.LeftParenttesis} //
	patern1 := []int{model.Number, model.LeftParenttesis}
	patern2 := []int{model.RightParentesis, model.Number}

	for i, v := range sliceOfToken {

		buffer = append(buffer, v)
		if GetPatern(patern, filteredTokenBytype[i:]) || GetPatern(patern1, filteredTokenBytype[i:]) || GetPatern(patern2, filteredTokenBytype[i:]) {

			token := model.Newtoken(model.Times).ToBlank("*")

			buffer = append(buffer, token)

		}

	}

	return buffer
}
