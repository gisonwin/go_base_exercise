package main

import "math"

type person struct {
	name string
	age  int
}

type student struct {
	person //匿名字段,默认Student包含了person的所有字段同
	specially string
}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p1.age

}

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func calcArea(r Rectangle) float64 {
	return r.width * r.height
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

/***
使用method的几点注意事项:
	虽然method名字一模一样,但如果接收者不一样,则method就不一样
	method里面可以访问接收者的字段
	调用method通过.访问,就像struct里面访问字段一样
	method可以定义在任何自定义的类型,内置类型,struct各种类型上面.
*/
type Color byte
type Box struct {
	width, height, depth float64
	color                Color
}
type BoxList [] Box

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}
func (b *Box) SetColor(c Color) {
	b.color = c
}

//method就是将java的顺序全倒过来了
//public int calc(int a,int b)
