# Go Programming(part3)

## 1. 배열

### 1) 배열

* 배열 : 연속적인 메모리 공간에 동일한 타임의 데이터를 순서대로 저장하는 자료구조

* 인덱스 - Zero based array(0, 1, 2, ... 순서대로 인덱스를 매김)

* 선언 : var (변수명) [ (배열크기) ] (데이터타입)

  ```go
  var num [3]int
  ```

* Go에서 배열의 크기는 Type을 구성하는 한 요소임

  -> [3]int와 [5]int는 서로 다른 타입으로 인식됨

* 초기화 - 배열을 정의할 때, 초기값 설정 가능

  [ (배열크기) ] (데이터타입) { (초기값1, 초기값2, ...) }

  ** [...] : 배열 크기 생략. 자동으로 초기화 요소 개수만큼 크기가 정해짐

  ```go
  var arr1 = [3]int{1, 2, 3}
  var arr2 = [...]int{1, 2, 3}
  ```

### 2) 다차원 배열

* 다차원 배열 : 배열 크기 부분을 복수 개로 설정하여 선언함

  ```go
  var multiArray [3][4][5]int //정의
  multiArray[0][1][2] = 10		//사용
  ```

* 다차원 배열 초기화 - 배열 초기값 안에 다시 배열값을 넣는 형태

  ```go
  var mArr = [2][3]int{
  		{1, 2, 3},
      {4, 5, 6},  //끝에 콤마 추가
  }
  ```



## 2. 슬라이스(Slice)

### 1) 슬라이스

* 슬라이스 : 배열과 달리 고정된 크기를 미리 지정하지 않을 수 있고, 차후 그 크기를 동적으로 변경할 수도 있음 + 부분 배열을 발췌할 수도 있음

* 선언 : var  (변수명) [] (타입)

  ```go
  var s []int	//선언
  s = []int{1, 2, 3}	//슬라이스에 리터럴 값 지정
  s[1] = 10
  ```

* make() - slice 생성 함수

  -> 슬라이스의 길이(Length), 용량(Capacity)를 임의로 지정할 수 있음

  -> len(), cap() 함수로 확인 가능

  -> 모든 요소가 Zero value인 슬라이스로 생성됨

  -> 만약 3번째 capacity 파라미터를 생략하면, capacity는 length와 같은 값을 가짐

  ```go
  s := make([]int, 5, 10)
  fmt.Println(len(s), cap(s))		//5, 10
  ```

* Nill Slice : 슬라이스에 별도의 길이, 용량을 지정하지 않으면 기본적으로 길이와 용량이 0인 슬라이스가 생성됨. 이때의 슬라이스를 Nill Slice라고 함

  -> nill과 비교하면 true를 리턴함

  ```go
  var nilSlice []int
  if nilSlice == nil {
  	fmt.Println("nll slice = nill")
  }
  ```

### 2) 부분 슬라이스(Sub-slice)

* 부분 슬라이스 : 슬라이스에서 일부를 발췌하여 생성

* 형식 : (슬라이스명)[ (처음인덱스) : (마지막인덱스) ]

* inclusive : 처음 인덱스, exclusive : 마지막 인덱스

  ```go
  s := []int{0, 1, 2, 3, 4, 5, 6, 7}
  subSlice := s[2:5]	//[2, 3, 4]
  ```

* 처음 인덱스, 마지막 인덱스 생략 가능. 처음 인덱스는 0, 마지막 인덱스는 그 슬라이스의 마지막 인덱스가 자동으로 대입 됨

  ```go
  s := []int{0, 1, 2, 3, 4, 5}
  s = s[2:5]	//[2 3 4]
  s = s[1:]		//[3 4]
  ```

### 3) 슬라이스 추가, 병합(append), 복사(copy)

* append() : 슬라이스에 새로운 요소 추가

* 형식 : append( (슬라이스객체), (추가할요소의값1), (추가할요소의값2), ... )

  ```go
  s := []int{0, 1}	//[0 1]
  s = append(s, 2)	//[0 1 2]
  s = append(s, 3, 4, 5)	//[0 1 2 3 4 5]
  ```

* 한 슬라이스에 다른 슬라이스 병합 가능

  ```go
  sliceA = append(sliceA, sliceB...)
  ```

* append 프로세스

  -> 슬라이스 용량(capacity)가 남아 있는 경우, 그 용량 내에서 슬라이스의 길이(length)를 변경하여 데이터를 추가

  -> 슬라이스 용량(capacity)를 초과하는 경우, 현재 용량의 2배에 해당하는 새로운 Underlying array를 생성 + 기존 배열 값들을 모두 새 배열에 복제한 후 다시 슬라이스를 할당

* copy() : 한 슬라이스를 다른 슬라이스로 복사

  ```go
  source := []int{0, 1, 2}
  target := make([]int, len(source), cap(source)*2)
  copy(target, source)
  ```

### 4) 슬라이스 내부구조

* 슬라이스는 내부적으로 사용하는 배열의 부분 역역인 세그먼트에 대한 메터 정보를 가지고 있음

* 3개의 필드로 구성됨

  ​	(1) 내부적으로 사용하는 배열에 대한 포인터 정보

  ​	(2) 세그먼트 길이

  ​	(3) 세그먼트의 최대 용량(capacity)

* 처음 슬라이드가 생성될 때, (만약 길이, 용량이 지정되었다면) 내부적으로 용량(capacity)만큼의 매열을 생성

  -> 슬라이드 1번째 필드에 그 배열의 처음 메모리 위치를 지정

  -> 2번째 길이 필드는 지정된 (첫 배열요소로부터의) 길이를 갖게 되고

  -> 3번째 용량 필드는 전체 배열 크기를 가짐

  ![image-20201231161511515](/Users/sunhapark/Library/Application Support/typora-user-images/image-20201231161511515.png)

## 3. Map

### 1) Map 선언 & 초기화

* Map : 키(Key)에 대응하는 값(Value)을 신속하게 찾는 해시테이블(Hash table)을 구현한 자료구조

* 선언 : map[ (Key타입) ] (Value타입)

  * Nil Map : 이때 선언된 변수는 nill 값을 갖음. 이때의 Map. 초기화 전에는 어떤 데이터도 쓸 수 없음

  * make() : Nil map 초기화 함수

    -> 해시테이블 자료구조를 메모리에 생성하고 그 메모리를 가리키는 map value를 리턴함

    -> map value : 내부적으로 runtime.hmap 구조체를 가리키는 포인터. (idMap 변수는 이 해시테이블을 가리키는 map을 가리킴)

  * 리터럴(literal)로 초기화 : map[ (Key타입) ] (Value타입) { (key) : (value) }

  ```go
  var idMap map[int]string	//선언
  idMap = make(map[int]string)	//초기화
  tickers := map[string]string {	//리터럴로 초기화
    "GOOG": "Google Inc",
    "MSFT": "Microsoft",
    "FB": "FaceBook",
  }
  ```

### 2) Map 값 할당

* 값 할당 : (map변수) [ (key) ] = (값)

* 값 삭제 : delete( (map변수), (key) )

  ```go
  var m map[int]string
  
  m = make(map[int]string)
  m[901] = "Apple"
  m[134] = "Grape"
  m[777] = "Tomato"
  
  str := m[134]
  noData := m[999]
  
  delete(m, 777)
  ```

### 3) Map key 체크

* map 안에 특정 키가 존재하는지 체크 : (map변수) [ (key) ]

  -> 리턴 값 2개 : 키에 상응하는 값, 그 키가 존자해는지 아닌지를 나타내는 bool

  ```go
  val, exists := tickers["MSFT"]
  if !exists {
  	fmt.Println("No MSFT ticker")
  }
  ```

### 4) Map 열거

* for문으로 map 열거

  -> Map은 unordered임. 따라서 출력 순서가 무작위.

  ```go
  myMap := map[string]string {
  	"A": "Apple",
  	"B": "Banana",
  	"C": "Charlie",
  }
  
  for key, val := range myMap {
  	fmt.Println(key, val)
  }
  ```

  

[참고] [http://golang.site/go/article/12-Go-%EC%BB%AC%EB%A0%89%EC%85%98---%EB%B0%B0%EC%97%B4]