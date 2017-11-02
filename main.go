package main

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

func (c *Circle) Sum() int {
	return c.X + c.Y
}

func main() {
	c := &Circle{
		Point:  Point{1, 2},
		Radius: 1,
	}

	if c.Sum() == 3 {
		println("Wow, cool")
	}
}
