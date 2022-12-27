package tempconv

import "fmt"

type KelvinScale float64

func (k KelvinScale) String() string { return fmt.Sprintf("%g°K", k) }

func CtoK(c Celsius) KelvinScale { return KelvinScale(c + AbsoluteZeroC) }

func KtoC(k KelvinScale) Celsius { return Celsius(k) - AbsoluteZeroC }

func FtoK(f Fahrenheit) KelvinScale { return CtoK(FtoC(f)) }

func KtoF(k KelvinScale) Fahrenheit { return CtoF(KtoC(k)) }
