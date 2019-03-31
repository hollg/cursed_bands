package words

import (
	"fmt"
	"testing"
)

func TestGetAdjectives(t *testing.T) {
	nouns := getAdjectives()
	want := "woeful"
	matched := false

	for _, got := range nouns {
		if got == want {
			matched = true
		}
	}

	if !matched {
		t.Error(fmt.Sprintf("Returned adjectives didn't include %s", want))
	}

}
