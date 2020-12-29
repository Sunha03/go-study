# Go Programming(part2)

## 1. 함수

* 함수 : 여러 문장을 묶어서 실행하는 코드 블럭의 단위.

* 형식 : func [ 함수명 ] ( 파라미터, ... ) [ 리턴타입, ... ] { ... }

  ​		ex) func say(name string) string { ... }

* 함수는 패키지 안에 정의되며 호출되는 함수가 호출하는 함수의 반드시 앞에 위치해야 할 필요는 없음



## 2. Pass By Value / Pass By Reference

* 파라미터 전달방식 - Pass By Value / Pass By Reference

### 1) Pass By Value

* 함수를 호출할 때 매개변수 값을 복사해서 함수 내부로 전달함

* 함수 내의 파라미터 변수가 변경되더라도 호출함수의 본래 변수는 변하지 않음

  ex) msg의 값 "Hello" 문자열이 복사되어 함수 say()에 전달됨. 만약 say() 내에서 msg 값이 변경된다 하더라도 호출함수 main()에서의 msg 변수는 변함이 없음

  ```go
  func main() {
      msg := "Hello"
      say(msg)
  }
   
  func say(msg string) {
      println(msg)
  }
  ```

### 2) Pass By Reference

* 함수를 호출할 때 매개변수를 메모리 주소로 전달함

* 함수 내의 파라미터 변수가 변경하면 호출함수의 본래 변수도 변함

* Dereferencing : 변수 앞에 '*'를 붙임

  ex) say()에서 *string과 같이 파라미터가 포인터임을 표히사고, 이때 say ()의 msg는 문자열이 아니라 문자열을 갖는 메모리 영역의 주소를 갖게 됨

  ```go
  func main() {
      msg := "Hello"
      say(&msg)
      println(msg) //변경된 메시지 출력
  }
   
  func say(msg *string) {
      println(*msg)
      *msg = "Changed" //메시지 변경
  }
  ```

  

## 3. Variadic Function(가변인자함수)

* variadic function : 함수에 고정된 수의 파라미터들을 전달하지 않고 다양한 숫자의 파라미터를 전달하는 함수

* 형식 : func [ 함수명 ] ([ 파라미터명 ] ...[ 파라미터타입 ]) { ... }

* 가변 파라미터를 갖는 함수를 호출할 때는 0, 1, 2, ..., n개의 동일타입 파라미터를 전달할 수 있음

  ```go
  func main() {   
      say("This", "is", "a", "book")
      say("Hi")
  }
   
  func say(msg ...string) {
      for _, s := range msg {
          println(s)
      }
  }
  ```



## 4. 함수 리턴값

* Go는 함수 리턴값이 0개, 1개, n개 일 수 있음

  ```go
  func main() {
      count, total := sum(1, 7, 3, 5, 9)
      println(count, total)   
  }
   
  func sum(nums ...int) (int, int) {
      s := 0      // 합계
      count := 0  // 요소 갯수
      for _, n := range nums {
          s += n
          count++
      }
      return count, s
  }
  ```

* Nameds Return Parameter : 리턴되는 값들을 (함수에 정의된) 리턴 파라미터들에 할당할 수 있는 기능

  -> 코드 가독성을 높임

  -> 실제 return 문에는 아무 값들을 리턴하지 않지만, 리턴되는 값이 있을 경우에는 빈 return 문을 반드시 써줘야 함

  ```go
  func sum(nums ...int) (count int, total int) {
      for _, n := range nums {
          total += n
      }
      count = len(nums)
      return
  }
  ```

  

## 5. 익명함수(Anonymous Funciton)

* 익명함수 : 함수명을 갖지 않는 함수. 일반적으로 그 함수 전체를 변수에 할당 or 다른 함수의 파라미터에 직접 정의되어 사용됨

  ```go
  func main() {
      sum := func(n ...int) int { //익명함수 정의
          s := 0
          for _, i := range n {
              s += i
          }
          return s
      }
   
      result := sum(1, 2, 3, 4, 5) //익명함수 호출
      println(result)
  }
  ```



## 6. 일급함수

* Go언어에서 함수는 일급함수로서 Go의 기본 타입과 동일하게 취급됨

* 따라서 다른 함수의 파라미터로 전달 or 다른 함수의 리턴 값으로도 사용 가능. 즉, 함수의 입력 파라미터 or 리턴 파라미터로서 함수 자체를 사용 가능

* 함수를 파라미터로 사용하는 2가지 방법

  ​	1) 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법

  ​	2) 직접 다른 함수 호출 파라미터에 함수를 적는 방법

  ```go
  func main() {
      //변수 add 에 익명함수 할당
      add := func(i int, j int) int {
          return i + j
      }
   
      // add 함수 전달
      r1 := calc(add, 10, 20)
      println(r1)
   
      // 직접 첫번째 파라미터에 익명함수를 정의함
      r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
      println(r2)
   
  }
   
  func calc(f func(int, int) int, a int, b int) int {
      result := f(a, b)
      return result
  }
  ```



## 7. type문을 사용한 함수 원형 정의

* 구조체(struct), 인터페이스 등 Custom Type(or User Defined Type)을 정의하기 위해 사용됨

* 또한, 함수 원형을 정의하는데 사용됨

* 델리게이트(Delegate) : 함수의 원형을 정의하고 함수를 타 메소드에 전달/리턴 받는 기능

  ```go
  // 원형 정의
  type calculator func(int, int) int
   
  // calculator 원형 사용
  func calc(f calculator, a int, b int) int {
      result := f(a, b)
      return result
  }
  ```

  



