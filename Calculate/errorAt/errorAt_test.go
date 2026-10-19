package errorAt

import (
	"fmt"
	"slices"
	"testing"
)

func TestA(t *testing.T) {

	buffer := [3][]int{

		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	bufferr := slices.Concat(buffer[0], buffer[1], buffer[2])
	fmt.Println(bufferr)
	// kalimat:= fmt.Sprintf("%v%v%v",buffer[0]...,)
	// fmt.Println(`
	// 	error at :

	// 				%v%v%v

	// 	`, buffer[0], buffer[1], buffer[2])

}

func TestKalimat(t *testing.T) {

	source := "0123456789"
	sliceIndex := []int{6, 7}
	prefix, main, posfix := partKalimat(sliceIndex, source)
	fmt.Println(prefix, main, posfix)

	fmt.Println([]byte(" "))
}
func TestPenanda(t *testing.T) {

	source := "0123456789"
	sliceIndex := []int{6, 7, 8, 9}
	prefix, main, posfix := partKalimat(sliceIndex, source)

	penanda := penanda(len(prefix), len(main), len(posfix), "^")
	kalimat := string(slices.Concat(prefix, main, posfix))

	fmt.Printf(`
		%v
		%v
		`, kalimat, penanda)
}
func TestErrorAt(t *testing.T) {

	source := "012345678+9"
	sliceIndex := []int{6, 7, 8, 9}

	err := ErrorAt(sliceIndex, source, fmt.Errorf("mantap"))

	t.Error(err)

}
