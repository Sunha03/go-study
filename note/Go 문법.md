# Go 문법

## Export / Private Function 사용

- say.go 파일 생성, export func 생성

  -> 프로젝트 폴더/say 폴더/say.go 파일 생성

  - <u>say.go</u>

  ```go
  package say
  
  import "fmt"
  
  //sayHi func
  // - export 가능한 function -> 대문자로 시작(소문자는 불가)
  // - private function임
  func sayHi() {
  	fmt.Println("Hi")
  }
  
  //SayHello func
  // - export 가능한 function -> 대문자로 시작
  func SayHello() {
  	fmt.Println("Hello")
  }
  ```

- export func 호출

  -> main.go에서 SayHello() 호출

  * <u>main.go</u>

  ```go
  package main
  
  import (
  	"fmt"
  	"test/say"
  )
  
  func main() {
  	fmt.Println("Hello World!")
  	say.SayHello() //export func인 SayHello() 호출
  	//say.sayHi()	//소문자로 시작하는 export func은 실행 불가
  }
  ```

- main.go 실행

  ```
  #terminal(/test)
  > go run main.go
  
  #Outputs
  Hello World!
  Hello
  ```

  

## Constants

- Constant -> 재할당 불가

  -- 축약형 코드 사용 불가

  ```go
  //Constants : 변경불가
  // - 한번 선언한 변수는 재할당 불가
  func Constants() {
  	const name string = "SH" //변수 앞에 const 선언
  	fmt.Println(name)
  }
  
  # Outputs
  SH
  ```



## Variables

- Variables

  => var [변수명] [타입]

  -- 한번 선언한 변수 재할당 가능

  -- var 선언 시 축약형으로 사용 가능함(go언어가 알아서 var 타입 확인함)

  -- 축약형 코드는 func 안에서만 사용 가능함

  ```go
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
  
# Outputs
str2
newStr2
  ```



## Type

- type 종류 - string, bool, numeric, array, slices, ...

  ** (a int, b int) = (a, b int)

  ```go
  //Type
  // - string, bool, numeric, array, slices, ...
  func multiply(a int, b int) int { //(a int, b int) = (a, b int)
  	return a * b
  }
  
  # main.go
  fmt.Println(multiply(3, 4))
  # Outputs
  12
  ```




## Multiple return

- func에서 return 여러개 반환

  -- return할 type을 () 안에 명시

  ```go
  //Multiple return -> (type, type, ...)
  func lenAndUpper(name string) (int, string) {
  	return len(name), strings.ToUpper(name)
  }
  
  # main.go
  totalLength, upperName := lenAndUpper("test")
  fmt.Println(totalLength, upperName)
  # Outputs
  4 TEST
  ```



## Multiple Auguments

- func에서 arguments 여러개 받기

  -- argument 타입 앞에 "..." 붙이기

  -- 원하는 개수만큼의 arguments를 주고 받을 수 있음

  -- "[arg1 arg2 arg3 ...]"와 같이 array 형태로 출력됨

  ```go
  //Multiple Auguments -> (var ...type)
  func numbers(words ...string) {
  	fmt.Println(words)
  }
  
  # main.go
  numbers("one", "two", "three")
  # Outputs
  [one two three]
  ```



## Naked Return

- naked return : return할 변수를 명시하지 않는 것

  -- func 함수 생성 시 return 부분에 변수와 타입을 함께 명시

  ​	-> 자동으로 이때의 변수가 return됨

  ```go
  //Naked Return
  func lenAndUpper2(name string) (length int, uppercase string) {
  	length = len(name) //변수 값 업데이트
  	uppercase = strings.ToUpper(name)
  	return
  }
  
  # main.go
  totalLength2, upperName2 := lenAndUpper2("test2")
  fmt.Println(totalLength2, upperName2)
  # Outputs
  5 TEST2
  ```

  

## Defer

- func 종료 후 실행되는 코드/기능

  -- defer를 사용하면 func이 끝난 뒤 필요한 것을 실행시킬 수 있음

  ​	ex) 오픈했던 이미지 닫기, 생성한 파일 삭제, API 요청 보내기, ...

  ```go
  //defer
  func lenAndUpper3(name string) (length int, uppercase string) {
  	defer fmt.Println("The End") //func 종료 후 실행됨
  
  	length = len(name)
  	uppercase = strings.ToUpper(name)
  	return
  }
  
  # main.go
  totalLength3, upperName3 := lenAndUpper3("test3")
  fmt.Println(totalLength3, upperName3)
  # Outputs
  The End
  5 TEST3
  ```



## for, range, ...args

- for - go에서 유일한 loop문

- range - array에 loop를 적용할 수 있도록 함

  -> 실행 시, index와 value를 같이 return 함

  ```go
  //for, range, ...args
  func superAdd(numbers ...int) int {
  	//range : array에 loop를 적용
  	for index, number := range numbers {
  		//(output) 0 10 / 1 20 / 2 30 / 3 40 / 4 50
  		fmt.Println(index, number)
  	}
  	return 1
  }
  
  # main.go
  superAdd(10, 20, 30, 40, 50)
  # Outputs
  0 10
  1 20
  2 30
  3 40
  4 50
  ```

  ```go
  //for
  func superAdd2(numbers ...int) int {
  	for i := 0; i < len(numbers); i++ {
  		fmt.Println(numbers[i])
  	}
  	return 1
  }
  
  # main.go
  superAdd2(10, 20, 30, 40, 50)
  # Outputs
  10
  20
  30
  40
  50
  ```

  ```go
  //range(_ : index 사용하지 않을 때)
  // -> for문 안에서만 작동함
  func superAdd3(numbers ...int) int {
  	total := 0
  
  	for _, number := range numbers {
  		total += number
  	}
  	return total
  }
  
  # main.go
  result := superAdd3(10, 20, 30, 40, 50)
  fmt.Println(result)
  # Outputs
  150
  ```

  

## if

- if문

  ```go
  //if + variable expression
  func canIDrinkKoreanAge(age int) bool {
  	if koreanAge := age+2; koreanAge < 18 {
  		return false
  	}
  	return true
  }
  
  # main.go
  fmt.Println(canIDrink(16))
  # Outputs
  false
  ```
  - variable expression : if문을 사용하는 순간 variable 생성 가능

    -> if문 안에서 variable을 작성하면 if-else 구문 안에서만 사용되는 변수

  ```go
  //if + variable expression
  func canIDrinkKoreanAge(age int) bool {
  	if koreanAge := age+2; koreanAge < 18 {
  		return false
  	}
  	return true
  }
  
  # main.go
  fmt.Println(canIDrinkKoreanAge(16))
  # Outputs
  true
  ```

  

## Switch

- switch문

  ```go
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
  
  # main.go
  fmt.Println(canIDrink2(16))
  # Outputs
  false
  ```

  ```go
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
  
  # main.go
  fmt.Println(canIDrink3(20))
  # Outputs
  true
  ```

  - variable expression

  ```go
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
  
  # main.go
  fmt.Println(canIDrinkKoreanAge2(17))
  # Outputs
  false
  ```

  

## Pointer

- &[ 출력 변수 ] -> 메모리 주소로 연결됨

- *[ 메모리주소가 담긴 변수 ] -> 해당 메모리 주소의 실제 값

- 포인터는 메모리 주소가 동일한 경우 바라보는 값을 변경할 수 있음

  ```go
  //Pointer1
  a := 2
  b := a
  //Pointer2
  a2 := 2
  b2 := a2
  a2 = 10
  //Pointer3
  a3 := 2
  b3 := &a3
  a3 = 10
  //Pointer4
  a4 := 2
  b4 := &a4
  *b4 = 20
  	
  # main.go
  fmt.Println(a, b)
  fmt.Println(a2, b2)
  fmt.Println(a3, b3)
  fmt.Println(a3, *b3)
  fmt.Println(a4, b4)
  # Outputs
  2 2
  10 2
  10 0xc0000b4028
  10 10
  20 0xc0000b4030
  ```

  

## Array

- 선언 : [ Array 변수 ] := [ 길이 ] 타입 { 값 }

- Go에서 Array는 꼭 길이를 명시해줘야 함

- array 값 변경 -> index를 지정해서 새로운 값 할당(기존의 인덱스 안에 있는 값만 변경 가능. 새로운 index를 추가로 생성 불가)

  ```go
  //Arrays
  names := [5]string{"one", "two", "three", "four"}
  names[3] = "change"
  
  # main.go
  fmt.Println(names)
  # Outputs
  [one two three change ]
  ```



## Slice

- Array와 동일하지만 길이가 정해져있지 않음

- slice 값 추가 가능 -> [새로운 변수] = append([기존의 slice 변수], [추가할 값])

  ([새로운 변수]에는 추가할 값이 추가된 slice를 반환함. [기존의 slice 변수]는 변경되지 않음)

  ```go
  //Slice
  names2 := []string{"1", "2", "3", "4", "5"}
  names2[3] = "change"
  newNames2 := append(names2, "adding")
  
  # main.go
  fmt.Println(names2)
  fmt.Println(newNames2)
  # Outputs
  [1 2 3 change 5]
  [1 2 3 change 5 adding]
  ```

  

## Map

- key : value 형태의 값

- 선언 : map [ 키타입 ] [ 값타입 ] { 실제_키&값 }

- iterate 사용 가능

  ```go
  //Map
  test := map[string]string{"name": "go", "age": "20"}
  
  # main.go
  fmt.Println(test)
  for key, value := range test {
  	fmt.Println(key, value)
  }
  # Outputs
  map[age:20 name:go]
  name go
  age 20
  ```



## Struct

- struct 생성 후 정의 방법

  1. struct 생성 -> 생성한 struct에 값 넣기
  2. 위에서 정의한 struct 값 순서대로 넣기
  3. key와 value를 지정해서 값 넣기

  ```go
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
  
  # main.go
  fmt.Println(golang)
  fmt.Println(java)
  fmt.Println(python)
  # Outputs
  {go 10 [pasta pizza]}
  {java 15 [pasta pizza]}
  {python 15 [pasta pizza]}
  ```

  



[참조] [https://soyoung-new-challenge.tistory.com/85?category=893866]

[참조] [https://soyoung-new-challenge.tistory.com/86?category=893866]