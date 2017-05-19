package vertex

import "math"

type Vertex struct {
	X, Y float64;
}

func (self *Vertex) Abs() float64  {
	return math.Sqrt2(self.X*self.X+self.Y*self.Y)
}

