package words

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/gobuffalo/packr"
)

type noun struct {
	Singular string `json:"singular"`
	Plural   string `json:"plural"`
}

// NewNoun returns a random noun, weighted in favour of plurals
func NewNoun() (noun string) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	shouldUsePlural := r.Intn(10) <= 5

	if shouldUsePlural {
		noun = NewNounPlural()
	} else {
		noun = NewNounSingular()
	}
	return
}

// NewNounPlural returns a random plural noun
func NewNounPlural() string {
	nouns := getNouns()

	return strings.Title(randomNounOf(nouns).Plural)
}

func NewNounSingular() string {
	nouns := getNouns()

	return strings.Title(randomNounOf(nouns).Singular)
}

func getNouns() (nouns []noun) {
	box := packr.NewBox("./data")
	data, err := box.Find("nouns.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &nouns)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
