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
package home

import (
	"strconv"

	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/logging"
	"github.com/ssoudan/hugo/scene"

	"github.com/andreaskoch/go.hue"
)

var log = logging.Log("home")

// Place is a collection of lights
type Place struct {
	Lights []*hue.Light
}

// Home is a collection of Places
type Home struct {
	Places map[string]Place
}

// New creates a new Home from its description
func New(desc types.HomeDescription, lights []*hue.Light) *Home {
	home := Home{
		Places: make(map[string]Place),
	}

	lightMap := make(map[string]*hue.Light)
	for _, l := range lights {
		lightMap[l.Name] = l
	}

	for name, place := range desc.Places {
		lls := []*hue.Light{}
		for _, lname := range place.Lights {
			light, ok := lightMap[lname]
			if ok {
				lls = append(lls, light)
			}
		}

		home.Places[name] = Place{
			Lights: lls,
		}
	}

	return &home
}

// LightPlaceOn turn On all the lights of a place
func (h Home) LightPlaceOn(placeName string) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			light.On()
		}
	}
}

// LightPlaceOff turn Off all the lights of a place
func (h Home) LightPlaceOff(placeName string) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			light.Off()
		}
	}
}

// SetPlaceBrightness sets the brightness of a Place, leaving the rest unchanged
func (h Home) SetPlaceBrightness(placeName string, brightness int) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			state := hue.SetLightState{
				Bri: strconv.Itoa(brightness * 255 / 100),
			}

			_, err := light.SetState(state)
			if err != nil {
				log.Errorf("Failed to set state: %v", err)
			}
		}
	}
}

// SetPlaceSaturation sets the saturation of a Place, leaving the rest unchanged
func (h Home) SetPlaceSaturation(placeName string, saturation float64) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			state := hue.SetLightState{
				Sat: strconv.Itoa(int(saturation * 255 / 100)),
			}

			_, err := light.SetState(state)
			if err != nil {
				log.Errorf("Failed to set state: %v", err)
			}
		}
	}
}

// SetPlaceHue sets the hue of a Place, leaving the rest unchanged
func (h Home) SetPlaceHue(placeName string, hh float64) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			state := hue.SetLightState{
				Hue: strconv.Itoa(int(hh * 65280 / 360.)),
			}

			_, err := light.SetState(state)
			if err != nil {
				log.Errorf("Failed to set state: %v", err)
			}
		}
	}
}

// SetPlaceAttributes sets all 3 attributes of a light
func (h Home) SetPlaceAttributes(placeName string, hh float64, saturation float64, brightness int) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {

			state := hue.SetLightState{
				Hue:            strconv.Itoa(int(hh * 65280 / 360.)),
				Sat:            strconv.Itoa(int(saturation * 255 / 100)),
				Bri:            strconv.Itoa(brightness * 255 / 100),
				TransitionTime: "0",
			}

			_, err := light.SetState(state)
			if err != nil {
				log.Errorf("Failed to set state: %v", err)
			}
		}
	}
}

// SetScene sets a Scene
func (h Home) SetScene(scene scene.Scene) {

	for _, c := range scene {
		h.SetPlaceAttributes(c.Place, c.Hue, c.Saturation, c.Brightness)
	}
}
