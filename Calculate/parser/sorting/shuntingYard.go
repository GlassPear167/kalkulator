package sorting

import (
	"kalkulator/Calculate/model"
	"slices"
)

const (
	LeftToRight = iota
	RightToLeft
)

// Struct untuk menyimpan properti setiap operator
type Operator struct {
	Precedence    int
	Associativity int
}

// Kamus (Map) untuk mengatur aturan Preseden dan Asosiativitas
var operators = map[int]Operator{
	model.Pluss:      {Precedence: 2, Associativity: LeftToRight},
	model.Minus:      {Precedence: 2, Associativity: LeftToRight},
	model.Times:      {Precedence: 3, Associativity: LeftToRight},
	model.Devide:     {Precedence: 3, Associativity: LeftToRight},
	model.UnaryMinus: {Precedence: 4, Associativity: RightToLeft}, // Pangkat bersifat Kanan-ke-Kiri
}

func ShuntingYardShorting(sliceOftoken []*model.Token) []*model.Token {

	result := []*model.Token{}
	stack := []*model.Token{}
	for _, v := range sliceOftoken {
		if v.Types == model.Number {
			result = append(result, v)
		} else if slices.Contains(model.Typefunc, v.Types) {

			o1 := v
			for len(stack) > 0 {
				o2 := stack[len(stack)-1]
				opr0 := operators[v.Types]
				opr2 := operators[stack[len(stack)-1].Types]

				// jika opr2 memiliki presedence yang lebih tinggi atau jika presedence nya sama tapi opr0 memiliki asosiasi leftToright maka pop dari stak dan pindahkan ke result
				//
				if (opr2.Precedence > opr0.Precedence) || (opr2.Precedence == opr0.Precedence && opr0.Associativity == LeftToRight) {
					result = append(result, o2)
					stack = stack[:len(stack)-1] // pop
				} else {
					break
				}
			}

			stack = append(stack, o1)

		} else if v.Types == model.LeftParenttesis {
			stack = append(stack, v)
		} else if v.Types == model.RightParentesis {

			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if top.Types == model.LeftParenttesis {

					break
				}
				result = append(result, top)
			}
		}
	}

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		result = append(result, top)
		stack = stack[:len(stack)-1]

	}

	return result
}
