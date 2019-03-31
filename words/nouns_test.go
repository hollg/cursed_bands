package words

import (
	"fmt"
	"testing"
)

func TestGetNouns(t *testing.T) {
	nouns := getNouns()
	want := "Rodeo"
	matched := false

	for _, got := range nouns {
		if got.Singular == want {
			matched = true
		}
	}

	if !matched {
		t.Error(fmt.Sprintf("Returned nouns didn't include %s", want))
	}

}
