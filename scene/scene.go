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
	"fmt"
	"io"
	"math"
)

// Element is a scene element
type Element struct {
	Place string `json:"place"`

	Saturation float64 `json:"saturation"`
	Brightness int     `json:"brightness"`
	Hue        float64 `json:"hue"`
}

// Scene is a collection of Element
type Scene []Element

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

// SetHSL sets values for all Elements
func (sc Scene) SetHSL(hue, saturation, brightness float32) Scene {
	r := make([]Element, len(sc))

	for i := 0; i < len(sc); i++ {
		r[i] = sc[i]
		r[i].Brightness = int(brightness * 255)
		r[i].Hue = float64(hue) * 360
		r[i].Saturation = float64(saturation) * 100
	}

	return r
}

func max(a, b float32) float32 {
	return float32(math.Max(float64(a), float64(b)))
}

func min(a, b float32) float32 {
	return float32(math.Min(float64(a), float64(b)))
}

func RGBtoHSL(r, g, b float32) (float32, float32, float32) {
	m := min(min(r, g), b)
	M := max(max(r, g), b)

	l := (m + M) / 2

	var s float32
	if l > 0.5 {
		s = (M - m) / (M + m)
	} else {
		s = (M - m) / (2.0 - M - m)
	}

	var h float32
	if s != 0 {
		if r >= g && r >= b {
			h = (g - b) / (M - m)
		} else if g >= r && g >= b {
			h = 2.0 + (b-r)/(M-m)
		} else if b >= r && b >= g {
			h = 4.0 + (r-g)/(M-m)
		}
		h = h * 60. / 360.
		if h > 1 {
			h = h - 1
		}
		if h < 0 {
			h = h + 1
		}
	}

	return h, s, l
}

// SetRGB sets values for all Elements
func (sc Scene) SetRGB(r, g, b float32) Scene {

	fmt.Printf("r=%.2f g=%.2f b=%.2f\n", r, g, b)

	h, s, l := RGBtoHSL(r, g, b)

	fmt.Printf("h=%.2f s=%.2f l=%.2f\n", h, s, l)

	return sc.SetHSL(h, s, l)
}

// Rotate scene
func (sc Scene) Rotate() Scene {

	r := make([]Element, len(sc))

	for i := 0; i < len(sc); i++ {
		r[i] = sc[i]
		r[i].Place = sc[(i+1)%len(sc)].Place
	}

	return r
}
