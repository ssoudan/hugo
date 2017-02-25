package main

import (
	"flag"
	"fmt"
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

	for {
		fmt.Println("=====")

		v, err := p.Read(0)
		check(err)
		fmt.Printf("p(0)=%f\n", v)

		v, err = p.Read(1)
		check(err)
		fmt.Printf("p(1)=%f\n", v)

		v, err = p.Read(2)
		check(err)
		fmt.Printf("p(2)=%f\n", v)

		v, err = p.Read(3)
		check(err)
		fmt.Printf("p(3)=%f\n", v)

		if v < 1.5 {
			fmt.Println("rotating")
			s = s.Rotate()
			home.SetScene(s)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
