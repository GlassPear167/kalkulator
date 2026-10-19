package model

import (
	"fmt"
	"testing"
)

// func TestModelMethodAddIndex(t *testing.T) {

// 	newtoken := Newtoken(Number, []rune("1"))
// 	newtoken.AddIndex(1)
// 	fmt.Println(*newtoken)
// }

// func TestModelMethodReplaceAllIndexWith(t *testing.T) {

// 	newtoken := Newtoken(Number, []rune("1"))
// 	newtoken.AddIndex(1)
// 	newtoken.ReplaceAllIndexWith([]int{2})
// 	fmt.Println(*newtoken)
// }

func TestLoop(t *testing.T) {

	sliceNum := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sliceNum); i++ {
		fmt.Println(sliceNum[i])
	}
}

func TestLoopvar(t *testing.T) {

	n := "adflkajsdf"
	m := n[0:4] // instan lookup, matnap

	// Numbers := []int32{44, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	// for _, v := range Numbers {
	// 	fmt.Print(v - 48)

	// }
	fmt.Println(m)
}
func TestGetRawIndex(t *testing.T) {

	m := Newtoken(Number).GetSliceIndex() // generate index berdasarkan nilai start dan end, bukan index asli
	fmt.Println(m)                        // []int{}
	m = Newtoken(Number).AddIndex(0, 3).GetSliceIndex()
	fmt.Println(m) //[]int{0,1,2,3}
	m = Newtoken(Number).AddIndex(3, 0).GetSliceIndex()
	fmt.Println(m) //[]int{}
	m = Newtoken(Number).ToBlank("12345").GetSliceIndex()
	fmt.Println(m) //[]int{-2,-2}

}
func TestGetRawData(t *testing.T) {

	n := "1234"

	m := Newtoken(Number).GetRawValue(n)               // instan lookup, matnap
	fmt.Println(m)                                     // ""
	m = Newtoken(Number).AddIndex(0, 3).GetRawValue(n) // instan lookup, matnap
	fmt.Println(m)                                     //"1234"
	m = Newtoken(Number).AddIndex(3, 0).GetRawValue(n) // instan lookup, matnap
	fmt.Println(m)                                     //""
	m = Newtoken(Number).ToBlank("12349999").GetRawValue(n)
	fmt.Println(m) //"12349999"

}
