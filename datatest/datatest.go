package datatest

import (
	"fmt"
	"reflect"
)

func DataStart() {
	makeSlice()
	makeMap()
	makeChan()
	newMap()
	appendSlice()
	copySlice()
	deleteMap()
}

/**
new 返回指针类型
make返回引用类型
*/
func newMap() {
	mNewMap := new(map[int]string)
	mMakeMap := make(map[int]string)
	fmt.Println("mNewMap", reflect.TypeOf(mNewMap))
	fmt.Println("mMakeMap", reflect.TypeOf(mMakeMap))
}

/**
切片
*/
func makeSlice() {
	mSlice := make([]string, 3)
	mSlice2 := make([]string, 5)
	mSlice[0] = "cat"
	mSlice[1] = "dog"
	mSlice[2] = "fish"
	fmt.Println(mSlice)
	fmt.Println(mSlice2)
}

/**
append切片,会自动增加长度
*/
func appendSlice() {
	mSlice := make([]string, 3)
	mSlice[0] = "cat"
	mSlice[1] = "dog"
	mSlice[2] = "fish"
	fmt.Println("mSlice 长度len:", len(mSlice))
	fmt.Println("mSlice 容量cap:", cap(mSlice))
	mSlice = append(mSlice, "tigger-4")
	fmt.Println("扩容后 mSlice 长度len:", len(mSlice))
	fmt.Println("扩容后 mSlice 容量cap:", cap(mSlice))
	fmt.Println(mSlice)
}

/**
copy测试
*/
func copySlice() {
	sourceSlice := make([]string, 2)
	sourceSlice[0] = "s-0"
	sourceSlice[1] = "s-1"

	targetSlice := make([]string, 3)
	targetSlice[0] = "t-0"
	targetSlice[1] = "t-1"
	targetSlice[2] = "t-2"

	copy(targetSlice, sourceSlice)
	fmt.Println(targetSlice)
}

/**
map
*/
func makeMap() {
	mMap := make(map[int]string)
	mMap[10] = "cat"
	mMap[100] = "dog"
	mMap[101] = "chilken"
	fmt.Println(mMap)
}

/**
delete test
*/
func deleteMap() {
	mMap := make(map[string]string)
	mMap["k1"] = "v1"
	mMap["k2"] = "v2"
	delete(mMap, "k1")
	fmt.Println(mMap)

}

/**
chan
*/
func makeChan() {
	mChan := make(chan string, 3)
	mChan <- "test1"
	mChan <- "test2"
	str, ok := <-mChan
	fmt.Println("str:", str)
	//ok = true
	fmt.Println("ok:", ok)
	//关闭channel后无法写数据
	close(mChan)

}
