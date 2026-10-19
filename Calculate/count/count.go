package count

var BufferUtama []rune

func Count(mathSentence string) {

	for _, v := range mathSentence {
		BufferUtama = append(BufferUtama, v)
	}

}
