package main

import (
	"fmt"
	"study/part"
)

func main() {
	fmt.Println("********** part1 **********")
	fmt.Println("== variables ==")
	part.Variables()

	fmt.Println("\n== consts ==")
	part.Consts()

	fmt.Println("\n== dataTypes ==")
	part.DataTypes()

	fmt.Println("\n== operators ==")
	part.Operators()

	fmt.Println("\n== conditional statements ==")
	part.ConditionalStatements()

	fmt.Println("\n== loop statements ==")
	part.LoopStatements()

	fmt.Println("\n\n********** part2 **********")
	fmt.Println("== functions ==")
	part.Functions()

	fmt.Println("\n== functions2 ==")
	part.Functions2()

	fmt.Println("\n\n********** part3 **********")
	fmt.Println("== arrays ==")
	part.Arrays()

	fmt.Println("\n== slices ==")
	part.Slices()
}
