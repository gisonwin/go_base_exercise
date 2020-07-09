package exercise

import "fmt"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/7 10:38
 */
type Rectangle struct {
	Width  float32 //宽
	Height float32 //高
	Color  string  //颜色
}

func (rect Rectangle) Area() float32 {
	return rect.Width * rect.Height
}
func (rect Rectangle) Painting() string {
	fmt.Println(rect.Color)
	return rect.Color
}
