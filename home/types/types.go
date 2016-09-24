package types

import (
	"encoding/json"
	"io/ioutil"
)

// PlaceDescription is the description of Place
type PlaceDescription struct {
	Lights []string `json:"lights"`
}

// HomeDescription is the description of a Home (made of Place(s))
type HomeDescription struct {
	Places map[string]PlaceDescription `json:"places"`
}

// ReadFromFile parses a home file into an Home object
func ReadFromFile(filename string) (*HomeDescription, error) {

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var h HomeDescription
	err = json.Unmarshal(b, &h)
	if err != nil {
		return nil, err
	}

	return &h, nil
}
