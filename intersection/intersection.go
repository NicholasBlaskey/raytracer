package intersection

import (
	"fmt"
)

type Intersection struct {
	T   float64
	Obj interface{} // Make a shape interface (eventually)
}

func New(t float64, obj interface{}) Intersection {
	fmt.Println("SEARCH", obj)
	return Intersection{t, obj}
}
