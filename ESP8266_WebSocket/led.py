import socket
import websocket
import websocket_helper
import time
from machine import Pin
from neopixel import NeoPixel


class Led():
    def __init__(self):
        self.np = NeoPixel(Pin(4), 16)
        self.s = None
        self.ws = None
        self.client = None

    def connect(self):
        self.s = socket.socket()
        self.s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.s.bind(('0.0.0.0', 4242))
        self.s.listen(1)
        self.client, ip = self.s.accept()
        self.client.setblocking(False)
        self.client.setsockopt(socket.SOL_SOCKET, 20, self.read_data)
        websocket_helper.server_handshake(self.client)
        self.ws = websocket.websocket(self.client, True)
        self.ws.write('Connected')

    def read_data(self, *args, **kwargs):
        msg = self.ws.read()
        print(msg)
        if not msg:
            return 'There is no message to read'
        split = msg.decode('utf-8').split(':')
        colours = []
        for i in split[1].split(','):
            colours.append(int(i))

        for i in range(16):
            self.np[i] = colours
            time.sleep(0.1)
            self.np.write()
            print('.')
        return 'It worked!'
