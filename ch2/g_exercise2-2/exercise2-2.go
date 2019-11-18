package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	conversions "../conversionsPkg"
)

func main() {
	meas, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	farPtr := flag.Bool("f", false, "a bool")
	celPtr := flag.Bool("c", false, "a bool")
	poundPtr := flag.Bool("lb", false, "a bool")
	kgPtr := flag.Bool("kg", false, "a bool")
	feetPtr := flag.Bool("ft", false, "a bool")
	meterPtr := flag.Bool("m", false, "a bool")

	flag.Parse()

	switch {
	case *farPtr:
		c := conversions.Celsius(meas)
		fmt.Printf("%v Celsius = %v Fahrenheit\n", meas, conversions.CToF(c))
	case *celPtr:
		f := conversions.Fahrenheit(meas)
		fmt.Printf("%v Fahrenheit = %v Celsius\n", meas, conversions.FToC(f))
	case *poundPtr:
		kg := conversions.Kilograms(meas)
		fmt.Printf("%v kg = %v lbs\n", meas, conversions.KgToLbs(kg))
	case *kgPtr:
		lbs := conversions.Pounds(meas)
		fmt.Printf("%v lbs = %v kg\n", meas, conversions.LbsToKg(lbs))
	case *feetPtr:
		m := conversions.Meters(meas)
		fmt.Printf("%v m = %v ft\n", meas, conversions.MToFt(m))
	case *meterPtr:
		ft := conversions.Feet(meas)
		fmt.Printf("%v ft = %v m\n", meas, conversions.FtToM(ft))
	}
}
