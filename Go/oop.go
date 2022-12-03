package main

import "fmt"

// Wheel Class: Check Tire Pressure, Fill Tire Pressure
type Wheel struct{
	name string
	tirePressure float32
}

func (tp Wheel) checkPressure(){
	fmt.Println("Current Level:", tp.tirePressure)
	if tp.tirePressure < 50.0{
		fmt.Println(tp.name, ": Refill Pressure")
	} else {
		fmt.Println(tp.name, ": Pressure in is at optimum capacity")
	}
}

func (ftp Wheel) fillPressure(){
	if ftp.tirePressure < 80.0{
		ftp.tirePressure = 100.0
		fmt.Println(ftp.name, ": Tire Pressure refilled!")
	} else {
		fmt.Println("No need for refilling in", ftp.name)
	}
}


// Engine Class: Check Coolant Level, Refill Coolant Level
type Engine struct{
	coolant float32
}

func (cl Engine) checkCoolant(){
	fmt.Println("Current Level:", cl.coolant)
	if cl.coolant <= 50.0{
		fmt.Println("Refill Coolant")
	} else {
		fmt.Println("Coolant is at optimum capacity")
	}
}

func (fcl Engine) fillCoolant(){
	if fcl.coolant <= 50.0{
		fcl.coolant = 100.0
		fmt.Println("Coolant refilled!")
	} else {
		fmt.Println("No need for refilling Coolant")
	}
}


// Vehicle Class: Check Wheel Stuff, Check Engine Stuff
type Vehicle struct{
	seats int
	eng Engine
	front Wheel
	back Wheel
}

func (v *Vehicle) pothole(){
	fmt.Println("\nHit a Pothole\n---------------------")
	v.front.tirePressure -= 10.0
	v.back.tirePressure -= 10.0

	v.front.checkPressure()
	v.back.checkPressure()
}

func (v *Vehicle) nitrous(){
	fmt.Println("\nActivating Nitrous\n---------------------")
	v.eng.coolant -= 30.0
	v.eng.checkCoolant()
}

func (v *Vehicle) mechanic(){
	fmt.Println("\nStops at the Mechanic\n---------------------")
	v.front.fillPressure()
	v.back.fillPressure()
	v.eng.fillCoolant()
}

func main(){

	engine:= Engine{coolant: 100.0}
	front:= Wheel{name: "Front Wheel", tirePressure: 100.0}
	back:= Wheel{name: "Back Wheel", tirePressure: 100.0}

	motorcycle := Vehicle{
		seats: 1, 
		eng: engine,
		front: front,
		back: back,
	}

	motorcycle.nitrous()
	motorcycle.pothole()
	motorcycle.pothole()
	motorcycle.nitrous()
	motorcycle.mechanic()
	

}