package errorAt

import (
	"fmt"
	"slices"
)

func ErrorAt(rawIndex []int, source string, message error) error {

	prefix, main, posfix := partKalimat(rawIndex, source)

	penanda := penanda(len(prefix), len(main), len(posfix), "^")
	kalimat := string(slices.Concat(prefix, main, posfix))

	err := fmt.Errorf(`
	%v
	%v

  error : %v
		`, kalimat, penanda, message)
	return err
}

// return prefix,main,posfix
func partKalimat(rawIndex []int, source string) ([]byte, []byte, []byte) {

	buffer := [3][]byte{}

	for _, v := range rawIndex {

		buffer[1] = append(buffer[1], source[v])

	}

	limit1 := 10

	for i := rawIndex[0] - 1; i >= 0; i-- {

		if limit1 < 0 {
			break
		}
		buffer[0] = append([]byte{source[i]}, buffer[0]...)
		limit1--
	}

	limit2 := 8
	count := 0
	for i := rawIndex[len(rawIndex)-1] + 1; i < len(source); i++ {
		if count == limit2 {
			break
		}

		buffer[2] = append(buffer[2], source[i])
		count++
	}

	return buffer[0], buffer[1], buffer[2]
}

func penanda(lenPrefix, lenMain, lenPosfix int, penanda string) string {
	result := []byte{}

	for range lenPrefix {
		result = append(result, 32)
	}
	penandaInbyte := []byte(penanda)
	for range lenMain {
		result = append(result, penandaInbyte...)
	}
	for range lenPosfix {
		result = append(result, 32)
	}

	return string(result)
}
