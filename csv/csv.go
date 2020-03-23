package csv

import (
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// ProcessLines processes each line with values separated by spaces
// and turns it into a comma separated line (csv like).
func ProcessLines(line string, isHeader bool) (string, error) {
	line = strings.TrimSpace(line)

	r, err := regexp.Compile("\\s+")

	if err != nil {
		return "", err
	}

	separatedByOneSpace := r.ReplaceAllString(line, " ")
	elements := strings.Split(separatedByOneSpace, " ")

	if len(elements[0]) > 0 {
		switch len(elements) {
		case 1:
			return elements[0] + ",,", nil
		case 2:
			switch len(elements[1]) {
			case 14:
				return elements[0] + "," + elements[1] + ",", nil
			case 9:
				return elements[0] + ",," + elements[1], nil
			}
			return "", fmt.Errorf("Element with %d characters can't be identified as NIT nor DUI", len(elements[1]))
		case 3:
			if !isHeader && elements[1] != "NULL" && len(elements[1]) != 14 {
				return "", fmt.Errorf("Element with %d characters can't be identified as NIT", len(elements[1]))
			} else if !isHeader && elements[2] != "NULL" && len(elements[2]) != 9 {
				return "", fmt.Errorf("Element with %d characters can't be identified as DUI", len(elements[2]))
			}
			return strings.Join(elements, ","), nil
		default:
			return "", errors.New("Line with more than 3 elements")
		}
	}

	return "", errors.New("Line with no data. Correlative, NIT and DUI are missing")
}

// LineCounter returns the number of lines read by a reader
func LineCounter(r io.Reader) (int, error) {
	return 0, nil
}
