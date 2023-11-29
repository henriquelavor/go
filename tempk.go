package main

import "fmt"

const ebulicaoK float64 = 373.15

func main() {

	tempK := ebulicaoK
	tempC := (tempK - 273.15)

	fmt.Println("A temperatura de ebulição da água  em °K = ", tempK)
	fmt.Println("A temperatura de ebulição da água  em °C = ", tempC)

}
