package main

import (
	"errors"
	"fmt"
	"main/exercise"
)

var inter interface{}

func main() {
	inter = 3
	//testBasic()
	//LogicFunc()
	//testStruct()
	//LearnInterface()
	//var tom person
	//tom.name, tom.age = "Tom", 27
	//print(tom)
	//bob := person{"bob", 25}
	//paul := person{name: "Paul", age: 18}
	//tb_Older, tb_diff := Older(tom, bob)
	//tp_Older, tp_diff := Older(tom, paul)
	//bp_Older, bp_diff := Older(bob, paul)
	//fmt.Printf("of %s and %s,%s is older by %d years\n", tom.name, bob.name, tb_Older.name, tb_diff)
	//fmt.Printf("of %s and %s,%s is older by %d years\n", tom.name, paul.name, tp_Older.name, tp_diff)
	//fmt.Printf("of %s and %s,%s is older by %d years\n", bob.name, paul.name, bp_Older.name, bp_diff)

	//sequence := getSequence()
	//fmt.Println(sequence(), sequence(), sequence())
	//index := getSequence()
	//fmt.Println(index(), index(), index())
	//SliceOperations()
	//PointOperation()
	//ExerciseInterface()

	InheritAndOverride()

}

func ExerciseInterface() {
	rect := exercise.Rectangle{
		Width:  123.4,
		Height: 567.8,
		Color:  "Red",
	}
	fmt.Printf("Rectangle 's width is %.1f,height is %.1f,and its Area is %.2f,Color is %s\n", rect.Width, rect.Height, rect.Area(), rect.Painting())

	r := exercise.Roundness{
		Radius: 234.5,
	}
	fmt.Println(r)
	r = exercise.Roundness{Radius: 567.8}
	fmt.Println(r)
	fmt.Printf("Roundness 's radius %.2f,and its Area is %.2f\n", r.Radius, r.Area())
}
func SliceOperations() {
	var slice1 = make([]int, 5, 10)
	println(slice1)
	slice2 := []int{4: 1}
	array := [5]int{4: 1}
	fmt.Println("slice:", slice2, ",array:", array)
	slice3 := []string{"a", "b", "c", "d", "e", "f", "g"}
	//for key, value := range slice3 {
	//	fmt.Printf("slice3==>key: %d,value: %s .\n", key, value)
	//}
	fmt.Println(slice3)
	fmt.Printf("%p\n", &slice3)
	slice4 := slice3[:]
	fmt.Println(slice4)
	fmt.Printf("%p\n", &slice4)
	//for key, value := range slice4 {
	//	fmt.Printf("slice4==>key: %d,value: %s .\n", key, value)
	//}
	slice5 := slice3[0:]
	fmt.Println(slice5)
	//for key, value := range slice5 {
	//	fmt.Printf("slice5==>key: %d,value: %s .\n", key, value)
	//}
	slice5[4] = "t"
	fmt.Println("slice5:", slice5)
	//for key, value := range slice5 {
	//	fmt.Printf("slice5==>key: %d,value: %s .\n", key, value)
	//}
	//slice6 := slice3[:7]
	//for key, value := range slice6 {
	//	fmt.Printf("slice6==>key: %d,value: %s .\n", key, value)
	//}
}

func testBasic() {
	var vn1, vn2, vn3 = 1, 2, 4
	println(vn1, vn2, vn3)
	vn4, vn5, vn6 := 4, 5, 6
	println(vn4, vn5, vn6)
	_, val := 34, 35
	println(val)
	//variable
	var a int = 100
	println(a)
	//const
	const Pi = 3.1415926
	println("Pi ==", Pi)
	const MaxThread = 16
	println("MaxThread ==", MaxThread)
	const prefix = "gisonwin_"
	println("const prefix ==", prefix)
	//embed basic types
	//Boolean,default is false
	var isAlived bool                   //global variable declare
	var enabled, disabled = true, false //ignore type declare
	println("isAlived is ", isAlived)
	println("enabled ==", enabled, "disabled==", disabled)
	//println(disabled)
	valid := false //short declare
	println(valid)
	var avar bool //common declare
	avar = true   // assignment operate,true
	println(avar)
	/*	number type
		unsigned 无符号 ,同时支持int,unit.这两种类型的长度相同 ,但具体长度取决于不同编译器的实现
		signed 有符号  Go里有直接定义好位数的类型 rune,int8,int16,int32,int64,byte,unit8,unit16,unit32,uint64
		其中rune是int32别称,byte是unit8的别称
		注意这些类型的变量之间不允许互相赋值或操作,不然会在编译时引起编译器报错
		尽管int长度是32bit,但int,int32并不可以互用.
		浮点数类型有float32,float64,default is float64
		Go还支持复数,default complex128(64位实数+64位虚数) complext64(32实数+32虚数).RE+IM i.RE是实数,IM是虚数,i是虚数单位.*/
	var aa int8
	var b int32
	c := aa + 3
	d := b + 8
	println(c)
	println(d)
	var cc complex64 = 5 + 5i
	fmt.Printf("Value is :%v", cc)
	//string types
	/***
	  Go中字符串都是采用UTF-8字符集编码.字符串是一对双引号或反引号括起来.它的类型是string
	  Go中字符串是不可变的.如果非要改变字符串,需要先将字符串改为数组类型,修改完数组再转回为string类型即可
	  Go中可以使用+操作符连接两个字符串
	  ``声明多行字符串
	*/
	var hello string //声明变量为字符串的一般方法
	println(hello)
	var emptyString string = "" //声明了一个字符串变量,初始化为空字符串
	println(emptyString)
	no, yes, maybe := "no", "yes", "maybe" //简短声明多个变量
	hello = "Bonjoure"                     //变量赋值
	println(hello)
	println(no)
	println(yes)
	println(maybe)
	cnt := `
            hello 
            world
            `
	println(cnt)

	// error types
	/****
	  Go中内置一个error类型,专门用来处理错误信息,Go的package中有一个包errors来处理错误:
	*/
	err := errors.New("emit macho dwarf : elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
	//some skills
	//group declare,在go中同时声明多个常量 ,变量,或者导入多个包时,可采用分组语言声明
	/****
	 import "fmt"
	import "os"
	import "time"
	may be written below:
	import (
		"fmt"
	    "os"
	    "time"
	)

	const i =100
	const pi=3.1415
	const pre = "Go_"
	may be written below :
	const (
		i =100
		pi= 3.1415
		pre="Go_"
	)

	var i int
	var pi float32
	var prefix string

	var (
		i int
	    pi float32
	    prefix string
	)
	*/
	//skills 2   iota enum
	/***
	 const (
	     x = iota //x == 0
	     y= iota  // y == 1
	     z= iota // z == 2
	     w //常量声明省略时默认和之前的字面相同即 w ==4
	)
	const v = iota //每遇到一个const关键字,iota就会重置,此时v == 0
	*/
	const (
		x = iota //x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        //常量声明省略时默认和之前的字面相同即 w ==3
	)
	const v = iota //每遇到一个const关键字,iota就会重置,此时v == 0
	println(x, y, z, w, v)
	//println(y)
	//println(z)
	//println(w)
	//println(v)
	/***
	Go程序设计的一些规则:
	 - 大写字母开头的变量是可导出的,就是其他包可以读取的,是公用变量
	- 小写字母开头的不可导出的,是私有变量
	- 大写字母开头的函数也是一样,相当于class中带public的公有函数
	- 小写字母开头的函数就是私有函数,private开头
	*/
	//array,slice,map

	var arr [10]int //10表示数组长度,int是存储元素的类型,对数组的操作是通过[]来进行读取或赋值
	fmt.Println(arr)
	arr[0] = 33
	arr[1] = 12
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])
	//长度是数组类型的一部分,因此[3]int,[4]int是不同的类型,数组也不能改变长度.数组之间的赋值是值的
	//赋值,当把一个数组作为参数传入函数时,传入的其实是该数组的副本,而不是它的指针.如果要使用指针就要用到后面介绍的slice
	ar := [3]int{1, 2, 3}   //声明了一个长度为3的int数组
	br := [10]int{1, 2, 3}  //声明一个长度为10的int数组,其中前三个元素初始值为1,2,3,其他默认为0
	cr := [...]int{4, 5, 6} //可以省略长度而采用...方式,Go会自动根据元素个数来计算 长度
	fmt.Println(ar)
	fmt.Println(br)
	fmt.Println(cr)
	//二维数组
	doubleArr := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArr)
	easyArr := [2][4]int{{1, 2, 3, 4}, {5, 7, 8, 9}}
	fmt.Println(easyArr)
	//slice.动态数组 ,它是一个引用类型.slice总是指向一个底层array.slice声明和array,但是不需要长度
	var fslice [] int //
	fmt.Println(fslice)
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Println("slice is ", slice)
	//slice可以从一个数组或一个已经存在的slice中再次声明,通过array[i:j]来获取,其中i是数组的开始位置,j是结束位置,但不包含array[j],它的长度是j-i
	//声明一个含有10个元素类型的byte数组
	byteArr := [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	fmt.Println(byteArr)
	var aaa, bbb [] byte
	aaa = byteArr[2:5]
	fmt.Println("aaa is ", aaa)
	bbb = byteArr[3:5]
	fmt.Println("bbb is ", bbb)

	/***
	  slice有一些简便操作
	  - slice默认开始位置是0,ar[:n] == ar[0:n]
	  - slice的第二个序列默认是数组的长度,ar[n:]==ar[n:len(ar)]
	  - 如果从一个数组里直接获取slice,可以这样ar[:],因为默认每个序列是0,第二个是数组长度,等价于arr[0:len(ar0}
	*/
	//从概念上来说slice像一个结构体,包含三个元素:一个指针,指向数组中slice指定的开始位置,长度,即slice的长度,最大长度,就是slice开始位置
	//至数组最后位置的长度.
	/***
	  slice 几个常用函数
	  len:获取长度
	  cap:获取最大容量
	  append:向slice追加一个或多个元素,然后一个和slice一样类型的slice
	  copy:copy从源slice的src中复制元素到目标dist,并且返回复制的元素个数
	*/
	//map ,format map[keyType]valueType.
	//声明一个key是字符串,值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var number map[string]int
	println("number is a map,and ==>", number)
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	println("map numbers is ==>", numbers)
	fmt.Println("the third number is ", numbers["three"])
	delete(numbers, "three")
	for key, value := range numbers {
		fmt.Printf("numbers map key: %s,value %d\n", key, value)
	}
	/***
	  map是无序的,每次打印出来 的map都会不一样,它不能通过index获取,必须通过key获取
	  map的长度是不固定的,和slice一样,也是一种引用类型
	  内置的len函数同样适用于map,返回map拥有的key数量
	  map的值可以很方便的修改.通过nubers["one"]=11很容易把key为one的值改为11
	  map的初始化可以通过key:val方式进行初始化,同时map内置有判断是否存在key的方式
	  通过delete删除map元素.
	*/
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	println("map rating ==>", rating)
	for k, v := range rating {
		fmt.Println(k, v, "   ==>")
	}
	key := "Go"
	keyRating, ok := rating[key]
	if ok {
		fmt.Printf("%s is in the map and its rating is %f\n\r", key, keyRating)
	} else {
		fmt.Printf("We have no rating associated with %s in the map\n\r", key)
	}
	delete(rating, "C")
	println("after delete , map rating ==>", rating)
	for k, v := range rating {
		fmt.Println(k, v, " <==")
	}
	//make ,new
	/***
	  make 用于内建类型(map,slice,channel)的内存分配,new用于各种类型的内存分配
	  new(T)分配了零值填充的T类型的内存空间,并且返回其地址,即一个*T类型的值.在Go中说的就是返回了一个指针,指向新分配的类型T的零值.
	  make只能创建slice,map,channel并且返回一个有初始值(非零)的T类型,而不是*T.本质来讲,导致这三个类型有所不同的原因是指向数据结构的引用在使用
	  前必须补初始化.make返回初始化后的(非零)值.
	*/

}

func PointOperation() {
	var age int = 80
	println(age)
	println(&age)
	var p *int = &age //指针p指向变量age的内存地址,如果修改*p的值,则变量age也会被修改
	println(*p)
	println(&*p)
	*p = 70
	println(&*p)
	println(age)

}

func LogicFunc() {
	//flow control
	//if :如果满足条件就做某事,否则做另一件事.Go中if条件判断不需要括号
	if 10 > 3 {
		fmt.Println("if condition is processed now")
	}
	//条件判断语句里面允许声明一个变量,而该变量作用域只能在该条件逻辑块内,其他地方就失效
	if x := computedValue(); x > 10 {
		fmt.Println("computedValue() return more than 10")
	}
	//fmt.Println(x)  x只在if 表达式范围内有效,此处已失效调用不到了
	if t := 1; t == 3 {
		fmt.Println("t==3,t is ", t)
	} else if (t < 3) {
		fmt.Println(" t<3 ,t is ", t)
	} else {
		fmt.Println("else t is ", t)
	}
	//goto,用goto跳转到必须当前函数内定义的标签.标签名是大小写敏感的.
	/*	func myFunc(){
		i:=0
		Here:
			println(i)
		    i++
		    goto Here
	}*/
	//for 用来循环读取数据,又可以当作while来控制逻辑,还能迭代操作.
	/*usage:
	     for expression1;expression2;expression3{
	     	//...
	     }
	  expression1,expression3是变量声明或函数调用返回值之类的.expression2用来条件判断.expression1在循环开始之前调用,
	  expression3在每轮循环结束之时调用.
	  break 跳出当前循环 .可以配合标签使用,跳转至标签所指定的位置.
	  continue 跳过本次循环
	*/
	sum := 1
	for sum < 1000 { //like while
		sum += sum
	}
	println(sum)
	//for 配合range 可用于读取slice,map的数据
	amap := map[string]string{"one": "1", "tow": "2", "three": "3", "four": "4", "five": "5"}
	for k, v := range amap {
		fmt.Printf("amap 's key is %s,value is %s\n", k, v)
	}
	for _, v := range amap { //将key丢掉
		fmt.Printf("amap 's value is %s\n", v)
	}
	//switch  用来代替多个if-else.
	/***
	switch expr {
		case expr1:
	          some instructions
	   case expr2:
	         some other instructions
	  case expr3:
	        some another instructions
	 default:
	       default instructions
	}
	 expr,expr1,expr2,expr3类型必须一致.
	每个case最后都带有break.匹配磖 不会自动向下执行其他case而是跳出整个switch.如果强制执行后续代码使用fallthrough.
	*/
	//函数.func来声明,多个参数后面带有类型,通过,分隔.函数可以返回多个值.
	//
	/**
	func funcName(input1 type1,input2 type2) (output type1,output2 type2) {
		//process logic code
	   //return multi value
		return value1,value2
	}
	*/
	//变参.接受参数的函数是有着不定数量的参数的.
	/****
	  func myfunc(arg ...int){}
	  arg ...int告诉Go这个函数接受不定数量的参数.这些参数全部是int.在函数体中,arg是一个int的slice.
	*/

	//传值与传指针
	//当们传一个参数值到被调用的函数里面时,实际上是传了这个值的一份copy.当在被调用函数中修改参数值的时候,
	//调用函数中相应的实参不会发生任何变化,因为数值变化只作用于copy上.
	te := 10
	fmt.Println(te)
	te1 := add(te)
	fmt.Println(te1)
	fmt.Println(te)
	//上面就是传值的例子,如果我们要改变这个值,需要使用指针.变量在内存中是存放于一个地址上的,修改变量实际上是修改变量地址处的内存.
	//只有add函数知道变量所在地址,才能修改变量的值.我们需要将a所所在地址&a传入函数,并将函数的参数类型由int 改为*int.即改为指针类型,
	//才能在函数中修改a变量的值.此时参数仍然是按copy传递,但传递copy的是一个指针.
	fmt.Println("======&a *point=======")
	te2 := add1(&te)
	fmt.Println(te2)
	fmt.Println(te)
	//传指针的好处是
	//- 使得多个函数能操作同一个对象
	//- 传指针比较轻量级(8 bytes),只是传内存地址,我们可以用指针传递体积较大的结构体.如果用参数传递的话,在每次copy
	//时就会花费较多的内存和时间.所以当我们传递一个较大的结构体时,用指针是明智的选择.
	//-Go中的string,slice,map这三种类型实现机制类似指针,可以直接传递,而不用取地址后传递指针(注意,如果函数要改变slice的长度,仍需要取地址传指针)

	//defer .Go特有的延迟语句,我们可以在函数中添加多个defer语句.当函数执行到最后时,这些defer语句会按照逆序执行,最后该函数返回.
	/***
	  特别是当你在进行一些打开资源的操作时,遇到错误需要提前返回,在返回前你需要关闭相应的资源,
	  不然很容易造成资源泄露等.
	  func ReadWrite() bool {
	  	file.Open("file")
	      if failureX{
	  		file.Close()
	          return false
	      }
	     if failureY{
	  	  file.Close()
	  	  return false
	     }
	      file.Close()
	      return true

	  }
	  我们使用Go中的defer修改后如下
	  func ReadWrite() bool {
	  	file.Open("file")
	  	defer file.Close()
	  	if failureX {
	      	return false
	     }
	  	if failureY {
	        	return false
	       }
	  	return true
	  }
	  如果有很多调用defer,则defer是采用后进先出的模式
	  for i:=0;i<5;i++ {
	  	defer fmt.Printf("%d ",i)
	  }
	  上述代码会输出4 3 2 1 0

	*/
	//函数也是一种变量,我们可以通过type来定义它,它的类型就是所有拥有相同的参数,相同的返回值的一种类型
	/***
	  type typeName func(input1 inputType1[,input2 inputType2] ) (result resultType)
	*/
	//我们可以把这个类型的函数当做值来传递.

	/***
	  package main
	  import "fmt"

	  type testInt func(int) bool
	  func isOdd(integer int ) bool {
	  	if integer %2 ==0 {
	  		return false
	     }
	  	return true
	  }
	  func isEven(integer int ) bool {
	  	if integer %2 ==0 {
	  		return true
	  	}
	  	return false
	  }
	  //声明函数类型在这个地方当做了一个参数
	  func filter(slice []int,f testInt) []int {
	  	var result []int
	  	for _,value := range slice{
	  		if f(value) {
	  			result = append(result,value)
	  		}
	  	}
	  }
	  func main(){
	  	slice := []int{1,2,3,4,5,6,7}
	  	fmt.Println("slice = ",slice)
	  	odd:= filter(slice,isOdd)//函数当作值来传递了
	  	fmt.Println("Odd element of slice are : ",odd)
	  	even := filter(slice,isEven)
	  	fmt.Println("Even elements of slice are : ",even)
	  }
	*/

	//Go没有像java那样的异常机制,不能抛出异常,而是使用了panic,recover机制.
	//panic
	//是一个内建函数,可以中断原有的控制流程.进入一个恐慌的流程国.当函数
	//F调用panic,函数F的执行被中断,但是F中的延迟函数会正常执行,然后F返回到调用它的地方.
	//在调用的地方,F的行为就像调用了panic.这一过程继续向上,直到发生panic的goroute中所有
	//调用的函数返回,此时程序退出.恐慌可以直接调用panic产生.也可以由支行时错误产生,
	//例如访问越界的数组

	//Recover
	//内建函数,可以让进入令人恐慌的流程中的goroute恢复过来.recover仅在延迟函数中有效.在正常执行过程中,调用recover会返回 nil
	//并且没有其他任何效果.如果当前goroutine陷入恐慌,调用recover可以捕获到panic的输入值,并恢复正常的执行.

	//main函数 init函数
	/****
	  Go里两个保留函数init(应用于所有函数),main(只能应用于package main).
	  这两个函数在定义时不能有任何的参数和返回值.虽然一个package里面可以写任意多个init函数,但这无论是对可读性
	  还是以后可维护性来说,我们都强烈建议用户在一个package中每个文件只写一个init函数.
	  Go程序会自动调用init,main函数,每个package中init函数都是可选的,但package main就必须包含一个main函数.

	  程序的初始化和执行都起始于main包.如果main包还导入了其他的包,则会在编译时将它们依次导入.有时一个包会被多个包同时导入
	  ,则它只会被导入一次.当一个包被导入时,如果该包还导入了其他包,则会先将其它包导入进来,然后再对这些包中的包级常量 和变量进行
	  初始化,接着执行init函数(如果有的话),依次类推,等所有被导入的包都加载完毕了就会开始对main包中的包级常量 和变量进行初始化,
	  然后执行main包中的init函数(如果有的话),最后执行main函数.
	*/
	//import
	//相对路径 import "./model"
	//绝对路径 import "shorturl/model"
	/***
	  一些骚操作
	  点操作
	  import (
	  	. "fmt" //这个包导入之后在你调用这个包的函数时,你可以省略前缀的包名,也就是前面
	              //你调用的fmt.Println("hello world"),可以省略写成Println("hello world")
	  )
	  别名操作
	  	import (
	  	f "fmt"//调用包函数时用我们定义的别名来操作即可.f.Println("hello world")
	  )
	  _操作
	  	import(
	  		"database/sql"
	  		_ "github.com/ziutek/mymysql/godrv" //_其实是引入该包,而不直接使用包里面的函数,而是调用了该包里里面的init函数.
	  )
	*/

}
func add(a int) int {
	a = a + 1
	return a
}
func add1(a *int) int {
	*a = *a + 1
	return *a
}
func computedValue() int {
	return 20
}
func testStruct() {
	//struct
	/***
	  Go中我们可以声明新的类型,作为其他类型的属性或字段的容器.比如创建一个自定义类型的person代表一个人的实体.
	  该实体拥有姓名和年龄.这样的类型我们称之为struct.

	*/
	//type person struct {
	//	name string
	//	age  int
	//}
	//var p person //P现在就是person类型的变量
	//p.name = "Gison win"
	//p.age = 33
	//fmt.Printf("The Person 's name is %s,age is %d\n", p.name, p.age)

	//P := person{"Tom", 20}
	//fmt.Println(P)

	//p1 := person{age: 19, name: "White"}
	//fmt.Println(p1)

	/***
	struct不仅能够将struct作为匿名字段 ,自定义类型,内置类型都可以作为匿名字段,而且可以在相应的字段上面进行函数
	操作.如果两个struct里都有同个字段如何破
	Go中使用最外层优先访问解决这个问题.
	这样就允许我们去重载通过匿名字段继承一些字段.当然如果我们想访问重载后对应匿名类型里面的字段,
	可以通过匿名字段来访问.
	type Human struct {
		name string
	 	age int
		phone string //Human holds on columns phone
	}
	type Employee struct {
		Human
		speciality string
		phone string //employee holds on columns phone
	}
	func main(){
		Bob := Employee{Human{"Bob",34,"444-333-222"},"Designer","111-444-333"}
		fmt.Println("Bob 's work phone is :",Bob.phone)
		fmt.Println("Bob 's personal phone is :",Bob.Human.phone)
	}
	*/

	/***
	  method的继承.如果匿名字段实现了一个method,则包含这个匿名字段的struct也能调用该method.
	  method的重写:如果子类要实现自己的方法,在子类定义一个method,重写了匿名字段的方法就是重写.

	*/
}

//func (p person) String() string{
//	return p.name+","+strconv.Itoa(p.age)
//}
func LearnInterface() {
	//Go语文里设计最精妙的应该就是interface.它让面向对象,内容组织实现非常 方便,
	/***
	什么是interface?
	就是一组method的组合,我们通过interface来定义对象的一组行为.
	比如Student实现了三个方法 SayHi,Sing,borrowMoney
		Employee实现了sayHi,Sing,SpendSalary
	上面这些方法的组合称为interface(被对象Student,Employee实现)
	interface类型:
		interface类型定义了一组方法,如果某个对象实现了某个接口的所有方法,则此对象就实现了该接口.
	*/

}

//Human对象实现sayhi 方法
func (h *Human) SayHi() {
	fmt.Printf("Hi,I am %s you can call me on %s\n", h.name, h.phone)
}

//Human 对象实现sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Printf("La la,la la la,lala....%s\n", lyrics)
}

//Human 对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle guzzle guzzle...", beerStein)
}

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human
	school string
	loan   float64
}
type Employee struct {
	Human
	company string
	money   float32
}

//Employee 重写sayhi
func (e *Employee) SayHi() {
	fmt.Printf("Hi,I am %s,I work at %s.Call me %s\n", e.name, e.company, e.phone)
}
func (s *Student) borrowMoney() (amount float64) {
	s.loan += amount
	return s.loan
}

type Men interface {
	//人类
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

//type YoungChap interface{//年轻小伙子
//	SayHi()
//	Sing(song string)
//	BorrowMoney(amount float32)
//}
//type ElderlyGen interface { //大年龄绅士
//	SayHi()
//	Sing(song string)
//	SpendSalary(amount float32)
//}

//Human实现了Men的接口.
func InheritAndOverride() {

	student := Student{Human{"Gison", 25, "111-333-555"}, "MIT", 123.5}
	employee := Employee{Human{"Win", 35, "999-333-888"}, "Google Inc", 12345.6}
	var men Men
	men = &student
	fmt.Println("This is Gison,a Student:")
	men.SayHi()
	men.Sing("Freezing rain")

	men = &employee
	fmt.Println(" This is Win,a Employee:")
	men.SayHi()
	men.Sing("Born to wild")

	x := make([]Men, 2)
	x[0], x[1] = &student, &employee
	for _, value := range x {
		value.SayHi()
	}
}

//所有都实现了一个空接口 interface{}
/**
传入两个int参数,返回这两个参数的最大值.
参数的值传递:指调用函数时将实际参数复制一份传递到函数中,这样在函数中如果对参数的修改将不会影响到实际的参数
参数的引用传递:指调用函数时将实际参数的地址传递到函数当中,则在函数中对参数的修改将会影响到实际参数.
Java中基础类型默认是使用值传递,引用类型会用的是地址传递
Go语言中默认都是值传递,如果需要地址传递,需要用*显式的声明.
*/
func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	}
	return num2
}

/**
Go 闭包,支持匿名函数,可以作为闭包,匿名函数是一个内联语句或表达式,它的特点是直接使用函数内的变量而不需要声明.
*/
var i = 0;

func getSequence() func() int {
	//i := 0
	return func() int {
		i += 1
		fmt.Printf("now is : %d\n", i)
		return i
	}
}
