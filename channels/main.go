package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var goRoutine sync.WaitGroup

func employee(projects chan string, employee int) {
	defer goRoutine.Done()

	for {
		project, result := <-projects

		if result == false {
			fmt.Printf("Employee: %d: Exit\n", employee)
			return
		}

		fmt.Printf("Employee :%d :Started %s\n", employee, project)

		sleep := rand.Int63n(50)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Println("\nTime to sleep", sleep, "ms")

		fmt.Printf("Employee : %d : Completed %s\n", employee, project)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	projects := make(chan string, 10)

	goRoutine.Add(5)

	for i := 0; i <= 5; i++ {
		go employee(projects, i)
	}

	for j := 0; j <= 10; j++ {
		projects <- fmt.Sprintf("Project :%d", j)
	}

	close(projects)
	goRoutine.Wait()
}
