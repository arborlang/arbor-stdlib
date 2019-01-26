package main

import arbor "github.com/arborlang/arbor-dev"

//Env is the environ
var Env = arbor.Resolver{
	ModuleName: "env",
	Execers:    map[string]arbor.Extension{},
}

func init() {
	Env.Register("__putch__", arbor.ExtensionFunc(Putch))
	Env.Register("__break__", arbor.ExtensionFunc(Breakpoint))
	Env.Register("__alloc__", arbor.ExtensionFunc(AllocateOnHeap))
}

func main() {}
