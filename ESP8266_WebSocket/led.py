import socket
import websocket
import websocket_helper
import time
from machine import Pin
from neopixel import NeoPixel

    
np = NeoPixel(Pin(4), 24)

def connect():
    s = socket.socket()
    s.bind(('0.0.0.0', 4242))
    s.listen(1)
    cl, ip = s.accept()
    websocket_helper.server_handshake(cl)
    ws = websocket.websocket(cl, True)
    return ws


def recebe(ws):
    global np
    msg = ws.read()
    print(msg)
    if not msg:
        return 'not msg'
    split = msg.decode('utf-8').split(':')
    colours = []
    for i in split[1].split(','):
        colours.append(int(i))
        
    for i in range(24):
        np[i] = colours
        time.sleep(0.5)
        np.write()
    return 'ok'
