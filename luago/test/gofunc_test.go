package test

import (
	"fmt"
	"io/ioutil"
	. "luago/api"
	"luago/state"
	"testing"
)

func TestGoFunc(t *testing.T) {
	data, err := ioutil.ReadFile("../lua/hello_world.out")
	if err != nil {
		panic(err)
	}
	ls := state.New()
	ls.Register("print", print) // 注册print函数
	ls.Load(data, "chunk", "b")
	ls.Call(0, 0)
}

func print(ls LuaState) int {
	nArgs := ls.GetTop()
	for i := 1; i <= nArgs; i++ {
		if ls.IsBoolean(i) {
			fmt.Printf("%t", ls.ToBoolean(i))
		} else if ls.IsString(i) {
			fmt.Print(ls.ToString(i))
		} else {
			fmt.Print(ls.TypeName(ls.Type(i)))
		}
		if i < nArgs {
			fmt.Print("\t")
		}
	}
	fmt.Println()
	return 0
}
