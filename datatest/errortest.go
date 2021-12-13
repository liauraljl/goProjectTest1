package datatest

import (
	"errors"
	"fmt"
)

func ErrorStart() {
	receivePanic()
}

func receivePanic() {
	fmt.Println("before panic")
	//defer会在该方法内最后执行 类似finally
	defer panicMessage()
	//发生panic时终止执行,并抛异常
	panic("i am string panic")
	panic(errors.New("i am error panic"))
	panic(1)
	fmt.Println("after panic")
}

func panicMessage() {
	//recover防止抛异常，铺捕获异常消息，但是程序仍然中断
	message := recover()
	switch message.(type) {
	case string:
		fmt.Println("string message:", message)
	case error:
		fmt.Println("error message:", message)
	default:
		fmt.Println("unknown message:", message)

	}
	fmt.Println("panic message:", message)
}
