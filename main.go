package main

import (
	"fmt"
	"time"
)

func main()  {
	d := NewDag()
	d.PipeLine(f1).OnComplete(f3).Then().Spawn(f2,f1).OnComplete(f4)
	d.Run()
}

func f1()  {
	time.Sleep(1*time.Second)
	fmt.Println("f1")
}

func f2()  {
	fmt.Println("f2")
}
func f3()  {
	fmt.Println("f3")
}
func f4()  {
	fmt.Println("f4")
}