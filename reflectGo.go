package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 0
	t := reflect.TypeOf(i) //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	fmt.Println(t)
	v:= reflect.ValueOf(i)//得到实际的值,通过v我们获取存储在里面的值,还可以改值.
	fmt.Println(v)

/**
	tag := t.Elem().Field(0).Tag
	fmt.Println(tag)
	name := v.Elem().Field(0).String()
	fmt.Println(name)
*/
 var x float64 = 3.1415926
  v = reflect.ValueOf(x)
  fmt.Println("type : ",v.Type())
  fmt.Println("kind is float64 : ",v.Kind() == reflect.Float64)
  fmt.Println("value ",v.Float())


  //通过反射修改值
  p:= reflect.ValueOf(&x)
  v = p.Elem()
  v.SetFloat(5.3211)
  fmt.Println("reflect and now value is ",v)
}
