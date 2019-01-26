package main

import (
	"fmt"
	"github.com/radding/arbor-dev"
)

//AllocateOnHeap allocates some data on the Heap
func AllocateOnHeap(vm *arbor.VM) int64 {
	size := vm.Life.GetCurrentFrame().Locals[0]
	tp := int64(64)
	if vm.Life.GetCurrentFrame().Locals[1] == 1 {
		tp = int64(32)
	}
	lastNdx := int64(len(vm.Life.Memory) - 1)

	newMem := make([]byte, (size+1)*tp/8)
	vm.Life.Memory = append(vm.Life.Memory, newMem...)
	fmt.Println(vm.Life.Memory)
	newLocation := lastNdx + 1
	vm.Life.Memory[newLocation] = byte(size * tp / 8)
	return newLocation
}
