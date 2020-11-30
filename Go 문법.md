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




## Multiple return

- func에서 return 여러개 반환

  -- return할 type을 () 안에 명시

  ```go
  //Multiple return -> (type, type, ...)
  func lenAndUpper(name string) (int, string) {
  	return len(name), strings.ToUpper(name)
  }
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
  ```

  





[참조] [https://soyoung-new-challenge.tistory.com/85?category=893866]

[참조] [https://soyoung-new-challenge.tistory.com/86?category=893866]