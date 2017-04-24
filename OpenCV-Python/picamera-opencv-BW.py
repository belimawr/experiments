#Sources: http://picamera.readthedocs.io/en/release-1.6/recipes1.html#capturing-to-an-opencv-object
#         https://raspberrypi.stackexchange.com/q/22241

import io

import picamera
import cv2

import numpy as np

with picamera.PiCamera() as camera:
	camera.resolution = (640, 480)
	camera.framerate = 24
	stream = io.BytesIO()
	
	while True:
		camera.capture(stream, format="jpeg", use_video_port=True)
		frame = np.fromstring(stream.getvalue(), dtype=np.uint8)
		stream.seek(0)
		frame = cv2.imdecode(frame, cv2.IMREAD_GRAYSCALE)
		frame = cv2.GaussianBlur(frame, (5, 5), 0)
		frame = cv2.threshold(frame, 0, 255, cv2.THRESH_BINARY | cv2.THRESH_OTSU)[1]

		cv2.imshow('Image', frame)
		if cv2.waitKey(1) & 0xFF == ord('q'):
			break

cv2.destroyAllWindows()
