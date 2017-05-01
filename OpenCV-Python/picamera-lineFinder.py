#Sources: http://picamera.readthedocs.io/en/release-1.6/recipes1.html#capturing-to-an-opencv-object
#         https://raspberrypi.stackexchange.com/q/22241

import io

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
	#HoughLinesP(image, rho, theta, threshold[, lines[, minLineLength[, maxLineGap]]]) -> lines
	lines = cv2.HoughLinesP(img, 1, np.pi/180, 180, 20, 15)
	draw_lines(img, lines)
	return img

def draw_lines(img, lines):
	try:
		cv2.line(img, (0,0), (640,480), [0,0,255], 15)
		for line in lines:
			coords = line[0]
			print(coords)
			cv2.line(img, (coords[0],coords[1]), (coords[2],coords[3]), [0,0,255], 15)
	except:
			cv2.line(img, (0,0), (640,480), (0,0,255), 15)
	return

with picamera.PiCamera() as camera:
	camera.resolution = (640, 480)
	camera.framerate = 24
	stream = io.BytesIO()
	
	while True:
		camera.capture(stream, format="jpeg", use_video_port=True)
		frame = np.fromstring(stream.getvalue(), dtype=np.uint8)
		stream.seek(0)
		
		frame = cv2.imdecode(frame, cv2.IMREAD_GRAYSCALE)
		#frame = roi(frame)
		cv2.imshow('Image- ROI', frame)
		
		frame = process_image(frame)
		
		cv2.imshow('Image', frame)
		if cv2.waitKey(1) & 0xFF == ord('q'):
			break

cv2.destroyAllWindows()
