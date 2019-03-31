package words

import (
	"fmt"
	"testing"
)

func TestGetVerbs(t *testing.T) {
	verbs := getVerbs()
	want := "boasted"
	matched := false

	for _, got := range verbs {
		if got.Past == want {
			matched = true
		}
	}

	if !matched {
		t.Error(fmt.Sprintf("Returned verbs didn't include '%s'", want))
	}
}
