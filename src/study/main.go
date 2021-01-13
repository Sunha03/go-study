package main

import (
	"fmt"

	"./gorm"
	"./part"
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

	fmt.Println("\n== maps ==")
	part.Maps()

	fmt.Println("\n\n********** part4 **********")
	fmt.Println("== packages ==")
	part.Packages()

	fmt.Println("\n== structs ==")
	part.Structs()

	fmt.Println("\n== go methods ==")
	part.GoMethods()

	fmt.Println("\n== go interfaces ==")
	part.GoInterfaces()

	fmt.Println("\n\n********** part5 **********")
	fmt.Println("== defer() ==")
	part.DeferFunc()

	fmt.Println("\n== panic() ==")
	//part.PanicFunc()

	fmt.Println("\n== recover() ==")
	part.RecoverFunc()

	fmt.Println("\n== goroutines() ==")
	part.Goroutines()

	fmt.Println("\n== anonymous goroutine() ==")
	part.AnonymousGoroutineFunc()

	fmt.Println("\n== go channels ==")
	part.GoChannels()

	fmt.Println("\n\n********** GORM **********")
	fmt.Println("== connect MySQL() ==")
	//gorm.ConnMySQL()

	fmt.Println("\n== create DB ==")
	//gorm.CreateDB()

	fmt.Println("\n== CRUD ==")
	//gorm.CRUD()

	fmt.Println("\n== select ==")
	//gorm.Select()

	fmt.Println("\n== insert ==")
	gorm.Insert()
}
