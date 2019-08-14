package models

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
)

type Joke struct {
  Template string
  Value map[string]interface{} `json: value`
  name Name
}

// GetJoke is a method which reaches out to the joke service
// After it reaches out, it will populate the template with a name from GetName
// Finally, it returns the value and an error if any
// TODO: Implement some sort of caching so we aren't hitting this
//   service everytime ;-)
func GetJoke() (*Joke, error) {
  resp, err := http.Get("http://api.icndb.com/jokes/random?firstName=John&lastName=Doe&limitTo=[nerdy]")
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
  joke.Template = strings.Replace(joke.Value["joke"].(string), "John Doe", "%s", 1)

  return joke, nil
}
