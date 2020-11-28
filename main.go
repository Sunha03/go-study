package main

import (
	"fmt"
	"test/say"
)

//Constants : 변경불가
// - 한번 선언한 변수는 재할당 불가
func Constants() {
	const name string = "SH" //변수 앞에 const 선언
	fmt.Println(name)
}

//Variables : 변경가능
func Variables() {
	var varName string = "str1"
	varName = "str2"
	fmt.Println(varName)

	//변수 선언 축약 코드
	// - 축약형 코드는 func 안에서만 가능
	// - 변수에만 사용 가능
	newVar := "newStr1"
	newVar = "newStr2"
	fmt.Println(newVar)
}

//Type
// - string, bool, numeric, array, slices, ...
func multiply(a int, b int) int { //(a int, b int) = (a, b int)
	return a * b
}

func main() {
	fmt.Println("Hello World!")

	//export func
	say.SayHello() //export func인 SayHello() 호출
	//say.sayHi()	//소문자로 시작하는 export func은 실행 불가

	//constants
	Constants()

	//variables
	Variables()

	//Type
	fmt.Println(multiply(3, 4))
}
