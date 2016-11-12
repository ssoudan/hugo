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
package main

import (
	"flag"
	"os"
	"time"

	"github.com/andreaskoch/go.hue"

	"github.com/ssoudan/hugo/home"
	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/logging"
	"github.com/ssoudan/hugo/scene"
)

var log = logging.Log("hugo")

var (
	homeFileName  = flag.String("h", "home.json", "Home description json file")
	sceneFileName = flag.String("s", "scene.json", "Scene description json file")

	partyMode = flag.Bool("p", false, "Party mode")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	log.Info("Using home description from %s", *homeFileName)
	log.Info("Using scene description from %s", *sceneFileName)

	desc, err := types.ReadFromFile(*homeFileName)
	check(err)

	bridge := hue.NewBridge(desc.Bridge.IP, desc.Bridge.APIKey)
	// bridge.Debug()
	lights, err := bridge.GetAllLights()
	check(err)

	home := home.New(*desc, lights)

	log.Debug("%v", home)

	sceneFile, err := os.Open(*sceneFileName)
	check(err)

	s, err := scene.Read(sceneFile)
	check(err)

	if *partyMode {
		log.Info("Party mode!")
		r := s
		for {
			home.SetScene(r)
			r = r.Rotate()
			time.Sleep(2 * time.Second)
		}
	} else {
		home.SetScene(s)
	}

}
