package lights

import (
	"fmt"
)

const (
	turnOnEvent  string = "turn_lights_on"
	turnOffEvent string = "turn_lights_off"
)

type Triggerer interface {
	TriggerEvent(event string) error
}

type Switcher struct {
	trig Triggerer
}

func NewSwitcher(trig Triggerer) *Switcher {
	return &Switcher{trig}
}

func (s *Switcher) TurnOn() error {
	if err := s.trig.TriggerEvent(turnOnEvent); err != nil {
		return fmt.Errorf("could not trigger turn on lights event: %v", err)
	}
	return nil
}

func (s *Switcher) TurnOff() error {
	if err := s.trig.TriggerEvent(turnOffEvent); err != nil {
		return fmt.Errorf("could not trigger turn off lights event: %v", err)
	}
	return nil
}
