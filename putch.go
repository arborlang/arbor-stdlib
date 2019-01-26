package main

import (
	// "fmt"
	"github.com/radding/arbor-dev"
	"strconv"
)

//Putch just puts a thing to the
func Putch(vm *arbor.VM) int64 {
	ptr := rune(vm.Life.GetCurrentFrame().Locals[0])
	// fmt.Printf("%X %d\n", ptr, ptr)
	log(string(ptr))
	return 0
}

//PutNum just puts a thing to the
func PutNum(vm *arbor.VM) int64 {
	ptr := vm.Life.GetCurrentFrame().Locals[0]
	log(strconv.FormatInt(ptr, 64))
	return 0
}
