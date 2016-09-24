package homekit

import (
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"

	"github.com/ssoudan/hugo/logging"
)

var log = logging.Log("homekit")

// Pin code for the registration of the Accessory
const Pin = "00102003"

// CreateSwitch creates a new Switch
func CreateSwitch(name string, callback func(bool)) *accessory.Accessory {

	info := accessory.Info{
		Name:         name,
		SerialNumber: "051AC-23AAM1",
		Manufacturer: "Apple",
		Model:        "AB",
	}
	acc := accessory.NewSwitch(info)

	acc.Switch.On.OnValueRemoteUpdate(callback)
	return acc.Accessory
}

// CreateLightBuld creates a new Switch
func CreateLightBuld(name string, onCallback func(bool), brightnessCallback func(int), saturationCallback func(float64), hueCallback func(float64)) *accessory.Accessory {

	info := accessory.Info{
		Name:         name,
		SerialNumber: "something",
		Manufacturer: "Seb",
		Model:        "AB",
	}
	acc := accessory.NewLightbulb(info)

	acc.Lightbulb.On.OnValueRemoteUpdate(onCallback)
	acc.Lightbulb.Brightness.OnValueRemoteUpdate(brightnessCallback)
	acc.Lightbulb.Saturation.OnValueRemoteUpdate(saturationCallback)
	acc.Lightbulb.Hue.OnValueRemoteUpdate(hueCallback)

	return acc.Accessory
}

// Start the homekit server
func Start(a *accessory.Accessory, accessories ...*accessory.Accessory) {

	config := hc.Config{Pin: Pin}
	t, err := hc.NewIPTransport(config, a, accessories...)
	if err != nil {
		log.Fatal(err)
	}

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()
}
