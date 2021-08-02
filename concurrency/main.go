package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var ph = []string{"Vi", "Vien", "Quy", "Lai", "Huan"}

const hunger = 3
const think = time.Second / 100
const eat = time.Second / 100

var fmt = log.New(os.Stdout, "", 0)

var dinner sync.WaitGroup

func dinnerProblem(ph string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println(ph, " Seated")
	h := fnv.New64a()
	h.Write([]byte(ph))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	for h := hunger; h > 0; h-- {
		fmt.Println(ph, "Hungry")
		dominantHand.Lock() // pick up forks
		otherHand.Lock()
		fmt.Println(ph, "Eating")
		rSleep(eat)
		dominantHand.Unlock() // put down forks
		otherHand.Unlock()
		fmt.Println(ph, "Thinking")
		rSleep(think)
	}
	fmt.Println(ph, "Satisfied")
	dinner.Done()
	fmt.Println(ph, "Left the table")
}

func main() {
	fmt.Println("Table empty")
	dinner.Add(5)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go dinnerProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go dinnerProblem(ph[0], fork0, forkLeft)
	dinner.Wait() // wait for philosphers to finish
	fmt.Println("Table empty")
}
