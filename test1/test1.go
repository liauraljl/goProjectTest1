package test1

import (
	"fmt"
	"reflect"
	"testing"
)

func init() {
	fmt.Println("test1 init................")
}

func TestName(t *testing.T) {
	b := 1.0
	var c = false

	fmt.Println("c的类型" + reflect.TypeOf(c).String())
	fmt.Println("b的类型" + reflect.TypeOf(b).String())
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("test1 testName")
	testName4()
	//testName5()

}

func testName2() string {
	switch 3 {
	case 1:
		fmt.Println("")
	case 2:
		fmt.Println("")
	case 3:
		fmt.Println("")
	default:
		fmt.Println("")
	}
	return "nil"
}

func testName3() {
	a := []string{"a", "b", "c"}
	for key, value := range a {
		fmt.Println(key)
		fmt.Println(value)
		fmt.Print("")
	}

	for _, value := range a {
		fmt.Println(value)
		fmt.Print("")
	}

	for i := 3; i < 10; i++ {
		fmt.Println("")
	}
}

func testName4() {
	fmt.Println("start")
	goto a
	fmt.Println("before")
a:
	fmt.Println("middle")
	fmt.Println("after")
}

func testName5() {
	fmt.Println("start")
	fmt.Println("before")
a:
	fmt.Println("middle")
	goto a
	fmt.Println("after")
}
