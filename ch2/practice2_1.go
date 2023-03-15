package main

import (
	"fmt"
	"gobook/ch2/tempconv"
)

type KelvinScale float64

func (k KelvinScale) String() string { return fmt.Sprintf("%g°K", k) }

func CtoK(c tempconv.Celsius) KelvinScale { return KelvinScale(c + tempconv.AbsoluteZeroC) }

func KtoC(k KelvinScale) tempconv.Celsius { return tempconv.Celsius(k) - tempconv.AbsoluteZeroC }

func FtoK(f tempconv.Fahrenheit) KelvinScale { return CtoK(tempconv.FtoC(f)) }

func KtoF(k KelvinScale) tempconv.Fahrenheit { return tempconv.CtoF(KtoC(k)) }
