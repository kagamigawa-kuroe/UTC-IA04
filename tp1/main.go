package main

import (
	"fmt"
	"mymodule/mypackage"

	// 包名和路径名保持一致
	"mymodule/package2"
)

func test(){
	sl := make([]int,10)
	mypackage.Fill(sl);
	a := mypackage.Moyenne(sl);
	fmt.Println("----------")
	fmt.Println(a)
	t := mypackage.ValeursCentrales(sl)
	fmt.Println("-----------")
	fmt.Println(t)
	mypackage.Plus1(sl)
    fmt.Println(sl)
    fmt.Println("-----------")

	// var m []int = []int{1,2,3,4,5} 
    mypackage.Compte(1,sl)
}

func test_Palindromes(){
	dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
	a := mypackage.Palindromes(dict[:])
	fmt.Println(a)
}

func test_file(){
	mypackage.DictFromFile("/Users/whz/utc/2022.automne/IA04/code/tp1/mypackage/file.txt")
}

func main() {
	package2.PrintHello2()
	test()
	// test_Palindromes()
	// test_file()
}


