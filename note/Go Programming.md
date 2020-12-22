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