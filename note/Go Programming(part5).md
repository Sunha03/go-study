# Go Programming(part5)

## 1. Go defer

* defer : 특정 문장 or 함수를 나중에 (deferf를 호출하는 함수가 리턴하기 직전에) 실행하게 함

* 일반적으로는 C#, Java 같은 언어에서의 finally 블럭처럼 마지막에 Clean-up 작업을 위해 사용됨

  ```go
  import "os"
  
  func main() {
  	f, err := os.Open("1.txt")
  	if err != nil {
  		panic(err)
  	}
  	
  	defer f.Close()	// main() 마지막에 파일 close 실행
  	
  	bytes := make([]byte, 1024)	// 파일 읽기
  	f.Read(bytes)
  	println(len(bytes))
  }
  ```

  -> 파일을 Open한 후 항상 파일을 close 함



## 2. panic()

* panic() : 현재 함수를 즉시 멈추고 현재 함수에 defer 함수들을 모두 실행한 후 즉시 리턴함

* panic 모드 실행 방식은 다시 상위함수에도 똑같이 적용되고, 계속 콜스택을 타고 올라가며 적용됨 + 마지막에는 프로그램이 에러를 내고 종료하게 됨

  ```go
  import "os"
  
  func main() {
  	openFile("Invalid.txt")		// 잘못된 파일명 입력
  	println("Done") // openFile() 안에서 panic이 실행되면 아래 print는 실행 안됨
  }
  
  func openFile(fn string) {
    f, err := os.Open(fn)
    if err != nil {
      panic(err)
    }
    
    defer f.Close()
  }
  ```



## 3. recover()

* recover() : panic()에 의한 패닉 상태를 다시 정상 상태로 되돌리는 함수

  ```go
  import (
  	"fmt"
  	"os"
  )
  
  func main() {
  	openFile("Invalid.txt")		// 잘못된 파일명 입력
  	println("Done")		// recover()에 의해 실행됨
  }
  
  func openFile(fn string) {
    defer func() {	// defer() : panic 호출 시 실행됨
  		if r := recover(); r != nil {
  			fmt.Println("OPEN ERROR", r)
  		}
    }()
  	
  	f, err := os.Open(fn)
  	if err != nil {
  		panic(err)
  	}
  	
  	defer f.Close()
  }
  ```

  -> recover()가 panic 상태를 제거하고 openFile()의 다음 문장인 println()을 호출함