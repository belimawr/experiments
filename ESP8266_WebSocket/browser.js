var ws = new WebSocket("ws://192.168.4.1:8080/")

ws.onerror = function (event) {console.log('ERROR');console.log(event.data)}

ws.onmessage = function (event) {console.log(event.data)}

ws.onclose = function (event) {console.log('CLOSE');console.log(event.data)}
