package main

import (
	"fmt"
	"time"
)

type CoffeeTeaMachine struct {
	Pay    chan int
	Coffee chan int
	Tea    chan int
}

func (self CoffeeTeaMachine) Run()  {
	var choice string
	for {
		<- self.Pay
		fmt.Println("CTM: got a coin")

		select {
		case <- self.Tea:
			choice = "Tea"
		case <- self.Coffee:
			choice = "KAFFEEE"
		}

		fmt.Println("CTM: selected", choice)
	}
}

func ComputerScientist(CTM *CoffeeTeaMachine) {
	var drink string
	for {
		fmt.Println("CS: work, work, got a coin")
		CTM.Pay <- 1
		select {
		case CTM.Coffee <- 1:
			fmt.Println("CS: I want a coffee")
			drink = "KAFEEEEEE!"
		case CTM.Tea <- 1:
			fmt.Println("CS: maybe i'm gay, and i want a tea")
			drink = "Tea"
		}

		fmt.Println("CS: i got a ", drink)
		time.Sleep(3e09);
	}
}

type ComputerBoy struct {
	HasMoney bool
	Machine CoffeeTeaMachine
}

func NewComputerBoy(machine CoffeeTeaMachine) ComputerBoy {
	return ComputerBoy{
		false,
		machine}
}

func (self *ComputerBoy) WorkWork()  {
	self.HasMoney = true
	fmt.Println("CS: workwork, got da money..")
}

func (self *ComputerBoy) InsertMoney() {
	self.Machine.Pay <- 1
	self.HasMoney = false
	fmt.Println("CS: clink! money in the machine..")
}

func (self *ComputerBoy) IWant(drink int) {
	switch drink {
	case TEA:
		self.Machine.Tea <- 1
		fmt.Println("CS: mmmmh shitty tea wtf i hate it so much")
	case COFFEE:
		self.Machine.Coffee <- 1
		fmt.Println("CS: uuuuuh KAFFEEEEEEE??!!")
	}
}


func main() {
	CTM := CoffeeTeaMachine{
		make(chan int),
		make(chan int),
		make(chan int)}

	go CTM.Run()

	//ComputerScientist(&CTM)

	var input string
	CS := NewComputerBoy(CTM)

	for {
		fmt.Scanln(&input)
		switch input {
		case "work":
			CS.WorkWork()
		case "pay":
			if CS.HasMoney {
				CS.InsertMoney()
			} else {
				fmt.Println("you got no money brah, try some workwork")
			}
		case "coffee":
			CS.IWant(COFFEE)
		case "tea":
			CS.IWant(TEA)
		}

	}
}

type SwaggerMachine struct {
	Pay 	chan int
	Coffee 	chan int
	Tea 	chan int
	Drink 	chan int
	Change	chan int
	TOut	chan int

	TeaPrice int
	CoffeePrice int
}

func NewSwaggerMachine() SwaggerMachine {
	swaggerMachine := SwaggerMachine{}
	swaggerMachine.Pay = make(chan int)
	swaggerMachine.Coffee = make(chan int)
	swaggerMachine.Tea = make(chan int)
	swaggerMachine.Drink = make(chan int)
	swaggerMachine.Change = make(chan int)
	swaggerMachine.TOut = make(chan int)

	swaggerMachine.TeaPrice = 40
	swaggerMachine.CoffeePrice = 50

	return swaggerMachine
}
func (self *SwaggerMachine) Run() {
	for {
		currency := <-self.Pay
		fmt.Println("CM: got %d cents", currency)

		select {
		case <-self.Coffee:
			if currency >= self.CoffeePrice {
				currency -= self.CoffeePrice
				self.Drink <- COFFEE
				fmt.Println("CM: here's your KAFFEEE")
				self.Change <- currency
				fmt.Println("CM: that's your change: %d centz", currency)
			} else {
				self.Drink <- NOTHING
				fmt.Println("not enought money you moron")
				self.Change <- currency
				fmt.Println("CM: that's your change: %d centz", currency)
			}
		case <-self.Tea:
			if currency >= self.TeaPrice {
				currency -= self.TeaPrice
				self.Drink <- TEA
				fmt.Println("CM: your tea you fag")
				self.Change <- currency
				fmt.Println("CM: that's your change: %d centz", currency)
			} else {
				self.Drink <- NOTHING
				fmt.Println("not enought money you moron")
				self.Change <- currency
				fmt.Println("CM: that's your change: %d centz", currency)
			}
		case <-time.After(10e09):
			self.TOut <- 1
			fmt.Println("CM: wait, wait, wait.. time out bro")
			self.Change <- currency
			fmt.Println("CM: get your money back and go home bwoy")
		}
	}

}










const (
	TEA = iota
	COFFEE = iota
	NOTHING = iota
)


























