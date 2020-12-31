# Go Programming(part1)

## 1. 변수(Variable)

- 선언 : var (변수명) (타입)

- 선언&초기화 : var (변수명) (타입) = (값)

  ```go
  var num int = 3
  ```

- 선언된 변수를 사용하지 않으면 -> ERROR 발생

- 복수 변수 - 동일한 타입의 변수가 여러 개 있을 경우, 변수들을 나열하고 마지막에 타입을 1번만 지정 가능. 초기값은 순서대로 변수에 할당됨

- default 값 - 숫자(0), bool(false), string("")

- Go에서는 할당되는 값을 보고 그 타입을 추론하는게 가능함

- Short Assignment Statement(:=)

  1) 함수(func) 내에서만 사용 가능

  2) 함수(func) 밖에서는 var를 사용해야 함



## 2. 상수(Const)

- 선언 : const (상수명) (타입) = (값)

  ```go
  const KEY string = "key"
  ```

- Go에서는 할당되는 값을 보고 그 타입을 추론하는게 가능함

- 복수 상수 - 괄호 안에 상수들을 나열하여 사용

- identifier "iota" - 상수 값을 0부터 순차적으로 부여



## 3. Go 키워드

- 예약 키워드 25개 - 변수명, 상수명, 함수명 등의 identifier로 사용할 수 없음

  -> break / case / chan / const / continue / default / defer / else / fallthrough / for / func / go / goto / if / import / interface / map / package / range / return / select / struct / swtich / type / var



## 4. Go Data Type

* bool 타입 - bool
* 문자열 타입 - string : 한 번 생성되면 수정될 수 없는 immutable 타입
* 정수형 타입 - int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
* Float 및 복소수 타입 - float32, float64, complex64, complex128
* 기타 - byte : uint8과 동일. 바이트 코드에 사용, rune : int32와 동일. 유니코드 코드포인트에 사용



## 5. 문자열

* Black Quote('') (= Raw String Literal) : 별도로 해석되지 않고, Raw String 그대로의 값을 가짐. 복수 라인의 문자열을 표현할 때 자주 사용됨.

  ex) 문자열 안에 \n이 있을 경우 New Line으로 인식되지 않음

* 이중인용부호("") (=Interpreted String Literal) : 복수 라인에 걸쳐 쓸 수 없으며, 인용부호 안의 Escape 문자열들은 특별한 의미로 해석됨. 이중인용부호를 이용해 문자열을 여러 라인에 걸쳐 쓰기 위해서는 '+' 연산자를 이용해 결합하여 사용함.

  ex) 문자열 안에 \n이 있을 경우 New Line으로 인식함



## 6. 데이터 타입 변환(Type Conversion)

* 데이터 타입 변환 : 다른 데이터 타입으로 변환하기 위해 T(v)와 같이 표현함. 이 때 T는 변환하고자 하는 타입을 표시, v는 변환될 값(value)를 지정함

  ```go
  varFloat := float32(100)	//int -> float
  varBytes := []byte("ABC")	//string -> []byte
  ```

* Go에서 타입 간 변환은 명시적으로 지정해주어야 함. 암묵적(=implicit) 변환이 없음.

  ex) int에서 uint로 암묵적 변환이 일어나지 않음



## 7. 연산자

### 1) 산술연산자

 * 사칙연산자(+, -, *, /, %(Modulus))
 * 증감연산자(++, --)

### 2) 관계연산자

* 서로의 크기를 비교하거나 동일함을 체크(==, !=, >=, <=, ...)

### 3) 논리연산자

* AND, OR, NOT, ...

### 4) Bitwise 연산자

* 비트단위 연산을 위해 사용됨.(바이너리 AND, OR, XOR, 바이너리 쉬프트 연산자, ...)

### 5) 할당연산자

* = 연산자, += 연산자, &= 연산자, <<= 연산자, ...

### 6) 포인터연산자

* C++과 같이 & or * 를 사용하여 해당 변수의 주소를 얻어내거나 이를 반대로 Dereference할 때 사용함
* 단, 포인터 산술(포인터에 더하고 빼는 기능)은 제공하지 않음



## 8. 조건문

### 1) if문

* if문 : ( ) 안의 조건이 맞으면 { } 블럭 안의 내용을 실행함

* 괄호 생략 가능(단, 조건 블럭 시작 브레이스 '{'를 if문과 같은 라인에 두어야 함)

* 조건식을 이용하기 전에 간단한 문장(Optional Statement)을 함께 실행할 수 있음. 이때 정의된 변수는 if문 블럭 안에서만 사용할 수 있음

  ```go
  i := 1
  max := 10
  if val := i * 2; val < max {
  	fmt.Println(val)
  }
  //var++		//ERROR - out of if statement
  ```

### 2) Switch문

* switch문 : switch문 뒤에 하나의 변수(or Expression)을 지정하고, case 문에 해당 변수가 가질 수 있는 값들을 지정하여, 각 경우에 다른 문장 블럭들을 실행할 수 있음. 여러 값을 비교해야 하는 경우 or 다수의 조건식을 체크해야 하는 경우 사용.

* case 값들이 복수개인 경우 콤마 사용

  ex) case 3, 4 :

* Go switch만의 특징

  - switch 뒤에 expression 생략 가능

    -> Go는 switch expression을 true로 생각하고 1번째 case문으로 이동하여 검사함

  - case문에 expression 사용 가능

    -> Go는 case문에 복잡한 expression을 가질 수 있음

  - no default fall through

    -> Go는 break가 없어도 다음 case로 이동하지 않음

    -> 다음 case문을 실행하려면 fallthrough문 명시 필요

  - type switch

    -> Go는 변수의 Type에 따라 case 분기가 가능함



## 9. 반복문

### 1) for문

* for문 : "for (초기값); (조건식); (증감) { ... }" 형식. 초기값, 조건식, 증감식 등은 경우에 따라 생략 가능.

* 초기값; 조건식; 증감을 둘러싸는 괄호를 쓰면 에러남. 항상 생략해야 함.

  ```go
  sum := 0	//for statements
  for i := 1; i <= 100; i++ {
  	sum += i
  }
  ```

* 조건식만

  ```go
  n := 1
  for n < 100 {
  	n *= 2
  }
  ```

* 무한루프

  ```go
  for {
  	println("Infinite loop")        
  }
  ```

* for range문  : 컬렉션으로부터 한 요소(element)씩 가져와서 차례로 for 블럭의 문장들을 실행함. foreach와 비슷함.

* "for 인텍스,요소값 := range 컬렉션" 형식

  -> range 키워드 다음의 컬렉션으로부터 하나씩 요소를 리턴해서 그 요소의 위치 인덱스와 값을 for 키워드 다음의 2개의 변수에 각각 할당함

### 2) Break, continue, goto문

* break문 : for문 내에서 즉시 빠져나옴. switch, select문에서 사용 가능

* break 레이블 : 지정된 레이블로 이동. 현재의 루프를 빠져나와 지정된 레이블로 이동 + break문의 직속 for문 전체의 다음 문장을 실행

  -> for문 바로 위에 작성

  ```go
  i := 0
  L1:
  	for {
  		if i == 0 {
  			break L1
  		}
  	}
  fmt.Println("OK")
  //(Outputs) OK
  ```

  

* continue문 : for문 중간에서 나머지 문장들을 실행하지 않고 for문 시작부분으로 바로 이동

* goto문 : 기타 임의의 문장으로 이동