package sugar_demo

import (
	"fmt"
	"reflect"
)

func SugarStart() {
	multiParam("a", "b", "c")
	typeSugar()
}

//多参数测试
func multiParam(values ...string) {
	for _, v := range values {
		fmt.Println("输入参数：", v)
	}
}

//类型推断 定义变量
func typeSugar() {
	c := false
	fmt.Println("变量类型:", reflect.TypeOf(c))
	fmt.Println("变量:", c)
}
