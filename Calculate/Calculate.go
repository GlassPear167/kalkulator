package Calculate

import (
	"kalkulator/Calculate/count"
	"kalkulator/Calculate/errorAt"
	"kalkulator/Calculate/evaluator"
	"kalkulator/Calculate/lexer"
	"kalkulator/Calculate/parser"
)

func Calculate(mathSentence string) ([]float64, error) {

	count.Count(mathSentence)

	bufferUtama := count.BufferUtama

	sliceOfToken, err, indexErr := lexer.LexsertoToken(bufferUtama)

	if err != nil {
		return []float64{}, errorAt.ErrorAt(indexErr, mathSentence, err)
	}

	sliceOfToken, err, indexErr = parser.Parser(sliceOfToken, mathSentence)

	if err != nil {
		return []float64{}, errorAt.ErrorAt(indexErr, mathSentence, err)
	}
	hasil, err, indexErr := evaluator.Evaluate(sliceOfToken, mathSentence)
	if err != nil {
		return []float64{}, errorAt.ErrorAt(indexErr, mathSentence, err)
	}

	return hasil, nil
}
