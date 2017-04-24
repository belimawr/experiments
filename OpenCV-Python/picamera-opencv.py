import io

import picamera
import cv2

import numpy as np

# (480, 848, 3)

with picamera.PiCamera() as camera:
	camera.resolution = (640, 480) # (480, 848)
	camera.framerate = 24
	stream = io.BytesIO()
	
	while True:
		camera.capture(stream, format="jpeg", use_video_port=True)
		frame = np.fromstring(stream.getvalue(), dtype=np.uint8)
		stream.seek(0)
		frame = cv2.imdecode(frame, 1)
		cv2.imshow('Imaet', frame)
		if cv2.waitKey(1) & 0xFF == ord('q'):
			break

cv2.destroyAllWindows()
