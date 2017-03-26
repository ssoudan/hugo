//
// Copyright 2017 Sebastien Soudan
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
package input

import "github.com/ssoudan/hugo/input/ads1015"

// Potentiometers is a group of potentiometers plugged to the ADS1015
type Potentiometers struct {
	adc     *ads1015.ADS1015
	channel byte
	min     map[byte]float32
	max     map[byte]float32
}

// Constants
const (
	I2CBus         = 1
	ADS1015Address = 0x48
)

// New creates a new group of Potentiometer
func New() (*Potentiometers, error) {
	adc, err := ads1015.New(I2CBus, ADS1015Address)
	if err != nil {
		return nil, err
	}

	err = adc.SetRange(ads1015.V4_096)
	if err != nil {
		return nil, err
	}

	return &Potentiometers{adc: adc, min: make(map[byte]float32), max: make(map[byte]float32)}, nil
}

// Read returns the value of the potentiometer
func (p *Potentiometers) Read(channel byte) (float32, error) {
	return p.adc.GetResult(channel)
}

// ReadAndScale reads and scales the value of the potentiometer between 0 and 1
func (p *Potentiometers) ReadAndScale(channel byte) (float32, error) {
	v, err := p.adc.GetResult(channel)
	if err != nil {
		return 0, err
	}

	max, okMax := p.max[channel]
	min, okMin := p.min[channel]

	if min > v || !okMin {
		p.min[channel] = v
		min = v
	}

	if max < v || !okMax {
		p.max[channel] = v
		max = v
	}

	if max == min {
		return 0.5, nil
	}

	// fmt.Printf("c=%d m=%f M=%f\n", channel, min, max)

	return (v - min) / (max - min), err
}
