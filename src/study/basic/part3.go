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

func Slices() {
	var s []int        //slice 선언
	s = []int{1, 2, 3} //리터럴 값 지정
	s[1] = 10
	fmt.Println(s)
	//(Outputs) [1 10 3]

	s2 := make([]int, 5, 10) //make() 사용하여 slice todtjd
	fmt.Println(len(s2), cap(s2))
	//(Outputs) 5, 10

	var nilSlice []int //nil slice
	fmt.Println(len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil slice = nil")
	}
	//(Outputs) 0 0
	// nil slice = nill

	oriSlice := []int{0, 1, 2, 3, 4, 5, 6}
	subSlice := oriSlice[2:5] //sub slice
	fmt.Println(subSlice)
	//(Outputs) [2 3 4]

	sub1 := oriSlice[3:]
	sub2 := oriSlice[:2]
	sub3 := oriSlice[:]
	fmt.Println(sub1, sub2, sub3)
	//(Outputs) [3 4 5 6] [0 1] [0 1 2 3 4 5 6]

	sa := []int{0, 1} //append() : 항목 추가
	sa = append(sa, 2)
	sa = append(sa, 3, 4, 5)
	fmt.Println(sa)
	//(Outputs) [0 1 2 3 4 5]

	sliceA := make([]int, 0, 3) //len:0, cap:3
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		fmt.Println(len(sliceA), cap(sliceA))
	}
	fmt.Println(sliceA)
	//(Outputs) 1 3 / 2 3 / 3 3 / 4 6 / 5 6 / 6 6 / 7 12 / 8 12
	// 9 12 / 10 12 / 11 12 / 12 12 / 13 24 / 14 24 / 15 24
	//(Outputs) [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15]

	sliceB := []int{1, 2, 3} //다중 슬라이스 병합
	sliceC := []int{4, 5, 6}
	sliceX := append(sliceB, sliceC...)
	fmt.Println(sliceX)
	//(Outputs) [1 2 3 4 5 6]

	source := []int{0, 1, 2} //copy() : 슬라이스 복사
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target, len(target), cap(target))
	//(Outputs)[0 1 2] 3 6
}

func Maps() {
	var idMap map[int]string      //map 선언
	idMap = make(map[int]string)  //초기화
	tickers := map[string]string{ //리터럴로 초기화
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
		"AMZN": "Amazon",
	}
	fmt.Println(idMap, tickers)
	//(Outputs) map[] map[FB:FaceBook GOOG:Google Inc MSFT:Microsoft]

	var m map[int]string
	m = make(map[int]string)
	m[901] = "Apple" //값 할당
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	noData := m[999] //값 없으면 -> nil or zero 리턴
	delete(m, 777)   //값 삭제
	fmt.Println(str, noData, m[777])
	//(Outputs) Grape

	val, exists := tickers["MSFT"] //map key 체크
	if !exists {
		fmt.Println("No MSFT ticker")
	}
	fmt.Println(val, exists)
	//(Outputs) Microsoft true

	for key, val := range tickers { //map 열거
		fmt.Println(key, val)
	}
	//(Outputs) GOOG Google Inc
	// MSFT Microsoft
	// FB FaceBook
	// AMZN Amazon
}
