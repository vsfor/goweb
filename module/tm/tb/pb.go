package tb

import (
	"fmt"
)

func BSayHi() {
	fmt.Println("say hi from package B")
}

var ASayHi func()
func RegisterASayHi(f func()) {
	ASayHi = f
}

func CallASayHi() {
	ASayHi()
}


