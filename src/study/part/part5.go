package part

import "os"

func Defers() {
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
