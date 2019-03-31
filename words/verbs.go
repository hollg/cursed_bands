package words

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
)

type verb struct {
	Present string `json:"present"`
	Past    string `json:"past"`
}

// NewVerbPresent returns a random present tense verb
func NewVerbPresent() string {
	verbs := getVerbs()

	return strings.Title(randomVerbOf(verbs).Present)
}

// NewVerbPast returns a new past tense verb
func NewVerbPast() string {
	verbs := getVerbs()

	return strings.Title(randomVerbOf(verbs).Past)
}

func getVerbs() (verbs []verb) {
	box := packr.NewBox("./data")
	data, err := box.Find("verbs.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &verbs)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
