package say

import "fmt"

//sayHi func
// - export 가능한 function -> 대문자로 시작(소문자는 불가)
// - private function임
func sayHi() {
	fmt.Println("Hi")
}

//SayHello func
// - export 가능한 function -> 대문자로 시작
func SayHello() {
	fmt.Println("Hello")
}
