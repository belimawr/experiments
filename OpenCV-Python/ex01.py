import numpy as np
import cv2

cap = cv2.VideoCapture(0)

i = 0
while(True):
    i = i + 1
    # Capture frame-by-frame
    ret, frame = cap.read()

    # Our operations on the frame come here
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    bw = cv2.threshold(gray, 128, 255, cv2.THRESH_BINARY | cv2.THRESH_OTSU)[1]

    # Display the resulting frame
    cv2.imshow('B&W Image', bw)
    cv2.imwrite('video' + str(i) + '.png', bw)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# When everything done, release the capture
cap.release()
cv2.destroyAllWindows()
