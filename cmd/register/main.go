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
	"github.com/andreaskoch/go.hue"

	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("register")

func main() {
	locators, _ := hue.DiscoverBridges(false)
	locator := locators[0] // find the first locator
	deviceType := "some app"

	// remember to push the button on your hue first
	bridge, _ := locator.CreateUser(deviceType)
	log.Infof("registered new device => %+v", bridge)
}
