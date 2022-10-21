package ta

import (
	"fmt"
	"vsfor/goweb/module/tm/tb"
)

func ASayHi() {
	fmt.Println("say hi from package A")
}

func CallBSayHi() {
	tb.BSayHi()
}

func init() {
	tb.RegisterASayHi(ASayHi)
}