package golang_united_school_homework

import (
	"errors"
	"fmt"
	"reflect"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	//
	errorFull = errors.New("the array is full")
	//
	erroNoElement = errors.New("there is no element")
	//
	erroNoCircles = errors.New("there are no circles")
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity > len(b.shapes) {
		b.shapes = append(b.shapes, shape)
		return nil
	} else {
		return fmt.Errorf("The array is full: %w", errorFull)
	}
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity || i >= len(b.shapes) {
		return nil, fmt.Errorf("There is no element: %w", erroNoElement)
	} else {
		return b.shapes[i], nil
	}
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity || i >= len(b.shapes) {
		return nil, fmt.Errorf("There is no element: %w", erroNoElement)
	} else {
		element := b.shapes[i]
		b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
		return element, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 0 || i >= b.shapesCapacity || i >= len(b.shapes) {
		return nil, fmt.Errorf("There is no element: %w", erroNoElement)
	} else {
		remShape := b.shapes[i]
		b.shapes[i] = shape
		return remShape, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	sum := 0.0
	for _, v := range b.shapes {
		sum += v.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	area := 0.0
	for _, v := range b.shapes {
		area += v.CalcArea()
	}
	return area
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	bFlag := false
	for i, v := range b.shapes {
		if reflect.TypeOf(v) == reflect.TypeOf(Circle) {
			b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)
			bFlag = true
		}
	}

	if bFlag {
		return nil
	} else {
		return fmt.Errorf("There are no circles: %w", erroNoCircles)
	}
}
