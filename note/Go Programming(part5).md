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



## 4. Go루틴(goroutine)

### 1) goroutine

* Go루틴 : Go 런타임이 관리하는 Lightweight 논리적 (or 가상적) 쓰레드

  ** goroutine은 OS 쓰레드보다 훨씬 가볍게 비동기 Concurrent 처리를 구현하기 위해 만든 것으로, 기본적으로 Go 런타임이 자체 관리함. Go 런타임 상에서 관리되는 작업 단위인 여러 goroutine들은 종종 하나의 OS 쓰레드 1개로도 실행되곤 함. 즉, goroutine들은 OS 쓰레드와 1:1로 대응되지 않고, Multiplexing으로 훨씬 적은 OS 쓰레드를 사용함. 메모리 측면에서도 OS 쓰레드가 1mb의 스택을 갖는 반면, goroutine은 이보다 훨씬 작은 몇 kb의 스택을 갖음(필요시 동적으로 증가). Go 런타임은 goroutine을 관리하면서 Go 채널을 통해 goroutine 간의 통신을 쉽게 할 수 있도록 함

* Go에서 "go" 키워드를 사용하여 함수를 호출하면, 런타임시 새로운 goroutine을 실행함

* goroutine은 비동기적(asynchronously)으로 함수루틴을 실행함

  -> 여러 코드를 동시에(Concurrently) 실행하는데 사용됨

  ```go
  import (
  	"fmt"
  	"time"
  )
  
  func say(s string) {
  	for i := 0; i < 10; i++ {
  		fmt.Println(s, "***", i)
  	}
  }
  
  func main() {
  	say("Sync")	// 함수를 동기적으로 실행
  	
  	go say("Async1")	// 함수를 비동기적으로 실행
  	go say("Async2")
  	go say("Async3")
  	
  	time.Sleep(time.Second * 3)		// 3초 대기
  }
  ```

  -> say()를 동기적으로 호출 + 동일한 say()를 비동기적으로 3번 호출함

  -> 1번째 동기적 호출은 say()가 완전히 끝났을 때, 다음 문장으로 이동

  -> 다은 3개의 go say() 비동기 호출은 별도의 goroutine들에서 동작하며, 메인루틴은 계속 다음 문장(여기서는 time.Sleep)을 실행함

  -> goroutine들은 실행순서가 일정하지 않으므로 프로그램 실행시 마다 다른 출력 결과가 나올 수 있음

### 2) 익명함수 goroutine

* goroutine은 익명함수(anonymous function)에 대해 사용할 수 있음

  -> 즉, go 키워드 뒤에 익명함수를 바로 정의하는 것으로, 이는 익명함수를 비동기로 실행하게 됨

  ```go
  import (
  	"fmt"
  	"sync"
  )
  
  func main() {
  	var wait sync.WaitGroup	// WaitGroup 생성
  	wait.Add(2)		// 2개의 goroutine을 기다림
  	
  	go func() {		// 익명함수를 사용한 goroutine
  		defer wait.Done()	// 끝나면 .Done() 호출
  		fmt.Println("Hello")
  	}()
  	
  	go func(msg string) {		// 익명함수에 파라미터 전달
  		defer wait.Done()		// 끝나면 .Done() 호출
  		fmt.Println(msg)
  	}("Hi")
  	
  	wait.Wait()		// goroutine 모두 끝날때까지 대기
  }
  ```

  -> 1번째 익명함수 - 간단히 "Hello" 문자열 출력. 이를 goroutine으로 실행하면 비동기적으로 그 익명함수를 실행하게 됨

  -> 2번째 익명함수 - 파라미터를 전달하는 예제. 익명함수에 파라미터가 있는 경우, go 익명함수 바로 뒤에 파라미터(Hi)를 함께 전달하게 됨

  -> sync.WaitGroup - 기본적으로 여러 goroutine들이 끝날때까지 기다리는 역할을 함. WaitGroup을 사용하기 위해서는 먼저 Add()에 몇 개의 goroutine을 기다릴 것인지 지정하고, 각 goroutine에서 Done()를 호출함(위에서는 defer를 사용). 그리고 메인 루틴에서는 Wait()를 호출하여 goroutine들이 모두 끝나기를 기다림.

### 3) 다중 CPU 처리

* Go는 디폴트로 1개의 CPU를 사용함. 즉, 여러 개의 goroutine을 만들더라도, 1개의 CPU에서 작업을 시분할하여 처리함(Concurrent 처리)

* 만약 머신이 복수 개의 CPU를 가진 경우, Go 프로그램을 다중 CPU에서 병렬처리(Parallel 처리)하게 할 수 있는데, 병렬처리를 위해서는 아래와 같이 runtime.GOMAXPROCS(CPU_수) 함수를 호출해야 함(여기서 CPU_수는 Logical CPU 수를 가리킴)

  ```go
  import "runtime"
  
  func main() {
  	runtime.GOMAXPROCS(4) 	// 4개의 CPU 사용
  	...
  }
  ```

  ** Concurrency(or Concurrent 처리)와 Parallelism(or Parallel 처리)는 비슷하게 들리지만, 전혀 다른 개념임

  -> In programming, concurrency is the composition of independently executing processes, while parallelism is the simultaneous execution of (possibly related) computations. Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.

  (프로그래밍에서 동시성은 독립적으로 실행되는 프로세스의 구성이고 병렬성은 (아마 관련이있는) 계산의 동시 실행입니다. 동시성은 한 번에 많은 것을 처리하는 것입니다. 병렬 처리는 한 번에 많은 작업을 수행하는 것입니다)





[참고] [http://golang.site/go/article/20-Go-defer%EC%99%80-panic]