package main

import ("fmt")

func main() {
	var a float32
	var b float32
	var media float32
	
	fmt.Scanf("%f", &a)
	fmt.Scanf("%f", &b)

	media = (a * 3.5 + b * 7.5) / 11

	fmt.Printf("MEDIA = %.5f\n", media)
}
