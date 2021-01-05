package part

import (
	"fmt"
	"os"
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

func PanicFunc() {
	openFile("Invalid.txt") // 잘못된 파일명 입력
	println("Done")         // openFile() 안에서 panic이 실행되면 아래 print는 실행 안됨
	//(logs) panic: open Invalid.txt: no such file or directory
	// goroutine 1 [running]:study/part.openFile(0x10d49a3, 0xb)
	// ...
}

func openFile(fn string) {
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
