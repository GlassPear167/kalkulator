package rules

import (
	"fmt"
	"kalkulator/Calculate/model"
	"slices"
)

func RuleLeftParenThesis(SliceOfToken []*model.Token, index int, source string) (error, []int) {
	//opr(unary,number
	rigthSide := []int{model.UnaryMinus, model.Number, model.RightParentesis, model.LeftParenttesis}
	err := fmt.Errorf(" code 1: Error : misplaced '%v' sign", string(SliceOfToken[index].GetRawValue(source)))

	if index == 0 && len(SliceOfToken) > 1 {

		if !slices.Contains(rigthSide, SliceOfToken[index].Types) {

			return err, SliceOfToken[index].GetSliceIndex()
		}
	}

	return nil, []int{}
}

func RuleRightParenThesis(SliceOfToken []*model.Token, index int, source string) (error, []int) {

	//number)opr
	leftside := []int{model.Number, model.LeftParenttesis, model.RightParentesis, model.UnaryMinus}
	err := fmt.Errorf(" code 2: Error : misplaced '%v' sign ", string(SliceOfToken[index].GetRawValue(source)))
	switch {
	case index < 2:
		return err, SliceOfToken[index].GetSliceIndex()
	case index >= 2:
		if !slices.Contains(leftside, SliceOfToken[index-1].Types) {

			return err, SliceOfToken[index].GetSliceIndex()
		}
	}

	return nil, []int{}
}
