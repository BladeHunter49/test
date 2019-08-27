package main

import "fmt"

type rectangle struct {
	length, width float32
}

type square struct {
	side float32
}

type triangle struct {
	bottom, heigh float32
}

type shaper interface {
	area() float32
}

func (r rectangle) area() float32 {
	return r.length * r.width
}

func (s *square) area() float32 {
	return s.side * s.side
}

func (t *triangle) area() float32 {
	return 0.5 * t.bottom * t.heigh
}

func main() {
	r := rectangle{10, 5}
	s := &square{4}
	t := &triangle{3, 4}

	//切片初始化
	//shapes := []shaper{r, s, t}
	var shapes []shaper = []shaper{r, s, t}

	for _, v := range shapes {
		//fmt.Println("detail:", shapes[k])
		//fmt.Println("algor area:", shapes[k].area())
		fmt.Println("detail:", v)
		fmt.Println("algor area:", v.area())
	}

	//类型断言
	if value, ok := shapes[0].(rectangle); ok {
		fmt.Printf("Type of shapes[0] is %T\n", value)
	}

	switch t := shapes[1].(type) {
	case *square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *triangle:
		fmt.Printf("Type triangle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}
