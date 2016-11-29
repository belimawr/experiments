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
# Sec-WebSocket-Accept must be: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=
# The "strip" is used because there is a '\n' at the end
# of the final string
hash = uhashlib.sha1('dGhlIHNhbXBsZSBub25jZQ==258EAFA5-E914-47DA-95CA-C5AB0DC85B11')
sec_accept = ubinascii.b2a_base64(hash.digest()).strip()

