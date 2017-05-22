package main

import (
	"fmt"
	"time"
)

type CoffeeMachine struct {
	Coin chan ComputerScientistWannaCoffee
}

func (self CoffeeMachine) Run()  {
	for {
		attendant := <-self.Coin
		attendant.Coffee<-true
		time.Sleep(time.Duration(500));
	}
}



type ComputerScientistWannaCoffee struct {
	Coffee chan bool
}

func (self ComputerScientistWannaCoffee) Run(CM CoffeeMachine)  {
	for {
		CM.Coin <- self
		myCoffee := <-self.Coffee
		fmt.Println("uuuuh a coffee ", myCoffee)
		time.Sleep(time.Duration(500));
	}
}



/////////////////////////////////////////



type CoffeeTeaMachine struct {
	Coin chan ComputerScientist
}

func (self CoffeeTeaMachine) Run() {
	for {
		CS := <-self.Coin
		//select {
		//case CS.Tea<-true:
		//	fmt.Println("your tea sir")
		//case CS.Coffee<-true:
		//	fmt.Println("coffe here")
		//default:
		//	fmt.Println("boh vecchio")
		//}
		CS.Coffee<-true
		//CS.Tea<-true
	}
}


type ComputerScientist struct {
	Coffee chan bool
	Tea chan bool
}

func (self ComputerScientist) Run(CTM CoffeeTeaMachine) {
	for {
		CTM.Coin<-self
		select {
		case <- self.Tea:
			fmt.Println("shitty tea, where's my COFFEINE??????")
		case <- self.Coffee:
			fmt.Println("nice sweet coffee, ready to do some Work work new Work")
		default:
			fmt.Println("what happend?")
		}
	}
}


func main()  {
	//CM := CoffeeMachine{
	//	make(chan ComputerScientistWannaCoffee)}
	//CSWC := ComputerScientistWannaCoffee{
	//	make(chan bool) }
	//
	//go CM.Run()
	//
	//CSWC.Run(CM)

	CTM := CoffeeTeaMachine{
		make(chan ComputerScientist) }
	CS := ComputerScientist{
		make(chan bool),
		make(chan bool) }

	go CTM.Run()

	CS.Run(CTM)

}
