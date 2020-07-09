package main

import (
	"fmt"
	"math"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/1 20:49
 */
type Shape interface {
	Area()
	Paint()
}
type Circle struct{}

var radius float32 = 4.5
var color string = "red"

func (c Circle) Area() {
	fmt.Println(radius * radius * math.Pi)
}
func (c Circle) Paint() {
	fmt.Println("painting :", color)
}
func calc(shape Shape) {
	shape.Area()
	shape.Paint()
}
func main() {
	calc(Circle{})
}
