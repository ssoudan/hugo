package input

import "github.com/ssoudan/hugo/input/ads1015"

type Potentiometers struct {
	adc *ads1015.ADS1015
	channel byte
}

const I2CBus = 1
const ADS1015Address = 0x48

func New() (*Potentiometers, error) {
	adc, err := ads1015.New(I2CBus, ADS1015Address)
	if err != nil {
		return nil, err
	}

	err = adc.SetRange(ads1015.V4_096)
	if err != nil {
		return nil, err
	}

	return &Potentiometers{adc: adc}, nil
}


// Read returns the value of the potentiometer
func (p *Potentiometers) Read(channel byte) (float32, error) {
	return p.adc.GetResult(channel)
}