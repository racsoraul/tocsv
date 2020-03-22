package csv

import (
	"io"
)

// ProcessLines processes each line with values separated by spaces
// and turns it into a comma separated line (csv like).
func ProcessLines(line string, isHeader bool) (string, error) {
	return "", nil
}

// LineCounter returns the number of lines read by a reader
func LineCounter(r io.Reader) (int, error) {
	return 0, nil
}
