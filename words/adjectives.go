package words

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
)

// NewAdjective returns a random adjective
func NewAdjective() string {
	adjectives := getAdjectives()

	return strings.Title(randomStringOf(adjectives))
}

func getAdjectives() (adjectives []string) {
	box := packr.NewBox("./data")
	data, err := box.Find("adjectives.json")

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &adjectives)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
