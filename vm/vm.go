package vm

import (
	"fmt"

	"monkey/code"
	"monkey/compiler"
	"monkey/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	// "Stack pointer"; always points to the next
	// value. The top of the stack is `stack[sp-1]`.
	sp int
}

func New(bytecode *compiler.Bytecode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}

	return vm.stack[vm.sp-1]
}

// Runs the VM's fetch-decode-execute cycle.
func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		// Fetch
		op := code.Opcode(vm.instructions[ip])

		switch op {
		// Decode
		case code.OpConstant:
			// Execute
			constIndex := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Our VM is a stack-based virtual machine.
func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}
