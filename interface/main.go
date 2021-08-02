package main

import (
    "fmt"
)

type Employee struct {
    name string
    daysWork int
}
type EmpSalary interface {
    CalculateTotalSalary() float64
}

func (e Employee) CalculateTotalSalary() float64{
    total := e.daysWork*1000
    return float64(total)
}

func main() {
    var emp1 EmpSalary
    emp1 = Employee{"Tran Chi Vi", 30}
    fmt.Println("Total salary: ", emp1.CalculateTotalSalary())
}