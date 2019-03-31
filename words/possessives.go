package words

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gobuffalo/packr"
)

// NewPossesive returns a random possesive
func NewPossessive() string {
	possesives := getPossessives()
	return strings.Title(randomStringOf(possesives))
}

func getPossessives() (possessives []string) {
	box := packr.NewBox("./data")
	data, err := box.Find("possessives.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &possessives)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
