package datatest

import "fmt"

type Annimal struct {
	Color string
}

type Cat struct {
	//继承关系
	Annimal
	Id   int
	Name string
	Age  int
}

type Dog struct {
	//继承关系
	Annimal
	Id   int
	Name string
	Age  int
}

/**
结构体方法
*/
func (cat *Cat) Run() string {
	fmt.Println("cat id:", cat.Id, "cat is run")
	return "cat is run"
}

func (dog *Dog) Run() string {
	fmt.Println("dog id:", dog.Id, "dog is run")
	return "dog is run"
}

func StructStart() {
	testStruct1()
}

func testStruct1() {
	var cat Cat
	cat.Id = 1
	cat.Name = "cat1"
	cat.Age = 4
	cat.Color = "black"
	fmt.Println(cat)

	cat2 := Cat{Id: 2, Name: "cat2", Age: 5}
	fmt.Println(cat2)

	//new 返回指针
	cat3 := new(Cat)
	cat3.Id = 3
	cat3.Name = "cat3"
	cat3.Age = 6
	fmt.Println(cat3)

	cat3.Run()

	//接口测试
	var interfaceTest Behavior
	interfaceTest = new(Dog)
	interfaceTest.Run()
}
