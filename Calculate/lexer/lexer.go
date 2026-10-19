package lexer

import (
	"fmt"
	"kalkulator/Calculate/lexer/lexeringNumber"
	"kalkulator/Calculate/model"
	"slices"
)

/*
output  int []int "error" = kumpulan index nilai yang error
*/
func LexsertoToken(input []rune) ([]*model.Token, error, []int) {
	sliceOfToken := []*model.Token{}

	token := model.Newtoken(model.START).ToBlank("START")
	sliceOfToken = append(sliceOfToken, token)

	for i := 0; i < len(input); i++ {
		v := input[i]
		if !slices.Contains(model.IncludeChar(), v) {

			return sliceOfToken[:0], fmt.Errorf("Ada karakter yang tidak sesuai: %v", string(v)), []int{i}
		}

		//Pengecekan anggka
		if slices.Contains(model.Numbers, v) {
			token, ofset, err := lexeringNumber.LexeringNumber(i, input)

			if err != nil {
				return sliceOfToken[:0], err, ofset

			}

			sliceOfToken = append(sliceOfToken, token)
			i += len(ofset) - 1
			continue

		}

		// pengecekan operator

		thisIsAOperator := slices.Contains(model.Oprt, v)

		if thisIsAOperator {
			switch v {
			case 42:
				token := model.Newtoken(model.Times).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)

			case 43:

				token := model.Newtoken(model.Pluss).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)
			//	sliceOfToken = append(sliceOfToken, model.Newtoken(model.Pluss, "+", i))

			case 45:

				token := model.Newtoken(model.Minus).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)
			//	sliceOfToken = append(sliceOfToken, model.Newtoken(model.Minus, "-", i))
			case 40:
				token := model.Newtoken(model.LeftParenttesis).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)
			//	sliceOfToken = append(sliceOfToken, model.Newtoken(model.LeftParenttesis, "(", i))
			case 41:
				token := model.Newtoken(model.RightParentesis).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)
				//	sliceOfToken = append(sliceOfToken, model.Newtoken(model.RightParentesis, ")", i))
			case 47:
				token := model.Newtoken(model.Devide).AddIndex(i, i)
				sliceOfToken = append(sliceOfToken, token)
				//	sliceOfToken = append(sliceOfToken, model.Newtoken(model.Devide, "/", i))

			}
		}

	}

	token = model.Newtoken(model.EOF).AddIndex(len(input), len(input))
	sliceOfToken = append(sliceOfToken, token)
	return sliceOfToken, nil, []int{-1}
}
