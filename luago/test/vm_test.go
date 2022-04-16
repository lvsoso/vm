package test

import (
	"fmt"
	"io/ioutil"
	"luago/binchunk"
	"luago/state"
	"testing"

	. "luago/vm"
)

// go test -timeout 30s -run ^TestVM$ luago/test
func TestVM(t *testing.T) {
	data, err := ioutil.ReadFile("../lua/call.out")
	if err != nil {
		panic(err)
	}
	// proto := binchunk.Undump(data)
	// luaMain(proto)
	ls := state.New()
	ls.Load(data, "../lua/call.out", "b")
	ls.Call(0, 0)
}

func luaMain(proto *binchunk.Prototype) {
	nRegs := int(proto.MaxStackSize)
	// ls := state.New(nRegs+8, proto)
	ls := state.New()
	ls.SetTop(nRegs)
	for {
		pc := ls.PC()
		inst := Instruction(ls.Fetch())
		if inst.Opcode() != OP_RETURN {
			inst.Execute(ls)
			fmt.Printf("[%02d] %s ", pc+1, inst.OpName())
			printStack(ls)
		} else {
			break
		}
	}
}
