package main
import "fmt"

func main() {
    strArr1 := [3]string{"JP", "CN", "VN"}
    strArr2 := strArr1
    strArr3 := &strArr1

    fmt.Println("Arr2:", strArr2);
    fmt.Println("Arr3:", *strArr3);
}