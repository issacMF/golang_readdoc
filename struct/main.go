package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
	id   int
}

func (s Student) myName() string {
	return s.name
}

func main() {
	var stu1 = Student{name: "Tran Chi Vi", id: 1}

	stu2 := &Student{}
	stu2.name = "Ho Thi Nhu Vien"
	stu2.id = 2

	fmt.Println("Student 1: ", stu1.name, "\tid: ", stu1.id)
	fmt.Println("Student 2: ", stu2.name, "\tid: ", stu2.id)

	fmt.Println("Date Type studennt1: ", reflect.TypeOf(stu1))
	fmt.Println("Kind of studennt1: ", reflect.ValueOf(stu1).Kind())
	fmt.Println("Date Type studennt2: ", reflect.TypeOf(stu2))
	fmt.Println("Kind of studennt2: ", reflect.ValueOf(stu2).Kind())

	fmt.Println(stu1.myName())

}
