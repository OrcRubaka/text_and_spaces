package main

import (
	"fmt"
	"math/rand"
)

func texeInRune(text string) (res []rune) {
	runes := []rune(text)

	var result []rune

	for i := 0; i < len(runes); i++ {
		if runes[i] >= 31 && runes[i] >= 33 {

			result = append(result, 1)
		}
		if runes[i] == 32 {
			result = append(result, 2)

		}

	}
	return result
}

func FuncLenRune(rune []rune) (lenInt int32) {
	var intLenX int32
	for x := range rune {
		if rune[x] == 2 {
			intLenX += 1

		}
	}

	return intLenX
}

func RandSpace(randInt32 int32, text string) (twospace string) {
	var raneText = []rune(text)

	var stringtwospace []rune
	var intSpase int32
	for i := 0; i < len(raneText); i++ {

		if raneText[i] == 32 {

			intSpase += 1
			stringtwospace = append(stringtwospace, raneText[i])

			if intSpase == randInt32 {
				stringtwospace = append(stringtwospace, raneText[i])
			}

		} else {
			stringtwospace = append(stringtwospace, raneText[i])
		}
	}
	twospace = string(stringtwospace)
	return twospace

}

func main() {

	for i := 0; i < 40; i++ {
		text := "Текст с одним пробелам"
		lenRune := texeInRune(text)
		int32LenRune := FuncLenRune(lenRune)

		randInt32 := rand.Int31n(int32LenRune) + 1

		final := RandSpace(randInt32, text)
		fmt.Println(final)
	}

}
