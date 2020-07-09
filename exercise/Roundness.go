package exercise

import "math"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/7 10:39
 */

type Roundness struct {
	Radius float32 //半径
}

func (round Roundness) Area() float32 {
	return math.Pi * round.Radius * round.Radius
}
