package tempconv

import (
	"flag"
	"fmt"
	basetempconv "gobook/ch2/tempconv"
)

type CelsiusFlag struct {
	basetempconv.Celsius
}

func (f *CelsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = basetempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = basetempconv.FtoC(basetempconv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func SetCelsiusFlag(name string, value basetempconv.Celsius, usage string) *basetempconv.Celsius {
	f := CelsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
