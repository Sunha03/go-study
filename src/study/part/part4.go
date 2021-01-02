package part

import (
	"fmt"
	"study/testlib"
	//_ "study/testlib" //패키지의 init()만 호출
)

func Packages() {
	song := testlib.GetMusic("Jogn Legend")
	fmt.Println(song)
	//(Outputs) All of Me
}
