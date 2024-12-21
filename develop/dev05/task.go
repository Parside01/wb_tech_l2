package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type GrepConfig struct {
	FilePath          string
	Pattern           string
	AfterN            int
	BeforeN           int
	ContextN          int
	Count             bool
	IgnoreCase        bool
	SelectNonMatching bool
	Fixed             bool
	PrintNumLine      bool
}

var config = &GrepConfig{}

const stdin = "stdin"

func initGrepConfig() {
	flag.IntVar(&config.AfterN, "A", 0, "Print NUM lines of trailing context")
	flag.IntVar(&config.BeforeN, "B", 0, "Print NUM lines of leading context")
	flag.IntVar(&config.ContextN, "C", 0, "print NUM lines of output context")
	flag.BoolVar(&config.Count, "c", false, "Print only a count of matching lines per FILE")
	flag.BoolVar(&config.IgnoreCase, "i", false, "Ignore case distinctions")
	flag.BoolVar(&config.SelectNonMatching, "v", false, "Select non-matching lines")
	flag.BoolVar(&config.Fixed, "F", false, "PATTERN is a set of newline-separated strings")
	flag.BoolVar(&config.PrintNumLine, "n", false, "Print line number with output lines")

	flag.Usage = func() {
		fmt.Println("Usage: grep [OPTION]... PATTERN [FILE]...")
	}

	flag.Parse()

	if !flag.Parsed() {
		flag.Usage()
		os.Exit(1)
	}
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if _, err := regexp.Compile(flag.Arg(0)); err != nil {
		fmt.Println("Invalid PATTERN")
		flag.Usage()
		os.Exit(1)
	}

	config.Pattern = flag.Arg(0)

	if flag.NArg() == 1 {
		config.FilePath = stdin
	} else {
		config.FilePath = flag.Arg(1)
	}
}

func organizeConfig() {
	if config.ContextN > config.AfterN {
		config.AfterN = config.ContextN
	}

	if config.ContextN > config.BeforeN {
		config.BeforeN = config.ContextN
	}

	if config.Fixed {
		replacer := strings.NewReplacer(
			`?`, `\?`,
			`*`, `\*`,
			`+`, `\+`,
			`{`, `\{`,
			`}`, `\}`,
			`|`, `\|`,
			`)`, `\)`,
			`(`, `\(`,
			`[`, `\[`,
			`]`, `\]`,
			`\`, `\\`,
			`^`, `\^`,
			`$`, `\$`,
		)
		config.Pattern = replacer.Replace(config.Pattern)
	}

	if config.IgnoreCase {
		config.Pattern = "(?i)" + config.Pattern
	}
}

func init() {
	initGrepConfig()
	organizeConfig()
}

func main() {
	var reader io.Reader
	if config.FilePath == stdin {
		reader = os.Stdin
	} else {
		file, err := os.Open(config.FilePath)
		if err != nil {
			fmt.Println("Error opening file:", err)
		}
		reader = file
	}

	var data []string
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
		os.Exit(1)
	}

	marks := MarkApproachLinesRegex(data)

	if config.SelectNonMatching {
		InvertMarks(marks)
	}

	indices := GetIndicesOfMarks(marks)

	if config.Count {
		fmt.Println(len(indices))
		return
	}

	// BuildResultString(data, indices)
	fmt.Println(BuildResultString(data, indices))
}

// ----------------------------------------------------------------- >>>

func BuildResultString(lines []string, indices []int) string {
	builder := strings.Builder{}     // Чтобы собрать выходную строку.
	tempBuilder := strings.Builder{} // Чтобы собирать одну строку.
	mini, maxi := 0, len(lines)
	lineSet := make(map[string]struct{}) // Чтобы убрать повторяющиеся строки.

	for _, index := range indices {
		minIndex := Clamp(mini, maxi, index-config.BeforeN)
		maxIndex := Clamp(mini, maxi, index+config.AfterN+1)

		toInsert := lines[minIndex:maxIndex]
		for i := minIndex; i < maxIndex; i++ {
			tempBuilder.Reset()

			if config.PrintNumLine {
				tempBuilder.WriteString(fmt.Sprintf("%d%s ", i+1, CompareAndGetChar(index, i)))
			}

			tempBuilder.WriteString(fmt.Sprintf("%s\n", toInsert[i-minIndex]))

			tempLine := tempBuilder.String()
			if _, ok := lineSet[tempLine]; ok {
				continue
			}

			lineSet[tempLine] = struct{}{}

			builder.WriteString(tempLine)
		}
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

func MarkApproachLinesRegex(lines []string) []bool {
	marks := make([]bool, len(lines))
	reg := regexp.MustCompile(config.Pattern)

	for i := range lines {
		if reg.MatchString(lines[i]) {
			marks[i] = true
		}
	}
	return marks
}

func InvertMarks(lines []bool) {
	for i := range lines {
		lines[i] = !lines[i]
	}
}

func GetIndicesOfMarks(marks []bool) []int {
	var indices []int

	for i, mark := range marks {
		if mark {
			indices = append(indices, i)
		}
	}
	return indices
}

func CompareAndGetChar(expect, got int) string {
	if expect == got {
		return ":"
	}
	return "-"
}
