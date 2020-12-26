# Go Programming

## 1. 변수(Variable)

- 선언 : var [변수명] [타입]

- 선언&초기화 : var [변수명] [타입] = [값]

- 선언된 변수를 사용하지 않으면 -> ERROR 발생

- 복수 변수 - 동일한 타입의 변수가 여러 개 있을 경우, 변수들을 나열하고 마지막에 타입을 1번만 지정 가능. 초기값은 순서대로 변수에 할당됨

- default 값 - 숫자(0), bool(false), string("")

- Go에서는 할당되는 값을 보고 그 타입을 추론하는게 가능함

- Short Assignment Statement(:=)

  1) 함수(func) 내에서만 사용 가능

  2) 함수(func) 밖에서는 var를 사용해야 함



## 2. 상수(Const)

- 선언 : const [상수명] [타입] = [값]
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

  ex) 정수 100을 float로 변경 -> float32(100), 문자열을 바이트배열로 변경 -> []byte("ABC")

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

