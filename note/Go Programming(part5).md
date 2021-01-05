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



## 5. Go 채널

### 1) Go 채널

* Go 채널 : 채널을 통해 데이터를 주고 받을 수 있는 통로

* 채널은 make()를 통해 미리 생성되어야 하며, 채널 연산자 '<-'을 통해 데이터를 보내고 받음

* 흔히 goroutine들 사이에서 데이터를 주고 받는데 사용되는데, 상대편이 준비될 때까지 채널에서 대기함으로써 별도의 lock을 걸지 않고 데이터를 동기화하는데 사용됨

  ```go
  func main() {
  	ch := make(chan int)	// 정수형 채널을 생성함
  	
  	go func() {
  		ch <- 123		// 채널에 123을 보냄
  	}()
  	
  	var i int
  	i = <- ch		// 채널로부터 123을 받음
  	println(i)
  }
  ```

  -> 실행과정 : 정수형 채널 생성 > 한 goroutine에서 그 채널에 123이란 정수 데이터를 보냄 > 이를 다시 메인 루틴에서 채널로부터 123 데이터를 받음

  -> 채널 생성 시, make()에 어떤 타입의 데이터를 채널에서 주고 받을지 미리 지정해 주어야 함

  -> 채널로 데이터 보낼 때 : (채널명) <- (데이터), 채널에서 데이터 받을 때 : <- (채널명)

  -> 위의 예제에서 메인 루틴은 마지막에서 채널로부터 데이터를 받고 있는데, 상대편 goroutine에서 데이터를 전송할 때까지는 계속 대기하게 됨. 따라서 이 예제에서는 time.Sleep() or fmt.Scanf() 같이 goroutine이 끝날 때까지 기다리는 코드는 필요 없음.

* Go 채널은 수신자와 송신자가 서로를 기다리는 속성 때문에, 이를 이용하여 goroutine이 끝날 때까지 기다리는 기능을 구현할 수 있음

  -> 즉, 익명함수를 사용한 한 Go 루틴에서 어떤 작업이 실행되고 있을 때, 메인 루틴은 '<-done'에서 계속 수신하며 대기하고 있게 됨

  -> 익명함수 Go 루틴에서 작업이 끝난 후, done 채널에 true를 보내면 메인 루틴은 이를 받고 프로그램을 끝냄

  ```go
  import "fmt"
  
  func main() {
  	done := make(chan bool)
  	go func() {
  		for i := 0; i< 10; i++ {
  			fmt.Println(i)
  		}
  		done <- true
  	}()
  	<-done		// 위의 goroutine이 끝날 때까지 대기함
  }
  ```

### 2) Go 채널 버퍼링

* Go 채널 2가지 - Unbuffered Channel, Buffered Channel

  * Unbuffered Channel - 이 채널에서는 하나의 수신자가 데이터를 받을 때까지 송신자가 데이터를 보내는 채널에 묶여 있게 됨
  * Buffered Channel - 비록 수신자가 받을 준비가 되어 있지 않을지라도 지정된 버퍼만큼 데이터를 보내고 계속 다른 일을 수행할 수 있음

* 버퍼 채널은 make(chan type, N)을 통해 생성되는데, 2번째 파라미터 N에 사용할 버퍼 개수를 넣음

  ex) make(chan int, 10) -> 10개의 정수형을 갖는 버퍼 채널을 만듦

* 버퍼 채널을 이용하지 않는 경우, 아래와 같은 코드는 에러를 발생시킴(fatal error: all goroutines are asleep-deadlock!)

  -> 왜냐하면 메인 루틴에서 채널에 1을 보내면서 상대편 수신자를 기다리고 있는데, 이 채널을 받는 수신자 goroutine이 없기 때문임

  ```go
  import "fmt"
  
  func main() {
  	c := make(chan int)
  	c <- 1		// 수신 루틴이 없음 -> deadlock
  	fmt.Println(<-c)		// 코멘트해도 deadlock (별도의 goroutine이 없기 대문)
  }
  ```

* 하지만, 아래와 같이 버퍼 채널을 사용하면, 수신자가 당장 없더라도 최대 버퍼 수까지 데이터를 보낼 수 있음

  ```go
  import "fmt"
  
  func main() {
  	ch := make(chan int, 1)
  	
  	ch <- 101		// 수신자가 없어도 보낼 수 있음
  	fmt.Println(<-ch)
  }
  ```

### 3) 채널 파라미터

* 함수의 파라미터로 채널을 전달할 대, 일반적으로 송수신을 모두 하는 채널을 전달하지만, 특별히 해당 채널로 송신만 할 것인지 or 수신만 할 것인지를 지정할 수도 있음

* 송산 파라미터 : chan<-, 수신 파라미터 : <-chan

  ex) (p chan<-int), (p <-chan int)

* 만약 송신 채널 파라미터에서 수신을 하거나, 수신 채널에서 송신을 하면 에러 발생함

  ```go
  import "fmt"
  
  func main() {
  	ch := make(chan string, 1)
  	sendChan(ch)
  	receiveChan(ch)
  }
  
  func sendChan(ch chan<- string) {
  	ch <- "Data"		// x := <-ch  -> 에러 발생
  }
  
  func receiveChan(ch <-chan string) {
  	data := <-ch
  	fmt.Println(data)
  }
  ```

  -> 만약 sendChan() 안에서 x := <-ch를 실행하면 송신전용 채널에서 수신을 시도하므로 에러 발생함

### 4) 채널 닫기

* 채널 오픈 & 데이터 수신 후, close()를 사용하여 채널을 닫을 수 있음

* 채널을 닫으면, 해당 채널로는 더이상 송신을 할 수 없지만, 채널이 닫힌 이후에도 계속 수신은 가능

* 채널 수신에 사용되는 '<-ch'은 1개의 리턴 값을 갖는데, 1번째는 채널 메세지, 2번째는 수신이 제대로 되었는지를 나타냄(만약 채널이 닫히면 해당 값은 false임)

  ```go
  func main() {
  	ch := make(chan int, 2)
  	
  	ch <- 1		// 채널에 송신
  	ch <- 2
  	close(ch)
  	
  	println(<-ch)		// 채널 수신
  	println(<-ch)
  	
  	if _, success := <-ch; !success {
  		println("더이상 데이터 없음")
  	}
  }
  ```

### 5) 채널 range문

* 채널에서 송신자가 송신한 후 채널을 닫을 수 있음. 그리고 수신자는 임의의 개수의 데이터를 채널이 닫힐 때까지 계속 수신할 수 있음

  -> 수신자는 채널이 닫히는 것을 체크하며 계속 루프를 될게 됨

  ```go
  func main() {
  	ch := make(chan int, 2)
  	ch <- 1		// 채널에 송신
  	ch <- 2
  	close(ch)
  	
  	// 방법1
  	for {		// 채널이 닫힌 것을 감지할 때까지 계속 수신
  		if i, success := <-ch; success {
  			println(i)
  		} else {
  			break
  		}
  	}
  	
  	// 방법2(range문)
  	for i := range ch {
  		println(i)
  	}
  }
  ```

### 6) 채널 select문

* Go select문 : 복수 채널들을 기다리면서 준비된 (데이터를 보내온) 채널을 실행하는 기능을 제공함

* select문은 여러 개의 case문에서 각각 다른 채널을 기다리다가 준비가 된 채널 case를 실행함

  -> select문은 case 채널들이 준비되지 않으면 계속 대기하게 되고, 가장 먼저 도착한 채널의 case를 실행함

  -> 만약 복수 채널에 신호가 오면, Go 런타임이 랜덤하게 그 중 1개를 선택함

  -> 하지만, select문에 default문이 있으면, case문 채널이 준비되지 않더라도 계속 대기하지 않고 바로 default문을 실행함

  ```go
  import "time"
  
  func main() {
  	done1 := make(chan bool)
  	done2 := make(chan bool)
  	go run1(done1)
  	go run2(done2)
  	
  EXIT:
  	for {
  		select {
  		case <-done1:
  			println("run1 완료")
  		case <-done2:
  			println("run2 완료")
  			break EXIT
  		}
  	}
  }
  
  func run1(done chan bool) {
  	time.Sleep(1 * time.Second)
  	done <- true
  }
  	
  func run2(done chan bool) {
  	time.Sleep(2 * time.Second)
  	done <- true
  }
  ```

  -> for문 안에 select문을 쓰며 2개의 goroutine이 모두 실행되기를 기다리고 있음

  -> 1번째 run1()을 1초간 실행 + done1 채널로부터 수신하여 해당 case 실행 + 다시 for문을 돎

  -> 다시 for문을 돌며 다시 select문이 실행되는데, 다음 run2()가 2초 후 실행 + done2 채널로부터 수신하여 해당 case 실행

  -> done2 채널 case문에 break EXIT이 있는데, 이 문장으로 인해 for문을 빠져나와 EXIT 레이블로 이동함

* Go의 "break 레이블" 문 - 해당 레이블로 이동한 후 자신이 빠져나온 루프 다음 문장을 실행함(C/C# 등의 언어에서의 goto 문과 다름)



[참고] [http://golang.site/go/article/20-Go-defer%EC%99%80-panic]