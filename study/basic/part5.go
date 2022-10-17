package part

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func DeferFunc() {
	f, err := os.Open("1.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close() // main() 마지막에 파일 close 실행

	bytes := make([]byte, 1024) // 파일 읽기
	f.Read(bytes)
	println(len(bytes))
	//(Outputs) 1024
}

func openFile(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}

func PanicFunc() {
	openFile("Invalid.txt") // 잘못된 파일명 입력
	println("Done")         // openFile() 안에서 panic이 실행되면 아래 print는 실행 안됨
	//(logs) panic: open Invalid.txt: no such file or directory
	// goroutine 1 [running]:study/part.openFile(0x10d49a3, 0xb)
	// ...
}

func openFile2(fn string) {
	defer func() { // defer() : panic 호출 시 실행됨
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

func RecoverFunc() {
	openFile2("Invalid.txt") // 잘못된 파일명 입력
	println("Done")          // recover()에 의해 실행됨
	//(Outputs) Done
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, "***", i)
	}
}

func Goroutines() {
	say("Sync") // 함수를 동기적으로 실행

	go say("Async1") // 함수를 비동기적으로 실행
	go say("Async2")
	go say("Async3")

	time.Sleep(time.Second * 3) // 3초 대기
	//(Outputs) Sync *** 0 / Sync *** 1 / Sync *** 2
	// / Async1 *** 0 / Async1 *** 1 / Async1 *** 2
	// / Async3 *** 0 / Async3 *** 1 / Async3 *** 2
	// / Async2 *** 0 / Async2 *** 1 / Async2 *** 2
}

func AnonymousGoroutineFunc() {
	var wait sync.WaitGroup // WaitGroup 생성
	wait.Add(2)             // 2개의 goroutine을 기다림

	go func() { // 익명함수를 사용한 goroutine
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	go func(msg string) { // 익명함수에 파라미터 전달
		defer wait.Done() // 끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() // goroutine 모두 끝날때까지 대기
	//(Outputs) Hello
	// Hi
}

func sendChan(ch chan<- string) {
	ch <- "Data" // x := <-ch  -> 에러 발생
}

func receiveChan(ch <-chan string) {
	data := <-ch
	fmt.Println(data)
}

func run1(done chan bool) {
	time.Sleep(1 * time.Second)
	done <- true
}

func run2(done chan bool) {
	time.Sleep(2 * time.Second)
	done <- true
}

func GoChannels() {
	ch := make(chan int) // 정수형 채널을 생성함

	go func() {
		ch <- 123 // 채널에 123을 보냄
	}()

	var i int
	i = <-ch // 채널로부터 123을 받음
	println(i)
	//(Outputs) 123

	done := make(chan bool) // Go 채널은 goroutine이 끝날 때까지 기다릴 수 있음
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	<-done // 위의 goroutine이 끝날 때까지 대기함
	//(Outputs) 0 / 1 / 2 / 3 / 4

	/*c := make(chan int) // Go 채널 버퍼링
	c <- 1              // 수신 루틴이 없음 -> deadlock
	fmt.Println(<-c)*/ // 코멘트해도 deadlock (별도의 goroutine이 없기 때문)
	//(logs) fatal error: all goroutines are asleep - deadlock!

	cha := make(chan int, 1)
	cha <- 101 // 수신자가 없어도 보낼 수 있음
	fmt.Println(<-cha)
	//(Outputs) 101

	ch1 := make(chan string, 1) // 채널 파라미터
	sendChan(ch1)
	receiveChan(ch1)
	//(Outputs) Data

	ch2 := make(chan int, 2) // 채널 닫기
	ch2 <- 1                 // 채널에 송신
	ch2 <- 2
	close(ch2) // 채널 닫기 : close()

	println(<-ch2) // 채널 수신
	println(<-ch2)
	if _, success := <-ch2; !success {
		println("더이상 데이터 없음")
	}
	//(Outputs) 1 / 2 / 더이상 데이터 없음

	ch3 := make(chan int, 2) // 채널 range문
	ch3 <- 1                 // 채널에 송신
	ch3 <- 2
	close(ch3)

	// 방법1
	/*for {		// 채널이 닫힌 것을 감지할 때까지 계속 수신
		if i, success := <-ch3; success {
			println(i)
		} else {
			break
		}
	}*/

	// 방법2(range문)
	for i := range ch3 {
		println(i)
	}
	//(Outputs) 1 / 2

	done1 := make(chan bool) // 채널 select문
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
	//(Outputs) run1 완료 / run2 완료
}
