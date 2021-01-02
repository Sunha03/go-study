# Go Programming(part4)

## 1. Go 패키지(Package)

* Go 패키지 - 코드의 모듈화, 코드의 재사용 기능을 제공함 (https://golang.org/pkg)

* Go는 패키지를 사용해서 작은 단위의 컴포넌트를 작성하고, 이러한 작은 패키지들을 활용해서 프로그램을 작성할 것을 권장함

* Go는 실제 프로그램 개발에 필요한 많은 패키지들을 표준 라이브러리로 제공함

  -> GOROOT/pkg 안에 존재함

* GOROOT 환경변수 : Go 설치 디렉토리를 가리킴. 보통 Go 설치하면 자동으로 추가됨.

### 1) Main 패키지

* main 패키지 : 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만듦
* main 패키지 안의 main() 는 프로그램의 시작점(Entry Point)가 됨
* 패키지를 공유 라이브러리로 만들 때에는, main 패키지 or main()를 사용하면 안됨

### 2) 패키지 import

* 다른 패키지를 프로그램에서 사용하기 위해서는 import 를 사용하여 패키지를 포함시킴

  ```go
  import "fmt"
  func main() {
  	fmt.Println("Hello World!")
  }
  ```

* 패키지를 import 할 때, Go 컴파일러는 GOROOT or GOPATH 환경변수를 검색함 (표준 패키지 -> GOROOT/pkg에서, 사용자 패키지 or 3rf party 패키지 -> GOPATH/pkg)

* GOROOT 환경변수는 Go 설치 시 자동으로 시스템에 설정되지만, GOPATH는 사용자가 지정해 주어야 함

### 3) 패키지 Scope

* public 패키지 - 패키지 내에는 함수, 구조체, 인터페이스, 메소드 등이 존재하는데, 이들의 이름(Identifier)이 첫 문자를 대문자로 시작하면 public으로 사용 가능
* 패키지 외부에서 이들을 호출/사용할 수 있음
* non-public 패키지 - 패키지 이름이 소문자로 시작됨

### 4) 패키지 init 함수, alias

* init() - 패키지가 로드되면서 실행되는 함수를 별도의 호출 없이 자동으로 호출됨

* 개발자가 패키지를 작성할 때, 패키지 실행 시 처음으로 호출되는 init() 함수 작성 가능

  ```go
  package testlib
  var pop map[string]string
  func init() {	//패키지 로드시 map 초기화
  	pop = make(map[string]string)
  }
  ```

* 경우에 따라 패키지를 Import 하면서 그 패키지 안의 init() 함수만 호출하고자 하는 케이스가 있음

  -> 이런 경우 패키지 import 할 때, '_'라는 alias를 지정함

  ```go
  package main
  import _ "other/xlib"
  ```

* 패키지 이름이 동일하지만, 서로 다른 버전 or 서로 다른 위치에서 로딩하고자 할 때는, 패키지 alias를 사용하여 구분 가능함

  ```go
  import (
  	mongo "other/mongo/db"
  	mysql "other/mysql/db"
  )
  func main() {
  	mondb := mongo.Get()
  	mydb := mysql.Get()
  	...
  }
  ```

### 5) 사용자 정의 패키지 생성

* 개발자는 사용자 정의 패키지를 만들어 재사용 가능한 컴포넌트를 만들어 사용할 수 있음

* 사용자 정의 라이브러리 패키지

  -> 일반적으로 폴더를 하나 만들고 그 폴더 안에 .go 파일들을 만들어 구성함

  -> 하나의 서브 폴더 안에 있는 .go 파일들은 동일한 패키지 명을 가지며, 패키지명은 해당 폴더의 이름과 같게함

  -> 해당 폴더에 있는 여러 *.go 파일들은 하나의 패키지로 묶임

* Optional

  -> 사이즈가 큰 복잡한 라이브러리 같은 경우, "go install" 명령을 사용하여 라이브러리를 컴파일하여 cache할 수 있음

  -> 이렇게 하면 다음 빌드 시 빌드 타임을 줄일 수 있음

  -> ex) Go 패키지 빌드 > /pkg 폴더에 인스톨 하기 위해 testlib 폴더 안에서 "go install" 실행 > /pkg/windows_amd64/24lab.net 안에 testlib.a 파일 생성됨

  ```go
  package main
  import (
  	"fmt"
  	"24lab.net/testlib"
  )
  func main() {
  	song := testlib.GetMusic("Alicia Keys")
  	fmt.Println(song)
  }
  ```

  * GOROOT > GOPATH 순서대로 패키지를 찾음