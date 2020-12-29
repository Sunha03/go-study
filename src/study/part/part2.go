package part

import "fmt"

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
