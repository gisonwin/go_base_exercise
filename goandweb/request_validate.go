package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/30 15:52
 */
var emailValidate = `^([\w\.\_]{2,10}@(\w{1,}).([a-z]{2,4}))$`

type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age       int `validate:"eq=10"`
	NestedVar Nested
}

func validateEmail(input string) bool {

	if pass, _ := regexp.MatchString(emailValidate, input); pass {
		return true
	}
	return false
}
func validate(v interface{}) (bool, string) {
	validateResult := true
	errmsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)

	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		k := fieldVal.Kind()
		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errmsg = "validate int failed,tag is" + strconv.FormatInt(tagVal, 10)
				validateResult = false
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "email":
				nestedResult := validateEmail(val)
				if !nestedResult {
					errmsg = " validate mail failed,field val is:" + val
					validateResult = false
				}
			}
		case reflect.Struct:
			// 如果有内嵌的struct,那么深度优先遍历
			//就是一个递归过程
			valInter := fieldVal.Interface()
			nestedResult, msg := validate(valInter)
			if !nestedResult {
				validateResult = false
				errmsg = msg
			}
		}
	}
	return validateResult, errmsg
}
func main() {
	t := T{Age: 10, NestedVar: Nested{Email: "abc@def.com"}}
	validateResult, errmsg := validate(t)
	fmt.Println(validateResult, errmsg)
}
/**
*  如果反射产生了性能问题,我们可以使用Go内置的Parser对源代码进行扫描,然后根据结构体的定义生成校验代码.可以将所有需要校验
的结构体放在单独的包内.
*/