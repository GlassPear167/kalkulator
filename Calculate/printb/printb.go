package printb

import (
	"fmt"
	"kalkulator/Calculate/model"
)

func PrintDetail(a []*model.Token, source string) {
	for _, v := range a {
		fmt.Println("value:", v.GetRawValue(source), "rawIndex:", v.GetSliceIndex(), "tipe token:", v.Types)
	}
	fmt.Println()

}
