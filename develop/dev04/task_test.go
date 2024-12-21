package main

import (
	"reflect"
	"testing"
)

func TestGetAnagramSets(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Тест из задания",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name:  "Повторяющиеся слова",
			input: []string{"кот", "ток", "кот", "окт"},
			expected: map[string][]string{
				"кот": {"кот", "окт", "ток"},
			},
		},
		{
			name:  "Отсеиваем множества из одного элемента",
			input: []string{"слон", "нос", "сон", "солнце"},
			expected: map[string][]string{
				"нос": {"нос", "сон"},
			},
		},
		{
			name:  "Приведение к нижнему регистру",
			input: []string{"СЛОН", "НОС", "СОН", "СОЛНЦЕ"},
			expected: map[string][]string{
				"нос": {"нос", "сон"},
			},
		},
		{
			name:     "Нет анаграмм",
			input:    []string{"а", "б", "в"},
			expected: map[string][]string{},
		},
		{
			name:  "Сортировка анаграмм",
			input: []string{"поток", "котоп", "топок", "кот", "ток"},
			expected: map[string][]string{
				"кот":   {"кот", "ток"},
				"поток": {"котоп", "поток", "топок"},
			},
		},
		{
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "пятак"},
			expected: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := GetAnagramSets(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For input %v, expected %v, but got %v", tc.input, tc.expected, result)
			}
		})
	}
}
