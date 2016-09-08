package main

import (
	"fmt"
	"test_lua/goluamodule"

	lua "github.com/yuin/gopher-lua"
)

func Double(L *lua.LState) int {
	lv := L.ToInt(1)            /* get argument */
	L.Push(lua.LNumber(lv * 2)) /* push result */
	return 1                    /* number of results */
}

func Sum(L *lua.LState) int {
	n0 := L.ToInt64(1)
	n1 := L.ToInt64(2)
	L.Push(lua.LNumber(n0 + n1))
	return 1
}

func main() {
	L := lua.NewState()
	defer L.Close()

	if err := L.DoString(`print("Hello World!")`); err != nil {
		panic(err)
	}

	var myVar lua.LString

	myVar = "MyVar is a string"
	L.Push(myVar)
	lv := L.Get(-1) // get the value at the top of the stack
	if str, ok := lv.(lua.LString); ok {
		// lv is LString
		fmt.Println(string(str))
	}
	if lv.Type() != lua.LTString {
		panic("string required.")
	}

	L.SetGlobal("double", L.NewFunction(Double)) /* Original lua_setglobal uses stack... */
	L.SetGlobal("sum", L.NewFunction(Sum))

	if err := L.DoString(`print(double(100))`); err != nil {
		panic(err)
	}

	if err := L.DoString(`print(sum(100, 11))`); err != nil {
		panic(err)
	}

	L.PreloadModule("mymodule", goluamodule.Loader)
	if err := L.DoFile("main.lua"); err != nil {
		panic(err)
	}
}
