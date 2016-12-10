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
	"fmt"
	"strings"
	"time"

	"github.com/andreaskoch/go.hue"
	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("follow")

var (
	homeFileName     = flag.String("h", "home.json", "Home description json file")
	masterLight      = flag.String("m", "stb - room 1", "Master light")
	slaveLightPrefix = flag.String("s", "stb - room", "Slave light name prefix")
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

	for {
		// bridge.Debug()
		lights, _ := bridge.GetAllLights()

		mainLight, err := bridge.FindLightByName(*masterLight)
		if err != nil {
			fmt.Println(err)
			return
		}
		la, err := mainLight.GetLightAttributes()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%v - %v\n", time.Now(), la.State)
		if la.State.Reachable && la.State.On {
			for _, light := range lights {
				if !strings.HasPrefix(light.Name, *slaveLightPrefix) {
					light.On()
				}
			}
		} else {
			for _, light := range lights {
				if !strings.HasPrefix(light.Name, *slaveLightPrefix) {
					light.Off()
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}
