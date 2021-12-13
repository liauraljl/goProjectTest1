package gorotine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GorotineStart() {
	//简单测试协程使用
	//testSimple()

	//测试协程通信
	go Send()
	go Receive()
	time.Sleep(time.Second * 3)

	//测试协程同步
	Read()
	go Write()
	WG.Wait()
	fmt.Println("All Done!")
	time.Sleep(time.Second * 10)
}

func testSimple() {
	fmt.Println("逻辑cpu个数num cpu", runtime.NumCPU())
	//设置go协程运行核心数
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//go 启动协程 协程比线程更高效
	go loopA()
	go loopB()
	time.Sleep(time.Second * 3)
}

func loopA() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Millisecond * 10)
		fmt.Printf("loopA:%d", i)
		fmt.Println()
	}
}

func loopB() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Millisecond * 15)
		fmt.Printf("loopB%d", i)
		fmt.Println()
	}
}

var chanInt chan int = make(chan int, 10)
var timeoutChan chan bool = make(chan bool, 1)

//协程同步
var WG sync.WaitGroup

func Send() {
	time.Sleep(time.Second * 1)

	chanInt <- 1
	time.Sleep(time.Second * 1)

	chanInt <- 2
	time.Sleep(time.Second * 1)

	chanInt <- 3
	time.Sleep(time.Second * 1)

	timeoutChan <- true
}

func Receive() {
	//num := <-chanInt
	//fmt.Println("receive num1:", num)
	//
	//num2 := <-chanInt
	//fmt.Println("receive num2:", num2)
	//
	//num3 := <-chanInt
	//fmt.Println("receive num3:", num3)

	for {
		select {
		case num := <-chanInt:
			fmt.Println("receive num:", num)
		case <-timeoutChan:
			fmt.Println("timeout.....")
		}
	}
}

/**
读取数据
*/
func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

/*
写入数据
*/
func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		WG.Done()
		fmt.Println("写入数据")
	}
}
