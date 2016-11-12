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
package scene

import (
	"encoding/json"
	"io"
)

// SceneElement is a scene element
type SceneElement struct {
	Place string `json:"place"`

	Saturation float64 `json:"saturation"`
	Brightness int     `json:"brightness"`
	Hue        float64 `json:"hue"`
}

// Scene is a collection of SceneElement
type Scene []SceneElement

// Read a scene
func Read(r io.Reader) (Scene, error) {
	scene := Scene{}

	decoder := json.NewDecoder(r)
	err := decoder.Decode(&scene)
	if err != nil {
		return nil, err
	}

	return scene, nil
}

// Rotate scene
func (s Scene) Rotate() Scene {

	r := make([]SceneElement, len(s))

	for i := 0; i < len(s); i++ {
		r[i] = s[i]
		r[i].Place = s[(i+1)%len(s)].Place
	}

	return r
}
