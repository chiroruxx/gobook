package main

import (
	"flag"
	"fmt"
)

const (
	AbsoluteZeroC Celsius2 = -273.15
)

type Celsius2 float64

func (c Celsius2) String() string { return fmt.Sprintf("%g°C", c) }

type Fahrenheit2 float64

func (f Fahrenheit2) String() string { return fmt.Sprintf("%g°F", f) }

type KelvinScale float64

func (k KelvinScale) String() string { return fmt.Sprintf("%g°K", k) }

type celsius2Flag struct {
	Celsius2
}

func (f *celsius2Flag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius2 = Celsius2(value)
		return nil
	case "F", "°F":
		f.Celsius2 = F2toC2(Fahrenheit2(value))
		return nil
	case "K", "°K":
		f.Celsius2 = KtoC2(KelvinScale(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func F2toC2(f Fahrenheit2) Celsius2 { return Celsius2((f - 32) * 5 / 9) }
func KtoC2(k KelvinScale) Celsius2  { return Celsius2(k) - AbsoluteZeroC }

func Celsius2Flag(name string, value Celsius2, usage string) *Celsius2 {
	f := celsius2Flag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius2
}

var temp2 = Celsius2Flag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp2)
}
