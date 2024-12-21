package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type CutConfig struct {
	SelectField     string
	Delimiter       string
	SelectSeparated bool
}

var config = &CutConfig{}

func initCutConfig() {
	flag.StringVar(&config.SelectField, "f", "", "Fields for print")
	flag.StringVar(&config.Delimiter, "d", " ", "Delimiter for print")
	flag.BoolVar(&config.SelectSeparated, "s", false, "Select separated fields")

	flag.Usage = func() {
		fmt.Println("cut OPTION... [AND input data in STDIN]")
		flag.PrintDefaults()
	}
	flag.Parse()

	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}

	if config.SelectField == "" {
		flag.Usage()
		fmt.Println("Be sure to specify -f")
		os.Exit(1)
	}
}

func init() {
	initCutConfig()
}

func main() {
	selectedFields := MustParseFieldsIndices(config.SelectField) // Написано indices, значит переводит в индексы.
	for _, val := range selectedFields {
		if val < 0 {
			fmt.Println("Use positive numbers for fields")
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if config.SelectSeparated && !strings.Contains(line, config.Delimiter) {
			continue
		}
		allFields := strings.Split(line, config.Delimiter)

		out := BuildWithSelected(allFields, selectedFields)
		if out != "" {
			fmt.Println(out)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// ------------------------------------------------------>>>>

func MustParseFieldsIndices(fields string) []int {
	var indices []int
	for _, field := range strings.Split(fields, ",") {
		index, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		indices = append(indices, index-1)
	}
	return indices
}

func BuildWithSelected(fields []string, indices []int) string {
	builder := strings.Builder{}
	m := make(map[string]struct{})
	for _, index := range indices {
		if index >= len(fields) {
			continue
		}
		if _, ok := m[fields[index]]; ok {
			continue
		}
		m[fields[index]] = struct{}{}
		builder.WriteString(fmt.Sprintf("%s%s", fields[index], config.Delimiter))
	}
	return builder.String()
}

func Clamp(min, max, value int) int {
	if value < min {
		value = min
	}
	if value > max {
		value = max
	}
	return value
}
