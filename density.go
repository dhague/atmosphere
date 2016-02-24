package atmosphere

type TemperatureC float64
type Pressure float64
type RelativeHumidity float64
type Density float64

func AirDensity(t TemperatureC, p Pressure, h RelativeHumidity) Density {
	const (
		Rd = 287.05 // Specific gas constant for dry air J/(KgK)
		Rv = 461.495 // Specific gas constant for water vapour J/(KgK)
	)
	waterVapourPressure := saturationPressure(t) * h
	dryAirPressure := p - waterVapourPressure
	temperatureK := t + 273.15
	return dryAirPressure/(Rd*temperatureK) + waterVapourPressure/(Rv*temperatureK)
}

func saturationPressure(t TemperatureC) Pressure {
	const (
		Eso = 6.1078
		c0 = 0.99999683
		c1 = -0.90826951e-2
		c2 = 0.78736169e-4
		c3 = -0.61117958e-6
		c4 = 0.43884187e-8
		c5 = -0.29883885e-10
		c6 = 0.21874425e-12
		c7 = -0.17892321e-14
		c8 = 0.11112018e-16
		c9 = -0.30994571e-19
	)
	p := (c0+t*(c1+t*(c2+t*(c3+t*(c4+t*(c5+t*(c6+t*(c7+t*(c8+t*(c9))))))))))

	return Eso / p^8
}
