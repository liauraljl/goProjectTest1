package pointtest

import "fmt"

func PointStart() {
	test1()
	test2()
}

func test1() {
	var count int = 20
	var countPoint *int = &count
	var countPoint2 *int
	fmt.Printf("count 地址：%x \n", &count)
	fmt.Printf("countPoint指向得到值：%d \n", *countPoint)
	fmt.Println("countPoint2地址：", countPoint2)
}

func test2() {
	//指针数组
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组pointArr：", pointArr)

	//数组的指针
	arr := [...]int{a, b}
	fmt.Println("数组arr的指针：", &arr)

}
