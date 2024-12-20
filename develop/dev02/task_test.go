package main

import "testing"

func TestUnpackString(t *testing.T) {
	tt := []struct {
		name    string
		input   string
		expect  string
		wantErr bool
	}{
		{
			name:    "Base sequence",
			input:   "a4bc2d5e",
			expect:  "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "Without numbers",
			input:   "abcd",
			expect:  "abcd",
			wantErr: false,
		},
		{
			name:    "Without letters",
			input:   "45",
			expect:  "",
			wantErr: true,
		},
		{
			name:    "Empty input string",
			input:   "",
			expect:  "",
			wantErr: false,
		},
		{
			name:    "With one letter + number",
			input:   "a10",
			expect:  "aaaaaaaaaa",
			wantErr: false,
		},
		{
			name:    "With zero",
			input:   "a0",
			expect:  "",
			wantErr: false,
		},
		{
			name:    "With simple escape",
			input:   "qwe\\4\\5",
			expect:  "qwe45",
			wantErr: false,
		},
		{
			name:    "With more symbols after escape",
			input:   "qwe\\45",
			expect:  "qwe44444",
			wantErr: false,
		},
		{
			name:    "Unicode with escape",
			input:   "кк\\л5",
			expect:  "ккллллл",
			wantErr: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := UnpackString(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("UnpackString() error = %v, wantErr = %v", err, tc.wantErr)
			}
			if got != tc.expect {
				t.Errorf("UnpackString() = %v, expect = %v", got, tc.expect)
			}
		})
	}
}
