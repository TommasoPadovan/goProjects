package main

import "fmt"
import "math"

func main() {
	//funzioneMaggica := func(x int) int {
	//	return x+1
	//}


	//fmt.Println(execute(funzioneMaggica, 1))



	//fmt.Println(fibonacciGrezzo(20));

	v := Vertex{3,4}
	fmt.Println(v.Abs())


	var v2 Vertex = new(Vertex)
	v2.X, v2.Y = 15, 18

	v3 := New()

}




type MyFunc func(int) int

func execute(fun MyFunc, operand int) int {
	return fun(operand);
}

func f() {
	for i:=0; i<10; i++ {
		defer fmt.Println(i)
	}
}


func fibonacci(i int) int {
	if i==0 {return 0}
	if i==1 {return 1}
	return fibonacci(i-1) + fibonacci(i-2)
}

func fibonacciGrezzo(i int) int  {
	defer fmt.Println("Dio lurido, ho finito, porcamadonna")
	fmt.Println("ora mi metto tiocca'")
	return fibonacci(i);
}








type Vertex struct {
	X, Y float64;
}

func (self *Vertex) Abs() float64  {
	return math.Sqrt(self.X*self.X+self.Y*self.Y)
}

func New(x, y float64) *Vertex {
	return &Vertex{x, y}
}







