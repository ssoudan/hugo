package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"

	hue "github.com/andreaskoch/go.hue"
	"github.com/ssoudan/hugo/home"
	"github.com/ssoudan/hugo/home/types"
	"github.com/ssoudan/hugo/input"
	"github.com/ssoudan/hugo/logging"
	"github.com/ssoudan/hugo/scene"
)

var (
	homeFileName  = flag.String("h", "home.json", "Home description json file")
	sceneFileName = flag.String("s", "scene.json", "Scene description json file")
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var log = logging.Log("hugotune")

func main() {

	flag.Parse()

	log.Infof("Using home description from %s", *homeFileName)
	log.Infof("Using scene description from %s", *sceneFileName)

	desc, err := types.ReadFromFile(*homeFileName)
	check(err)

	bridge := hue.NewBridge(desc.Bridge.IP, desc.Bridge.APIKey)
	// bridge.Debug()
	lights, err := bridge.GetAllLights()
	check(err)

	home := home.New(*desc, lights)

	log.Debugf("%v", home)

	sceneFile, err := os.Open(*sceneFileName)
	check(err)

	s, err := scene.Read(sceneFile)
	check(err)

	p, err := input.New()
	check(err)

	home.SetScene(s)

	fmt.Println("Starting")

	var values [4]float32
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			valueChanged := false
			for i := byte(0); i < 4; i++ {

				v, err := p.ReadAndScale(i)
				check(err)

				if v != values[i] {
					valueChanged = math.Abs(float64(v-values[i])) > 0.001
					values[i] = v
				}
			}
			if valueChanged {
				fmt.Println("=====")
				fmt.Printf("p(0)=%.2f p(1)=%.2f p(2)=%.2f\n", values[0], values[1], values[2])
				s = s.SetRGB(values[0], values[1], values[2])
				home.SetScene(s)
			}
		}
	}

}
