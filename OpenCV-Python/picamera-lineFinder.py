import io
import math

import picamera
import cv2

import numpy as np

vertices = np.array([[0,0],[0,320], [640,320], [640,0]], np.int32)

def roi(image):
	mask = np.zeros_like(image)
	cv2.fillPoly(mask, [vertices], 255)
	masked = cv2.bitwise_and(image, mask)
	return masked

def process_image(image):
	img = cv2.Canny(image, threshold1=100, threshold2=200)
	img = cv2.GaussianBlur(img, (3,3), 0 )
	lines = cv2.HoughLinesP(img, 1, np.pi/180, 180, minLineLength=150, maxLineGap=10)
	return img, lines

def draw_lines(img, lines):
	try:
		for i, line in enumerate(lines):
			coords = line[0]
			if coords[1] >= 315:
				continue

			print(i, angle(coords))
			cv2.line(img, (coords[0],coords[1]), (coords[2],coords[3]), 255, 3)
	except TypeError as e:
		print('No Line :(')
	return

def angle(line):
	x1, y1, x2, y2 = line
	m = abs(y1 - y2)/abs(x1 - x2)
	return math.degrees(math.atan(m))

with picamera.PiCamera() as camera:
	camera.resolution = (640, 480)
	camera.framerate = 24
	stream = io.BytesIO()
	
	while True:
		camera.capture(stream, format="jpeg", use_video_port=True)
		frame = np.fromstring(stream.getvalue(), dtype=np.uint8)
		stream.seek(0)
		
		frame = cv2.imdecode(frame, cv2.IMREAD_GRAYSCALE)
		frame = roi(frame)
		
		processed_img, lines = process_image(frame)
		draw_lines(frame, lines)
		
		cv2.imshow('Image - ROI', frame)

		if cv2.waitKey(1) & 0xFF == ord('q'):
			break

cv2.destroyAllWindows()
