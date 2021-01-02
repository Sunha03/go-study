package part

import (
	"fmt"
	"study/testlib"
	//_ "study/testlib" // 패키지의 init()만 호출
)

type person struct { // struct 정의
	name string
	age  int
}

type dick struct {
	data map[int]string
}

type Rect struct {
	width, height int
}

func Packages() {
	song := testlib.GetMusic("John Legend")
	fmt.Println(song)
	//(Outputs) All of Me
}

func newDict() *dick { // 생성자 함수 정의
	d := dick{}
	d.data = map[int]string{}
	return &d // 포인터 전달
}

func Structs() {
	p := person{}  // person 객체 생성
	p.name = "Lee" // 필드 값 설정
	p.age = 10
	fmt.Println(p)
	//(Outputs) {Lee 10}

	var p1 person // struct 생성 & 초기화
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	p3 := new(person)
	p3.name = "Lee"  //p3가 포인터여도 '.'사용 가능
	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(p1, p2, p3, dic)
	//(Outputs) {Bob 20} {Sean 50} &{Lee 0} &{map[1:A]}
}

func (r Rect) area() int { // Value receiver
	return r.width * r.height
}

func (r *Rect) area2() int { // Pointer receiver
	r.width++
	return r.width * r.height
}

func GoMethods() {
	rect := Rect{10, 20}
	area := rect.area() //메소드 호출
	fmt.Println(area)
	//(Outputs) 200

	rect2 := Rect{10, 20}
	area2 := rect.area2()
	fmt.Println(rect2.width, area2)
	//(Outputs) 11 220
}
