package goluamodule

import (
	"bytes"
	"net/http"

	"github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	// register functions to the table
	mod := L.SetFuncs(L.NewTable(), exports)
	// register other stuff
	L.SetField(mod, "name", lua.LString("some value to attr name"))

	L.SetField(mod, "version", lua.LNumber(42))

	// returns the module
	L.Push(mod)
	return 1
}

var exports = map[string]lua.LGFunction{
	"myfunc":  myfunc,
	"tworet":  tworet,
	"httpGet": httpGet,
}

func myfunc(L *lua.LState) int {
	L.Push(lua.LString("Myfunc works!"))
	return 1
}

func tworet(L *lua.LState) int {
	L.Push(lua.LString("ret1"))
	L.Push(lua.LString("ret2"))
	return 2
}

func httpGet(L *lua.LState) int {
	url := L.ToString(1)

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	s := buf.String()

	if err != nil {
		panic(err)
	}

	L.Push(lua.LString(s))
	L.Push(lua.LNumber(response.StatusCode))

	return 2
}
