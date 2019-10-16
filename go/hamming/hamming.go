package hamming

import "errors"

// Distance refers to the number of points at which two lines of string differ
func Distance(a, b string) (int, error) {
	countMistake := 0
	ar, br := []rune(a), []rune(b)
	if len(ar) != len(br) {
		return countMistake, errors.New("error: the two secuences are differents sizes")
	}
	for i := range ar {
		if ar[i] != br[i] {
			countMistake++
		}
	}
	return countMistake, nil
}
