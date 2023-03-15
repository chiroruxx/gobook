package practice7_6

import (
	"flag"
	"fmt"
	"gobook/ch2/tempconv"
)

type KelvinScale float64

func (k KelvinScale) String() string { return fmt.Sprintf("%g°K", k) }

type celsius2Flag struct {
	tempconv.Celsius
}

func (f *celsius2Flag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FtoC(tempconv.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = KtoC2(KelvinScale(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func KtoC2(k KelvinScale) tempconv.Celsius { return tempconv.Celsius(k) - tempconv.AbsoluteZeroC }

func Celsius2Flag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsius2Flag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
