package polygo

// Polygon is the intern representation of a polygon
type Polygon struct {
	corners []Point
}

// NewPolygon creates a new Polygon
func NewPolygon(points ...Point) Polygon {
	var poly Polygon
	for _, p := range points {
		poly.corners = append(poly.corners, p)
	}
	return poly
}

// Point consists of X and Y
type Point struct {
	X float64
	Y float64
}

// Contains checks if the point q is in the polygon or at least at the edge.
// source: https://de.wikipedia.org/wiki/Punkt-in-Polygon-Test_nach_Jordan
func (p Polygon) Contains(q Point) bool {
	// res := polyTest(p, q)
	// if res == 1 || res == 0 {
	// 	return true
	// }
	// return false
	res := pnpoly(p, q)
	return res

}

// https://de.wikipedia.org/wiki/Diskussion:Punkt-in-Polygon-Test_nach_Jordan (2019-07-10 22:44:24)
func pnpoly(p Polygon, q Point) bool {
	var i int
	var c bool
	c = false
	i = 0
	npol := len(p.corners)
	for j := npol - 1; j < npol; j = i + 1 {
		if ((p.corners[i].Y <= q.Y && q.Y < p.corners[j].Y) || (p.corners[j].Y <= q.Y && q.Y < p.corners[i].Y)) && (q.X < (p.corners[j].X-p.corners[i].X)*(q.Y-p.corners[i].Y)/(p.corners[j].Y-p.corners[i].Y+p.corners[i].X)) {
			c = !c
		}
	}
	return c
}

func polyTest(p Polygon, q Point) int {
	var c []Point
	l := len(p.corners) - 1
	c = append(c, p.corners[l])
	c = append(c, p.corners...)
	t := -1
	for i := 0; i < len(c); i++ {
		t = t * crossProductCheck(q, c[i], c[i+1])
		if t == 0 {
			break
		}
	}
	return t
}

func crossProductCheck(a Point, b Point, c Point) int {
	if b.Y > c.Y {
		b, c = c, b
	}
	if a.Y <= b.Y || a.Y > c.Y {
		return 1
	}
	delta := (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
	if delta > 0 {
		return 1
	} else if delta < 0 {
		return -1
	} else {
		return 0
	}
}
