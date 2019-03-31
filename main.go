package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/garyhollandxyz/cursed-bands/twitter"
	"github.com/garyhollandxyz/cursed-bands/words"
)

func Handler(request events.APIGatewayProxyRequest) {
	fmt.Println("Starting ...")

	client, err := twitter.New()
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	randomName := newName()
	_, _, err = client.Statuses.Update(randomName, nil)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Successfully tweeted %s", randomName)

}

func newName() (name string) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	n := r.Intn(40)

	switch {
	case n <= 10:
		name = newAdjectiveNounName()
	case n <= 15:
		name = newPrepositionAdjectiveNounName()
	case n <= 22:
		name = newPossessiveAdjectiveNounName()
	case n <= 29:
		name = newPossessiveNounName()
	case n <= 33:
		name = newPrepositionNounName()
	case n <= 39:
		name = newVerbTheNounName()
	default:
		name = newVerbName()
	}
	fmt.Println("Generated name is: %s", name)
	return
}

func main() {
	lambda.Start(Handler)
}

func newNounName() (name string) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	shouldStartWithThe := r.Intn(4) <= 3

	if shouldStartWithThe {
		name = fmt.Sprintf("The %s", words.NewNoun())
	} else {
		name = words.NewNoun()
	}

	return
}

func newAdjectiveNounName() (name string) {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	shouldStartWithThe := r.Intn(4) <= 3

	if shouldStartWithThe {
		name = fmt.Sprintf("The %s %s", words.NewAdjective(), words.NewNoun())
	} else {
		name = fmt.Sprintf("%s %s", words.NewAdjective(), words.NewNoun())
	}

	return
}

func newPrepositionNounName() string {

	return fmt.Sprintf("%s %s", words.NewPreposition(), words.NewNoun())
}

func newPrepositionAdjectiveNounName() string {
	return fmt.Sprintf("%s The %s %s", words.NewPreposition(), words.NewAdjective(), words.NewNoun())

}

func newPossessiveNounName() string {
	return fmt.Sprintf("%s %s", words.NewPossessive(), words.NewNoun())
}

func newPossessiveAdjectiveNounName() string {
	return fmt.Sprintf("%s %s %s", words.NewPossessive(), words.NewAdjective(), words.NewNoun())
}

func newVerbTheNounName() string {
	return fmt.Sprintf("%s The %s", words.NewVerbPresent(), words.NewNoun())
}

func newVerbNounName() string {
	return fmt.Sprintf("The %s %s", words.NewVerbPresent(), words.NewNoun())
}

func newVerbName() string {
	return fmt.Sprintf("%s!", words.NewVerbPresent())
}
