package words

import (
	"math/rand"
	"time"
)

func randomStringOf(s []string) string {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source) // initialize local pseudorandom generator

	return s[r.Intn(len(s))]
}

func randomVerbOf(s []verb) verb {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source) // initialize local pseudorandom generator

	return s[r.Intn(len(s))]
}

func randomNounOf(s []noun) noun {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source) // initialize local pseudorandom generator

	return s[r.Intn(len(s))]
}

func randomIntOf(s []int) int {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source) // initialize local pseudorandom generator

	return s[r.Intn(len(s))]
}
