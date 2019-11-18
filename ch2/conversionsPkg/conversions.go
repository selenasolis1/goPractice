package conversions

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meters float64
type Pounds float64
type Kilograms float64
type Miles float64
type Kilometers float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FtToM(ft Feet) Meters {
	return Meters(ft / 3.2808)
}

func MToFt(m Meters) Feet {
	return Feet(m * 3.2808)
}

func LbsToKg(lbs Pounds) Kilograms {
	return Kilograms(lbs / 2.2046)
}

func KgToLbs(kg Kilograms) Pounds {
	return Pounds(kg * 2.2046)
}
