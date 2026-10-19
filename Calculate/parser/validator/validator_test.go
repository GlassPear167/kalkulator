package validator

import (
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/lexer"
	"testing"
)

func TestParenthesisBalanceChecker(t *testing.T) {
	//())((()()
	count.Count("()(()())") // pass
	count.BufferUtama = count.BufferUtama[:0]
	count.Count("//())((()()") // failed

	input, err, errIndex := lexer.LexsertoToken(count.BufferUtama)
	if err != nil {
		t.Error(err)
		t.Errorf("err index: %v", errIndex)

	}
	err, errIndex = parenthesisBalanceChecker(input)
	if err != nil {
		t.Error(err)
		t.Errorf("err index: %v", errIndex)

	}

}
