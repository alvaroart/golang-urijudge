package main

import ("fmt")

func main(){
	var raio float64
	var a  float64
	var pi = 3.14159
	fmt.Scanf("%f", &raio)
	a = (raio * raio) * pi
	fmt.Printf("A=%.4f\n", a)
}
