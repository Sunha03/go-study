package part

import "fmt"

func Arrays() {
	var numbers [3]int //int형 변수 3개를 갖는 배열
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println(numbers[1])
	//(Outputs) 2

	var arr1 = [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3, 4} //배열크기 자동지정
	fmt.Println(arr1, arr2)
	//(Outputs) [1 2 3] [1 2 3 4]

	var multiArray [3][4][5]int //정의
	multiArray[0][1][2] = 10    //사용
	fmt.Println(multiArray[0][1][2])
	//(Outputs) 10

	var mArr = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(mArr)
	//(Outputs) [[1 2 3] [4 5 6]]
}
