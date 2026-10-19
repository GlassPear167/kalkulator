package evaluator

import (
	"fmt"
	"kalkulator/Calculate/model"
	"strconv"
)

// evaluatiing shuntingyard sorted
func Evaluate(a []*model.Token, source string) ([]float64, error, []int) {

	stack := []float64{}

	// var result float64

	operator := make(map[int]func(stack []float64) []float64)

	operator[model.Pluss] = pluss
	operator[model.Minus] = minus
	operator[model.Times] = times
	operator[model.Devide] = devide
	operator[model.UnaryMinus] = UnaryMinus

	for _, v := range a {
		if v.Types == model.Number {

			angka, err := convert(v, source)
			if err != nil {
				return []float64{}, err, v.GetSliceIndex()
			}
			stack = append(stack, angka)

			continue
		}

		function, bolean := operator[v.Types]

		if !bolean {

			return []float64{}, fmt.Errorf("Operator Tidak di temukan: %v", v.Types), v.GetSliceIndex()

		}

		stack = function(stack)

	}

	if len(stack) != 1 {
		return []float64{0}, fmt.Errorf("internal error!"), []int{}
	}

	return stack, nil, []int{}
}

func convert(a *model.Token, source string) (float64, error) {

	b, err := strconv.ParseFloat(a.GetRawValue(source), 64)
	if err != nil {
		return 0, fmt.Errorf("konvert teks menjadi angka tidak berhasil: %v", a.GetRawValue(source))
	}

	return b, nil
}

// mendapatkan 2 angka dari paling belakang end-1 dan end-2, dan mengosongkan sebanyak 2
func pop(stack []float64) (float64, []float64) {

	A := stack[len(stack)-1]

	stack = stack[:len(stack)-1] //pop

	return A, stack

}

func pluss(stack []float64) []float64 {

	A, stack := pop(stack) //pop

	B, stack := pop(stack)

	result := B + A

	stack = append(stack, result)

	return stack
}
func minus(stack []float64) []float64 {

	A, stack := pop(stack) //pop

	B, stack := pop(stack)

	result := B - A

	stack = append(stack, result)

	return stack
}
func times(stack []float64) []float64 {

	A, stack := pop(stack) //pop

	B, stack := pop(stack)
	result := B * A

	stack = append(stack, result)

	return stack
}
func devide(stack []float64) []float64 {

	A, stack := pop(stack) //pop

	B, stack := pop(stack)

	result := B / A

	stack = append(stack, result)

	return stack
}
func UnaryMinus(stack []float64) []float64 {

	B, stack := pop(stack)

	result := B * -1

	stack = append(stack, result)

	return stack
}
