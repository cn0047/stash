import cv2
import numpy

f= './z.png'

frame = cv2.imread(f)
height, width, *_ = frame.shape
top    = int(0.14 * height)
bottom = int(0.25 * height)
left   = int(0.60 * width)
right  = int(0.71 * width)
blur_bbox = frame[top:bottom, left:right]
sigmaX = 50
sigmaX = 20
sigmaX = 10
ksize = (11, 11)
ksize = (99, 99)
blur_bbox = cv2.GaussianBlur(blur_bbox, ksize, sigmaX, sigmaY=40)
frame[top:bottom, left:right] = blur_bbox
cv2.imwrite(f+'.r.png', frame)
