package service

import (
	"fmt"

	basic "go-study/study/basic"

	"go-study/study/gorm"
)

func ExeGobasic() {
	fmt.Println("********** basic1 **********")
	fmt.Println("== variables ==")
	basic.Variables()

	fmt.Println("\n== consts ==")
	basic.Consts()

	fmt.Println("\n== pointers ==")
	basic.Pointers()

	fmt.Println("\n== dataTypes ==")
	basic.DataTypes()

	fmt.Println("\n== operators ==")
	basic.Operators()

	fmt.Println("\n== conditional statements ==")
	basic.ConditionalStatements()

	fmt.Println("\n== loop statements ==")
	basic.LoopStatements()

	fmt.Println("\n\n********** basic2 **********")
	fmt.Println("== functions ==")
	basic.Functions()

	fmt.Println("\n== functions2 ==")
	basic.Functions2()

	fmt.Println("\n\n********** basic3 **********")
	fmt.Println("== arrays ==")
	basic.Arrays()

	fmt.Println("\n== slices ==")
	basic.Slices()

	fmt.Println("\n== maps ==")
	basic.Maps()

	fmt.Println("\n\n********** basic4 **********")
	fmt.Println("== packages ==")
	basic.Packages()

	fmt.Println("\n== structs ==")
	basic.Structs()

	fmt.Println("\n== go methods ==")
	basic.GoMethods()

	fmt.Println("\n== go interfaces ==")
	basic.GoInterfaces()

	fmt.Println("\n\n********** basic5 **********")
	fmt.Println("== defer() ==")
	basic.DeferFunc()

	fmt.Println("\n== panic() ==")
	//basic.PanicFunc()

	fmt.Println("\n== recover() ==")
	basic.RecoverFunc()

	fmt.Println("\n== goroutines() ==")
	basic.Goroutines()

	fmt.Println("\n== anonymous goroutine() ==")
	basic.AnonymousGoroutineFunc()

	fmt.Println("\n== go channels ==")
	basic.GoChannels()
}

func ExeGorm() {
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
	//gorm.Insert()

	fmt.Println("\n== prepared statement ==")
	//gorm.PreparedStatement()

	fmt.Println("\n== transaction ==")
	gorm.Transaction()
}
