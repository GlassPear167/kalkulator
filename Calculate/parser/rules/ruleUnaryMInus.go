package rules

import (
	"fmt"
	"kalkulator/Calculate/model"
	"kalkulator/Calculate/parser/getpatern"
	"slices"
)

func RuleUnaryMinus(SliceOfToken []*model.Token, index int, source string) (error, []int) {

	leftSide := []int{model.Pluss, model.Minus, model.Times, model.Devide, model.LeftParenttesis, model.START}

	//-(-(-5)+7) >> kondisi
	rigthSide := []int{model.Number, model.LeftParenttesis}

	err := fmt.Errorf("code 8: Error : misplaced '%v' sign", SliceOfToken[index].GetRawValue(source))
	switch {
	case index == 0 && len(SliceOfToken) > 0:
		if !slices.Contains(rigthSide, SliceOfToken[index+1].Types) {
		}
		return err, SliceOfToken[index].GetSliceIndex()
	case index > 0:
		if !getpatern.LeftRightCheck(SliceOfToken, index, leftSide, rigthSide) {

			return err, SliceOfToken[index].GetSliceIndex()
		}

	}

	return nil, []int{}
}
