package main

func ConvertTemperature(unit float64, from string, to string) float64 {
	if from == "celcius" {
		if to == "farenheit" {
			return unit*(9.0/5.0) + 32.0
		} else if to == "kelvin" {
			return unit + 273.15
		}
	} else if from == "farenheit" {
		if to == "celcius" {
			return (unit - 32.0) * (5.0 / 9.0)
		} else if to == "kelvin" {
			return (unit-32.0)*(5.0/9.0) + 273.15
		}
	} else if from == "kelvin" {
		if to == "celcius" {
			return unit - 273.15
		} else if to == "farenheit" {
			return (unit-273.15)*(9/5) + 32
		}
	}

	return 0.0
}
