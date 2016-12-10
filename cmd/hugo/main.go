//
// Copyright 2015 Sebastien Soudan
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
package main

import (
	"flag"

	"github.com/andreaskoch/go.hue"
	"github.com/brutella/hc/accessory"

	"github.com/ssoudan/hugo/home"
	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/homekit"
	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("hugo")

var (
	homeFileName = flag.String("h", "home.json", "Home description json file")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	log.Info("Using home description from %s", *homeFileName)
	desc, err := types.ReadFromFile(*homeFileName)
	check(err)

	bridge := hue.NewBridge(desc.Bridge.IP, desc.Bridge.APIKey)

	lights, _ := bridge.GetAllLights()
	home := home.New(*desc, lights)

	accessories := []*accessory.Accessory{}

	for name := range home.Places {
		a := homekit.CreateLightBuld(name+" room",
			func(on bool) {
				if on == true {
					log.Debug("Client changed switch to on")
					home.LightPlaceOn(name)
				} else {
					log.Debug("Client changed switch to off")
					home.LightPlaceOff(name)
				}
			},
			func(brightness int) {
				log.Debug("Client changed lightbulb brightness %d", brightness)
				home.SetPlaceBrightness(name, brightness)
			},
			func(saturation float64) {
				log.Debug("Client changed lightbulb saturation %f", saturation)
				home.SetPlaceSaturation(name, saturation)
			},
			func(hue float64) {
				log.Debug("Client changed lightbulb hue %f", hue)
				home.SetPlaceHue(name, hue)
			})

		accessories = append(accessories, a)
	}

	master := homekit.CreateSwitch("Master", func(on bool) {
		for name := range home.Places {
			if on == true {
				log.Debug("Client changed switch to on")
				home.LightPlaceOn(name)
			} else {
				log.Debug("Client changed switch to off")
				home.LightPlaceOff(name)
			}
		}
	})

	homekit.Start(master, accessories...)

}
