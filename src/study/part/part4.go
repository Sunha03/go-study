package part

import (
	"fmt"
	"math"
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

type Rectangle struct {
	width, height int
}

type Rect struct { // Rect 정의
	width, height float64
}

type Circle struct { // Circle 정의
	radius float64
}

type Shape interface {
	area() float64
	perimeter() float64
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

func (r Rectangle) area() int { // Value receiver
	return r.width * r.height
}

func (r *Rectangle) area2() int { // Pointer receiver
	r.width++
	return r.width * r.height
}

func GoMethods() {
	rect := Rectangle{10, 20}
	area := rect.area() //메소드 호출
	fmt.Println(area)
	//(Outputs) 200

	rect2 := Rectangle{10, 20}
	area2 := rect.area2()
	fmt.Println(rect2.width, area2)
	//(Outputs) 11 220
}

// Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // 인터페이스 메소드 호출
		fmt.Println(a)
	}
}

func printIt(v interface{}) {
	fmt.Println(v)
}

func GoInterfaces() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c) // interface Shape
	//(Outputs) 200
	// 314.1592653589793

	var x interface{} // Empty interface
	x = 1
	x = "Tom"
	printIt(x)
	//(Outputs) Tom

	var a interface{} = 1
	i := a       // a와 i는 dynamic type, 값은 1
	j := a.(int) // j는 int 타입, 값은 1
	println(i, j)
	//(Outputs) (0x10b7cc0,0x10f2be0)(포인터주소) 1
}
