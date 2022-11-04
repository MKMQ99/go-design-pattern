package simple_factory

import (
	"fmt"
)

type ship struct{}
type truck struct{}

type API interface {
	deliver(transport string) string
}

func (*ship) deliver(transport string) string {
	return fmt.Sprintf("SeaLogistics using %s", transport)
}

func (*truck) deliver(transport string) string {
	return fmt.Sprintf("RoadLogistics using %s", transport)
}

func NewAPI(i int) API {
	if i == 1 {
		return &ship{}
	} else if i == 2 {
		return &truck{}
	}
	return nil
}
