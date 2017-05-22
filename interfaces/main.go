package main

import "fmt"



/**
 * Valentina Nappi Example
 */
//import "fmt"
//
//type ButtSexable interface {
//	LikeButtSex() bool
//	DoButtSex()
//}
//
//
//type ValentinaNappi struct {
//	DiametroAno int
//}
//
//func (self ValentinaNappi) LikeButtSex() bool  {
//	return true;
//}
//
//func (self ValentinaNappi) DoButtSex() {
//	for i:=0; i<self.DiametroAno; i++ {
//		fmt.Println("aaaah, stick it in my ass babeeee")
//	}
//}
//
//func NewValentinaNappi(i int) ValentinaNappi  {
//	return ValentinaNappi{i}
//}
//
//func main() {
//	var b ButtSexable;
//
//	b = NewValentinaNappi(10);
//
//
//	fmt.Println("hey Valentina, do you like anal?")
//	fmt.Println(b.LikeButtSex());
//	fmt.Println("lets go then")
//	b.DoButtSex();
//
//}
















/**
 * Dining philosophers
 */
//import (
//	"fmt"
//	"time"
//	"math/rand"
//)
//type Fork struct {
//	name string
//	free bool
//
//	Take chan int
//	Leave chan int
//}
//
//func (self Fork) Run() {
//	for {
//		<- self.Take
//		fmt.Println("fork %s taken", self.name)
//		<- self.Leave
//		self.free = true
//		fmt.Println("fork %s released", self.name)
//	}
//}
//
//
//
//
//func Philosopher(id int, leftFork *Fork, rightFork *Fork) {
//	for {
//		fmt.Println("philosopher %d is thinking...", id);
//		time.Sleep(time.Duration(rand.Int63n(2*1e9)))
//		leftFork.Take <- 1
//		time.Sleep(time.Duration(rand.Int63n(2*1e9)))
//		rightFork.Take <- 1
//		fmt.Println("philosopher %d is eating...", id)
//		leftFork.Leave <- 1
//		rightFork.Leave <- 1
//	}
//}
//
//func main() {
//	const NPhil = 5
//
//	var forks [NPhil]Fork
//
//	for i := 0; i < NPhil; i++ {
//		forks[i] = Fork{
//			fmt.Sprintf("Fork%d", i),
//			true,
//			make(chan int),
//			make(chan int)}
//		go forks[i].Run()
//	}
//
//	for i := 0; i < NPhil - 1; i++ {
//		go Philosopher(i, &forks[i], &forks[(i+1)%NPhil])
//	}
//
//	Philosopher(NPhil-1, &forks[0], &forks[NPhil-1])
//}







func FiBROnacci() chan int {
	channel := make(chan int)
	go func() {
		prepre := 0
		pre := 1
		for {
			prepre, pre = pre, prepre + pre
			channel <- pre
		}
	}()
	return channel
}

func main() {
	broam := FiBROnacci()
	for i := 0; i < 10; i++ {
		fmt.Println(<-broam)
	}
}





























































