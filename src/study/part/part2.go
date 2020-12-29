package part

import "fmt"

type calculator func(int, int) int //type문 - 원형 정의

func say_val(msg string) { //pass by value
	fmt.Println(msg)
}

func say_ref(msg *string) { //pass by reference
	*msg = "Hi"
	fmt.Println(msg)
}

func say_variadic(msg ...string) { //variadic function
	for _, s := range msg {
		fmt.Print(s)
	}
	fmt.Println("")
}

func sum(nums ...int) (int, int) { //다중 return
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return count, s
}

func sum_namedReturn(nums ...int) (count int, total int) { //Named Return Parameter
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}

func Functions() {
	msg := "Hello"
	say_val(msg)
	fmt.Println("pass by value :", msg)
	say_ref(&msg)
	fmt.Println("pass by reference", msg)
	//(Outputs) Hello
	// pass by value : Hello
	// 0xc00008e270

	say_variadic("A", "B", "CD", "EFG")
	//(Outputs) ABCDEFG

	count, total := sum(1, 3, 5, 7)
	fmt.Println(count, total)
	//(Outputs) 4 16

	count2, total2 := sum_namedReturn(2, 4, 6, 8)
	fmt.Println(count2, total2)
	//(Outputs) 4, 20
}

func calc(f func(int, int) int, a int, b int) int { //일급함수
	//func calc(f calculator, a int, b int) int {	//type문 - calculator 원형 사용
	result := f(a, b)
	return result
}

func Functions2() {
	sum := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) //익명함수 호출
	fmt.Println(result)
	//(Outputs) 15

	add := func(i int, j int) int { //익명함수 할당
		return i + j
	}

	r1 := calc(add, 10, 20) //add 함수 전달
	fmt.Println(r1)
	//(Outputs) 30

	//파라미터에 익명함수를 정의
	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	fmt.Println(r2)
	//(Outputs) -10
}
