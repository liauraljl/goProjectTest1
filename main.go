package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProjectTest1/json_demo"
	"goProjectTest1/pointtest"
	"goProjectTest1/sugar_demo"
)

func main() {
	//log.Print("*******************************")
	//test12.TestName(nil)
	//test22.TestName(nil)
	//fmt.Println("test1 start!!！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！!")
	//test1()
	//fmt.Println("test1 end!!!！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！")
	//fmt.Println("test2 start!!!！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！")
	//test2()
	//fmt.Println("test2 end!!!！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！！")

	//datatest.DataStart()

	//datatest.ErrorStart()

	//datatest.StructStart()

	//gorotine.GorotineStart()

	pointtest.PointStart()

	json_demo.JsonStart()

	sugar_demo.SugarStart()
}

func test1() {
	//布尔类型
	var a bool = true
	if a {
		fmt.Println("it is true")
	}

	//字符类型
	var b string = "Runoob"
	fmt.Println(b)

	//根据值自动判断变量类型
	var c = "Runoob b"
	var d = true
	//var e int = 2
	e := 2
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("Hello, World!")

}

func test2() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
