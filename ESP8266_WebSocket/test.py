import asyncio
import websockets

@asyncio.coroutine
def hello(websocket, path):
    websocket.send("HTTP/1.0 200 \n\n")
    name = yield from websocket.recv()
    print("< {}".format(name))

    greeting = "Hello {}!".format(name)
    yield from websocket.send(greeting)
    print("> {}".format(greeting))

start_server = websockets.serve(hello, '0.0.0.0', 8080)


asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
