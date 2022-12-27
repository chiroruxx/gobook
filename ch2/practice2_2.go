package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Feet float64
type Meeter float64

const feetPerMeeter = 3.28084

func main() {
	args := getItems()

	for _, arg := range args {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "practice2_2: %v\n", err)
			os.Exit(1)
		}

		feet := Feet(value)
		meeter := Meeter(value)

		fmt.Printf("%vf = %vm\t", feet, feetToMeeter(feet))
		fmt.Printf("%vm = %vf\n", meeter, meeterToFeet(meeter))

	}
}

func getItems() []string {
	items := os.Args[1:]
	if len(items) != 0 {
		return items
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		items = append(items, input.Text())
	}

	return items
}

func feetToMeeter(feet Feet) Meeter {
	value := float64(feet) / feetPerMeeter
	return Meeter(value)
}

func meeterToFeet(meeter Meeter) Feet {
	value := float64(meeter) * feetPerMeeter
	return Feet(value)
}
