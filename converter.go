package main

func ConvertLength(unit float64, from string, to string) float64 {
	switch from {
	case "millimeter":
		switch to {
		case "centimeter":
			return unit / 10
		case "meter":
			return unit / 1_000
		case "kilometer":
			return unit / 1_000_000
		case "inch":
			return unit / 25.4
		case "foot":
			return unit / 304.8
		case "yard":
			return unit / 914.4
		case "mile":
			return unit / 1.609e+6
		}
	case "centimeter":
		switch to {
		case "milimeter":
			return unit * 10
		case "meter":
			return unit / 100
		case "kilometer":
			return unit / 100_000
		case "inch":
			return unit / 2.54
		case "foot":
			return unit / 30.48
		case "yard":
			return unit / 91.44
		case "mile":
			return unit / 160_900
		}
	case "meter":
		switch to {
		case "milimeter":
			return unit * 1_000
		case "centimeter":
			return unit * 100
		case "kilometer":
			return unit / 1_000
		case "inch":
			return unit * 39.37
		case "foot":
			return unit * 3.281
		case "yard":
			return unit * 1.094
		case "mile":
			return unit / 1_609
		}
	case "kilometer":
		switch to {
		case "milimeter":
			return unit * 1_000_000
		case "centimeter":
			return unit * 100_000
		case "meter":
			return unit * 1_000
		case "inch":
			return unit * 39370
		case "foot":
			return unit * 3281
		case "yard":
			return unit * 1_094
		case "mile":
			return unit / 1.609
		}
	case "inch":
		switch to {
		case "milimeter":
			return unit * 25.4
		case "centimeter":
			return unit * 2.54
		case "meter":
			return unit / 39.37
		case "kilometer":
			return unit / 39_370
		case "foot":
			return unit / 12
		case "yard":
			return unit / 16
		case "mile":
			return unit / 63_360
		}
	case "foot":
		switch to {
		case "milimeter":
			return unit * 304.8
		case "centimeter":
			return unit * 30.48
		case "meter":
			return unit / 3.281
		case "kilometer":
			return unit / 3281
		case "inch":
			return unit / 12
		case "yard":
			return unit / 3
		case "mile":
			return unit / 5_280
		}
	case "yard":
		switch to {
		case "milimeter":
			return unit * 914.4
		case "centimeter":
			return unit * 91.44
		case "meter":
			return unit / 1.094
		case "kilometer":
			return unit / 1094
		case "inch":
			return unit * 36
		case "foot":
			return unit * 3
		case "mile":
			return unit / 1_760
		}
	case "mile":
		switch to {
		case "milimeter":
			return unit * 1.609e+6
		case "centimeter":
			return unit * 160_900
		case "meter":
			return unit * 1_609
		case "kilometer":
			return unit * 1.609
		case "inch":
			return unit * 63_360
		case "foot":
			return unit * 5_280
		case "yard":
			return unit * 1_760
		}
	}

	return unit
}

func ConvertWeight(unit float64, from string, to string) float64 {
	switch from {
	case "milligram":
		switch to {
		case "gram":
			return unit / 1_000
		case "kilogram":
			return unit / 1e+6
		case "ounce":
			return unit / 28_350
		case "pound":
			return unit / 453_600
		}
	case "gram":
		switch to {
		case "milligram":
			return unit * 1_000
		case "kilogram":
			return unit / 1_000
		case "ounce":
			return unit / 28.35
		case "pound":
			return unit / 453.6
		}
	case "kilogram":
		switch to {
		case "milligram":
			return unit * 1e+6
		case "gram":
			return unit * 1_000
		case "ounce":
			return unit * 35.274
		case "pound":
			return unit * 2.205
		}
	case "ounce":
		switch to {
		case "milligram":
			return unit * 28_350
		case "gram":
			return unit * 28.35
		case "kilogram":
			return unit / 35.274
		case "pound":
			return unit / 16
		}
	case "pound":
		switch to {
		case "milligram":
			return unit * 453_600
		case "gram":
			return unit * 453.6
		case "kilogram":
			return unit / 2.205
		case "ounce":
			return unit * 16
		}
	}

	return unit
}

func ConvertTemperature(unit float64, from string, to string) float64 {
	switch from {
	case "celcius":
		switch to {
		case "farenheit":
			return unit*(9.0/5.0) + 32.0
		case "kelvin":
			return unit + 273.15
		}
	case "farenheit":
		switch to {
		case "celcius":
			return (unit - 32.0) * (5.0 / 9.0)
		case "kelvin":
			return (unit-32.0)*(5.0/9.0) + 273.15
		}
	case "kelvin":
		switch to {
		case "celcius":
			return unit - 273.15
		case "farenheit":
			return (unit-273.15)*(9/5) + 32
		}
	}

	return unit
}
