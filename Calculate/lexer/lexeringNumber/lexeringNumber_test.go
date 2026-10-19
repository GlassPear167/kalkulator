package lexeringNumber

import (
	"fmt"
	"kalkulator/Calculate/count"
	"testing"
)

func TestLexeringNumber(t *testing.T) {

	// utils.Count("11,")
	count.Count("-11,1")
	input := count.BufferUtama
	// fmt.Println(input)
	newtoken, ofset, err := LexeringNumber(0, input)
	if err != nil {
		t.Error(err, ", err offset:", ofset)
	}

	fmt.Println(*newtoken, "", "offset:", ofset)
}

func TestLooptest(t *testing.T) {

}
