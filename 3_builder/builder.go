package builder

type Car struct {
	ret string
}

type Manual struct {
	ret string
}

type Builder interface {
	setSeats()
	setEngine()
	setGPS()
}

type CarBuilder struct {
	car Car
}

func (carBuilder *CarBuilder) setSeats() {
	carBuilder.car.ret += "Seat"
}

func (carBuilder *CarBuilder) setEngine() {
	carBuilder.car.ret += "Engine"
}

func (carBuilder *CarBuilder) setGPS() {
	carBuilder.car.ret += "GPS"
}

type CarManualBuilder struct {
	manual Manual
}

func (carManualBuilder *CarManualBuilder) setSeats() {
	carManualBuilder.manual.ret += "Seat"
}

func (carManualBuilder *CarManualBuilder) setEngine() {
	carManualBuilder.manual.ret += "Engine"
}

func (carManualBuilder *CarManualBuilder) setGPS() {
	carManualBuilder.manual.ret += "GPS"
}

type Director struct {
}

type Construct interface {
	ConstructSportsCar(builder Builder)
	ConstructSUV(builder Builder)
}

func (Director) ConstructSportsCar(builder Builder) {
	builder.setEngine()
}

func (Director) ConstructSUV(builder Builder) {
	builder.setEngine()
	builder.setGPS()
}
