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

func DataTypes() {
	rawLiteral := `Hello\nHi` //Raw String Literal

	interLiteral := "Hi\nHello" //Interpreted String Literal
	interLiteral2 := "Hi\n" + "Hello"

	fmt.Println(rawLiteral)
	//(Outputs) Hello\nHi
	fmt.Println(interLiteral)
	//(Outputs) Hi
	// Hello
	fmt.Println(interLiteral2)
	//(Outputs) Hi
	// Hello

	var i int = 100            //Type Conversion
	var u uint = uint(i)       //int -> uint
	var f float32 = float32(i) //uint -> float32

	str := "ABC"
	bytes := []byte(str)  //string -> byte arrays
	str2 := string(bytes) //byte arrays -> string

	println(f, u)
	//(Outputs) +1.000000e+002 100
	println(bytes, str2)
	//(Outputs) [3/32]0xc00006ef08 ABC
}

func Operators() {
	a := 5
	b := 10
	var c int
	i := 0

	c = (a + b) / 5 //산술연산자
	i++
	fmt.Println(a, b, c, i)
	//(Outputs)5 10 3 1

	fmt.Println(a == b, a != c, a >= b) //관계연산자
	//(Outputs) false true false

	boolA := true //논리연산자
	boolB := false
	boolC := true
	fmt.Println(boolA && boolB, boolA || !(boolC && boolB))
	//(Outputs) false true

	a = 100
	fmt.Println(a) //(Outputs) 100
	a *= 10
	fmt.Println(a) //(Outputs) 1000
	a >>= 2
	fmt.Println(a) //(Outputs) 250
	a |= 2
	fmt.Println(a) //(Outputs) 250

	var k int = 10
	var p = &k
	fmt.Println(*p)
	//(Outputs) 10
}

func main() {
	Variables()
	Consts()
	DataTypes()
	Operators()
}
