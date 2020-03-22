package csv

import (
	"strings"
	"testing"
)

var testLines = []struct {
	actual, expected string
}{
	{
		"0   NULL                        NULL",
		"0,NULL,NULL",
	},
	{
		"0                          NULL",
		"",
	},
	{
		"1   06142512111014",
		"1,06142512111014,",
	},
	{
		"3   21213010123018              122451223",
		"3,21213010123018,122451223",
	},
	{
		"     0   NULL                        NULL",
		"0,NULL,NULL",
	},
	{
		"     1   06142512111014",
		"1,06142512111014,",
	},
	{
		"      3   21213010123018              122451223    ",
		"3,21213010123018,122451223",
	},
	{
		"      3                122451223    ",
		"3,,122451223",
	},
	{
		"      3234567   21213010123018              122451223    ",
		"3234567,21213010123018,122451223",
	},
	{
		"            ",
		"",
	},
	{
		"783   2121301012301 8              122451223    ",
		"",
	},
	{
		"783   2121301012301            ",
		"",
	},
	{
		"783   2121301012301              122451223    ",
		"",
	},
}

func TestProcessLine(t *testing.T) {
	for _, test := range testLines {
		actual, err := ProcessLines(test.actual, false)

		if err != nil {
			t.Log(err)
		}
		if actual != test.expected {
			t.Fatalf(`
			Error!

			Actual: %s
			Expected: %s
			`, actual, test.expected)
		}
	}
}

func BenchmarkProcessLine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ProcessLines("      3   21213010123018              122451223    ", false)
	}
}

var lines = strings.NewReader(
	`asd sadas asdsd
0   NULL                        NULL
1   06142512111014
3   21213010123018              122451223
	0   NULL                        NULL
	1   06142512111014
	3   21213010123018              122451223    
	3                122451223    
`)

func TestLineCounter(t *testing.T) {
	const expected = 8
	actual, err := LineCounter(lines)
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Fatalf(`
		Actual: %d
		Expected: %d
		`, actual, expected)
	}
}

func BenchmarkLineCounter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LineCounter(lines)
	}
}
