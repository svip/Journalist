package journalist

import (
	"encoding/json"
	"os"
	"log"
)

type phraseCollection struct {
	NormalMessage []string
	ToMessage []string
	Joins []string
	Leaves []string
}

var phrases phraseCollection

func collectPhrases() {
	// Hardcode Danish for now.
	filename := "../content/phrases-da.json"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(file)
	if err = dec.Decode(&phrases); err != nil {
		log.Fatal(err)
	}
}
