import uhashlib
import ubinascii

s = socket.socket()
s.bind(('0.0.0.0', 4242))
s.listen(1)
conn, ip = s.accept()

l = b''

# Print all headers
while l != b'\r\n':
    l = conn.readline()
    print(l)


# Example of how to calculate Sec-WebSocket-Accept
# Based on Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==
hash = uhashlib.sha1('VZGaO6XKTfVnMQsALKnikw==258EAFA5-E914-47DA-95CA-C5AB0DC85B11')
sec_accept = ubinascii.b2a_base64(hash.digest())

