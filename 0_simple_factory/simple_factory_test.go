package simple_factory

import "testing"

func TestShipAPI(t *testing.T) {
	api := NewAPI(1)
	s := api.deliver("ship")
	if s != "SeaLogistics using ship" {
		t.Fatal("Ship test fail")
	}
}

func TestTruckAPI(t *testing.T) {
	api := NewAPI(2)
	s := api.deliver("truck")
	if s != "RoadLogistics using truck" {
		t.Fatal("Ship test fail")
	}
}
