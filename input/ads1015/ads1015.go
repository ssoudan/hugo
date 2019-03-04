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
package ads1015

import (
	"github.com/ssoudan/edisonIsThePilot/drivers/gpio"
	"github.com/ssoudan/hugo/logging"

	"bitbucket.org/gmcbay/i2c"

	"encoding/binary"
	"errors"
	"time"
)

var log = logging.Log("ads1015")

type ADS1015 struct {
	bus     byte
	address byte
	i2c     *i2c.I2CBus
	scaler  float32
}

const (
	// i2c6SCL is the pin number of SCL line for i2c bus number 6
	i2c6SCL = 27
	// i2c6SDA is the pin number of SDA line for i2c bus number 6
	i2c6SDA = 28
	// i2c1SCL is the pin number of SCL line for i2c bus number 1
	i2c1SCL = 19
	// i2c1SDA is the pin number of SDA line for i2c bus number 1
	i2c1SDA = 20
)

// New creates a new ADS1015 driver on a i2c bus of the Edison
func New(bus byte, address byte) (*ADS1015, error) {

	switch bus {
	case 1:
		gpio.EnableI2C(i2c1SCL)
		gpio.EnableI2C(i2c1SDA)
		gpio.EnableFastI2C(1)
	case 6:
		gpio.EnableI2C(i2c6SCL)
		gpio.EnableI2C(i2c6SDA)
		gpio.EnableFastI2C(6)
	default:
		log.Panic("Unknown i2c bus")
	}

	i2c, err := i2c.Bus(bus)
	if err != nil {
		return nil, err
	}

	return &ADS1015{bus: bus, address: address, i2c: i2c}, nil
}

func (a ADS1015) writeRegister16(reg uint8, value uint16) error {
	log.Debug("Writing [", value, "]@", reg)
	b := toBytes(value)
	return a.i2c.WriteByteBlock(a.address, reg, b)
}

func (a ADS1015) readRegister16(reg byte) (uint16, error) {
	log.Debug("Reading 2@", reg)
	d, err := a.i2c.ReadByteBlock(a.address, reg, 2)
	log.Debug("->", d)
	return fromBytes(d), err
}

func fromBytes(value []uint8) uint16 {
	return binary.BigEndian.Uint16(value)
}

func toBytes(value uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, value)

	return b
}

type VoltageRange uint16

const (
	V6_144 = iota
	V4_096 = iota
	V2_048 = iota
	V1_024 = iota
	V0_512 = iota
	V0_256 = iota
)

const (
	RangeMask  = 0x0E00
	RangeShift = 9
)

// SetRange sets the voltage range
func (a *ADS1015) SetRange(r VoltageRange) error {
	cfgRegVal, err := a.getConfigRegister()
	if err != nil {
		return err
	}
	cfgRegVal &^= RangeMask
	cfgRegVal |= (uint16(r) << RangeShift) & RangeMask
	a.setConfigRegister(cfgRegVal)
	switch r {
	case V6_144:
		a.scaler = 3.0 // each count represents 3.0 mV

	case V4_096:
		a.scaler = 2.0 // each count represents 2.0 mV

	case V2_048:
		a.scaler = 1.0 // each count represents 1.0 mV

	case V1_024:
		a.scaler = 0.5 // each count represents 0.5mV

	case V0_512:
		a.scaler = 0.25 // each count represents 0.25mV

	case V0_256:
		a.scaler = 0.125 // each count represents 0.125mV

	default:
		return errors.New("Invalid voltage range")
	}

	return nil
}

const (
	ConversionRegister = 0x0
	ConfigRegister     = 0x1
)

func (a ADS1015) getConfigRegister() (uint16, error) {
	return a.readRegister16(ConfigRegister)
}

func (a ADS1015) setConfigRegister(conf uint16) error {
	return a.writeRegister16(ConfigRegister, conf)
}

const (
	StartReadFlag = 0x8000
	BusyMask      = 0x8000
	SingleEnded   = 0x4000
	ChannelMask   = uint16(0x3000)
	ChannelShift  = 12
)

func (a ADS1015) readADC() (int16, error) {

	configRegVal, err := a.getConfigRegister()
	if err != nil {
		return 0, err
	}
	configRegVal |= StartReadFlag
	err = a.setConfigRegister(configRegVal)
	if err != nil {
		return 0, err
	}

	busyDelay := 0

	configRegVal, err = a.getConfigRegister()
	if err != nil {
		return 0, err
	}
	for configRegVal&BusyMask == 0 {
		time.Sleep(100 * time.Microsecond)
		if busyDelay > 100 {
			return 0, errors.New("Failed to read value from ADC")
		}
		busyDelay++

		configRegVal, err = a.getConfigRegister()
		if err != nil {
			return 0, err
		}
	}

	result, err := a.readRegister16(ConversionRegister)

	log.Debug("Value=", result)

	return int16(result >> 4), nil
}

// getRawResult returns single ended raw result. Single-ended results are effectively unsigned 11-bit values, from 0 to 2047.
func (a ADS1015) getRawResult(channel byte) (int16, error) {
	cfgRegVal, err := a.getConfigRegister()
	if err != nil {
		return 0, err
	}

	cfgRegVal &= ^ChannelMask                                    // clear existing channel settings
	cfgRegVal |= SingleEnded                                     // set the SE bit for a s-e read
	cfgRegVal |= (uint16(channel) << ChannelShift) & ChannelMask // put the channel bits in
	cfgRegVal |= StartReadFlag                                   // set the start read bit

	err = a.setConfigRegister(cfgRegVal)
	if err != nil {
		return 0, err
	}

	return a.readADC()
}

// GetResult returns the current reading on a channel, scaled by the current scaler and
// presented as a floating point number.
func (a ADS1015) GetResult(channel byte) (float32, error) {
	rawVal, err := a.getRawResult(channel)
	log.Debug("rawVal=", rawVal)
	return float32(rawVal) * a.scaler / 1000., err
}
