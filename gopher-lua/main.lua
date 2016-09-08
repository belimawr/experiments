local m = require("mymodule")

ret = m.myfunc()
a, b = m.tworet()

body, status = m.httpGet("https://github.com/yuin/gopher-lua")
print(ret)
print(m.name)
print(m.version)
print(a, b)
print(status)
print(string.sub(body, 0, 990))
