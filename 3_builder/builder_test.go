package builder

import "testing"

func TestBuilder(t *testing.T) {
	direct := Director{}
	carBuilder := CarBuilder{}
	direct.ConstructSportsCar(&carBuilder)
	s := carBuilder.car.ret
	if s != "Engine" {
		t.Fatal("Construct Sportscar Error")
	}
	manualBuilder := CarManualBuilder{}
	direct.ConstructSUV(&manualBuilder)
	s = manualBuilder.manual.ret
	if s != "EngineGPS" {
		t.Fatal("Construct SUV Error")
	}
}
