package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	Value map[string]interface{} `json: value`
	name  Name
}

// GetJoke is a method which reaches out to the joke service
// After it reaches out, it will populate the template with a name from GetName
// Finally, it returns the value and an error if any
// TODO: Implement some sort of caching so we aren't hitting this
//   service everytime ;-)
func GetJoke(name *Name) (*Joke, error) {
	url := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s&limitTo=[nerdy]", name.First, name.Last)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var joke *Joke
	err = json.Unmarshal([]byte(body), &joke)
	if err != nil {
		return nil, err
	}

	return joke, nil
}
