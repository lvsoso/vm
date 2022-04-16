package test

import (
	"io/ioutil"
	"luago/binchunk"
	"testing"
)

func TestTable(t *testing.T) {
	data, err := ioutil.ReadFile("../lua/test.out")
	if err != nil {
		panic(err)
	}

	proto := binchunk.Undump(data)
	luaMain(proto)
}
