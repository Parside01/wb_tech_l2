package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и glint.
*/

type SortConfig struct {
	FilePath            string
	Column              uint
	ByNumeric           bool
	ByHumanNumeric      bool
	ByMonth             bool
	Reverse             bool
	Unique              bool
	IgnoreLeadingBlanks bool
	Check               bool
}

var config = &SortConfig{}

func initConfig() {
	flag.UintVar(&config.Column, "k", 1, "Column to sort by. 1 - based index)")
	flag.BoolVar(&config.ByNumeric, "n", false, "Compare according to string numerical value")
	flag.BoolVar(&config.ByMonth, "M", false, "Compare according to string month value")
	flag.BoolVar(&config.Reverse, "r", false, "Reverse the result of comparisons")
	flag.BoolVar(&config.Unique, "u", false, "Output only the first of an equal run")
	flag.BoolVar(&config.Check, "c", false, "Check the results of comparisons")
	flag.BoolVar(&config.IgnoreLeadingBlanks, "b", false, "Ignore leading blanks")
	flag.BoolVar(&config.ByHumanNumeric, "h", false, "Compare human readable numbers")

	flag.Usage = func() {
		fmt.Println("Usage: sort [OPTION]... [FILE]")
	}

	flag.Parse()

	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}

	if len(flag.Args()) == 0 || len(flag.Args()) > 1 {
		flag.Usage()
		os.Exit(1)
	}

	config.FilePath = flag.Arg(0)
}

func validateConfig() {
	if config.Column < 1 {
		config.Column = 1
	}
}

func init() {
	initConfig()
	validateConfig()
}

func main() {
	file, err := os.Open(config.FilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]string
	m := make(map[string]struct{})

	for scanner.Scan() {
		text := scanner.Text()

		if config.Unique {
			if _, ok := m[text]; !ok {
				m[text] = struct{}{}
			} else {
				text = ""
			}
		}

		if text == "" {
			continue
		}

		if config.IgnoreLeadingBlanks {
			text = strings.TrimSpace(text)
		}

		data = append(data, strings.Fields(text))
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	switch {
	case config.ByNumeric:
		var sortIn sort.Interface = &SortNumeric{SortData{data, config.Column}}
		if config.Reverse {
			sortIn = sort.Reverse(sortIn)
		}
		CheckSortedData(config.Check, sortIn)
		sort.Sort(sortIn)
	case config.ByMonth:
		var sortIn sort.Interface = &SortMonth{SortData{data, config.Column}}
		if config.Reverse {
			sortIn = sort.Reverse(sortIn)
		}
		CheckSortedData(config.Check, sortIn)
		sort.Sort(sortIn)
	case config.ByHumanNumeric:
		var sortIn sort.Interface = &SortHumanNumeric{SortData{data, config.Column}}
		if config.Reverse {
			sortIn = sort.Reverse(sortIn)
		}
		CheckSortedData(config.Check, sortIn)
		sort.Sort(sortIn)
	default:
		var sortIn sort.Interface = &SortString{SortData{data, config.Column}}
		if config.Reverse {
			sortIn = sort.Reverse(sortIn)
		}
		CheckSortedData(config.Check, sortIn)
		sort.Sort(sortIn)
	}

	fmt.Println(BuildString(data))
}

// <<<< -------------------------------------------------------------------------- >>>>

func CheckSortedData(check bool, s sort.Interface) {
	if !check {
		return
	}
	sorted := sort.IsSorted(s)
	if sorted {
		fmt.Println("Data is sorted")
	} else {
		fmt.Println("Data is not sorted")
	}
}

func BuildString(data [][]string) string {
	builder := &strings.Builder{}

	for _, row := range data {
		builder.WriteString(fmt.Sprintf("%s\n", strings.Join(row, " ")))
	}
	return builder.String()
}

// По сути то, что мы будем сортировать и как.
type SortData struct {
	data   [][]string
	column uint
}

// Частичная реализация sort.Interface.
func (s *SortData) Len() int {
	return len(s.data)
}

func (s *SortData) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

// Для каждой сортировки есть отдельная структура,
// которая по своему реализует sort.Interface.
type SortNumeric struct {
	SortData
}

func (s *SortNumeric) Less(i, j int) bool {
	a, err := strconv.ParseFloat(s.data[i][s.column-1], 64)
	if err != nil {
		return false // Такая логика нужна, чтобы сначала шли строки, а потом числа.
	}

	b, err := strconv.ParseFloat(s.data[j][s.column-1], 64)
	if err != nil {
		return true // Такая логика нужна, чтобы сначала шли строки, а потом числа.
	}

	return a < b
}

type SortString struct {
	SortData
}

func (s *SortString) Less(i, j int) bool {
	return strings.Compare(s.data[i][s.column-1], s.data[j][s.column-1]) == -1
}

type SortMonth struct {
	SortData
}

func (s *SortMonth) Less(i, j int) bool {
	months := map[string]int{
		"jan": 1, "feb": 2, "mar": 3, "apr": 4,
		"may": 5, "jun": 6, "jul": 7, "aug": 8,
		"sep": 9, "oct": 10, "nov": 11, "dec": 12,
	}
	a := strings.ToLower(s.data[i][s.column-1])
	b := strings.ToLower(s.data[j][s.column-1])

	return months[a] < months[b]
}

type SortHumanNumeric struct {
	SortData
}

func (s *SortHumanNumeric) Less(i, j int) bool {
	a := parseHumanNumeric(s.data[i][s.column-1])
	b := parseHumanNumeric(s.data[j][s.column-1])
	fmt.Println(a, b)
	return a < b
}

// Пока такая реализация, может быть нужно еще что-то добавить, но вроде все работает.
func parseHumanNumeric(value string) float64 {
	value = strings.TrimSpace(value)
	var multiplier float64 = 1

	if strings.HasSuffix(value, "$") {
		value = strings.TrimSuffix(value, "$")
	}

	if strings.HasSuffix(value, "K") {
		multiplier = 1e3
		value = strings.TrimSuffix(value, "K")
	} else if strings.HasSuffix(value, "M") {
		multiplier = 1e6
		value = strings.TrimSuffix(value, "M")
	} else if strings.HasSuffix(value, "G") {
		multiplier = 1e9
		value = strings.TrimSuffix(value, "G")
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}

	return num * multiplier
}
