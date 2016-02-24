package atmosphere

import "testing"
import (
	"flag"
	"math"
)

// Returns true if a and b are within err of each other
func moreOrLessEqual(a, b, err float64) bool {
	return math.Abs(a-b) < err
}

func TestCalculation(t *testing.T) {
	flag.Parse()

	temp := TemperatureC(21.0)
	pressure,err := NewPressure(1007.0, "hPa")
	if err != nil	{
		t.Errorf("Error: %v", err)
	}
	humidity := RelativeHumidity(50.0)

	density := AirDensity(temp, pressure, humidity)
	t.Logf("Air density is %v", density)

	if !moreOrLessEqual(float64(density), 1.187, 0.001) {
		t.Errorf("Expected density %f, got %v", 1.187, density)
	}
}
