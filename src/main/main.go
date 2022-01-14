package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("main init...")
}

const (
	n1 = iota
	n2
	_
	n3
)

func main() {
	fmt.Println("demo ...")
	bytes := make([]byte, 100)
	f, _ := os.Open("D:/123/使用前必看.txt")
	defer f.Close()
	for {
		n, _ := f.Read(bytes)
		if n == 0 {
			break
		}
		os.Stdout.Write(bytes[:n])
	}
	fmt.Println(n1, n2, n3)
	fmt.Println("-----")
	arr()
	//panic("os is panic...")
	makeT()

}
func arr() {
	array1 := [...]int{1, 2, 5, 9, 7}
	for i, v := range array1 {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
	array2 := array1[1:3]
	for i, v := range array2 {
		fmt.Println(i, v)
	}
	fmt.Println("-----")
}
func makeT() {
	fmt.Println("-----")
	b := make(map[string]int, 10)
	b["test"] = 100
	fmt.Println(b)
	fmt.Println("-----")
}
