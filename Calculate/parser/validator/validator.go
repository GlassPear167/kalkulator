package validator

import (
	"kalkulator/Calculate/model"
	"kalkulator/Calculate/parser/rules"
)

type listValidator map[int]func(SliceOfToken []*model.Token, index int, source string) (error, []int)

func (a listValidator) doRuleChecker(SliceOfToken []*model.Token, source string) (error, []int) {

	for i, v := range SliceOfToken {
		validation, exist := a[v.Types]

		if exist {

			err, indexErrors := validation(SliceOfToken, i, source)
			if err != nil {
				return err, indexErrors
			}
		}

	}

	return nil, []int{}
}

func Check(SliceOfToken []*model.Token, source string) (error, []int) {

	list := make(listValidator)
	list[model.Pluss] = rules.RulePlus
	list[model.Minus] = rules.RuleMinus
	list[model.Times] = rules.RuleTimes
	list[model.Devide] = rules.RuleDevide
	list[model.LeftParenttesis] = rules.RuleLeftParenThesis
	list[model.RightParentesis] = rules.RuleRightParenThesis
	list[model.UnaryMinus] = rules.RuleUnaryMinus

	err, indexes := parenthesisBalanceChecker(SliceOfToken)
	if err != nil {
		return err, indexes
	}

	err, indexes = list.doRuleChecker(SliceOfToken, source)
	if err != nil {
		return err, indexes
	}
	return nil, []int{}
}
