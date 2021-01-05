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
