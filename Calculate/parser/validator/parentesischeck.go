package validator

import (
	"fmt"
	"kalkulator/Calculate/model"
)

func parenthesisBalanceChecker(SliceOfToken []*model.Token) (error, []int) {

	buffer := [][]int{} // len buffer = counter

	for _, v := range SliceOfToken {

		if v.Types == model.RightParentesis && len(buffer) == 0 { // memeriksa awal pengulangan
			return fmt.Errorf("ada ) yang tidak memiliki pasangan"), v.GetSliceIndex()
		}
		if v.Types == model.LeftParenttesis {
			buffer = append(buffer, v.GetSliceIndex())
			continue
		}

		if v.Types == model.RightParentesis && len(buffer) != 0 {
			buffer = buffer[:len(buffer)-1]
			continue
		}

		if v.Types == model.EOF && len(buffer) != 0 {
			return fmt.Errorf("ada (  yang tidak memiliki pasangan"), buffer[len(buffer)-1]
		}
	}

	return nil, []int{}
}
