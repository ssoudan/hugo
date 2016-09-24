package home

import (
	"fmt"
	"strconv"

	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/logging"

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

	for _, light := range lights {
		fmt.Printf("%+v\n", light.Name)
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
			attr, err := light.GetLightAttributes()
			if err != nil {
				log.Error("Failed to get light attribtues: %v", err)
				continue
			}
			state := hue.SetLightState{
				Hue:            strconv.Itoa(attr.State.Hue),
				On:             fmt.Sprintf("%v", attr.State.On),
				Effect:         attr.State.Effect,
				Alert:          attr.State.Alert,
				Bri:            strconv.Itoa(brightness),
				Sat:            strconv.Itoa(attr.State.Sat),
				Ct:             strconv.Itoa(attr.State.Ct),
				Xy:             attr.State.Xy,
				TransitionTime: "0",
			}

			_, err = light.SetState(state)
			if err != nil {
				log.Error("Failed to set state: %v", err)
			}
		}
	}
}

// SetPlaceSaturation sets the saturation of a Place, leaving the rest unchanged
func (h Home) SetPlaceSaturation(placeName string, saturation float64) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			attr, err := light.GetLightAttributes()
			if err != nil {
				log.Error("Failed to get light attribtues: %v", err)
				continue
			}
			state := hue.SetLightState{
				Hue:            strconv.Itoa(attr.State.Hue),
				On:             fmt.Sprintf("%v", attr.State.On),
				Effect:         attr.State.Effect,
				Alert:          attr.State.Alert,
				Bri:            strconv.Itoa(attr.State.Bri),
				Sat:            strconv.Itoa(int(saturation)),
				Ct:             strconv.Itoa(attr.State.Ct),
				Xy:             attr.State.Xy,
				TransitionTime: "0",
			}

			_, err = light.SetState(state)
			if err != nil {
				log.Error("Failed to set state: %v", err)
			}
		}
	}
}

// SetPlaceHue sets the hue of a Place, leaving the rest unchanged
func (h Home) SetPlaceHue(placeName string, hh float64) {
	place, ok := h.Places[placeName]
	if ok {
		for _, light := range place.Lights {
			attr, err := light.GetLightAttributes()
			if err != nil {
				log.Error("Failed to get light attribtues: %v", err)
				continue
			}
			state := hue.SetLightState{
				Hue:            strconv.Itoa(int(hh)),
				On:             fmt.Sprintf("%v", attr.State.On),
				Effect:         attr.State.Effect,
				Alert:          attr.State.Alert,
				Bri:            strconv.Itoa(attr.State.Bri),
				Sat:            strconv.Itoa(attr.State.Sat),
				Ct:             strconv.Itoa(attr.State.Ct),
				Xy:             attr.State.Xy,
				TransitionTime: "0",
			}

			_, err = light.SetState(state)
			if err != nil {
				log.Error("Failed to set state: %v", err)
			}
		}
	}
}
