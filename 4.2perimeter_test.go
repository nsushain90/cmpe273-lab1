package perimeter

import (
	"fmt"
	"math"
	"testing"
)

/*func CirclePerimeter(t *testing.T) {

	expected := 62.83185307179586
	actual := c.perimeter()
	if actual != expected {
		fmt.Println(actual)
		t.Error("Test failed")
	}
}*/
func TestCirclePerimeter(t *testing.T) {
	r := Rectangle{22, 33, 12, 5}
	var expected float64 = 36
	actual := r.perimeter()
	if actual != expected {
		fmt.Println("The expected perimeter is", actual)
		t.Error("Test failed")
	}
}
func TestRectanglePerimeter(t *testing.T) {
	c := Circle{1, 1, 10}
	var expected float64 = 62.83185307179586
	actual := c.perimeter()
	if actual != expected {
		fmt.Println("The expected perimeter is ", actual)
		t.Error("Test failed")
	}
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type shape interface {
	/*area() float64*/
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

/*func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}*/

func (c *Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

/*func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}*/

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func (c *Circle) circle() float64 {
	return math.Pi * c.r * c.r
}

/*func main() {
	c := Circle{1, 1, 10}
	r := Rectangle{22, 33, 12, 5}
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}*/
