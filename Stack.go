package main

// import (
// 	"github.com/perlin-network/life/exec"
// )

// // MemStack is the stack manager for Arbor
// type MemStack struct {
// 	StackTop      int64
// 	StackPointers []int64
// }

// func (m *MemStack) PushStack(vm *exec.VM) int64 {
// 	m.StackPointers = append(m.StackPointers, m.StackTop)
// 	return m.StackTop
// }

// func (m *MemStack) PopStack(vm *exec.VM) int64 {
// 	last := len(m.StackPointers) - 1
// 	stackTop, newStack := m.StackPointers[last], m.StackPointers[:last]
// 	m.StackTop = stackTop
// 	m.StackPointers = newStack
// 	return m.StackTop
// }
