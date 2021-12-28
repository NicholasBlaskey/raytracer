package intersection

type Intersection struct {
	T   float64
	Obj interface{} // Make a shape interface (eventually)
}

func New(t float64, obj interface{}) *Intersection {
	return &Intersection{t, obj}
}

// See if this function is useful as syntatic sugar?
func Aggregate(intersections ...*Intersection) []*Intersection {
	return intersections
}

func Hit(intersections []*Intersection) *Intersection {
	minT := float64(9999999)
	var minIntersect *Intersection
	for _, i := range intersections {
		if i.T < minT && i.T >= 0.0 {
			minIntersect = i
			minT = i.T
		}
	}

	return minIntersect
}
