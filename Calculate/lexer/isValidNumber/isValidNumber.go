package isValidNumber

import (
	"slices"
)

func IsValidNumber(numberInRune []rune) bool {

	ProsesCather := []int{}

	number := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	karakterPertama := []rune{}
	karakterPertama = append(karakterPertama, number...)
	// karakterPertama = append(karakterPertama, 45)

	var koma rune = 46

	// memeriksa karakter lain selain yang diperlukan (tab dan space)

	if slices.Contains(numberInRune, 9) || slices.Contains(numberInRune, 32) {

		return false
	}
	//=================
	if len(numberInRune) == 1 {

		if slices.Contains(number, numberInRune[0]) {
			ProsesCather = append(ProsesCather, 1)

		}
	}

	//=================
	if len(numberInRune) == 2 {

		// pemeriksaan karakter peratama
		if slices.Contains(karakterPertama, numberInRune[0]) {
			ProsesCather = append(ProsesCather, 1)

		}

		// Pemeriksaan karakters terakhir
		if slices.Contains(number, numberInRune[1]) {
			if slices.Contains(number, numberInRune[len(numberInRune)-1]) {

				ProsesCather = append(ProsesCather, 2)

			}
		}
		if numberInRune[0] == 48 {

		}
	}

	// ==============
	if len(numberInRune) > 2 {

		// pemeriksaan karakter peratama
		if slices.Contains(karakterPertama, numberInRune[0]) {
			ProsesCather = append(ProsesCather, 1)

		}

		// pemeriksaan 0 di awal

		// Pemeriksaan karakters terakhir
		if len(numberInRune) >= 2 {
			if slices.Contains(number, numberInRune[len(numberInRune)-1]) {

				ProsesCather = append(ProsesCather, 2)

			}

		}
		// pemeriksaan koma

		if slices.Contains(numberInRune, koma) {
			indexKoma := []int{}
			for i, v := range numberInRune {
				if v == koma {
					indexKoma = append(indexKoma, i)
				}

			}
			if len(indexKoma) == 1 {

				if int(indexKoma[0]) >= 1 || int(indexKoma[0]) <= int(len(numberInRune)-2) {
					ProsesCather = append(ProsesCather, 3)

				}

				if numberInRune[0] == 45 && !slices.Contains(number, numberInRune[1]) {
					ProsesCather = append(ProsesCather, 4)

				}
			}

			if len(indexKoma) != 1 {
				return false
			}
		}

		//pemeriksan angka 0 didepan tanpa didampingi koma setelahnya

		if numberInRune[0] == 48 && numberInRune[1] != koma {
			ProsesCather = append(ProsesCather, 5)

		}
		// //pemeriksan angka -0 didepan tanpa didampingi koma setelahnya

		// if numberInRune[0] == 45 && numberInRune[1] == 48 && numberInRune[2] != koma {
		// 	ProsesCather = append(ProsesCather, 6)
		// 	slog.Debug("tag 6")
		// }
	}

	// validating bilangan
	//=============
	if len(numberInRune) == 0 {
		return false
	}

	// satu bilangan
	if len(numberInRune) == 1 && len(ProsesCather) == 1 && ProsesCather[0] == 1 {
		return true

	}
	// dua bilangan
	if len(numberInRune) == 2 && len(ProsesCather) == 2 && ProsesCather[0] == 1 && ProsesCather[1] == 2 {
		return true
	}
	// tiga bilangan
	if len(numberInRune) > 2 {
		if len(ProsesCather) == 2 && ProsesCather[0] == 1 && ProsesCather[1] == 2 {
			return true
		}
		if len(ProsesCather) == 3 && ProsesCather[0] == 1 && ProsesCather[1] == 2 && ProsesCather[2] == 3 {

			return true
		}
	}

	return false
}
