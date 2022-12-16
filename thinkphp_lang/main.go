package main

import (
	"thinkphp_lang/Check"
	"thinkphp_lang/Common"
)

func main() {
	var Cmd_instruction Common.Cmd
	//fmt.Println(&Cmd_instruction)
	Common.Flag(&Cmd_instruction)
	Check.Check_url(Cmd_instruction)

}
