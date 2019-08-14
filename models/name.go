package models

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "strings"
)

type Name struct {
  First string
  Last string
}

// GetName is a method which reaches out to the name service
// After it reaches out, it will parse a name into a Name struct
// Finally, it returns the value and an error if any
// TODO: Implement some sort of caching so we aren't hitting this
//   service everytime ;-)
func GetName() (*Name, error) {
  resp, err := http.Get("http://names.drycodes.com/1?nameOptions=boy_names&separator=space&format=json")
  if err != nil {
		return nil, err
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  var arr []interface{}
  err = json.Unmarshal([]byte(body), &arr)
  if err != nil {
		return nil, err
  }
  names := strings.Split(arr[0].(string), " ")
  return &Name{
    First: names[0],
    Last: names[1],
  }, nil
}

func (n *Name) String() string {
  return n.First + " " + n.Last
}
