WebSocket using ESP8266 and JavaScript
======================================

This is a small experiment/study to understand better how WebSocket
works and also how it is implemented.

The code here was mostly developed at [Garoa Hacker clube](http://garoa.net.br/)

The source files are:
* **esp8266.py:** A small snippet that shows how to calculate the *Sec-WebSocket-Accept* header. This is not a fully working script yet. This code was developed to run on a ESP8266 with microPython.
* **test.py:** A example of how to use ``asyncio`` and ``websockets`` libraries to accept websocket connections. This was used to test the JavaScript/Browser code.
* **browser.js:** JavaScript code to create a websocket and connect to the server (ESP8266).
* **led.py:** A small snippet of how to connect, receive and parse messages using WebSocket. It assumes that there is a strip of 24 LEDs (WS2812) connected to pin 4.
* **led_control.html:** A still working in progress HMTL/JavaScript code to use WebSocket to control some WS2812

TODO
----
* Make **esp8266.py** a fully working script
* Write an script showing how to use microPython's ``websocket``

References:
* https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers
* https://github.com/micropython/micropython/blob/master/esp8266/modules/webrepl.py
