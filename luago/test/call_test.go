package test

import (
	"io/ioutil"
	"luago/state"
	"testing"
)

func TestCall(t *testing.T) {
	data, err := ioutil.ReadFile("../lua/call.out")
	if err != nil {
		panic(err)
	}
	ls := state.New()
	ls.Load(data, "../lua/call.out", "b")
	ls.Call(0, 0)
}
