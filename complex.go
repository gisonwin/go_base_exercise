package main

import "fmt"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/7 14:51
 */
const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type ComplexColor byte
type ComplexBox struct {
	width, height, depth float64
	color                ComplexColor
}
type ComplexBoxList []ComplexBox

func (b ComplexBox) Volume() float64 {
	return b.height * b.width * b.depth
}
func (b *ComplexBox) SetColor(c ComplexColor) {
	b.color = c
	println("SetColor is:", b.color.String())
}
func (bl ComplexBoxList) PaintItBlack() {
	for key, _ := range bl {
		(&bl[key]).SetColor(BLACK)
	}
}
func (c ComplexColor) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}
func RunComplext()  {
	boxes := ComplexBoxList{
		ComplexBox{4, 4, 4, RED},
		ComplexBox{10, 10, 1, YELLOW},
		ComplexBox{1, 1, 20, BLACK},
		ComplexBox{10, 10, 1, BLUE},
		ComplexBox{10, 30, 1, WHITE},
		ComplexBox{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxes in list.\n The Volume of the first one is %v cm3.\n", len(boxes), boxes[0].Volume())
	fmt.Printf("The color of the last one is %s.\n Let's paint the all black \n", boxes[len(boxes)-1].color.String())
	boxes.PaintItBlack()
	fmt.Printf("The color of the second one is %s\n", boxes[1].color.String())
	boxes[1].SetColor(RED)
	fmt.Printf("Now the color of second one is %s\n", boxes[1].color.String())
}
func main() {
	RunComplext()
}

