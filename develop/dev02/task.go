package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str := "a4bc2d5e"
	fmt.Println(UnpackString(str))
}

var (
	ErrNotAllowedSequence = errors.New("Sequence is not allowed for unpack")
)

func UnpackString(str string) (string, error) {
	result := &strings.Builder{}
	runes := []rune(str)

	if len(str) == 0 {
		return "", nil
	}
	if unicode.IsDigit(runes[0]) {
		return "", ErrNotAllowedSequence
	}

	var insertLetter rune
	var insertCount int
	numBuilder := &strings.Builder{}

	for i := 0; i < len(runes); i++ {
		// Блок с обработкой escape последовательности.
		if runes[i] == '\\' {
			i++
			if i < len(runes) {
				insertLetter = runes[i]
			}

			// По сути код повторяется с тем, что ниже ->
			// TODO: Вынести эти операции в отдельную функцию. (Если успею конечно)
			// А в целом, все работает как надо.
			numBuilder.Reset()
			for i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
				numBuilder.WriteRune(runes[i+1])
				i++
			}

			if numBuilder.Len() > 0 {
				num, err := strconv.ParseInt(numBuilder.String(), 10, 32)
				if err != nil {
					return "", err
				}
				insertCount = int(num)
			}

			if insertCount > 0 {
				result.WriteString(strings.Repeat(string(insertLetter), insertCount))
			}
			continue
		}

		if unicode.IsLetter(runes[i]) {
			insertLetter = runes[i]
			insertCount = 1
		}

		numBuilder.Reset()
		for i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			numBuilder.WriteRune(runes[i+1])
			i++
		}

		if numBuilder.Len() > 0 {
			num, err := strconv.ParseInt(numBuilder.String(), 10, 32)
			if err != nil {
				return "", err
			}
			insertCount = int(num)
		}

		if insertCount > 0 {
			result.WriteString(strings.Repeat(string(insertLetter), insertCount))
		}
	}
	return result.String(), nil
}
