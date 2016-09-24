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

	"github.com/ssoudan/hugo/home"
	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/homekit"
	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("hugo")

var (
	redPtr    = flag.Bool("red", false, "red color")
	greenPtr  = flag.Bool("green", false, "green color")
	bluePtr   = flag.Bool("blue", false, "blue color")
	yellowPtr = flag.Bool("yellow", false, "yellow color")
	pinkPtr   = flag.Bool("pink", false, "pink color")
	offPtr    = flag.Bool("off", false, "off")
	whitePtr  = flag.Bool("white", false, "white")
)

func main() {
	flag.Parse()

	const bridgeIP = "192.168.1.100"
	const apiKey = "rRgZYFkvVS0hOKAAHscxOM5gx3RPMBOGN3VTBloV"

	bridge := hue.NewBridge(bridgeIP, apiKey)
	// bridge.Debug()
	lights, _ := bridge.GetAllLights()

	// if !*redPtr && !*greenPtr && !*bluePtr && !*yellowPtr && !*pinkPtr && !*offPtr && !*whitePtr {
	// 	fmt.Println("You need to specify one flag!")
	// 	flag.Usage()
	// 	return
	// }

	// count := 0
	// for _, ptr := range []*bool{redPtr, greenPtr, bluePtr, yellowPtr, pinkPtr, offPtr, whitePtr} {
	// 	if *ptr {
	// 		count = count + 1
	// 	}
	// }
	// if count > 1 {
	// 	fmt.Println("You can't have more than one flag!")
	// 	flag.Usage()
	// 	return
	// }

	desc, err := types.ReadFromFile("home.json")

	if err != nil {
		log.Fatal(err)
	}

	home := home.New(*desc, lights)

	log.Debug("%v", home)

	home.LightPlaceOn("salon")

	a := homekit.CreateSwitch("Lamp", func(on bool) {
		if on == true {
			log.Debug("Client changed switch to on")
			home.LightPlaceOn("salon")
		} else {
			log.Debug("Client changed switch to off")
			home.LightPlaceOff("salon")
		}
	})

	b := homekit.CreateLightBuld("Salon",
		func(on bool) {
			if on == true {
				log.Debug("Client changed switch to on")
				home.LightPlaceOn("salon")
			} else {
				log.Debug("Client changed switch to off")
				home.LightPlaceOff("salon")
			}
		},
		func(brightness int) {
			log.Debug("Client changed lightbulb brightness %d", brightness)
			home.SetPlaceBrightness("salon", brightness)
		},
		func(saturation float64) {
			log.Debug("Client changed lightbulb saturation %f", saturation)
			home.SetPlaceSaturation("salon", saturation)
		},
		func(hue float64) {
			log.Debug("Client changed lightbulb hue %f", hue)
			home.SetPlaceHue("salon", hue)
		})

	homekit.Start(a, b)

	//
	// for _, light := range lights {
	// 	fmt.Printf("%+v\n", light.Name)
	//
	// 	// attributes, err := light.GetLightAttributes()
	// 	// if err != nil {
	// 	// 	fmt.Fprintf(os.Stderr, "%s", err)
	// 	// 	continue
	// 	// }
	// 	// fmt.Printf("%#v\n", attributes.Name)
	//
	// 	// light.ColorLoop()
	// 	// if *offPtr == true {
	// 	// 	light.Off()
	// 	// } else if *redPtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		Hue:    "0", // Red
	// 	// 		Sat:    "255",
	// 	// 		Bri:    "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// } else if *yellowPtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		Hue:    "12750", // Yellow
	// 	// 		Sat:    "255",
	// 	// 		Bri:    "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// } else if *greenPtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		// Hue:    "36210", // Green
	// 	// 		Hue: "25500",
	// 	// 		Sat: "255",
	// 	// 		Bri: "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// } else if *bluePtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		Hue:    "46920", // Blue
	// 	// 		Sat:    "255",
	// 	// 		Bri:    "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// } else if *pinkPtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		Hue:    "56100", // Pink
	// 	// 		Sat:    "255",
	// 	// 		Bri:    "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// } else if *whitePtr == true {
	// 	// 	state := hue.SetLightState{
	// 	// 		On:     "true",
	// 	// 		Effect: "none",
	// 	// 		Ct:     "153", // White 6500k
	// 	// 		Sat:    "255",
	// 	// 		Bri:    "255",
	// 	// 	}
	// 	// 	light.SetState(state)
	// 	// }
	// }

}
