//
// Copyright 2016 Sebastien Soudan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package types

import (
	"encoding/json"
	"io/ioutil"
)

// PlaceDescription is the description of Place
type PlaceDescription struct {
	Lights []string `json:"lights"`
}

// BridgeDescription is the description of a bridge
type BridgeDescription struct {
	IP     string `json:"ip"`
	APIKey string `json:"api-key"`
}

// HomeDescription is the description of a Home (made of Place(s))
type HomeDescription struct {
	Places           map[string]PlaceDescription `json:"places"`
	Bridge           BridgeDescription           `json:"bridge"`
	LightingSequence []string                    `json:"lighting-sequence"`
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
