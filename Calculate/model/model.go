package model

type Token struct {
	Types         int
	rawIndexStart int
	rawIndexEnd   int
	debugName     string
}

// membuat token baru

func Newtoken(typeToken int) *Token {

	return &Token{
		Types:         typeToken,
		rawIndexStart: -1,
		rawIndexEnd:   -1,
	}
}

// untuk membuat token menjadi tidak valid
func (a *Token) ToNotValid() *Token {

	a.rawIndexStart = -1
	a.rawIndexEnd = -1
	return a
}

// menambahkan index awal dan akhir dari token
func (a *Token) AddIndex(start, end int) *Token {

	a.rawIndexStart = start
	a.rawIndexEnd = end
	return a
}

// harus berupa Operator,karena nilainya tidak dipakai dalam hitungan.
func (a *Token) ToBlank(BlankValue string) *Token {
	a.rawIndexStart = -2
	a.rawIndexEnd = -2

	a.debugName = BlankValue
	return a
}

// mendapatkan slice dari raw index
// return slice kosong jika tiidak valid : []int{}
// return []int{-2,-2}, untuk karakter tambahan atau baru
func (a *Token) GetSliceIndex() []int {

	if a.rawIndexStart == -2 && a.rawIndexEnd == -2 {
		return []int{a.rawIndexStart, a.rawIndexEnd}
	} else if a.rawIndexStart < 0 && a.rawIndexEnd < 0 || a.rawIndexStart > a.rawIndexEnd {
		return []int{}
	}

	buffer := []int{}
	for i := a.rawIndexStart; i <= a.rawIndexEnd; i++ {

		buffer = append(buffer, i)
	}
	return buffer
}

// mendapatkan string dari data asli,
// return string kosong jika tidak valid : ""
func (a *Token) GetRawValue(source string) string {

	if a.rawIndexStart == -2 && a.rawIndexEnd == -2 {

		return a.debugName
	}
	if a.rawIndexStart < 0 && a.rawIndexEnd < 0 || a.rawIndexStart > a.rawIndexEnd {
		return ""
	} else if a.Types == EOF {
		return "EOF"
	}
	return source[a.rawIndexStart : a.rawIndexEnd+1]

}

const (
	Number = iota
	Pluss
	Minus
	Times
	Devide
	LeftParenttesis
	RightParentesis
	UnaryMinus
	EOF
	START
)

var Typefunc = []int{Pluss, Minus, Times, Devide, UnaryMinus}
var Numbers = []rune{46, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57} // 46 = koma (titik)

var Oprt = []rune{42, 47, 43, 45, 40, 41}
var SpaceAndTab = []rune{32, 9}

func IncludeChar() []rune {

	var IncludeChar = []rune{}
	IncludeChar = append(IncludeChar, Numbers...)
	IncludeChar = append(IncludeChar, Oprt...)
	IncludeChar = append(IncludeChar, SpaceAndTab...)
	return IncludeChar

}
