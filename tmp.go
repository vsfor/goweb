package main

import (
	"fmt"
	"vsfor/goweb/module/tm/ta"
	"vsfor/goweb/module/tm/tb"
)

func main() {
	fmt.Println("call bSayHi in a:")
	ta.CallBSayHi()

	fmt.Println("call aSayHi in b:")
	tb.CallASayHi()
}
