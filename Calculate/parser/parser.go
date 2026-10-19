package parser

import (
	"kalkulator/Calculate/model"
	"kalkulator/Calculate/parser/getpatern"
	"kalkulator/Calculate/parser/sorting"
	"kalkulator/Calculate/parser/validator"
)

func Parser(SliceOfToken []*model.Token, source string) ([]*model.Token, error, []int) {

	// tahap 1 :get semua patern yang ada
	tokens := getpatern.GetImplisitTimesPatern(SliceOfToken)

	final := getpatern.GetUnaryMinusPatern(tokens)

	// tahap 2 :validasi semua patern
	err, indexes := validator.Check(final, source)
	if err != nil {
		return []*model.Token{}, err, indexes
	}

	final = sorting.ShuntingYardShorting(final)
	// tahap 3
	return final, nil, []int{}
}
