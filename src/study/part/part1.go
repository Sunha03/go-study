package part

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

func ConditionalStatements() {
	k := 1
	if k == 1 { //if statement
		fmt.Println("One")
	} else if k == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("Other")
	}
	//(Outputs) One

	i := 1
	max := 10
	if val := i * 2; val < max {
		fmt.Println(val)
	}
	//var++		//ERROR - out of if statement
	//(Outputs) 2

	var name string
	var category = 1
	switch category { //switch statement
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	fmt.Println(name)
	//(Outputs) Paper Book

	fmt.Print("x is ")
	switch x := category << 2; x - 1 {
	case 3:
		fmt.Println("Three")
	}
	//(Outputs) x is Three

	score := 88
	switch { //no default fall through
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("No Hope")
	}
	//(Outputs) B

	val := 3
	switch val {
	case 1:
		fmt.Println("1 or less")
		fallthrough
	case 2:
		fmt.Println("2 or less")
		fallthrough
	case 3:
		fmt.Println("3 or less")
		fallthrough
	default:
		fmt.Println("dafault")
	}
	//(Outputs) 3 or less
	// default

	n := "str"
	switch v := interface{}(n).(type) { //type switch
	case int:
		fmt.Println("int :", v)
	case bool:
		fmt.Println("bool :", v)
	case string:
		fmt.Println("string :", v)
	default:
		fmt.Println("unknown")
	}
	//(Outputs) string : str
}

func LoopStatements() {
	sum := 0 //for statements
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//(Outputs) 5050

	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n)
	//(Outputs) 128

	/*for { //Infinite loop
		fmt.Println("Infinite loop")
	}*/

	names := []string{"James", "John", "Amy"}
	for index, name := range names {
		fmt.Println(index, name)
	}
	//(Outputs) 0 James
	// 1 John
	// 2 Amy

	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	fmt.Println(a)
END:
	fmt.Println("The End")
	//(Outputs) The End

	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	fmt.Println("OK")
	//(Outputs) OK
}
