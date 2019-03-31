package words

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
)

// NewPreposition returns a random preposition
func NewPreposition() string {
	prepositions := getPrepositions()
	return strings.Title(randomStringOf(prepositions))
}

func getPrepositions() (prepositions []string) {
	box := packr.NewBox("./data")
	data, err := box.Find("prepositions.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &prepositions)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
