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
  
  #실행결과
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
  ```

  

## Variables

- Variables

  => var [변수명] [타입]

  -- 한번 선언한 변수 재할당 가능

  -- var 선언 시 축약형으로 사용 가능함(go언어가 알아서 var 타입 확인함)

  -- 축약형 코드는 func 안에서만 사용 가능함

  ```go
  /Variables : 변경가능
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
  ```

  







[참조] [https://soyoung-new-challenge.tistory.com/85?category=893866]