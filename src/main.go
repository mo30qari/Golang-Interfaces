package main

import("fmt")

func main(){

	var w IO = Monitor{}
	w.write("Outputs")
	w.read("Inputs")

}

type IO interface{

	write (string)
	read (string)

}

type Printer struct{}

func (p Printer) write(data string){

	fmt.Println("I write on paper: " + data)

}

type Monitor struct{}

func (m Monitor) write(data string){

	fmt.Println("I write on screen: " + data)

}

func (m Monitor) read(data string){

	fmt.Println("I read from inputs: " + data)

}

