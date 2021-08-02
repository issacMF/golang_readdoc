package main

import "fmt"

func main(){
	var (
		name string = "Tran Chi Vi"
		age  int = 23
		height float64 = 60.2
	)
	super_var := "Another way to create Variables of Golang"
	fmt.Println("Name: ", name, " Age: ", age, " Height: ", height)

	fmt.Println("Super Var: ", super_var)
}
