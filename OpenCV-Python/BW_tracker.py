import numpy as np
import cv2

cap = cv2.VideoCapture(0)

# (480, 848, 3)

while(True):
    # Capture frame-by-frame
    ret, frame = cap.read()

    # Our operations on the frame come here
    bw = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)
    bw = cv2.threshold(bw, 128, 255, cv2.THRESH_BINARY | cv2.THRESH_OTSU)[1]

    frame = cv2.rectangle(bw,(0,350),(848,400),(0,255,0),3)

    # Display the resulting frame
    cv2.imshow('B&W Image', frame)
    if cv2.waitKey(1) & 0xFF == ord('q'):
        break

# When everything done, release the capture
cap.release()
cv2.destroyAllWindows()
