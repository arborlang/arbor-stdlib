package main

import (
	"bufio"
	"encoding/binary"

	// "bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	arbor "github.com/arborlang/arbor-dev"
	"github.com/perlin-network/life/exec"
)

// logs logs a formated message to the open stdout HANDLER
func log(msg string, args ...interface{}) {
	formatedMsg := fmt.Sprintf(msg, args...)
	os.Stdout.Write([]byte(formatedMsg))
}

func logLn(msg string, args ...interface{}) {
	log(msg+"\n", args...)
}

//Breakpoint inserts a break point for the run time
func Breakpoint(v *arbor.VM) int64 {
	vm := v.Life
	logLn("Execution paused for breakpoint")
	reader := bufio.NewReader(os.Stdin)

	for {
		log(">> ")
		exit := false
		// buff := bytes.NewBuffer())
		// _, err := fmt.Scanln(buff)
		// convert CRLF to LF
		cmd, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		cmd = strings.Replace(cmd, "\n", "", -1)
		// cmd := getInput()
		exploded := strings.Split(cmd, " ")
		switch exploded[0] {
		case "quit", "q":
			os.Exit(0)
		case "continue", "c", "next", "n":
			exit = true
		case "show":
			handleShow(exploded[1:], vm, v)
		case "access":
			handleMemAccess(exploded[1:], vm)
		default:
			logLn("unknown command: %s", cmd)
		}
		if exit {
			break
		}
	}
	return 0
}

func handleMemAccess(cmdParts []string, vm *exec.VirtualMachine) {
	num, err := strconv.Atoi(cmdParts[0])
	if err != nil {
		logLn("Expected second argument to be a number, got %q instead", cmdParts[0])
		return
	}
	data := vm.Memory[num]
	cmdParts = cmdParts[1:]
	log("VALUE AT %d: %X", num, data)
	if len(cmdParts) == 0 {
		return
	}
	if cmdParts[0] == "as" {
		cmdParts = cmdParts[1:]
		if len(cmdParts) == 0 {
			logLn("expected a type, but got nothing")
		}
		for _, part := range cmdParts {
			switch part {
			case "char":
				log(" AS CHAR: %q", rune(data))
			case "number":
				log(" AS NUMBER: %d", data)
			case "le":
				var le uint64
				binary.LittleEndian.PutUint64([]byte{data}, le)
				log(" AS LE: %d", le)
			default:
				logLn("unknown type %s", part)
			}
		}
	}
	log("\n")
}

func handleShow(cmdParts []string, vm *exec.VirtualMachine, v *arbor.VM) {
	switch part := cmdParts[0]; part {
	case "mem", "memory", "m":
		printMemory(cmdParts[1:], vm)
	case "stack", "s":
		PrintStack(vm)
	case "trace", "t":
		vm.PrintStackTrace()
	case "stacktop":
		logLn("STACK TOP: %d", v.StackTop)
	default:
		logLn("unknown command %d", part)
	}
}

//PrintMemory just puts a thing to the
func printMemory(parts []string, vm *exec.VirtualMachine) int64 {
	start := 0
	end := 20
	if len(parts) != 0 {
		var err error
		sliceString := parts[0]
		sliceParts := strings.Split(sliceString, ":")
		if len(sliceParts) == 2 {
			start, err = strconv.Atoi(sliceParts[0])
			if err != nil {
				log("malformed slice")
				return 0
			}
			sliceParts = sliceParts[1:]
		}
		end, err = strconv.Atoi(sliceParts[0])
		if err != nil {
			log("malformed slice")
			return 0
		}
	}

	logLn("BYTES OF MEMORY IN RANGE %d -> %d", start, end)
	for _, mem := range vm.Memory[start:end] {
		log("%X ", mem)
	}
	logLn("")
	return 0
}

//PrintStack just puts a thing to the
func PrintStack(vm *exec.VirtualMachine) int64 {
	frame := vm.GetCurrentFrame()
	logLn("CURRENT STACK:")
	for index, stackVal := range frame.Regs {
		logLn("%d: %v", index, stackVal)
	}
	return 0
}
