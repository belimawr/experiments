import io
import math

import picamera
import cv2

import numpy as np

x_offset = 120
y_offset = 90

x_size = 400
y_size = 300

vertices = np.array([
                      [x_offset, y_offset],
                      [x_offset, y_offset + y_size],
                      [x_offset + x_size, y_offset + y_size],
                      [x_offset + x_size, y_offset]
                    ], np.int32)

def roi(image):
    mask = np.zeros_like(image)
    cv2.fillPoly(mask, [vertices], 255)
    masked = cv2.bitwise_and(image, mask)
    return masked

def process_image(image):
    img = cv2.Canny(image, threshold1=100, threshold2=200)
    img = cv2.GaussianBlur(img, (3,3), 0 )

    lines = cv2.HoughLinesP(img, 1, np.pi/180, 100, minLineLength=150, maxLineGap=10)
    lines = filter(valid_line, map(lambda x: x[0], lines))

    return img, lines

def valid_line(line):
    x1, y1, x2, y2 = line

    if (y1 >= (y_offset + y_size-5)) or (y2 >= (y_offset + y_size-5)) or (y1 <= (y_offset+5)) or (y2 <= (y_offset+5)):
        if angle(line) < 1:
            return False

    if (x1 >= (x_offset + x_size-5)) or (x2 >= (x_offset + x_size-5)) or (x1 <= (x_offset+5)) or (x2 <= (x_offset+5)):
        if angle(line) > 85:
            return False

    return True

def draw_lines(img, lines):
    if lines is None:
        print('No Line :(')
        return

    for i, line in enumerate(lines):
        cv2.line(img, (line[0], line[1]), (line[2], line[3]), 255, 3)


def angle(line):
    x1, y1, x2, y2 = line

    dx = abs(x1 - x2)
    dy = abs(y1 - y2)

    if dx == 0:
        return 0

    m = dy/dx
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
