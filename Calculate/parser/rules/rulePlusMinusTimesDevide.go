package rules

import (
	"fmt"
	"kalkulator/Calculate/model"
	"kalkulator/Calculate/parser/getpatern"
)

func innerFunc(SliceOfToken []*model.Token, index int, err error) (error, []int) {

	var leftside = []int{model.Number, model.RightParentesis}
	var rigthSide = []int{model.Number, model.LeftParenttesis, model.UnaryMinus}

	// aturan pernulisan minus:
	//

	switch {
	case index == 0:
		return err, SliceOfToken[index].GetSliceIndex()
	case index >= 1:
		if !getpatern.LeftRightCheck(SliceOfToken, index, leftside, rigthSide) {

			return err, SliceOfToken[index].GetSliceIndex()
		}

	}

	return nil, []int{}

}

// aturan untu +-*/
func RulePlus(SliceOfToken []*model.Token, index int, source string) (error, []int) {

	err := fmt.Errorf("code 7: code Error : misplaced '%v' sign", SliceOfToken[index].GetRawValue(source))
	err, indexErr := innerFunc(SliceOfToken, index, err)
	return err, indexErr
}
func RuleMinus(SliceOfToken []*model.Token, index int, source string) (error, []int) {
	err := fmt.Errorf("code 4: code Error : misplaced '%v' sign", SliceOfToken[index].GetRawValue(source))
	err, indexErr := innerFunc(SliceOfToken, index, err)
	return err, indexErr
}
func RuleTimes(SliceOfToken []*model.Token, index int, source string) (error, []int) {
	err := fmt.Errorf("code 5: code Error : misplaced '%v' sign", SliceOfToken[index].GetRawValue(source))
	err, indexErr := innerFunc(SliceOfToken, index, err)
	return err, indexErr
}
func RuleDevide(SliceOfToken []*model.Token, index int, source string) (error, []int) {
	err := fmt.Errorf("code 6: code Error : misplaced '%v' sign", SliceOfToken[index].GetRawValue(source))
	err, indexErr := innerFunc(SliceOfToken, index, err)
	return err, indexErr
}
