package main

import (
	"fmt"
	"strings"
	"test/say"
)

//Struct
type person struct {
	name    string
	age     int
	favFood []string
}

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

//Multiple return -> (type, type, ...)
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//Multiple Auguments -> (var ...type)
func numbers(words ...string) {
	fmt.Println(words)
}

//Naked Return
func lenAndUpper2(name string) (length int, uppercase string) {
	length = len(name) //변수 값 업데이트
	uppercase = strings.ToUpper(name)
	return
}

//defer
func lenAndUpper3(name string) (length int, uppercase string) {
	defer fmt.Println("The End") //func 종료 후 실행됨

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

//for, range, ...args
func superAdd(numbers ...int) int {
	//range : array에 loop를 적용
	for index, number := range numbers {
		//(output) 0 10 / 1 20 / 2 30 / 3 40 / 4 50
		fmt.Println(index, number)
	}
	return 1
}

//for
func superAdd2(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return 1
}

//range(_ : index 사용하지 않을 때)
// -> for문 안에서만 작동함
func superAdd3(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

//if
func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

//if(+variable expression)
func canIDrinkKoreanAge(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true
}

//switch
func canIDrink2(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

//switch
func canIDrink3(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return true
}

//switch(+variable expression)
func canIDrinkKoreanAge2(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
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

	//Multiple return
	totalLength, upperName := lenAndUpper("test")
	fmt.Println(totalLength, upperName)

	//Multiple Auguments
	numbers("one", "two", "three")

	//Naked Return
	totalLength2, upperName2 := lenAndUpper2("test2")
	fmt.Println(totalLength2, upperName2)

	//defer
	totalLength3, upperName3 := lenAndUpper3("test3")
	fmt.Println(totalLength3, upperName3)

	//for, range, ...args
	superAdd(10, 20, 30, 40, 50)
	superAdd2(10, 20, 30, 40, 50)
	result := superAdd3(10, 20, 30, 40, 50)
	fmt.Println(result)

	//if
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkKoreanAge(16))

	//switch
	fmt.Println(canIDrink2(16))
	fmt.Println(canIDrink3(20))
	fmt.Println(canIDrinkKoreanAge2(17))

	//Pointer1
	a := 2
	b := a
	fmt.Println(a, b)

	//Pointer2
	a2 := 2
	b2 := a2
	a2 = 10
	fmt.Println(a2, b2)

	//Pointer3
	a3 := 2
	b3 := &a3
	a3 = 10
	fmt.Println(a3, b3)
	fmt.Println(a3, *b3)

	//Pointer4
	a4 := 2
	b4 := &a4
	*b4 = 20
	fmt.Println(a4, b4)

	//Array
	names := [5]string{"one", "two", "three", "four"}
	names[3] = "change"
	fmt.Println(names)

	//Slice
	names2 := []string{"1", "2", "3", "4", "5"}
	names2[3] = "change"
	newNames2 := append(names2, "adding")
	fmt.Println(names2)
	fmt.Println(newNames2)

	//Map
	test := map[string]string{"name": "go", "age": "20"}
	fmt.Println(test)
	for key, value := range test {
		fmt.Println(key, value)
	}

	//Struct
	favFood := []string{"pasta", "pizza"}
	//struct 생성1
	golang := person{}
	golang.name = "go"
	golang.age = 10
	golang.favFood = favFood
	//struct 생성2
	java := person{"java", 15, favFood}
	//struct 생성3
	python := person{name: "python", age: 15, favFood: favFood}
	fmt.Println(golang)
	fmt.Println(java)
	fmt.Println(python)
}
