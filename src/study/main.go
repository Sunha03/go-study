package main

import "fmt"

func Variables() {
	var num int               //int형 변수 선언
	var floatNum float32 = 11 //float32형 변수 선언&초기화
	num = 10                  //변수 선언 후 값 할당
	floatNum = 12.0

	//var doNotUse string		//사용하지 않는 변수 > error

	var i, j, k int = 1, 2, 3 //복수 변수 선언&초기화

	var one = 1 //변수 타입 자동 추론
	str := "Hi" //Short Assignment Statement(:=)

	fmt.Println(num, floatNum, i, j, k, one, str)
	//(Outputs) 10 12 1 2 3 1 Hi
}

func Consts() {
	const num int = 10      //int형 상수 선언&초기화
	const str string = "Hi" //str형 상수 선언&초기화

	const num2 = 20 //변수 타입 자동 추론
	const str2 = "Hello"

	const ( //여러개 상수들의 나열
		Visa   = "Visa"
		Master = "Master Card"
		Amex   = "American Express"
	)

	const ( //identifier "iota"
		zero = iota //0 - 1씩 증가
		one         //1
		two         //2
	)

	fmt.Println(num, str, num2, str2, Visa, Master, Amex, zero, one, two)
	//(Outputs) 10 Hi 20 Hello Visa Master Card American Express 0 1 2
}

func main() {
	Variables()
	Consts()
}
