package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	input := []string{"пятак", "пятка", "тяпка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(GetAnagramSets(input))
}

// Так как принимаем слайс == принимаем ссылку на массив.
// Возвращать ссылку на мапу, когда тав внутри указатель на hmap... сомнение.
func GetAnagramSets(words []string) map[string][]string {
	result := make(map[string][]string)
	sortAnagramMap := make(map[string][]string)
	wordsMap := make(map[string]struct{})

	for i, word := range words {
		words[i] = strings.ToLower(word) // 1. Приводим все к нижнему регистру
	}

	for _, word := range words {
		if _, ok := wordsMap[word]; ok {
			continue // 2. Если слово уже было, то не добавляем
		}
		var runes Runes = []rune(word)
		sort.Sort(runes)
		sortedWord := string(runes)
		sortAnagramMap[sortedWord] = append(sortAnagramMap[sortedWord], word)
		wordsMap[word] = struct{}{}
	}

	for _, anagrams := range sortAnagramMap {
		if len(anagrams) == 1 {
			continue // Множества из одного элемента не идет в результат.
		}
		result[anagrams[0]] = append(result[anagrams[0]], anagrams...)
	}

	for _, sets := range result {
		sort.Strings(sets) // сортируем по возрастанию.
	}
	return result
}

type Runes []rune

func (r Runes) Len() int {
	return len(r)
}

func (r Runes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Runes) Less(i, j int) bool {
	return r[i] < r[j]
}
